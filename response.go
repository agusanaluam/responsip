package responsip

import (
	"net/http"
	"time"

	"github.com/agusanaluam/responsip/constant"
	"github.com/agusanaluam/responsip/schema"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Responsip struct {
	Lang   language.Tag
	Module string
	bundle *i18n.Bundle
}

type Context interface {
	JSON(statusCode int, v interface{}) error
	GetHeader(key string) string
}

type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

func (r *Responsip) SuccessResponse(ctx Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, schema.Base{
		Status:     constant.StatusSuccessText,
		StatusCode: http.StatusOK,
		Message:    r.GetLocalizedString(ctx, message),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	})
}

// CreatedResponse returns
func (r *Responsip) CreatedResponse(ctx Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusCreated, schema.Base{
		Status:     constant.StatusCreatedText,
		StatusCode: http.StatusCreated,
		Message:    r.GetLocalizedString(ctx, message),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	})
}

// NoContentResponse returns
func (r *Responsip) NoContentResponse(ctx Context, message string) error {
	return ctx.JSON(http.StatusNoContent, schema.Base{
		Status:     constant.StatusNoContent,
		StatusCode: http.StatusNoContent,
		Message:    r.GetLocalizedString(ctx, message),
		Timestamp:  time.Now().UTC(),
	})
}
