package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan"},
    {ID: 2, Title: "Learning Go", Author: "Jon Bodner"},
}

func getNextID() int {
    if len(books) == 0 {
        return 1
    }
    return books[len(books)-1].ID + 1
}

func main() {
    router := gin.Default()
    router.Use(cors.Default()) // Enable CORS with default settings

    router.GET("/books", func(c *gin.Context) {
        c.JSON(http.StatusOK, books)
    })

    router.POST("/books", func(c *gin.Context) {
        var newBook Book
        if err := c.BindJSON(&newBook); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newBook.ID = getNextID()
        books = append(books, newBook)
        c.JSON(http.StatusCreated, newBook)
    })

	// Delete a book by ID
	router.DELETE("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		for i, book := range books {
			if book.ID == id {
				books = append(books[:i], books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	})

// Update a book by ID
router.PUT("/api/books/:id", func(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }
    for i, book := range books {
        if book.ID == id {
            var updatedBook Book
            if err := c.BindJSON(&updatedBook); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            books[i].Title = updatedBook.Title
            books[i].Author = updatedBook.Author
            c.JSON(http.StatusOK, books[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
})


    router.Run(":8080")
}
