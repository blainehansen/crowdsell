package main

import (
	// "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/blainehansen/goqu"
)

var _ r = route(GET, "/projects", func(c *gin.Context) {
	projectFilters := struct {
		Search string `form:"q"`
		UserId int64 `form:"u"`
	}{}
	c.ShouldBindQuery(&projectFilters)

	projectQuery := Projects.Table.Join(
		goqu.I("users"), goqu.On(goqu.Ex{ "projects.user_id": goqu.I("users.id") }),
	).Select(
		Projects.UrlSlug.I(), Projects.Name.I(), Projects.Description.I(),
		Users.UrlSlug.As("user_url_slug"),
		Users.Name.As("user_name"),
	)

	var filterExpressions []goqu.Expression
	if projectFilters.Search != "" {
		filterExpressions = append(filterExpressions, Projects.GeneralSearchVector.Search(projectFilters.Search))
	}
	if projectFilters.UserId != 0 {
		filterExpressions = append(filterExpressions, Projects.UserId.Eq(projectFilters.UserId))
	}
	if len(filterExpressions) > 0 {
		projectQuery = projectQuery.Where(filterExpressions...)
	}


	type PublicProject struct {
		UrlSlug string
		Name string
		Description string
		UserUrlSlug string
		UserName string
	}
	var projects []PublicProject

	err := projectQuery.Limit(10).ScanStructs(&projects)

	if err != nil {
		c.AbortWithError(500, err); return
	}

	c.JSON(200, &projects)
})


var _ r = authRoute(POST, "/projects", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	project := struct {
		Name string
		Description string
		UrlSlug string
	}{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithError(422, err); return
	}

	var projectSlug string
	found, err := Projects.Query.Returning(Projects.Slug).Insert(
		Projects.Name.Set(project.Name),
		Projects.Description.Set(project.Description),
		Projects.UrlSlug.Set(project.UrlSlug),
		Projects.UserId.Set(userId),
	).ScanVal(&projectSlug)
	if err != nil {
		c.AbortWithError(500, err); return
	}
	if !found {
		c.AbortWithError(500, fmt.Errorf("projectSlug not found? %s", projectSlug)); return
	}

	c.JSON(200, &projectSlug)
})


var _ r = authRoute(PATCH, "/projects/:projectSlug", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	projectSlug := c.Param("projectSlug")
	projectId, decodeError := decodeSlug(projectSlug)
	if decodeError != nil {
		c.AbortWithStatus(400); return
	}

	var projectMap map[string]interface{}
	if err := c.ShouldBindJSON(&projectMap); err != nil {
		c.AbortWithError(422, err); return
	}

	patchQuery := Projects.Query.Where(
		Projects.Id.Eq(projectId), Projects.UserId.Eq(userId),
	).Patch(projectMap)

	if !doPatch(c, patchQuery) { return }

	c.Status(204)
})


// they can vote:
// - no, this project doesn't fulfill its promises (with a list of specifically which ones), or a flag of "fraudulent" or "nothing"
// - yes, this project does fulfill its promises
// - almost, this project tried but I expected more (with a list of promises that are lacking and a text field of commentary)

// they can give feedback of almost and say either yes or no to fulfills (depending on whether they think the almosts are big enough to warrant a no)
type projectConfirmation struct {
	Fulfills struct {
		Proceed bool
		AlmostPromises []string
	} `binding:"-"`

	Unacceptable struct {
		FraudulentFlag bool
		BrokenPromises []string
	} `binding:"-"`

	Commentary string `binding:"-"`
}

// what should the table structure be?
// everything gets way more complex if we don't inline
//


// this route will take the feedback of a particular user in about a project
var _ r = authRoute(POST, "/projects/:projectSlug/confirmation", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	projectSlug := c.Param("projectSlug")
	projectId, decodeError := decodeSlug(projectSlug)
	if decodeError != nil {
		c.AbortWithStatus(400); return
	}

	var confirmation projectConfirmation
	if err := c.ShouldBindJSON(&confirmation); err != nil {
		c.AbortWithError(422, err); return
	}

	fmt.Println(confirmation)


	fulfills := confirmation.Fulfills
	lenAlmostPromises := len(fulfills.AlmostPromises)
	fulfillsHas := fulfills.Proceed || lenAlmostPromises != 0

	unacceptable := confirmation.Unacceptable
	lenBrokenPromises := len(unacceptable.BrokenPromises)
	unacceptableHas := unacceptable.FraudulentFlag || lenBrokenPromises != 0

	// if they are both full or both not full
	if fulfillsHas == unacceptableHas {
		c.AbortWithError(422, fmt.Errorf("can't both or neither fulfill and unacceptable %s", confirmation)); return
	}

	lenCommentary := len(confirmation.Commentary)
	if fulfills.Proceed && lenAlmostPromises == 0 && lenCommentary != 0 {
		c.AbortWithError(422, fmt.Errorf("can't simply proceed and give commentary", confirmation)); return
	}

	sets := []SetExpression {
		ProjectConfirmations.Proceed.Set(fulfills.Proceed),
		ProjectConfirmations.FraudulentFlag.Set(unacceptable.FraudulentFlag),
	}


	if lenAlmostPromises != 0 {
		sets = append(sets, ProjectConfirmations.AlmostPromises.Set(fulfills.AlmostPromises))
	} else {
		sets = append(sets, ProjectConfirmations.AlmostPromises.Empty())
	}
	if lenBrokenPromises != 0 {
		sets = append(sets, ProjectConfirmations.BrokenPromises.Set(unacceptable.BrokenPromises))
	} else {
		sets = append(sets, ProjectConfirmations.BrokenPromises.Empty())
	}
	if lenCommentary != 0 {
		sets = append(sets, ProjectConfirmations.Commentary.Set(confirmation.Commentary))
	} else {
		sets = append(sets, ProjectConfirmations.Commentary.Clear())
	}
	updateRecord := makeRecord(sets)

	sets = append(sets, ProjectConfirmations.ProjectId.Set(projectId))
	sets = append(sets, ProjectConfirmations.UserId.Set(userId))
	insertRecord := makeRecord(sets)

	exec := ProjectConfirmations.Table.InsertConflict(
		goqu.DoUpdate("ON CONSTRAINT project_confirmations_unique_project_user", updateRecord).Where(
			ProjectConfirmations.ProjectId.Eq(projectId),
			ProjectConfirmations.UserId.Eq(userId),
		),
		insertRecord,
	)

	if !doExec(c, exec) { return }

	c.Status(204)
})
