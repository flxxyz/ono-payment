package trade

import (
	"github.com/flxxyz/ono"
	"net/http"
)

type Trade interface {
	App() ono.HandleFunc
	Wap() ono.HandleFunc
	JSAPI() ono.HandleFunc
	Native() ono.HandleFunc
	Notify() ono.HandleFunc
	Query() ono.HandleFunc
}

type Controller struct{}

func (c *Controller) Ohhhhhh(ctx *ono.Context) {
	ctx.String(http.StatusOK, "ohhhhhhhhhhhhhhhhhhhhhhh")
}
