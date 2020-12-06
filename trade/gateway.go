package trade

import (
	"github.com/flxxyz/ono"
	"log"
	"net/http"
	"payment/utils"
	"strings"
)

var Getaway = &getawayController{}

type getawayController struct {
	*Controller
}

type actionGroup map[string]ono.HandleFunc
type httpMethodGroup map[string]actionGroup
type controllerGroup map[string]httpMethodGroup

var controllers = make(controllerGroup, 0)

func init() {
	controllers["alipay"] = httpMethodGroup{
		http.MethodPost: actionGroup{
			"App":    AlipayController.App,
			"Wap":    AlipayController.Wap,
			"Notify": AlipayController.Notify,
		},
		http.MethodGet: actionGroup{
			"Query": AlipayController.Query,
		},
	}
}

func (getaway *getawayController) Index(ctx *ono.Context) {
	typeParam := strings.ToLower(ctx.Query("type"))
	methodParam := utils.UpperFirstChar(ctx.Query("method"))

	log.Println(
		"http method:", ctx.Request.Method, ",",
		typeParam, ",",
		methodParam,
	)

	if controller, ok := controllers[typeParam]; ok {
		if httpMethods, ok := controller[ctx.Request.Method]; ok {
			if fn, ok := httpMethods[methodParam]; ok {
				fn(ctx)
				return
			}
		}
	}

	ctx.Fail("the method does not exist")
}
