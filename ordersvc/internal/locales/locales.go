package locales

import (
	"context"
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle *i18n.Bundle

	localizerCache sync.Map
)

func init() {
	bundle = i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// ./internal/locales/
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.zh.toml")
}

func GetLocaleMsg(ctx context.Context, lang string, messageID LocalesMsgId) string {
	return GetLocaleMessageWith(ctx, lang, messageID, nil)
}

func GetLocaleMessageWith(ctx context.Context, lang string, messageID LocalesMsgId, templateData map[string]interface{}) string {
	localizer := GetLocalizer(lang)
	if localizer == nil {
		return ""
	}
	// return localizer.MustLocalize(&i18n.LocalizeConfig{
	// 	MessageID:    string(messageID),
	// 	TemplateData: templateData,
	// })
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    string(messageID),
		TemplateData: templateData,
	})
	if err != nil {
		defaultMsg := GetDefaultLocaleMsg(messageID)
		return defaultMsg
	}
	return msg
}

func GetLocalize(ctx context.Context, lang string, messageId LocalesMsgId) (string, error) {
	localizer := GetLocalizer(lang)
	if localizer == nil {
		return "", fmt.Errorf("lang not exist")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: string(messageId),
	})
	g.Log().Warningf(ctx, "msg: %s", msg)
	if err != nil {
		defaultMsg := GetDefaultLocaleMsg(messageId)
		return defaultMsg, nil
	}
	return msg, nil
}

func GetLocalizer(lang string) *i18n.Localizer {
	if v, ok := localizerCache.Load(lang); ok {
		return v.(*i18n.Localizer)
	}
	localizer := i18n.NewLocalizer(bundle, lang)
	localizerCache.Store(lang, localizer)
	return localizer
}
