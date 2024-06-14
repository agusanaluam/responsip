package responsip

import (
	"encoding/json"

	"github.com/agusanaluam/responsip/constant"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// var langDir string = "constant/lang"

type Localizer struct{}

func (r *Responsip) InitLocalizer() error {
	r.bundle = i18n.NewBundle(r.Lang)
	r.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	constant.InitMessage()
	for _, msg := range constant.ListMessage {
		r.bundle.AddMessages(language.Make(msg.Lang), msg.Msg...)
	}

	// TEMPORARY COMMENTED - CANT READ PACKAGE FILES
	// files, err := os.ReadDir("./" + langDir)
	// if err != nil {
	// 	return err
	// }

	// for _, file := range files {
	// 	r.bundle.LoadMessageFile(fmt.Sprintf("%s/%s", langDir, file.Name()))
	// }
	return nil
}

func getLocalizedString(ctx Context, bundle *i18n.Bundle, messageID, module string) string {
	accept := ctx.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, accept)
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
		TemplateData: map[string]interface{}{
			"module": module,
		},
	})
	if err != nil {
		return messageID
	}
	return msg
}
