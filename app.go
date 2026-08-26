package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In search of lost time", Author: "Author 1", Quantity: 2},
	{ID: "2", Title: "The Great gatsby", Author: "Author 2", Quantity: 5},
	{ID: "3", Title: "War and peace", Author: "Author 3", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

type Employee interface {
	GetAll() (string, error)
	CreateEmployee(employee struct {
		ID        int64
		FirstName string
		LastName  string
	}) (struct {
		ID        int64
		LastName  string
		FirstName string
	}, error)
}

type EmployeeRepo struct{}

func (e *EmployeeRepo) GetAll() (string, error) {
	return "", nil
}

func (e *EmployeeRepo) CreateEmployee(emp struct {
	ID        int64
	FirstName string
	LastName  string
}) struct {
	ID        int64
	FirstName string
	LastName  string
} {
	return emp
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/books", getBooks)
// 	router.POST("/books", createBook)
// 	router.Run("localhost:8080")
// }
