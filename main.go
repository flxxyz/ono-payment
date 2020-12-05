package main

import (
	"log"
	"payment/trade"
	"payment/trade/alipay"
)

func main() {
	app, _ := New()

	app.GET("/", app.WithContext(&trade.Controller{}, "Homepage"))
	app.GET("/alipay/query", app.WithContext(&alipay.Trade{}, "Query"))
	app.POST("/alipay/notify", app.WithContext(&alipay.Trade{}, "Notify"))
	app.POST("/alipay/app/pay", app.WithContext(&alipay.Trade{}, "App"))
	app.POST("/alipay/h5/pay", app.WithContext(&alipay.Trade{}, "Wap"))

	go func() {
		log.Println("[Server]", app.Addr())
		log.Fatal(app.Run())
	}()
	select {}

}
