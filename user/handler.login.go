package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetLogin(ctx echo.Context) error{
	username, password, ok := ctx.Request().BasicAuth()
	failed := map[string]interface{}{
		"error":true,
		"message":"unauthorized",
	}
	
	if !ok{
		return ctx.JSON(http.StatusNotAcceptable, failed)
	}
	u, err := Auth(username, password)
	if err == nil && u.Username != ""{
		success := map[string]interface{}{
			"message": "login berhasil",
		}
		return ctx.JSON(http.StatusOK, success)
	}
	
	return ctx.JSON(http.StatusNotAcceptable, failed)
}