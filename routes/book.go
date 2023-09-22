package routes

import (
	"books_api/controllers"
	"books_api/middlewares"
	"github.com/labstack/echo/v4"
)

func BooksRoutes(e *echo.Group) {
	e.GET("/:book_id", controllers.GetBook, middlewares.ValidateToken)
	e.POST("", controllers.AddBook, middlewares.ValidateToken)
	//e.PUT("/:book_id", controllers.UpdateBook, middlewares.ValidateToken)
	//e.DELETE("/:book_id", controllers.DeleteBook, middlewares.ValidateToken)
}
