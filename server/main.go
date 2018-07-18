package main

import (
	// "os"
	// "log"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type RouteMethod int

const (
	GET RouteMethod = 0
	POST RouteMethod = 1
	PUT RouteMethod = 2
	PATCH RouteMethod = 3
	DELETE RouteMethod = 4
)

type Route struct {
	Path string
	Method RouteMethod
	Handler func(c *gin.Context)
}

var routes []Route = []Route{}
var authRoutes []Route = []Route{}

type r struct {}

func route(method RouteMethod, path string, handler gin.HandlerFunc) r {
	routeStruct := Route { Path: path, Method: method, Handler: handler }
	routes = append(routes, routeStruct)
	var emptyR r
	return emptyR
}
func authRoute(method RouteMethod, path string, handler gin.HandlerFunc) r {
	routeStruct := Route { Path: path, Method: method, Handler: handler }
	authRoutes = append(authRoutes, routeStruct)
	var emptyR r
	return emptyR
}


var db *gorm.DB

func addRoutesToGroup(router gin.IRouter, routesArray []Route) {
	for _, routeStruct := range routesArray {
		switch routeStruct.Method {
			case GET:
				router.GET(routeStruct.Path, routeStruct.Handler)
			case POST:
				router.POST(routeStruct.Path, routeStruct.Handler)
			case PUT:
				router.PUT(routeStruct.Path, routeStruct.Handler)
			case PATCH:
				router.PATCH(routeStruct.Path, routeStruct.Handler)
			case DELETE:
				router.DELETE(routeStruct.Path, routeStruct.Handler)
			default:
				fmt.Println(routeStruct)
				panic("created a Route with an invalid Method")
		}
	}
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AddAllowHeaders("Authorization")
	config.MaxAge = 24 * time.Hour

	router.Use(cors.New(config))

	var connectionError error
	db, connectionError = gorm.Open("postgres", "host=go-database port=5432 dbname=dev_database user=user password=asdf sslmode=disable")
	if connectionError != nil {
		fmt.Println(connectionError)
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Project{}, &User{})

	addRoutesToGroup(router, routes)

	authorized := router.Group("/secure")
	authorized.Use(VerifyTokenMiddleWare)
	addRoutesToGroup(authorized, authRoutes)

	routes = nil
	authRoutes = nil

	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.Run(":5050")
}
