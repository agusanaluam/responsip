package context

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	Context *gin.Context
}

func (g GinContext) JSON(statusCode int, v interface{}) {
	g.Context.JSON(statusCode, v)
	return
}

func (g GinContext) GetHeader(key string) string {
	return g.Context.GetHeader(key)
}
