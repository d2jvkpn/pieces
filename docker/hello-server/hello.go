package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		releaseMode      bool
		mode, port, prog string
		err              error

		router *gin.Engine
		srv    *http.Server
		quit   chan os.Signal
		ctx    context.Context
		cancel context.CancelFunc
	)

	flag.BoolVar(&releaseMode, "release", false, "use release mode")
	flag.StringVar(&port, "port", ":8080", "specify service port")
	flag.Parse()

	logStderr := func(msg string, a ...interface{}) {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		fmt.Fprintf(os.Stderr, time.Now().Format("["+time.RFC3339+"] ")+msg, a...)
		return
	}

	if releaseMode {
		mode = "release"
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	} else {
		mode = "default"
		router = gin.Default()
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	prog = path.Base(os.Args[0])

	//// http server
	router.GET("/:name", func(c *gin.Context) {
		var name string = c.Param("name")

		if name == "" {
			name = "world"
		}

		c.String(200, fmt.Sprintf("Hello, %s!\n", name))
		return
	})

	// router.Run(port)
	srv = &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Graceful restart or stop
	// https://chenyitian.gitbooks.io/gin-web-framework/docs/38.html
	quit = make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2)

	go func() {
		var err error

		logStderr("Satrting %s in %q mode using port %q", prog, mode, port)

		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logStderr("Server error: %v", err)
			quit <- syscall.SIGUSR2
		}
	}()

	//// graceful shutdown
	<-quit
	logStderr("Shutdown server...")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	if err = srv.Shutdown(ctx); err != nil {
		logStderr("Server shutdown: %v", err)
	}

	cancel()
	logStderr("Server exit")
}
