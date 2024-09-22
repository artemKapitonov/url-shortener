package dto

import (
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	ctx echo.Context
	Msg string
}

func SendError(ctx echo.Context, msg string, code int) error {
	return ctx.JSON(code, ErrorResponse{Msg: msg})
}
