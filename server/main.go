package main

import (
	// "os"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
// 	return defaultTableName;
// }

var db *gorm.DB

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Project struct {
	Model
	Name string `json:"name"`
}


func main() {
	r := gin.Default()

	// os.Getenv("FOO")

	var connectionError error
	db, connectionError = gorm.Open("postgres", "host=localhost port=5432 dbname=crowdsell user=user password=Password1 sslmode=disable")
	if connectionError != nil {
		fmt.Println(connectionError)
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Project{})

	r.GET("/projects", ListProjects)
	r.POST("/projects", CreateProject)

	r.Run(":5050")
}

func ListProjects(c *gin.Context) {
	var projects []Project

	if err := db.Find(&projects).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.JSON(200, projects)
}

func CreateProject(c *gin.Context) {
	var project Project
	if err := c.BindJSON(&project); err != nil {
		c.AbortWithError(422, err)
		fmt.Println(err)
		return
	}

	if err := db.Create(&project).Error; err != nil {
		c.AbortWithError(500, err)
		fmt.Println(err)
		return
	}

	c.JSON(200, &project)
}

// https://github.com/gin-gonic/gin/issues/198
