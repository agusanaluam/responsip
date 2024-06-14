package responsip

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinContext struct {
	Context *gin.Context
}

func (g GinContext) JSON(statusCode int, v interface{}) error {
	g.Context.JSON(statusCode, v)
	return nil
}

func (g GinContext) GetHeader(key string) string {
	return g.Context.Request.Header[key][0]
}

func (g GinContext) SetCookie(cookie *http.Cookie) {
	g.SetCookie(cookie)
}
