package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/doug-martin/goqu.v4"
)

var _ r = route(GET, "/projects", func(c *gin.Context) {
	type PublicProject struct {
		UrlSlug string
		Name string
		Description string
		UserUrlSlug string
		UserName string
	}
	var projects []PublicProject
	err := db.From("projects").LeftJoin(
		goqu.I("users"), goqu.On(goqu.Ex{ "projects.user_id": goqu.I("users.id") }),
	).Select(
		"url_slug", "name", "description",
		goqu.I("users.url_slug").As("user_url_slug"),
		goqu.I("users.name").As("user_name"),
	).Limit(5).ScanStructs(&projects)

	if err != nil {
		c.AbortWithError(500, err); return
	}

	c.JSON(200, &projects)
})

var _ r = authRoute(POST, "/projects", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	type EditableProject struct {
		Slug string
		Name string
		Description string
		UrlSlug string
		UserId int64
	}
	var project EditableProject
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithError(422, err); return
	}
	if project.Slug != "" {
		c.AbortWithError(422, errors.New("bad field: slug")); return
	}
	if project.UserId != 0 {
		c.AbortWithError(422, errors.New("bad field: user_id")); return
	}
	project.UserId = userId

	var slug string
	if found, err := db.From("projects").Returning("slug").Insert(&project).ScanVal(&slug); err != nil || !found {
		c.AbortWithError(500, err); return
	}
	project.Slug = slug

	c.JSON(200, &project)
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
	// TODO need to validate the contents of projectMap

	result, updateError := db.From("projects").Where(goqu.Ex{ "id": projectId, "user_id": userId }).Update(projectMap).Exec()
	if updateError != nil {
		c.AbortWithError(500, updateError); return
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		c.AbortWithStatus(404)
	}

	c.Status(204)
})
