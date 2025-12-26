package locales

type LocalesMsgId string

const (
	OrderSkuNotExist   LocalesMsgId = "order.SkuNotExist"
	OrderServerError   LocalesMsgId = "order.ServerError"
	OrderCreateSuccess LocalesMsgId = "order.CreateSuccess"
)

var (
	defaultLocaleMap = map[LocalesMsgId]string{}
)

func RegisterDefaultLocale(msgId LocalesMsgId, defaultMsg string) {
	defaultLocaleMap[msgId] = defaultMsg
}

func init() {
	RegisterDefaultLocale(OrderSkuNotExist, "order sku not exist")
	RegisterDefaultLocale(OrderServerError, "order server error")
	RegisterDefaultLocale(OrderCreateSuccess, "order create success")
}

func GetDefaultLocaleMsg(msgId LocalesMsgId) string {
	msg, ok := defaultLocaleMap[msgId]
	if !ok {
		return "internal error"
	}
	return msg
}
