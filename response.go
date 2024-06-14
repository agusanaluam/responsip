package responsip

import (
	"errors"
	"net/http"
	"time"

	"github.com/agusanaluam/responsip/constant"
	"github.com/agusanaluam/responsip/helper"
	"github.com/agusanaluam/responsip/schema"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "go.uber.org/zap"
	"golang.org/x/text/language"
)

type Responsip struct {
	Lang   language.Tag
	Module string
	bundle *i18n.Bundle
	Ctx    Context
}

type Context interface {
	JSON(statusCode int, v interface{}) error
	GetHeader(key string) string
	SetCookie(cookie *http.Cookie)
}

type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

// ================RESPONSE 2XX======================
// Status Code : 200
func (r *Responsip) OK(message string, data interface{}) error {
	log.S().Info("success response")

	return r.Ctx.JSON(http.StatusOK, schema.Base{
		Status:     constant.ResponseSuccess,
		StatusCode: http.StatusOK,
		Message:    getLocalizedString(r.Ctx, r.bundle, message, r.Module),
		Timestamp:  time.Now(),
		Data:       data,
	})
}

// Status Code : 201
func (r *Responsip) Created(message string, data interface{}) error {
	log.S().Info("created response")

	return r.Ctx.JSON(http.StatusCreated, schema.Base{
		Status:     constant.ResponseCreated,
		StatusCode: http.StatusCreated,
		Message:    getLocalizedString(r.Ctx, r.bundle, message, r.Module),
		Timestamp:  time.Now(),
		Data:       data,
	})
}

// Status Code : 202
func (r *Responsip) Accepted(message string, data interface{}) error {
	log.S().Info("accepted response")

	return r.Ctx.JSON(http.StatusAccepted, schema.Base{
		Status:     constant.ResponseSuccess,
		StatusCode: http.StatusAccepted,
		Message:    getLocalizedString(r.Ctx, r.bundle, message, r.Module),
		Timestamp:  time.Now(),
		Data:       data,
	})
}

// Status Code : 204
func (r *Responsip) NoContent(message string) error {
	log.S().Info("no content response")

	return r.Ctx.JSON(http.StatusNoContent, schema.Base{
		Status:     constant.ResponseNoContent,
		StatusCode: http.StatusNoContent,
		Message:    getLocalizedString(r.Ctx, r.bundle, message, r.Module),
		Timestamp:  time.Now(),
	})
}

// ErrorResponse returns
func (r *Responsip) ErrorResponse(err error, data interface{}) error {
	statusCode, errCode := helper.ErrorType(err)
	switch statusCode {
	case http.StatusConflict:
		return r.Conflict(errCode, data)
	case http.StatusBadRequest:
		return r.BadRequest(errCode, data)
	case http.StatusNotFound:
		return r.NotFound(errCode, data)
	case http.StatusUnauthorized:
		return r.Unauthorized(errCode, data)
	case http.StatusForbidden:
		return r.Forbidden(errCode, data)
	}
	return r.InternalServerError(errCode, data)
}

// ====================RESPONSE 4XX=======================
// Status Code : 400
func (r *Responsip) BadRequest(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)

	responseData := schema.Base{
		Status:     constant.ResponseBadRequest,
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Timestamp:  time.Now(),
		Data:       data,
	}

	log.S().Errorf("bad request error : %s ", msg)

	return r.Ctx.JSON(http.StatusBadRequest, responseData)
}

// Status Code : 401
func (r *Responsip) Unauthorized(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)
	responseData := schema.Base{
		Status:     constant.ResponseUnauthorized,
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("unauthorized error : %s ", msg)

	return r.Ctx.JSON(http.StatusUnauthorized, responseData)
}

// Status Code : 403
func (r *Responsip) Forbidden(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)
	responseData := schema.Base{
		Status:     constant.ResponseForbidden,
		StatusCode: http.StatusForbidden,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("forbidden error : %s ", msg)

	return r.Ctx.JSON(http.StatusForbidden, responseData)
}

// Status Code : 404
func (r *Responsip) NotFound(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)

	responseData := schema.Base{
		Status:     constant.ResponseNotFound,
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("error not found : %s ", msg)

	return r.Ctx.JSON(http.StatusNotFound, responseData)
}

// Status Code : 409
func (r *Responsip) Conflict(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)

	responseData := schema.Base{
		Status:     constant.ResponseConflict,
		StatusCode: http.StatusConflict,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("conflict data error : %s ", msg)

	return r.Ctx.JSON(http.StatusConflict, responseData)
}

// Status Code 422
func (r *Responsip) ErrorParsing(errCode string, data interface{}) error {

	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)

	responseData := schema.Base{
		Status:     constant.ResponseUnprocessableEntity,
		StatusCode: http.StatusUnprocessableEntity,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("parsing data error : %s ", msg)

	return r.Ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ===========================RESPONSE 5XX===========================
// Status Code 500
func (r *Responsip) InternalServerError(errCode string, data interface{}) error {
	msg := getLocalizedString(r.Ctx, r.bundle, errCode, r.Module)

	responseData := schema.Base{
		Status:     constant.ResponseInternalServerError,
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("internal server error : %s ", msg)

	return r.Ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorValidate returns
func (r *Responsip) ErrorValidate(errCode string, data interface{}) error {
	message := helper.SwitchErrorValidation(errors.New(errCode))
	responseData := schema.Base{
		Status:     constant.ResponseBadRequest,
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("validate data error : %s ", errCode)

	return r.Ctx.JSON(http.StatusBadRequest, responseData)
}

// SetCookie returns
func SetCookie(ctx Context, data map[string]string) {
	for key, value := range data {
		cookie := &http.Cookie{}
		cookie.Name = key
		cookie.Value = value
		cookie.HttpOnly = true
		cookie.Path = "/"
		ctx.SetCookie(cookie)
	}
}
