package main

import (
	// "errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/tools/go/analysis/passes/ifaceassert"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return //IndentedJSON is going to take care of this return statement
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Book not available"})
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
func returnBook(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book not found"})
		return
	}
  book.Quantity += 1 
  c.IndentedJSON(http.StatusOK,book)

}

func createBook(c *gin.Context) {
	//newBook here is the variable of type book
	var newBook book
	//&newBook is the pointer to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return //BindJSON will handle the return statement for us
	}
	books = append(books, newBook) //this means books will be appended to newbook
	c.IndentedJSON(http.StatusCreated, newBook)

}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
  router.PATCH("/return",returnBook)
	router.Run("localhost:8080")
}
