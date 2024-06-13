package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var CnMsg = []*i18n.Message{
	{
		ID:    "hello",
		Other: "Hello, {{.Name}}!",
	},
	{
		ID:    "success_message",
		Other: "success insert {{.module}}",
	},
}
