package lang

const (
	SuccessInsert   = "success_create"
	SuccessUpdate   = "success_update"
	SuccessDelete   = "success_delete"
	SuccessRead     = "success_read"
	SuccessDownload = "success_download"
	SuccessSend     = "success_send"
	SuccessTransfer = "success_transfer"
)

const (
	ErrNoResult              = "err_no_result"
	ErrInvalidDate           = "err_invalid_date"
	ErrInvalidDateRange      = "err_date_range"
	ErrWalletSuspended       = "err_wallet_suspended"
	ErrMinimumTransfer       = "err_minimum_transfer"
	ErrMaximumTransfer       = "err_maximum_transfer"
	ErrNominalNotMatch       = "err_nominal_not_match"
	ErrInvalidMutationID     = "err_mutation_id"
	ErrInvalidMutationStatus = "err_mutation_status"
	ErrId                    = "err_invalid_id"
	ErrDebitNoteType         = "err_debit_note_type"
	ErrCreditNoteType        = "err_credit_note_type"
	ErrWeight                = "err_weight"
	ErrCourier               = "err_courier"
	ErrExist                 = "err_exist"
	ErrContactExist          = "err_contact_exist"
	ErrConvertCoin           = "err_convert_coin"
	ErrSuspended             = "err_suspended"

	// Authorization
	ErrInvalidOrEmptyToken       = "err_invalid_token"
	ErrorHttpInvalidServiceToken = "err_invalid_service_token"
	ErrTokenIsExpired            = "err_token_expired"
	ErrInvalidSignature          = "err_invalid_signature"
	ErrAccountSuspended          = "err_account_suspended"

	ErrDuplicateElement = "err_duplicate_element"
)
