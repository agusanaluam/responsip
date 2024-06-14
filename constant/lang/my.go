package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var CnMsg = []*i18n.Message{
	{
		ID:    SuccessRead,
		Other: "success get data {{.module}}!",
	},
	{
		ID:    SuccessInsert,
		Other: "success insert {{.module}}",
	},
	{
		ID:    SuccessUpdate,
		Other: "success update {{.module}}",
	},
	{
		ID:    SuccessDelete,
		Other: "success delete {{.module}}",
	},
	{
		ID:    ErrNoResult,
		Other: "resource {{ .module }} not found",
	},
	{
		ID:    ErrInvalidDate,
		Other: "invalid date range",
	},
	{
		ID:    ErrInvalidDateRange,
		Other: "start date must be less than end date",
	},
	{
		ID:    ErrId,
		Other: "{{.module}} id is required",
	},
	{
		ID:    ErrSuspended,
		Other: "{{.module}} is suspended",
	},
	{
		ID:    ErrMinimumTransfer,
		Other: "minimum transfer is Rp. 1",
	},
	{
		ID:    ErrMaximumTransfer,
		Other: "maximum transfer is Rp. 10.000.000",
	},
	{
		ID:    ErrNominalNotMatch,
		Other: "nominal not match",
	},
	{
		ID:    ErrInvalidMutationID,
		Other: "mutation unique id is required",
	},
	{
		ID:    ErrInvalidMutationStatus,
		Other: "invalid mutation status id",
	},
	{
		ID:    ErrDebitNoteType,
		Other: "invalid debit note type",
	},
	{
		ID:    ErrCreditNoteType,
		Other: "invalid credit note type",
	},
	{
		ID:    ErrWeight,
		Other: "invalid weight value",
	},
	{
		ID:    ErrCourier,
		Other: "invalid courier name",
	},
	{
		ID:    ErrExist,
		Other: "{{ .module }} already exists",
	},
	{
		ID:    ErrContactExist,
		Other: "contact already exist",
	},
	{
		ID:    ErrConvertCoin,
		Other: "can't convert cash to coin",
	},
	{
		ID:    ErrDuplicateElement,
		Other: "duplicate key value {{.module}}",
	},
}
