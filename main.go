package main

import (
	"books_api/routes"
	"books_api/storage"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()

	booksRoutes := e.Group("/books")

	//authorsRoutes := e.Group("/authors")

	routes.BooksRoutes(booksRoutes)

	//routes.AuthorRoutes(authorsRoutes)

	storage.ConnectDB()

	e.Logger.Fatal(e.Start(":3000"))
}
