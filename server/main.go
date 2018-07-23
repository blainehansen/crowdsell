package main

import (
	// "os"
	// "log"
	"fmt"
	"time"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"

	"github.com/json-iterator/go/extra"
)


var environment *map[string]string = func() *map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		panic("error reading .env file")
	}
	return &env
}()


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

var dbUserStore *UserStore
var dbProjectStore *ProjectStore

func main() {
	// CONNECTING TO DATABASE
	db, connectionError := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			environment["DATABASE_HOST"],
			environment["DATABASE_PORT"],
			environment["DATABASE_DB_NAME"],
			environment["DATABASE_USER"],
			environment["DATABASE_PASSWORD"],
			environment["DATABASE_SSL"],
		),
	)
	if connectionError != nil {
		fmt.Println(connectionError)
		panic("failed to connect to database")
	}
	defer db.Close()

	// INSTANTIATING STORES
	dbUserStore = NewUserStore(db)
	dbProjectStore = NewProjectStore(db)

	// CHANGING JSON NAMING CONVENTION
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)

	// SETTING UP ROUTER
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AddAllowHeaders("Authorization")
	config.MaxAge = 24 * time.Hour

	router.Use(cors.New(config))

	// SETTING UP ROUTES
	addRoutesToGroup(router, routes)

	authorized := router.Group("/secure")
	authorized.Use(VerifyTokenMiddleWare)
	addRoutesToGroup(authorized, authRoutes)

	routes = nil
	authRoutes = nil

	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.Run(":5050")
}
