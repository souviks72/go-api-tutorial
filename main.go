package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "1", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) { c.IndentedJSON(http.StatusOK, books) }

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)
	if err!= nil{
		c.JSON(http.StatusNotFound, map[string]string{"msg": err.Error()})
	}else{
		c.IndentedJSON(http.StatusOK, book)
	}	
}

func getBookById(id string) (*book, error) {
	for  _, b := range books{
		if id == b.ID{
			return &b, nil
		}
	}

	return nil, errors.New("Book not found")
}

func createBook(c *gin.Context){
	var newBook book 

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/book/:id", bookById)
	router.POST("/book", createBook)
	router.Run("localhost:8080")
}