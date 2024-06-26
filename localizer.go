package responsip

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var langDir string = "constant/lang"

func (r *Responsip) InitLocalizer() error {
	r.bundle = i18n.NewBundle(r.Lang)
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
	accept := ctx.GetHeader("Accept-Language")
	matcher := language.NewMatcher(r.bundle.LanguageTags())
	_, i, _ := matcher.Match(language.Make(accept))

	lang := r.bundle.LanguageTags()[i]
	localizer := i18n.NewLocalizer(r.bundle, lang.String())
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
		TemplateData: map[string]interface{}{
			"module": r.Module,
		},
	})
	if err != nil {
		return messageID
	}
	return msg
}
