package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

var _ r = route(GET, "/projects", func(c *gin.Context) {
	var projects []Project
	if err := db.Find(&projects).Error; err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, &projects)
})

var _ r = authRoute(POST, "/projects", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithError(422, err)
		return
	}
	project.UserId = userId

	if err := db.Create(&project).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, &project)
})

var _ r = authRoute(PATCH, "/projects/:projectId", func(c *gin.Context) {
	userId := c.MustGet("userId").(int64)

	projectId, parseErr := strconv.ParseUint(c.Param("projectId"), 10, 32)
	if parseErr != nil {
		c.AbortWithError(400, parseErr); return
	}

	var project map[string]interface{}
	bindErr := BindJSONWithTemplate(c, &project, Project{})
	if bindErr != nil {
		c.AbortWithStatus(404); return
	}

	queryResult := db.Model(Project{}).Where("id = ?", projectId).Where("user_id = ?", userId).Updates(project)
	if err := queryResult.Error; err != nil {
		c.AbortWithError(500, err); return
	}
	if queryResult.RowsAffected == 0 {
		c.AbortWithStatus(404); return
	}

	c.Status(204)
})


