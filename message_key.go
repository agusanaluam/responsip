package responsip

import "github.com/agusanaluam/responsip/constant/lang"

const (
	SuccessInsert   = lang.SuccessInsert
	SuccessUpdate   = lang.SuccessUpdate
	SuccessDelete   = lang.SuccessDelete
	SuccessRead     = lang.SuccessRead
	SuccessDownload = lang.SuccessDownload
	SuccessSend     = lang.SuccessSend
	SuccessTransfer = lang.SuccessTransfer

	ErrNoResult              = lang.ErrNoResult
	ErrInvalidDate           = lang.ErrInvalidDate
	ErrInvalidDateRange      = lang.ErrInvalidDateRange
	ErrWalletSuspended       = lang.ErrWalletSuspended
	ErrMinimumTransfer       = lang.ErrMinimumTransfer
	ErrMaximumTransfer       = lang.ErrMaximumTransfer
	ErrNominalNotMatch       = lang.ErrNominalNotMatch
	ErrInvalidMutationID     = lang.ErrInvalidMutationID
	ErrInvalidMutationStatus = lang.ErrInvalidMutationStatus
	ErrId                    = lang.ErrId
	ErrDebitNoteType         = lang.ErrDebitNoteType
	ErrCreditNoteType        = lang.ErrCreditNoteType
	ErrWeight                = lang.ErrWeight
	ErrCourier               = lang.ErrCourier
	ErrExist                 = lang.ErrExist
	ErrContactExist          = lang.ErrContactExist
	ErrConvertCoin           = lang.ErrConvertCoin

	// Authorization
	ErrInvalidOrEmptyToken     = lang.ErrInvalidOrEmptyToken
	ErrHttpInvalidServiceToken = lang.ErrorHttpInvalidServiceToken
	ErrTokenIsExpired          = lang.ErrTokenIsExpired
	ErrInvalidSignature        = lang.ErrInvalidSignature
	ErrSuspended               = lang.ErrSuspended

	ErrDuplicateElement = lang.ErrDuplicateElement
)
