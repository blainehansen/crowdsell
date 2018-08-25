package main

import (
	// "net/http"
	"fmt"
	"github.com/dghubble/sling"

	"github.com/gin-gonic/gin"

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


var _ r = authRoute(POST, "/pledge/:projectSlug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	payment := struct {
		AmountWhole int64
		AmountDecimal int64
		IpAddress string
		DeviceId string
	}{}
	// account_id
	// CurrencyCode string
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.AbortWithError(422, err); return
	}

	projectId, decodeError = decodeSlug(c.Param("projectSlug"))
	if decodeError != nil {
		c.AbortWithError(400, decodeError); return
	}


	// create a pledge, returning id, with a bunch of default information

	c.Status(204)
})
