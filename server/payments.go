package main

import (
	// "net/http"
	"fmt"
	"github.com/dghubble/sling"

	"github.com/gin-gonic/gin"

	"github.com/blainehansen/goqu"
)

type ProjectPledgesStateEnum string
const (
	UNPAID ProjectPledgesStateEnum = "UNPAID"
	PAID ProjectPledgesStateEnum = "PAID"
	RELEASED ProjectPledgesStateEnum = "RELEASED"
)

var assemblyBaseUrl string = environment["ASSEMBLY_ENDPOINT"]

var AssemblyClient *sling.Sling = sling.New().Base(assemblyBaseUrl).Set(
	"Authorization",
	fmt.Sprintf("Basic %s", environment["ASSEMBLY_AUTH"]),
)

type AssemblyStatus struct {
	Status string `json:"status"`
}

var _ r = route(GET, "/assembly/status", func(c *gin.Context) {
	status := AssemblyStatus{}
	_, err := AssemblyClient.New().Get("/status").Receive(&status, nil)
	if err != nil {
		c.AbortWithError(503, err); return
	}

	c.JSON(200, &status)
})


// // create a payment user
// var _ r = authRoute(PATCH, "user/payment/user", func(c *gin.Context) {
// 	userId := c.MustGet("userId").(string)

// })

// var _ r = authRoute(POST, "/user/payment/account/:type", func(c *gin.Context) {
// 	userId := c.MustGet("userId").(string)

// 	// account
// })


// var _ r = route(GET, "/project/:projectSlug", func(c *gin.Context) {
// 	projectId, decodeError := decodeSlug(c.Param("projectSlug"))
// 	if decodeError != nil {
// 		c.AbortWithError(400, decodeError); return
// 	}

// 	Projects.Table.Where(
// 		Projects.Id.Eq(projectId),
// 	).Select(
// 		Projects.UrlSlug.I(), Projects.Name.I(), Projects.Description.I(),
// 		Users.UrlSlug.As("user_url_slug"),
// 		Users.Name.As("user_name"),
// 		goqu.L("(select sum(amount) from project_pledges where project_id = projects.id) as amount"),
// 	).Join(
// 		goqu.I("users"), goqu.On(goqu.Ex{ "projects.user_id": goqu.I("users.id") }),
// 	)


// 	Projects.Table.Select(
// 		Projects.UrlSlug.I(), Projects.Name.I(), Projects.Description.I(),
// 		Users.UrlSlug.As("user_url_slug"),
// 		Users.Name.As("user_name"),
// 		goqu.L("sum(project_pledges.amount) as amount"),
// 	).Join(
// 		goqu.I("users"), goqu.On(goqu.Ex{ "projects.user_id": goqu.I("users.id") }),
// 	).LeftJoin(
// 		goqu.I("project_pledges"), goqu.On(goqu.Ex{ "projects.id": goqu.I("project_pledges.project_id") })
// 	).GroupBy(
// 		goqu.L("1, 2, 3, 4, 5"),
// 	)

// 	// select projects.url_slug, projects.name, users.url_slug, users.name, sum(project_pledges.amount) from projects
// 	// join users on projects.user_id = users.id
// 	// left join project_pledges on projects.id = project_pledges.project_id and project_pledges.state = 'PAID'
// 	// group by 1, 2, 3, 4;

// })

type AssemblyTokenResponse struct {
	TokenAuth struct {
		TokenType string `json:"token_type"`
		UserId string `json:"user_id"`
		Token string `json:"token"`
	} `json:"token_auth"`
}

var _ r = route(POST, "/user/card-token", func(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	createTokenBody := struct {
		TokenType string `json:"token_type"`
		UserId string `json:"user_id"`
	}{
		TokenType: "card",
		UserId: userId,
	}


	createTokenReceive := AssemblyTokenResponse{}
	_, createItemErr := AssemblyClient.New().Post("/token_auths").BodyJSON(&createTokenBody).Receive(&createTokenReceive, nil)
	if createItemErr != nil {
		c.AbortWithError(500, createItemErr); return
	}

	c.JSON(200, &createTokenReceive.TokenAuth.Token)
})

// var _ r = authRoute(POST, "/user/bank-token", func(c *gin.Context) {
// 	//
// })

var _ r = authRoute(POST, "/pledge/:projectId", func(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	payment := struct {
		// this will be in some non-fractional unit, like cents
		Amount int64
		AccountId string
		IpAddress string
		DeviceId string
		// CurrencyCode string
	}{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.AbortWithError(422, err); return
	}

	projectId := c.Param("projectId")

	var projectUserId string
	projectUserIdFound, projectUserIdErr := db.From("project").Select(
		goqu.I("person_id"),
	).Where(
		goqu.I("id").Eq(projectId),
	).ScanVal(&projectUserId)

	if projectUserIdErr != nil {
		c.AbortWithError(500, projectUserIdErr); return
	}
	if !projectUserIdFound {
		c.AbortWithError(500, fmt.Errorf("projectUserId not found? %s", projectUserId)); return
	}


	var pledgeId string
	pledgeIdFound, pledgeIdErr := db.From("project_pledge").Returning(
		goqu.I("id"),
	).Insert(
		goqu.Record {
			"project_id": projectId,
			"person_id": userId,
			"amount": payment.Amount,
		},
	).ScanVal(&pledgeId)

	if pledgeIdErr != nil {
		c.AbortWithError(500, pledgeIdErr); return
	}
	if !pledgeIdFound {
		c.AbortWithError(500, fmt.Errorf("pledgeId not found? %s", pledgeId)); return
	}


	// create an item
	createItemBody := struct {
		Id string
		Amount int64
		BuyerId string
		SellerId string
	}{
		Id: pledgeId,
		Amount: payment.Amount,
		BuyerId: userId,
		SellerId: projectUserId,
	}
	_, createItemErr := AssemblyClient.New().Post("/items").BodyJSON(&createItemBody).ReceiveSuccess(nil)
	if createItemErr != nil {
		c.AbortWithError(500, createItemErr); return
	}

	// make a payment on that item
	makePaymentBody := struct {
		AccountId string
		IpAddress string
		DeviceId string
	}{
		AccountId: payment.AccountId,
		IpAddress: payment.IpAddress,
		DeviceId: payment.DeviceId,
	}
	_, makePaymentErr := AssemblyClient.New().Patch(
		fmt.Sprintf("/items/%s/make_payment", pledgeId),
	).BodyJSON(&makePaymentBody).ReceiveSuccess(nil)
	if makePaymentErr != nil {
		c.AbortWithError(500, makePaymentErr); return
	}

	// else update its state to PAID
	updateQuery := db.From("project_pledge").Where(
		goqu.I("id").Eq(pledgeId),
	).Update(
		goqu.Record{ "state": PAID },
	)

	if !doExec(c, updateQuery) { return }

	c.Status(204)
})


func ReleaseProjectFunds(projectId string) []string {
	var pledges []string
	db.From("project_pledge").Where(
		goqu.I("project_id").Eq(projectId),
		goqu.I("state").Eq(PAID),
	).Select(
		goqu.I("id"),
	).ScanVals(&pledges)

	unsuccessfulPledges := []string{}
	for _, pledgeId := range pledges {
		_, makePaymentErr := AssemblyClient.New().Patch(
			fmt.Sprintf("/items/%s/release_payment", pledgeId),
		).ReceiveSuccess(nil)

		if makePaymentErr != nil {
			unsuccessfulPledges = append(unsuccessfulPledges, pledgeId)
		}
	}

	return unsuccessfulPledges
}
