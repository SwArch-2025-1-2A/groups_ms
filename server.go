package main

import (
	"github.com/gin-gonic/gin"
)

// const defaultPort = "8080"

func main() {
	// Setup gin
	r := gin.Default()
	// r.POST("/query", graphqlHandler())
	// r.GET("/", playgroundHandler())
	r.Run()
}
