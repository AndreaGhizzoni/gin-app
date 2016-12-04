// route.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRoute(route *gin.Engine) {
	route.GET("/", showIndex)
	route.GET("/article/view/:article_id", showArticle)
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data)
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data)
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

// Define the route for the index page and display the index.html template
// To start with, we'll use an inline route handler. Later on, we'll create
// standalone functions that will be used as route handlers.
func showIndex(c *gin.Context) {
	articles := getAllArticles()

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles},
		"index.html")
}

func showArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			render(
				c,
				gin.H{
					"title":   article.Title,
					"payload": article},
				"article.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
