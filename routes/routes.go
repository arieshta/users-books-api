package routes

import (
	"users-api/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUserController)
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)

	e.POST("/books", controllers.AddBookController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookByIdController)
	e.PUT("/books/:id", controllers.UpdateBookByIdController)
	e.DELETE("/books/:id", controllers.DeleteBookByIdController)

	return e
}