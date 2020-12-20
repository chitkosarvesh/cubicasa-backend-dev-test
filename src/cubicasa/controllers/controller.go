package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var app *gin.Engine

//Setup - Starts initilizing the controllers
func Setup() {

	if isRelease := viper.GetBool("release"); isRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	app = gin.Default()
	app.GET("/", home)
	// HUB OPERATIONS
	hubGroup := app.Group("/hub")
	{
		hubGroup.GET("", hubList)
		hubGroup.GET("/:id", hubReadOne)
		hubGroup.POST("", hubCreate)
		hubGroup.POST("/:id", hubUpdate)
		hubGroup.PUT("/:id", hubUpdate)
		hubGroup.DELETE("/:id", hubDelete)
	}
	// TEAM OPERATIONS
	teamGroup := app.Group("/team")
	{
		teamGroup.GET("", teamList)
		teamGroup.GET("/:id", teamReadOne)
		teamGroup.POST("", teamCreate)
		teamGroup.POST("/:id", teamUpdate)
		teamGroup.PUT("/:id", teamUpdate)
		teamGroup.DELETE("/:id", teamDelete)
	}
	// USER OPERATIONS
	userGroup := app.Group("/user")
	{
		userGroup.GET("", userList)
		userGroup.GET("/:id", userReadOne)
		userGroup.POST("", userCreate)
		userGroup.POST("/:id", userUpdate)
		userGroup.PUT("/:id", userUpdate)
		userGroup.DELETE("/:id", userDelete)
	}
	// SEARCH OPERATIONS
	app.GET("search/:query", search)

	app.Run()
}
func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Cubicasa",
	})
}
