package constant

import (
	"github.com/agusanaluam/responsip/constant/lang"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Message struct {
	Lang string
	Msg  []*i18n.Message
}

var ListMessage []Message

func InitMessage() {

	idMessage := &Message{
		Lang: "id",
		Msg:  lang.IdMsg,
	}
	enMessage := &Message{
		Lang: "en",
		Msg:  lang.EnMsg,
	}
	cnMessage := &Message{
		Lang: "cn",
		Msg:  lang.CnMsg,
	}

	ListMessage = append(ListMessage, *idMessage)
	ListMessage = append(ListMessage, *enMessage)
	ListMessage = append(ListMessage, *cnMessage)

}
