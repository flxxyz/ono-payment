package service

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"net/url"
	"payment/conf"
)

var (
	AlipayClient *alipay.Client
	alipayConfig struct {
		appPrivateKey string
		notifyURL     string
		publicKey     string
		appId         string
		sellerId      string
		sellerName    string
		isProduction  bool
	}
)

func init() {
	alipayConfig.notifyURL = conf.AlipayNotifyURL
	alipayConfig.appPrivateKey = conf.AppPrivateKey
	alipayConfig.publicKey = conf.AlipayPublicKey
	alipayConfig.appId = conf.AlipayAppid
	alipayConfig.sellerId = conf.AlipaySellerId
	alipayConfig.sellerName = conf.AlipaySellerName
	alipayConfig.isProduction = true

	if conf.DefaultAppConfig().Debug {
		alipayConfig.publicKey = conf.AlipaySandboxPublicKey
		alipayConfig.appId = conf.AlipaySandboxAppid
		alipayConfig.sellerId = conf.AlipaySandboxSellerId
		alipayConfig.sellerName = conf.AlipaySandboxSellerName
		alipayConfig.isProduction = false
	}

	AlipayClient, _ = alipay.New(alipayConfig.appId, alipayConfig.appPrivateKey, alipayConfig.isProduction)
	_ = AlipayClient.LoadAliPayPublicKey(alipayConfig.publicKey)
}

// AlipayParseNotify 支付宝解析回调消息
func AlipayParseNotify(req *http.Request) (*alipay.TradeNotification, error) {
	return AlipayClient.GetTradeNotification(req)
}

// AlipayNewTradeApp 支付宝生成app交易结构体
func AlipayNewTradeApp(orderId string, amount int64) alipay.TradeAppPay {
	trade := alipay.TradeAppPay{}
	trade.OutTradeNo = orderId
	trade.TotalAmount = fmt.Sprintf("%.2f", float64(amount)/100)
	trade.NotifyURL = alipayConfig.notifyURL
	trade.SellerId = alipayConfig.sellerId
	trade.ProductCode = "QUICK_MSECURITY_PAY"
	return trade
}

// AlipayTradeAppPay 支付宝app支付接口
// 更多参数请查看下方链接
// @link https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
func AlipayTradeAppPay(trade alipay.TradeAppPay, title, returnURL string) (string, error) {
	trade.ReturnURL = returnURL
	trade.Subject = title
	//trade.GoodsType = goodsType //商品主类型 :0-虚拟类商品,1-实物类商品
	//trade.TimeExpire = expire //绝对超时时间，格式为yyyy-MM-dd HH:mm。
	//trade.PassbackParams = ""
	return AlipayClient.TradeAppPay(trade)
}

// AlipayTradeWapPay 支付宝手机网站支付接口
// 更多参数请查看下方链接
// @link https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
func AlipayTradeWapPay(trade alipay.TradeWapPay, title, returnURL string) (*url.URL, error) {
	trade.ReturnURL = returnURL
	trade.Subject = title
	//trade.GoodsType = goodsType //商品主类型 :0-虚拟类商品,1-实物类商品
	//trade.TimeExpire = expire //绝对超时时间，格式为yyyy-MM-dd HH:mm。
	return AlipayClient.TradeWapPay(trade)
}

// AlipayNewTradeWap 支付宝生成交易结构体
func AlipayNewTradeWap(orderId string, amount int64) alipay.TradeWapPay {
	trade := alipay.TradeWapPay{}
	trade.OutTradeNo = orderId
	trade.TotalAmount = fmt.Sprintf("%.2f", float64(amount)/100)
	trade.NotifyURL = alipayConfig.notifyURL
	trade.SellerId = alipayConfig.sellerId
	trade.ProductCode = "QUICK_WAP_WAY"
	return trade
}

// AlipayNewTradeQuery 支付宝生成交易查询结构体
func AlipayNewTradeQuery(orderId, tradeNo string) alipay.TradeQuery {
	trade := alipay.TradeQuery{}
	trade.OutTradeNo = orderId
	if tradeNo != "" {
		trade.TradeNo = tradeNo
	}
	return trade
}

// AlipayTradeQuery 支付宝查询交易订单
// tradeNo 支付宝的订单号，建议优先使用
// orderId 内部订单号
// @link https://opendocs.alipay.com/apis/api_1/alipay.trade.query
func AlipayTradeQuery(tradeNo, orderId string) (*alipay.TradeQueryRsp, error) {
	trade := AlipayNewTradeQuery(orderId, tradeNo)
	return AlipayClient.TradeQuery(trade)
}
