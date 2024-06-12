package context

import "github.com/gofiber/fiber/v2"

type FiberContext struct {
	Context fiber.Ctx
}

func (f FiberContext) JSON(statusCode int, v interface{}) error {
	return f.Context.JSON(v)
}

func (f FiberContext) GetHeader(key string) string {
	return string(f.Context.Request().Header.Peek(key))
}
