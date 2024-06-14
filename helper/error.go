package helper

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/agusanaluam/responsip/constant"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func calculateMin(fieldName, num string) string {
	fieldsName := map[string]int{
		"phone number": 2,
	}

	i, err := strconv.Atoi(num)
	if err != nil {
		return num
	}

	for key := range fieldsName {
		if key == fieldName {
			return strconv.Itoa(i - fieldsName[key])
		}
	}
	return num
}

func oneofErrorSplit(choices string) (result string) {
	s := strings.Split(choices, " ")
	for idx := range s {
		message := fmt.Sprintf("%s or ", s[idx])
		if (idx + 1) == len(s) {
			message = s[idx]
		}
		result += message
	}
	return
}

func ErrorType(err error) (int, string) {
	switch {
	case isPqError(err):
		return pqError(err)
	}
	return commonError(err)
}

func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

func SwitchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok { //nolint
		for _, err := range castedObject {
			field := setLowerAndAddSpace(err.Field())

			// Change Field Name
			switch field {
			case "msisdn":
				field = "phone number"
			}

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					field)
			case "number":
				message = fmt.Sprintf("%s must be numbers only",
					field)
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					field, err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					field, err.Param())
			case "min":
				minimum := calculateMin(field, err.Param())
				message = fmt.Sprintf("%s at least %s characters long",
					field, minimum)
			case "max":
				message = fmt.Sprintf("the length of %s must be %s characters or fewer",
					field, err.Param())
			case "startswith":
				message = fmt.Sprintf("%s must starts with %s",
					field, err.Param())
			case "len":
				message = fmt.Sprintf("%s length must %s characters",
					field, err.Param())
			case "oneof":
				choices := oneofErrorSplit(err.Param())
				message = fmt.Sprintf("%s must specify one of %s",
					field, choices)
			default:
				message = err.Error()
			}

			// break
		}
	}
	return
}

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func pqError(err error) (int, string) {
	re := regexp.MustCompile(`\\((.*?)\\)`) // nolint
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)
		// Change Field Name
		switch match[1] {
		case "msisdn":
			match[1] = "phone number"
		}

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Sprintf("%s already exists", match[1])
		}
	}

	return http.StatusInternalServerError, err.Error()
}

// CommonError is
func commonError(err error) (int, string) {
	if strings.HasPrefix(err.Error(), "wait") || err.Error() == "invalid pin" {
		return http.StatusForbidden, err.Error()
	}

	switch err.Error() {
	case constant.ErrInvalidOrEmptyToken.Error():
		return http.StatusUnauthorized, err.Error()
	case constant.ErrorHttpInvalidServiceToken.Error():
		return http.StatusUnauthorized, err.Error()
	case constant.ErrTokenIsExpired.Error():
		return http.StatusUnauthorized, err.Error()
	case constant.ErrInvalidSignature.Error():
		return http.StatusUnauthorized, err.Error()
	case "token signature is invalid":
		return http.StatusUnauthorized, err.Error()
	case constant.ErrSuspended.Error():
		return http.StatusForbidden, err.Error()
	case constant.ErrNoResult.Error():
		return http.StatusNotFound, err.Error()
	case constant.ErrInvalidDate.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrInvalidDateRange.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrMinimumTransfer.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrMaximumTransfer.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrNominalNotMatch.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrInvalidMutationID.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrInvalidMutationStatus.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrUserId.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrDebitNoteType.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrCreditNoteType.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrWeight.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrCourier.Error():
		return http.StatusBadRequest, err.Error()
	case constant.ErrWalletExist.Error():
		return http.StatusConflict, err.Error()
	case constant.ErrContactExist.Error():
		return http.StatusConflict, err.Error()
	case constant.ErrConvertCoin.Error():
		return http.StatusBadRequest, err.Error()
	case "pq: null value in column \"active_balance\" violates not-null constraint":
		return http.StatusBadRequest, "insufficient balance"
	case "pq: null value in column \"active_coin\" violates not-null constraint":
		return http.StatusBadRequest, "can't convert cash to coin"
	// case "pq: duplicate key value violates unique constraint \"favorite_contacts_un\"":
	// 	return http.StatusConflict, constant.ErrContactExist
	case "sql: no rows in result set":
		return http.StatusNotFound, constant.ResponseNotFound
	case "cannot transfer to yourself":
		return http.StatusBadRequest, err.Error()
	case "pin incorrect":
		return http.StatusForbidden, err.Error()
	case "order already paid":
		return http.StatusBadRequest, err.Error()
	case "order expired":
		return http.StatusBadRequest, err.Error()
	}

	if strings.HasPrefix(err.Error(), "pq:") {
		return http.StatusBadRequest, "bad request"
	}

	return http.StatusInternalServerError, err.Error()
}
