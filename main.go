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
	app.Any("/getaway.do", app.WithContext(trade.Getaway.Index))

	go func() {
		log.Println("[Server]", app.Addr())
		log.Fatal(app.Run())
	}()
	select {}

}
