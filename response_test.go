package responsip

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestEchoResponseSuccess(t *testing.T) {
	// Set up Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "en")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	resp := Responsip{
		Lang:   language.Indonesian,
		Module: "image-api",
	}
	resp.InitLocalizer()
	ctx := EchoContext{Context: c}

	// Call JSONSuccess function
	err := resp.SuccessResponse(ctx, "success_message", "success_data")
	fmt.Println(rec.Body.String())

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, fmt.Sprintf(`{"status":"success", "statusCode":200,"message":"%s", "timestamp":"%s","data":"success_data"}`, resp.GetLocalizedString(ctx, "success_message"), time.Now().UTC().Format("2006-01-02T15:04:05.00000000Z")), rec.Body.String())
	}
}

func TestGinResponseSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("Accept-Language", "en")

	resp := Responsip{
		Lang:   language.Indonesian,
		Module: "image-api",
	}
	resp.InitLocalizer()
	gctx := GinContext{Context: ctx}

	// Call JSONSuccess function
	err := resp.SuccessResponse(gctx, "success_message", "success_data")
	fmt.Println(w.Body.String())

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, fmt.Sprintf(`{"status":"success", "statusCode":200,"message":"%s", "timestamp":"%s","data":"success_data"}`, resp.GetLocalizedString(gctx, "success_message"), time.Now().UTC().Format("2006-01-02T15:04:05.00000000Z")), w.Body.String())
	}
}

func TestFiberResponseSuccess(t *testing.T) {
	app := fiber.New()

	// Define a test route that calls JSONSuccess
	app.Get("/test", func(c *fiber.Ctx) error {
		resp := Responsip{
			Lang:   language.Indonesian,
			Module: "image-api",
		}
		resp.InitLocalizer()
		fctx := FiberContext{Context: c}

		// Call JSONSuccess function
		return resp.SuccessResponse(fctx, "success_message", "success_data")
	})

	// Create a test request
	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)
	b, _ := io.ReadAll(resp.Body)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, fmt.Sprintf(`{"status":"success", "statusCode":200,"message":"%s", "timestamp":"%s","data":"success_data"}`, "success_message", time.Now().UTC().Format("2006-01-02T15:04:05.00000000Z")), string(b))
	}
}
