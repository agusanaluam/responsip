package responsip

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var langDir string = "constant/lang"

func (r *Responsip) InitLocalizer() error {
	r.bundle = i18n.NewBundle(r.lang)
	r.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, err := os.ReadDir(langDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		r.bundle.LoadMessageFile(fmt.Sprintf("%s/%s", langDir, file))
	}
	return nil
}

func (r *Responsip) GetLocalizedString(ctx Context, messageID string) string {
	localizer := i18n.NewLocalizer(r.bundle, ctx.GetHeader("Accept-Language"))
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	if err != nil {
		return messageID
	}
	return msg
}
