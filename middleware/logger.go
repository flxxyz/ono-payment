package middleware

import (
	"github.com/gin-gonic/gin"
	"payment/conf"
)

func Logger() gin.HandlerFunc {
	if conf.DefaultAppConfig().Debug {
		return gin.Logger()
	} else {
		return func(context *gin.Context) {
			context.Next()
		}
	}
}
