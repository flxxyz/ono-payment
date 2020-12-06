package trade

import (
	"github.com/flxxyz/ono"
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
