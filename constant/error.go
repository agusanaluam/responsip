package constant

import "fmt"

var (
	ErrNoResult              = fmt.Errorf("resource not found")
	ErrInvalidDate           = fmt.Errorf("invalid date range")
	ErrInvalidDateRange      = fmt.Errorf("start date must be less than end date")
	ErrWalletSuspended       = fmt.Errorf("wallet is suspended")
	ErrMinimumTransfer       = fmt.Errorf("minimum transfer is Rp. 1")
	ErrMaximumTransfer       = fmt.Errorf("maximum transfer is Rp. 10.000.000")
	ErrNominalNotMatch       = fmt.Errorf("nominal not match")
	ErrInvalidMutationID     = fmt.Errorf("mutation unique id is required")
	ErrInvalidMutationStatus = fmt.Errorf("invalid mutation status id")
	ErrUserId                = fmt.Errorf("user id is required")
	ErrDebitNoteType         = fmt.Errorf("invalid debit note type")
	ErrCreditNoteType        = fmt.Errorf("invalid credit note type")
	ErrWeight                = fmt.Errorf("invalid weight value")
	ErrCourier               = fmt.Errorf("invalid courier name")
	ErrWalletExist           = fmt.Errorf("user already have wallet")
	ErrContactExist          = fmt.Errorf("contact already exist")
	ErrConvertCoin           = fmt.Errorf("can't convert cash to coin")

	// Authorization
	ErrInvalidOrEmptyToken       = fmt.Errorf("unauthorized")
	ErrorHttpInvalidServiceToken = fmt.Errorf("invalid service token")
	ErrTokenIsExpired            = fmt.Errorf("token is expired")
	ErrInvalidSignature          = fmt.Errorf("invalid signature key")
	ErrSuspended                 = fmt.Errorf("your account is suspended")
)

const (
	ErrDuplicateContact = `pq: duplicate key value violates unique constraint "favorite_contacts_un"`
	ErrDuplicateWallet  = `pq: duplicate key value violates unique constraint "balances_pkey"`
)
