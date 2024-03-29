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
		showHelp bool
		command  string
		err      error
		flagSet  *flag.FlagSet
	)

	if len(os.Args) < 2 {
		log.Fatalln("commands: serve, client")
	}
	command = os.Args[1]
	flagSet = flag.NewFlagSet(command, 1)

	switch command {
	case "serve":
		showHelp, err = runServe(flagSet, os.Args[2:])
	case "client":
		showHelp, err = runClient(flagSet, os.Args[2:])
	default:
		log.Fatalf("unknown command: %s\n", command)
	}

	if err != nil {
		log.Fatalln(err)
	} else if showHelp {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", flagSet.Name())
		flagSet.PrintDefaults()
	}
}

func runServe(flagSet *flag.FlagSet, args []string) (showHelp bool, err error) {
	var (
		addr           string
		proxies        string
		trustedProxies []string
		debug          bool
		engine         *gin.Engine
		router         *gin.RouterGroup
	)

	flagSet.StringVar(&addr, "addr", ":8080", "http server address")
	flagSet.StringVar(&proxies, "proxies", "", "trusted proxies, separated by comma")
	flagSet.BoolVar(&debug, "debug", false, "run in debug mode")

	if err = flagSet.Parse(args); err != nil {
		return false, err
	}
	// flagSet.Usage: func()
	// fmt.Println("~~~", flagSet.NArg(), flagSet.Args())
	if flagSet.NArg() > 1 {
		return true, nil
	}

	if debug {
		engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}

	trustedProxies = strings.Fields(strings.Replace(proxies, ",", " ", -1))
	if err = engine.SetTrustedProxies(trustedProxies); err != nil {
		return false, err
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
	err = engine.Run(addr)
	return false, err
}

func headers(h http.Header) (strs []string) {
	strs = make([]string, 0, len(h))
	for k, v := range h {
		bts, _ := json.Marshal(v)
		strs = append(strs, fmt.Sprintf("HEAD::%s: %s", k, bts))
	}

	return
}

func runClient(flagSet *flag.FlagSet, args []string) (showHelp bool, err error) {
	var (
		addr  string
		start time.Time
		resp  *http.Response
	)

	flagSet.StringVar(&addr, "addr", "http://localhost:8080", "request http address")
	if err = flagSet.Parse(args); err != nil {
		return false, err
	}
	if flagSet.NArg() > 1 {
		return true, nil
	}

	start = time.Now()
	if resp, err = http.Get(addr); err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(
		"<== %s\n    Status: %d, Proto: %q\n    %s\n    Body: %q\n",
		start.Format(time.RFC3339), resp.StatusCode, resp.Proto,
		strings.Join(headers(resp.Header), "\n    "),
		body,
	)

	return false, nil
}

func inspect(ctx *gin.Context) {
	start := time.Now()
	req := ctx.Request

	record := fmt.Sprintf(
		"ClientIP: %q, RemoteAddr: %q\n    Method: %q, Path: %q, Query: %q, Proto: %q\n    %s",
		ctx.ClientIP(), req.RemoteAddr, req.Method,
		req.URL.Path, req.URL.RawQuery, req.Proto,
		strings.Join(headers(ctx.Request.Header), "\n    "),
	)

	ctx.Next()

	fmt.Printf(
		"<== %s\n    %s\n    Status: %d, Elapsed: %s\n",
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
