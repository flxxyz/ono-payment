package dto

type TradeRespWap struct {
	OrderId string `json:"order_id"`
	MWebURL string `json:"mweb_url"`
}

type TradeRespApp struct {
	OrderId string `json:"order_id"`
	PayInfo string `json:"pay_info"`
}

type PayTradeInfo struct {
	OrderId     string `json:"orderId"`
	TradeNo     string `json:"tradeNo"`
	Amount      int64  `json:"amount"`
	Account     string `json:"account"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	IP          string `json:"ip"`
	Nickname    string `json:"nickname"`
	Channel     string `json:"channel"`
	URL         string `json:"url"`
	PayType     int    `json:"payType"`
}
