package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const TIME_FORMAT = "2006-01-02T15:04:05.000Z07:00"

func init() {
	SetLogRFC3339ms()
}

func main() {
	var (
		releaseMode      bool
		port             int
		mode, addr, prog string
		cer, key         string
		err              error

		router *gin.Engine
		srv    *http.Server
		quit   chan os.Signal
		ctx    context.Context
		cancel context.CancelFunc
	)

	flag.BoolVar(&releaseMode, "release", false, "use release mode")
	flag.IntVar(&port, "port", 8080, "service port")
	flag.StringVar(&cer, "cer", "", "certFile path for https")
	flag.StringVar(&key, "key", "", "keyFile path for https")
	flag.Parse()

	prog = path.Base(os.Args[0])

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
		mode, router = "release", gin.New()
	} else {
		mode, router = "default", gin.Default()
	}

	addr = fmt.Sprintf(":%d", port)

	//// http server
	router.GET("", func(c *gin.Context) {
		// name := c.Param("name")

		c.String(200, fmt.Sprintf("Hello, %s!\n", time.Now().Format(TIME_FORMAT)))
		return
	})

	// router.Run(addr)
	srv = &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Graceful restart or stop
	// https://chenyitian.gitbooks.io/gin-web-framework/docs/38.html
	quit = make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2)

	log.Printf("starting %s in %q mode: %q\n", prog, mode, addr)

	go func() {
		var err2 error

		if cer == "" {
			err2 = srv.ListenAndServe()
		} else {
			err2 = srv.ListenAndServeTLS(cer, key)
		}
		if err2 != nil && err2 != http.ErrServerClosed {
			log.Printf("server error: %v\n", err2)
			quit <- syscall.SIGUSR2
		}
	}()

	//// graceful shutdown
	<-quit
	log.Printf("shutdown server...")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	if err = srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %v", err)
	}

	cancel()
	log.Printf("server exit")
}

func SetLogRFC3339ms() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

type logWriter struct{}

func (writer *logWriter) Write(bts []byte) (int, error) {
	// time.RFC3339
	return fmt.Print(time.Now().Format(TIME_FORMAT) + " " + string(bts))
}
