package utils

import "github.com/labstack/echo/v4"

type SuccessContent struct {
	Code    int         `json:"code"`
	Success bool        `json:"success" default:"true"`
	Data    interface{} `json:"data"`
}

func HttpResponse(c echo.Context, statusCode int, data interface{}) error {
	content := SuccessContent{}
	content.Code = statusCode
	content.Data = data
	content.Success = true

	return c.JSON(statusCode, content)
}
