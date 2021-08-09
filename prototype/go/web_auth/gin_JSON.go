package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

const (
	CONN_ADDRESS = "localhost"
	CONN_PORT    = ":9000"
)

func main() {
	engine := gin.Default()
	{
		engine.GET("/hello", hello)
		engine.GET("/Hello", Hello)
	}

	server := &http.Server{
		Addr:         CONN_ADDRESS + CONN_PORT,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Fprintf(os.Stderr, ">>> Service %s%s\n", CONN_ADDRESS, CONN_PORT)
	server.ListenAndServe()
}

func hello(c *gin.Context) {
	bts := []byte(`{"hello": "world"}`)
	RespJsonBytes(c, bts)
}

func Hello(c *gin.Context) {
	d := map[string]string{"hello": "wold"}
	c.JSON(200, d)
}

func RespJsonBytes(c *gin.Context, bts []byte) (int, error) {
	c.Writer.Header().Add("StatusCode", "200")
	c.Writer.Header().Add("Status", "ok")
	c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")

	return c.Writer.Write(bts)
}
