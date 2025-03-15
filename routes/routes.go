package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rotinoo/sortrithm/controllers"
)

func InitializeRoutes(router *gin.Engine) {
	// Define your routes here
	router.GET("/", controllers.ShowIndexPage)
	router.GET("/bogorand", controllers.BogoSortRand)
	router.POST("/bogo", controllers.BogoSort)
	router.POST("/selection", controllers.SelectionSort)
}
