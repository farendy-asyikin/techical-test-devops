package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struktur Book untuk mewakili data buku
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Fungsi untuk membuat buku baru
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO books (title, author) VALUES (?, ?)"
	_, err := db.Exec(query, book.Title, book.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book created successfully!"})
}

// Fungsi untuk mendapatkan semua buku
func GetBooks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan book"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// Fungsi untuk mendapatkan satu buku berdasarkan ID
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book Book

	query := "SELECT id, title, author FROM books WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Fungsi untuk memperbarui buku berdasarkan ID
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE books SET title = ?, author = ? WHERE id = ?"
	_, err := db.Exec(query, book.Title, book.Author, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully!"})
}

// Fungsi untuk menghapus buku berdasarkan ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM books WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully!"})
}
