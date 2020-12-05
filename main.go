package main

import (
	"github.com/flxxyz/ono"
	"log"
	"net/http"
	"payment/trade"
)

func main() {
	app, _ := ono.New()

	app.GET("/", app.WithContext(func(ctx *ono.Context) {
		ctx.String(http.StatusOK, "hello payment")
	}))
	app.GET("/alipay/query", app.WithContext(trade.AlipayController.Query))
	app.POST("/alipay/notify", app.WithContext(trade.AlipayController.Notify))
	app.POST("/alipay/app/pay", app.WithContext(trade.AlipayController.App))
	app.POST("/alipay/h5/pay", app.WithContext(trade.AlipayController.Wap))

	go func() {
		log.Println("[Server]", app.Addr())
		log.Fatal(app.Run())
	}()
	select {}

}
