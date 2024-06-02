package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/this-is-mjk/SPO/p3/pkg/handlers"
	"github.com/this-is-mjk/SPO/p3/pkg/helpers"
)

func main() {
	// load env file
	helpers.LoadEnv()

	// connect to the database
	users := helpers.Connect()

	// routes of the server
	router := gin.Default()
	router.Use(cors.Default())

	// signup
	router.POST("/signup", func(c *gin.Context) {
		handlers.Signup(c, users)
	})

	// login
	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, users)
	})

	// run at port 8080
	router.Run(":8080")
}
