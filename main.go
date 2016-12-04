// main.go

package main

import "github.com/gin-gonic/gin"

var route *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	route = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	route.LoadHTMLGlob("static/templates/*")

	initRoute(route) // in route.go

	// Start serving the application
	route.Run()
}
