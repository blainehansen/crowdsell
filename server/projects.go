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
