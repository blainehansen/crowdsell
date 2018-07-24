package main

import (
	// "strconv"
	"github.com/gin-gonic/gin"
)

var _ r = route(GET, "/projects", func(c *gin.Context) {
	projects, err := dbProjectStore.FindAll(
		NewProjectQuery().WithUser().Select(
			Schema.Project.Slug, Schema.Project.InternalSlug, Schema.Project.Name, Schema.Project.Description,
			Schema.User.InternalSlug, Schema.User.Name,
		).Limit(5),
	)

	if err != nil {
		c.AbortWithError(500, err); return
	}

	c.JSON(200, &projects)
})

var _ r = authRoute(POST, "/projects", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithError(422, err); return
	}
	project.User = User { Id: userId }

	if err := dbProjectStore.Insert(&project); err != nil {
		c.AbortWithError(500, err); return
	}

	c.JSON(200, &project)
})

// var _ r = authRoute(PATCH, "/projects/:projectId", func(c *gin.Context) {
// 	userId := c.MustGet("userId").(int64)

// 	projectId, parseErr := strconv.ParseInt(c.Param("projectId"), 10, 32)
// 	if parseErr != nil {
// 		c.AbortWithError(400, parseErr); return
// 	}

// 	var project map[string]interface{}
// 	bindErr := BindJSONWithTemplate(c, &project, Project{})
// 	if bindErr != nil {
// 		c.AbortWithStatus(404); return
// 	}

// 	queryResult := db.Model(Project{}).Where("id = ?", projectId).Where("user_id = ?", userId).Updates(project)
// 	rowsUpdated, err := dbProjectStore.Update(&project, Schema.User.Slug)
// 	if err := queryResult.Error; err != nil {
// 		c.AbortWithError(500, err); return
// 	}
// 	if rowsUpdated == 0 {
// 		c.AbortWithStatus(404); return
// 	}

// 	c.Status(204)
// })
