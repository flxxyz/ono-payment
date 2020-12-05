package context

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type HandleFunc func()

type Context struct {
	TimeoutCtx context.Context
	*gin.Context
}

func (c *Context) RequestBody(data interface{}) interface{} {
	body, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, data)
	return data
}
