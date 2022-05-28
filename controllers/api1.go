package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Api1(ctx echo.Context) error {
	ctx.Logger().Debug("test api1")
	return ctx.String(http.StatusOK, "success")
}
