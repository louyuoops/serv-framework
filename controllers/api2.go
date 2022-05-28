package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Api2(ctx echo.Context) error {
	ctx.Logger().Debug("this is test api2")
	return ctx.String(http.StatusOK, "success")
}
