package conf

const (
	// 支付类型
	PayTypeWechatH5 = iota + 1
	PayTypeWechatJSAPI
	PayTypeWechatMini
	PayTypeWechatApp
	PayTypeAliapyH5
	PayTypeAliapyApp
	PayTypeAliapyJSAPI
	PayTypeAliapyMini
)

const (
	// 应用密钥
	AppPrivateKey = ""
	AppPublicKey  = ""

	AlipayNotifyURL      = "https://example.com/alipay/notify"
	AlipayEncryptContent = ""
	AlipayAppPrivateKey  = ""
	AlipayPublicKey      = ""
	AlipayAppid          = ""
	AlipaySellerId       = ""
	AlipaySellerName     = ""
	// 支付宝的沙盒账号
	AlipaySandboxPublicKey  = ""
	AlipaySandboxAppid      = ""
	AlipaySandboxSellerId   = ""
	AlipaySandboxSellerName = ""
)
