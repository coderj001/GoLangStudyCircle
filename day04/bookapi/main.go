package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var books []Book = [{Book{Id: 1, Title: "Some Title1", Price: 311.22}]

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBooks)
	router.DELETE("/books", deleteBooks)

	router.Run()
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
func addBooks(c *gin.Context)    {}
func deleteBooks(c *gin.Context) {}
