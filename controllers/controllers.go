package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sort "github.com/rotinoo/sortrithm/sort_algo"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Sortrithm",
	})
}

func BogoSortRand(c *gin.Context) {
	array_lenght := 10
	Data := sort.BogoSortRand(array_lenght)
	c.IndentedJSON(http.StatusOK, Data)
}

type RequestData struct {
	Array []int `json:"array"`
}

func BogoSort(c *gin.Context) {
	var requestData RequestData

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	result := sort.BogoSort(requestData.Array)
	c.IndentedJSON(http.StatusOK, result)
}
