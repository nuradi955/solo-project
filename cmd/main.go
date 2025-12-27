package cmd
package main

import (
	config "solo-project/data-base"
	"solo-project/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()

	r.POST("/authors", handlers.CreateAuthorHTTP)
	r.GET("/authors", handlers.GetAuthorsHTTP)
	r.GET("/authors/:id", handlers.GetAuthorByID)

	r.POST("/books", handlers.CreateBookHTTP)
	r.GET("/books/:id", handlers.GetBookByID)
	r.PATCH("/books/:id", handlers.UpdateBookHTTP)
	r.GET("/books/available", handlers.GetAvailableBooks)
	r.GET("/books", handlers.GetBooksQuery)

	r.POST("/readers", handlers.CreateReader)
	r.GET("/readers", handlers.GetReaders)
	r.GET("/readers/:id", handlers.GetReaderByID)
	r.GET("/readers/:id/history", handlers.GetRearedHistory)

	r.POST("/borrowings", handlers.BorrowBookHTTP)
	r.PATCH("/borrowings/:id/return", handlers.BorrowingReturn)
	r.GET("/borrowings", handlers.GetBorrowings)
	r.GET("/borrowings/overdue", handlers.GetBorrowingsOverdueHTTP) 
	r.Run()
}