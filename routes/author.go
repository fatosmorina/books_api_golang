package routes

import (
	"books_api/controllers"
	"books_api/middlewares"
	"github.com/labstack/echo/v4"
)

func AuthorRoutes(e *echo.Group) {
	e.GET("/:author_id", controllers.GetAllBooksForAuthor, middlewares.ValidateToken)
}
