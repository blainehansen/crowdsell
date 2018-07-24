package main

import (
	// "net/http"
	"fmt"
	"github.com/dghubble/sling"

	"github.com/gin-gonic/gin"

)

var assemblyBaseUrl string = environment["ASSEMBLY_ENDPOINT"]

var AssemblyClient *sling.Sling = sling.New().Base(assemblyBaseUrl).Set(
	"Authorization",
	fmt.Sprintf("Basic %s", environment["ASSEMBLY_AUTH"]),
)

type AssemblyStatus struct {
	Status string `json:"status"`
}

var _ r = route(GET, "/assembly/status", func(c *gin.Context) {
	status := AssemblyStatus{}
	_, err := AssemblyClient.New().Get("/status").Receive(&status, nil)
	if err != nil {
		c.AbortWithError(503, err)
		return
	}

	c.JSON(200, &status)
})
