package responsip

import (
	"encoding/json"
	"net/http"
)

type MuxContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (m MuxContext) JSON(statusCode int, v interface{}) error {
	m.Writer.Header().Set("Content-Type", "application/json")
	m.Writer.WriteHeader(statusCode)
	return json.NewEncoder(m.Writer).Encode(v)
}

func (m MuxContext) GetHeader(key string) string {
	return m.Request.Header.Get(key)
}

func (m MuxContext) SetCookie(cookie *http.Cookie) {
	m.SetCookie(cookie)
}
