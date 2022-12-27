package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var books []Book = []Book{Book{Id: 1, Title: "Some Title1", Price: 311.22}}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBooks)
	router.DELETE("/books/:id", deleteBooks)

	router.Run()
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func addBooks(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

func deleteBooks(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.Id == bookID {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}
