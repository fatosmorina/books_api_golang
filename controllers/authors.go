package controllers

import (
	"books_api/services"
	"books_api/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllBooksForAuthor(c echo.Context) error {
	fmt.Println("Get all books")
	authorId := c.Param("author_id")
	books, err := services.GetBooksForAuthor(authorId)
	if err != nil {
		fmt.Println("Error while getting all books", err)
		return utils.HttpResponse(c, http.StatusBadRequest, err)
	}
	return utils.HttpResponse(c, http.StatusOK, books)
}
