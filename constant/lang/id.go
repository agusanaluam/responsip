package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var IdMsg = []*i18n.Message{
	{
		ID:    "hello",
		Other: "Halo, {{.Name}}!",
	},
	{
		ID:    "success_message",
		Other: "sukses menyimpan {{.module}}",
	},
}
