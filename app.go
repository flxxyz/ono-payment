package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"payment/conf"
	myContext "payment/context"
	"payment/middleware"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

type App struct {
	*gin.Engine
	Conf *conf.AppConfig
}

func (app *App) Run() (err error) {
	srv := &http.Server{
		Addr:         app.Conf.Addr,
		Handler:      app.Engine,
		ReadTimeout:  time.Duration(app.Conf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(app.Conf.WriteTimeout) * time.Millisecond,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errMsg := fmt.Sprintf("启动server失败：%+v, %+v ", srv, err)
			panic(errMsg)
		}
	}()

	app.waitGraceExit(srv)
	return
}

func (app *App) init() {
	if app.Conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app.Engine = gin.New()
}

func (app *App) registerMiddleware() {
	// Todo: 一些自带的中间件
	app.Use(middleware.Logger())
}

func (app *App) WithContext(route interface{}, method string) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		// 请求超时控制
		withTimeout := time.Duration(app.Conf.WithTimeout) * time.Millisecond
		timeoutCtx, cancelFunc := context.WithTimeout(ginCtx, withTimeout)
		defer cancelFunc()

		// 反射设置上下文
		e := reflect.ValueOf(route).Elem()
		e.FieldByName("Context").Set(reflect.ValueOf(&myContext.Context{
			TimeoutCtx: timeoutCtx,
			Context:    ginCtx,
		}))

		m := reflect.ValueOf(route).MethodByName(method)
		m.Call(make([]reflect.Value, 0))
	}
}

func (app *App) waitGraceExit(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Fprintf(os.Stdout, "收到信号: %s, 服务正在退出... \n", s.String())
			srv.Close()
			return
		case syscall.SIGHUP:
		default:
		}
	}
}

func (app *App) Addr() string {
	return app.Conf.Addr
}

func init() {
	// cpu效率最大化
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func New() (app *App, err error) {
	app = &App{
		Conf: conf.DefaultAppConfig(),
	}
	// 初始化框架
	app.init()
	// 注册初始的中间件
	app.registerMiddleware()
	return
}
