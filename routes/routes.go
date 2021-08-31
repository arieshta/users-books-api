package routes

import (
	"users-api/config"
	"users-api/controllers"

	// m "users-api/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserController)

	e.POST("/users", controllers.CreateUserController)
	// e.GET("/users", controllers.GetUsersController)
	// e.GET("/users/:id", controllers.GetUserByIdController)
	// e.PUT("/users/:id", controllers.UpdateUserByIdController)
	// e.DELETE("/users/:id", controllers.DeleteUserByIdController)

	// JWT Auth Group
	eJWT := e.Group("/jwt")
	eJWT.Use(middleware.JWT([]byte(config.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserByIdController)
	eJWT.PUT("/users/:id", controllers.UpdateUserByIdController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserByIdController)

	// BasicAuth Group
	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(m.BasicAuthDB))
	// eAuth.GET("/users", controllers.GetUsersController)
	// eAuth.GET("/users/:id", controllers.GetUserByIdController)
	// eAuth.PUT("/users/:id", controllers.UpdateUserByIdController)
	// eAuth.DELETE("/users/:id", controllers.DeleteUserByIdController)

	e.POST("/books", controllers.AddBookController) //
	eJWT.GET("/books", controllers.GetBooksController)
	eJWT.GET("/books/:id", controllers.GetBookByIdController)
	eJWT.PUT("/books/:id", controllers.UpdateBookByIdController)    //
	eJWT.DELETE("/books/:id", controllers.DeleteBookByIdController) //

	return e
}
