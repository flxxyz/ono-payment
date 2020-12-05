package trade

import (
	"net/http"
	"payment/context"
)

type Trade interface {
	App() context.HandleFunc
	Wap() context.HandleFunc
	JSAPI() context.HandleFunc
	Native() context.HandleFunc
	Notify() context.HandleFunc
	Query() context.HandleFunc
}

type Controller struct {
	*context.Controller
	Trade
}

func (c *Controller) Homepage() {
	c.String(http.StatusOK, "hello payment")
}
