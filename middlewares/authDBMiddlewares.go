package middlewares

import (
	"users-api/config"
	"users-api/models"

	"github.com/labstack/echo/v4"
)

// var db = config.DB

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.Users
	tx := db.Where("email = ? AND password = ?", username, password).First(&user)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}