package main

import (
	// "fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		isRelease bool
		port      string = ":8080"

		router *gin.Engine
	)

	if len(os.Args) > 1 {
		isRelease, _ = strconv.ParseBool(os.Args[1])
	}

	if len(os.Args) > 2 {
		if port = os.Args[2]; !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
	}

	if isRelease {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	} else {
		router = gin.Default()
	}

	router.GET("/", Hello)
	router.Run(port)
}

func Hello(c *gin.Context) {
	c.String(200, "Hello, world!\n")
	return
}
