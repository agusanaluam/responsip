package responsip

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestAddCustomMessage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "id")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	IdMsg := []*i18n.Message{
		{
			ID:    "world",
			Other: "dunia, ini {{.Name}}!",
		},
		{
			ID:    "info_message",
			Other: "ini data {{.module}}",
		},
	}

	resp := Responsip{
		Lang:   language.Indonesian,
		Module: "localizer",
		Ctx:    EchoContext{Context: c},
	}
	resp.InitLocalizer()
	resp.bundle.AddMessages(language.Make("id"), IdMsg...)

	err := resp.OK("success_message", "success_data")
	fmt.Println(rec.Body.String())

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.JSONEq(t, fmt.Sprintf(`{"status":"success", "statusCode":200,"message":"%s", "timestamp":"%s","data":"success_data"}`, "success_message", time.Now().UTC().Format("2006-01-02T15:04:05.00000000Z")), rec.Body.String())
	}
}
