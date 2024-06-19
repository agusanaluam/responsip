package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var IdMsg = []*i18n.Message{
	{
		ID:    SuccessRead,
		Other: "sukses mengambil data {{.module}}!",
	},
	{
		ID:    SuccessInsert,
		Other: "sukses menyimpan {{.module}}",
	},
	{
		ID:    SuccessUpdate,
		Other: "sukses mengubah {{.module}}",
	},
	{
		ID:    SuccessDelete,
		Other: "sukses menghapus {{.module}}",
	},
	{
		ID:    ErrNoResult,
		Other: "tidak dapat menemukan {{ .module }}",
	},
	{
		ID:    ErrInvalidDate,
		Other: "tanggal tidak valid",
	},
	{
		ID:    ErrInvalidDateRange,
		Other: "tanggal mulai harus lebih awal dari tanggal berakhir",
	},
	{
		ID:    ErrId,
		Other: "{{.module}} id di butuhkan",
	},
	{
		ID:    ErrSuspended,
		Other: "{{.module}} telah dibekukan",
	},
	{
		ID:    ErrMinimumTransfer,
		Other: "minimum transfer adalah Rp. 1",
	},
	{
		ID:    ErrMaximumTransfer,
		Other: "maximum transfer adalah Rp. 10.000.000",
	},
	{
		ID:    ErrNominalNotMatch,
		Other: "nominal tidak cocok",
	},
	{
		ID:    ErrInvalidMutationID,
		Other: "id mutasi dibutuhkan",
	},
	{
		ID:    ErrInvalidMutationStatus,
		Other: "id status mutasi tidak valid",
	},
	{
		ID:    ErrDebitNoteType,
		Other: "debit note type tidak valid",
	},
	{
		ID:    ErrCreditNoteType,
		Other: "credit note type tidak valid",
	},
	{
		ID:    ErrWeight,
		Other: "berat tidak valid",
	},
	{
		ID:    ErrCourier,
		Other: "nama kurir tidak valid",
	},
	{
		ID:    ErrExist,
		Other: "{{ .module }} sudah terdaftar",
	},
	{
		ID:    ErrContactExist,
		Other: "kontak sudah terdaftar",
	},
	{
		ID:    ErrConvertCoin,
		Other: "tidak bisa konversi ke koin",
	},
	{
		ID:    ErrDuplicateElement,
		Other: "duplikat data {{.module}}",
	},
	{
		ID:    ErrInvalidOrEmptyToken,
		Other: "unauthorized",
	},
	{
		ID:    ErrorHttpInvalidServiceToken,
		Other: "invalid service token",
	},
	{
		ID:    ErrTokenIsExpired,
		Other: "token is expired",
	},
	{
		ID:    ErrInvalidSignature,
		Other: "invalid signature key",
	},
	{
		ID:    ErrAccountSuspended,
		Other: "your account is suspended",
	},
}
