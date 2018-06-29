package main

import (
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

var _ r = route(POST, "/projects", func(c *gin.Context) {
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithError(422, err)
		return
	}

	if err := db.Create(&project).Error; err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, &project)
})
