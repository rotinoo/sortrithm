package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rotinoo/sortrithm/routes"
)

func main() {
	router := gin.Default()

	routes.InitializeRoutes(router)

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")

	router.Run(":8080")
}
