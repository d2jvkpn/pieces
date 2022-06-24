package main

//go:generate bash go_build.sh

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		command string
		err     error
		flagSet *flag.FlagSet
	)

	if len(os.Args) < 2 {
		log.Fatalln("subcommands: serve, client")
	}
	command = os.Args[1]
	flagSet = flag.NewFlagSet(command, 1)

	switch command {
	case "serve":
		err = runServe(flagSet, os.Args[1:])
	case "client":
		err = runClient(flagSet, os.Args[1:])
	default:
		log.Fatalf("unknown commands: %s\n", command)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func runServe(flagSet *flag.FlagSet, args []string) (err error) {
	var (
		addr           string
		proxies        string
		trustedProxies []string
		release        bool
		engine         *gin.Engine
		router         *gin.RouterGroup
	)

	flagSet.StringVar(&addr, "addr", ":8080", "http server address")
	flagSet.StringVar(&proxies, "proxies", "", "trusted proxies, separated by comma")
	flagSet.BoolVar(&release, "release", false, "run in release mode")

	if err = flagSet.Parse(args); err != nil {
		return err
	}

	if release {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	} else {
		engine = gin.Default()
	}

	trustedProxies = strings.Fields(strings.Replace(proxies, ",", " ", -1))
	if len(trustedProxies) > 0 {
		fmt.Println("~~~ Trusted Proxies:", trustedProxies)
		engine.SetTrustedProxies(trustedProxies)
	}
	router = &engine.RouterGroup

	engine.NoRoute(inspect, func(ctx *gin.Context) {
		// ctx.AbortWithStatus(http.StatusNotFound)
		ctx.String(http.StatusNotFound, "Sorry, not found!\n")
	})

	irouters := router.Use(inspect)

	irouters.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!\n")
	})

	fmt.Printf(">>> Http service listen on %s\n", addr)
	return engine.Run(addr)
}

func runClient(flagSet *flag.FlagSet, args []string) (err error) {
	var (
		addr  string
		start time.Time
		resp  *http.Response
	)

	flagSet.StringVar(&addr, "addr", "http://localhost:8080", "request http address")
	if err = flagSet.Parse(args); err != nil {
		return err
	}

	start = time.Now()
	if resp, err = http.Get(addr); err != nil {
		return err
	}
	defer resp.Body.Close()

	bts, _ := json.Marshal(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(
		"<=> Status: %d, Elapsed: %v\n    Headers: %s\n    Body: %q\n",
		resp.StatusCode, time.Since(start), bts, body,
	)

	return nil
}

func inspect(ctx *gin.Context) {
	start := time.Now()
	bts, _ := json.Marshal(ctx.Request.Header)

	record := fmt.Sprintf(
		"ClientIP: %q, RemoteAddr: %q, Method: %q, Path: %q, Query: %q, Headers: %s",
		ctx.ClientIP(), ctx.Request.RemoteAddr, ctx.Request.Method,
		ctx.Request.URL.Path, ctx.Request.URL.RawQuery, bts,
	)

	ctx.Next()

	fmt.Printf(
		"<=> %s %s, Status: %d, Elapsed: %v\n",
		start.Format(time.RFC3339), record,
		ctx.Writer.Status(), time.Since(start),
	)
}

//type arrayFlags []string

//func (i *arrayFlags) String() string {
//    return "my string representation"
//}

//func (i *arrayFlags) Set(value string) error {
//    *i = append(*i, value)
//    return nil
//}

//var myFlags arrayFlags

//func main() {
//    flag.Var(&myFlags, "list1", "Some description for this param.")
//    flag.Parse()
//}
