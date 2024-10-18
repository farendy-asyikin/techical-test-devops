package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB() // Inisialisasi database

	r := gin.Default()

	// Define routes for CRUD operations
	r.POST("/books", CreateBook)
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)

	r.Run(":8080") // Jalankan server di port 8080
}
