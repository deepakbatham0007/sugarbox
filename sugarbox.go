package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/deepakbatham0007/sugarbox/handler"
)

var router *gin.Engine

func indexProcess(c *gin.Context) {
	handler.Logout(c)
	session := sessions.Default(c)
	sessUser:=session.Get("user")
	if sessUser != nil {
		session.Set("user", sessUser)
		session.Save()
	}

	

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		gin.H{
			"title": "",
		},
	)

}
func main() {

	// Set the router as the default
	router = gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html
	router.GET("/", indexProcess)
	router.POST("/login", handler.Login)
	router.POST("/logout", handler.Logout)
	router.POST("/add", handler.AddItemType)
	//router.POST("/search", handler.Middlew, handler.SearchItem)
	router.POST("/search", handler.SearchItem)
	router.POST("/searchuseritems", handler.SearchUserItems)
	router.POST("/rating", handler.RateItem)
	router.POST("/comment", handler.CommentItem)

	// Start serving the application
	router.Run()

}
