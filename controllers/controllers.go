package controllers

import (
	"net/http"
	"strconv"
	"users-api/lib/database"
	"users-api/models"

	"github.com/labstack/echo/v4"
	// "github.com/thedevsaddam/govalidator"
)

// USERS CONTROLLERS
func CreateUserController(c echo.Context) error {
	var user models.Users
	c.Bind(&user)

	// rules := govalidator.MapData{
	// 	"name":     []string{"required"},
	// 	"email":    []string{"required", "email"},
	// 	"password": []string{"required"},
	// }
	// opts := govalidator.Options{
	// 	Request: c.Request(),
	// 	Data:    &user,
	// 	Rules:   rules,
	// }
	// v := govalidator.New(opts)
	// e := v.Validate()
	// if e != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
	// 		"validationError": e,
	// 	})
	// }
	if e := database.CreateUser(&user); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, e := database.GetUserById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	var user models.Users
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	if e := database.UpdateUserById(id, &user); e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success update user",
		"user":   user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DeleteUserById(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user",
	})
}

func LoginUserController(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	users, e := database.LoginUser(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user": users,
	})
}

func GetUserDetailByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := database.GetDetailUserById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user": user,
	})
}

// BOOKS CONTROLLERS
func AddBookController(c echo.Context) error {
	var book models.Books
	c.Bind(&book)

	// rules := govalidator.MapData{
	// 	"title":  []string{"required"},
	// 	"author": []string{"required"},
	// 	"year":   []string{"required", "numeric"},
	// }
	// opts := govalidator.Options{
	// 	Request: c.Request(),
	// 	Data:    &book,
	// 	Rules:   rules,
	// }
	// v := govalidator.New(opts)
	// // err := map[string]interface{}{"validationError": e}
	// if e := v.ValidateStruct(); e != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
	// 		"validationError": e,
	// 	})
	// }

	if e := database.AddBook(&book); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new book",
		"book":    book,
	})
}

func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, e := database.GetBookById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func UpdateBookByIdController(c echo.Context) error {
	var book models.Books
	c.Bind(&book)

	id, _ := strconv.Atoi(c.Param("id"))

	if e := database.UpdateBookById(id, &book); e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success update book",
		"book":   book,
	})
}

func DeleteBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DeleteBookById(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete book",
	})
}
