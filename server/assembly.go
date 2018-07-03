package main

import (
	// "net/http"
	"github.com/dghubble/sling"

	"github.com/gin-gonic/gin"

)

const assemblyBaseUrl string = "https://test.api.promisepay.com"
var AssemblyClient *sling.Sling = sling.New().Base(assemblyBaseUrl).Set(
	"Authorization",
	"Basic ZmFpY2hlbnNoaW5nQGdtYWlsLmNvbTpZVFk0TURKalpEaGpNelpsWXpVMk5qTXhaRGsxTlRKaVlURmpZalU0WlRjPQ==",
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
