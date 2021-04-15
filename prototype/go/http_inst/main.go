package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		release bool
		addr    string
		engi    *gin.Engine
	)

	flag.StringVar(&addr, "addr", "localhost:8080", "http listen address")
	flag.BoolVar(&release, "release", false, "run in release mode")
	flag.Parse()

	///
	SetLogRFC3339()

	if release {
		gin.SetMode(gin.ReleaseMode)
		engi = gin.New()
		engi.Use(gin.Recovery())
	} else {
		engi = gin.Default()
	}

	engi.GET("/hello", Hello)

	server := &http.Server{
		Addr:         addr,
		Handler:      engi,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	///
	shutdown := func() {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		server.Shutdown(ctx) // ignore error
		cancel()
	}

	quit, exch := ListenOSIntr(shutdown)

	go func() {
		var err error
		log.Printf("HTTP serve listen on %q\n", addr)
		if err = server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("ListenAndServe(): %v\n", err)
			quit <- syscall.SIGUSR1
		}
	}()

	<-exch
}

func Hello(ctx *gin.Context) {
	data := map[string]string{"key": "hello", "value": "wold"}
	ctx.JSON(200, data)
	return
}

func RespJSON(c *gin.Context, bts []byte) (int, error) {
	c.Writer.Header().Add("StatusCode", "200")
	c.Writer.Header().Add("Status", "ok")
	c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")

	return c.Writer.Write(bts)
}

func fatalf(format string, v ...interface{}) {
	log.Fatalf(strings.TrimRight(format, "\n ")+"\n", v...)
}

func SetLogRFC3339() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

type logWriter struct{}

func (writer *logWriter) Write(bts []byte) (int, error) {
	// time.RFC3339
	var RFC3339ms = "2006-01-02T15:04:05.000Z07:00"
	return fmt.Print(time.Now().Format(RFC3339ms) + "  " + string(bts))
}

//
func ListenOSIntr(do func(), sgs ...os.Signal) (quit chan os.Signal, exch chan struct{}) {
	// linux support syscall.SIGUSR2
	quit, exch = make(chan os.Signal), make(chan struct{})

	if len(sgs) == 0 {
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)
	} else {
		signal.Notify(quit, sgs...)
	}

	go func() {
		sig := <-quit
		if sig != syscall.SIGUSR1 && do != nil {
			do()
		}
		if sig == os.Interrupt || sig == syscall.SIGTERM {
			log.Println("receive interrupt or sigterm, exit program")
		}
		exch <- struct{}{}
	}()

	return quit, exch
}
