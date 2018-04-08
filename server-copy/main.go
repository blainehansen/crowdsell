package main

import (
	"os"
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
var err error

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

type Thing struct {
	NotName string
}

func main() {
	r := gin.Default()

	// os.Getenv("FOO")

	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=crowdsell user=user password=Password1")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Project{})

	thing := Thing{"thing"}
	db.Create(&thing)
	fmt.Println(db.Error)

	r.GET("/projects", ListProjects)
	r.POST("/project", CreateProject)

	r.Run(":5050")
}

func ListProjects(c *gin.Context) {
	var projects []Project

	if err = db.Find(&projects).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, projects)
	}
}

func CreateProject(c *gin.Context) {
	var project Project
	// project = ReceiveAndCreate(c, &project).(*Project)
	if err = c.BindJSON(&project); err != nil {
		c.AbortWithError(422, err)
		fmt.Println(err)
	}

	db.Create(&project)
	fmt.Println(db.Error)
	c.JSON(200, &project)
}

// func ReceiveAndCreate(c *gin.Context, s interface{}) interface{} {
// 	if err = c.BindJSON(s); err != nil {
// 		c.AbortWithError(422, err)
// 		fmt.Println(err)
// 	}

// 	db.Create(s)
// 	fmt.Println(db.Error)
// 	return s
// }

// https://github.com/gin-gonic/gin/issues/198
