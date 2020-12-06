package trade

import (
	"encoding/json"
	"github.com/flxxyz/ono"
	"github.com/smartwalle/alipay/v3"
	"log"
	"payment/conf"
	"payment/dto"
	"payment/service"
)

var AlipayController = &alipayController{}

type alipayController struct {
	*Controller
	Trade
}

func (c *alipayController) App(ctx *ono.Context) {
	raw := ctx.RequestBody(&dto.BodyApp{}).(*dto.BodyApp)

	//data, _ := json.Marshal(raw)
	//_, _ = ctx.Writer.Write(data)

	// 内部订单号
	orderId := "xxxxxxxxxxxxxxx"

	// 支付金额
	var amount int64 = 1
	tradeApp := service.AlipayNewTradeApp(orderId, amount)
	result, err := service.AlipayTradeAppPay(
		tradeApp,
		"商品名称[app]",
		raw.URL,
	)
	if err != nil {
		log.Println("[支付宝] [App] 生成支付链接失败:", err.Error())
		return
	}

	tradeVal, _ := json.Marshal(dto.PayTradeInfo{
		OrderId:     orderId,
		Amount:      amount,
		Phone:       raw.Phone,
		CountryCode: "86",
		IP:          ctx.ClientIP(), // Todo: 支付时的用户IP
		URL:         raw.URL,
		PayType:     conf.PayTypeAliapyApp,
	})

	log.Println(tradeVal)

	// Todo: 创建支付请求

	ctx.JSON(0, dto.TradeRespApp{
		OrderId: orderId,
		PayInfo: result,
	})
}

func (c *alipayController) Wap(ctx *ono.Context) {
	raw := ctx.RequestBody(&dto.BodyWap{}).(*dto.BodyWap)

	//data, _ := json.Marshal(raw)
	//_, _ = t.Writer.Write(data)

	// 内部订单号
	orderId := "xxxxxxxxxxxxxxx"

	// 支付金额
	var amount int64 = 1
	tradeWap := service.AlipayNewTradeWap(orderId, amount)
	uri, err := service.AlipayTradeWapPay(
		tradeWap,
		"商品名称[wap]",
		raw.URL,
	)
	if err != nil {
		log.Println("[支付宝] [Wap] 生成支付链接失败:", err.Error())
		ctx.Fail("生成支付链接失败")
		return
	}

	tradeVal, _ := json.Marshal(dto.PayTradeInfo{
		OrderId:     orderId,
		Amount:      amount,
		Phone:       raw.Phone,
		CountryCode: "86",
		IP:          ctx.ClientIP(),
		URL:         raw.URL,
		PayType:     conf.PayTypeAliapyH5,
	})

	log.Println(tradeVal)

	// Todo: 创建支付请求

	ctx.Response(0, dto.TradeRespWap{
		OrderId: orderId,
		MWebURL: uri.String(),
	})
}

func (c *alipayController) Notify(ctx *ono.Context) {
	notice, _ := service.AlipayParseNotify(ctx.Request)
	if notice != nil {
		switch notice.TradeStatus {
		case alipay.TradeStatusSuccess:
			// Todo: 处理交易订单
			break
		case alipay.TradeStatusFinished:
			log.Println("[支付宝] [Notify] 交易结束", notice.OutTradeNo)
			break
		case alipay.TradeStatusWaitBuyerPay:
			log.Println("[支付宝] [Notify] 交易创建，等待买家付款", notice.OutTradeNo)
			break
		case alipay.TradeStatusClosed:
		default:
			log.Println("[支付宝] [Notify] 未付款交易超时关闭，或支付完成后全额退款", notice.OutTradeNo)
			break
		}
	}

	alipay.AckNotification(ctx.Writer)
}

func (c *alipayController) Query(ctx *ono.Context) {
	tradeNo := ctx.Query("tradeNo")
	orderId := ctx.Query("orderId")
	//if orderId == "" {
	//	log.Println("[支付宝] [Query] 内部订单号不允许为空！")
	//	c.Fail(ctx, "内部订单号不允许为空!")
	//	return
	//}

	resp, err := service.AlipayTradeQuery(tradeNo, orderId)
	if err != nil {
		log.Println("[支付宝] [Query] 查询订单出错！", err.Error())
		ctx.Fail("查询订单出错!")
		return
	}

	if !resp.IsSuccess() {
		log.Println("[支付宝] [Query] 查询订单接口调用失败！", resp)
		ctx.Fail("查询订单出错!")
	} else {
		switch resp.Content.TradeStatus {
		case alipay.TradeStatusSuccess:
			log.Println("[支付宝] [Query] 查询订单成功")
			ctx.Success("查询订单成功")
			break
		case alipay.TradeStatusFinished:
			log.Println("[支付宝] [Query] 交易结束，不可退款")
			ctx.Fail("交易结束，不可退款")
			break
		case alipay.TradeStatusWaitBuyerPay:
			log.Println("[支付宝] [Query] 交易创建，等待买家付款")
			ctx.Fail("交易创建，等待买家付款")
			break
		case alipay.TradeStatusClosed:
		default:
			log.Println("[支付宝] [Query] 未付款交易超时关闭，或支付完成后全额退款")
			ctx.Fail("查询订单成功!")
			break
		}
	}
}
