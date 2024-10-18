# Bookshelf API

A simple CRUD (Create, Read, Update, Delete) API built with Go, using Gin framework and MySQL as the database. This API is designed to manage a collection of books, providing endpoints to add, retrieve, update, and delete books. Configuration for the MySQL connection is managed via a `.env` file.

## Features

- **Create** a new book entry
- **Read** all books or a specific book by ID
- **Update** an existing book entry
- **Delete** a book entry

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (version 1.16 or higher)
- MySQL (version 5.7 or higher)
- Postman or any API testing tool (optional for testing the API)

## Installation and Setup

# Clone the Repository
- git clone https://github.com/your-username/bookshelf-api.git
  cd bookshelf-api

## Set Up the .env File

## Install Dependencies
- run : go mod tidy

## Running the Application
- run : go run .

## Project Structure

```bash
bookshelf/
├── main.go           # The entry point for the Go application
├── database.go       # Database connection setup
├── book.go           # CRUD operations for the Book model
├── .env              # Environment variables (MySQL credentials)
└── go.mod            # Go module file for dependency management
