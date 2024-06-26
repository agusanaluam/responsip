package context

import "github.com/labstack/echo"

type EchoContext struct {
	Context echo.Context
}

func (e EchoContext) JSON(statusCode int, v interface{}) error {
	return e.Context.JSON(statusCode, v)
}

func (e EchoContext) GetHeader(key string) string {
	return e.Context.Request().Header.Get(key)
}
