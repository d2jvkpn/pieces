package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "os"
	// "strings"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		addr   string
		engine *gin.Engine
		router *gin.RouterGroup
	)

	flag.StringVar(&addr, "addr", ":8080", "http server address")
	flag.Parse()

	engine = gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	// engine = gin.New()
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	router = &engine.RouterGroup

	engine.NoRoute(inspect, func(ctx *gin.Context) {
		// ctx.AbortWithStatus(http.StatusNotFound)
		ctx.String(http.StatusNotFound, "Sorry, not found!\n")
	})

	irouters := router.Use(inspect)
	irouters.GET("/", hello)

	fmt.Printf(">>> Http service listen on %s\n", addr)
	engine.Run(addr)
}

func hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, world!\n")
	return
}

func inspect(ctx *gin.Context) {
	fmt.Printf(
		"~~~ %s ClientIP: %q, RemoteAddr: %q, Method: %q\n    Path: %q, Query: %q\n",
		time.Now().Format(time.RFC3339), ctx.ClientIP(), ctx.Request.RemoteAddr,
		ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.URL.RawQuery,
	)
	bts, _ := json.Marshal(ctx.Request.Header)
	fmt.Printf("    Headers: %s\n", bts)

	ctx.Next()
	fmt.Printf("    Status: %d\n", ctx.Writer.Status())
}
