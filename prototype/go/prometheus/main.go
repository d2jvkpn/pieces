package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"x/api/web"
	"x/pkg/prom"

	"github.com/gin-gonic/gin"
)

var (
	server *http.Server
)

func init() {
	server = &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 23, // 8M
	}
}

func main() {
	var (
		addr string
		err  error
		engi *gin.Engine
	)

	addr = ":8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	engi = gin.Default() // with the Logger and Recovery middleware attached
	// gin.SetMode(gin.ReleaseMode)
	// engi = gin.New()
	// engi.Use(gin.Recovery())

	web.LoadAPI(engi, prom.NewPromMiddleware())

	moni := engi.Group("/api/monitor/v1")
	moni.GET("/prometheus", prom.NewPromHandler())

	server.Addr, server.Handler = addr, engi
	log.Printf("Http server listening on address: %q\n", addr)
	if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server.ListenAndServe: %v\n", err)
	}
}
