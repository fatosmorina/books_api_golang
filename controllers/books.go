package controllers

import (
	"books_api/services"
	"books_api/types"
	"books_api/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetBook(c echo.Context) error {
	id := c.Param("book_id")
	book, err := services.GetBook(id)
	if err != nil {
		fmt.Println("Book was not found")
		return utils.HttpResponse(c, http.StatusNotFound, "Book was not found")
	}
	return utils.HttpResponse(c, http.StatusOK, book)
}

func AddBook(c echo.Context) error {
	bookInputBody := &types.BookInputBody{}
	err := c.Bind(bookInputBody)
	if err != nil {
		fmt.Println("Input not valid", err)
		return utils.HttpResponse(c, http.StatusBadRequest, err)
	}
	errorWhileAddingABook := services.AddBook(bookInputBody)
	if errorWhileAddingABook != nil {
		return utils.HttpResponse(c, http.StatusBadRequest, errorWhileAddingABook)
	}
	return utils.HttpResponse(c, http.StatusOK, "Book was added successfully")
}

func UpdateBook(c echo.Context) error {
	bookInputBody := &types.BookInputBody{}
	err := c.Bind(bookInputBody)
	if err != nil {
		fmt.Println("Input not valid", err)
		return utils.HttpResponse(c, http.StatusBadRequest, err)
	}

	updateError := services.UpdateBook(bookInputBody)

	if updateError != nil {
		fmt.Println("Error while updating a book", updateError)
		return utils.HttpResponse(c, http.StatusBadRequest, updateError)
	}
	return utils.HttpResponse(c, http.StatusOK, "Book was updated successfully")
}

func DeleteBook(c echo.Context) error {
	id := c.Param("book_id")
	err := services.DeleteBook(id)
	if err != nil {
		fmt.Println("Error while deleting a book", err)
		return utils.HttpResponse(c, http.StatusBadRequest, err)
	}
	return utils.HttpResponse(c, http.StatusOK, "Book was deleted successfully")
}
