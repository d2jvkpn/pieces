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
		addr    string
		release bool
		engine  *gin.Engine
		router  *gin.RouterGroup
	)

	flag.StringVar(&addr, "addr", ":8080", "http server address")
	flag.BoolVar(&release, "release", false, "run in release mode")
	flag.Parse()

	if release {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	} else {
		engine = gin.Default()
	}
	// engine.SetTrustedProxies([]string{"192.168.1.2"})
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
	now := time.Now()
	bts, _ := json.Marshal(ctx.Request.Header)

	record := fmt.Sprintf(
		"<=> %s ClientIP: %q, RemoteAddr: %q, Method: %q, Path: %q, Query: %q, Headers: %s",
		now.Format(time.RFC3339), ctx.ClientIP(), ctx.Request.RemoteAddr,
		ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.URL.RawQuery, bts,
	)

	ctx.Next()
	fmt.Printf(
		"%s, Status: %d, Elapsed: %v\n",
		record, ctx.Writer.Status(), time.Since(now),
	)
}
