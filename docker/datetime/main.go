package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TimeFormat = "2006-01-02T15:040:05.000Z07:00"
)

var (
	logFields = strings.Join(
		[]string{
			"%s", // time
			"%s", // ip
			"%s", // method
			"%s", // path
			"%d", // content length
			"%s", // handler
			"%d", // status code
			"%d", // response bytes
			"%s", // error
			"%s", // user agent
		}, "\t",
	) + "\n"
)

func main() {
	var (
		address           string
		certFile, keyFile string
		err               error
	)

	flag.StringVar(&address, "address", ":8080", "http server address")
	flag.StringVar(&certFile, "certFile", "", "TLS cert file location")
	flag.StringVar(&keyFile, "keyFile", "", "TLS key file location")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Datetime)

	server := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 23, // 8M
		Handler:           mux,
	}

	if certFile == "" || keyFile == "" {
		log.Printf(">>> http listen on address: %q\n", address)
		err = server.ListenAndServe()
	} else {
		log.Printf(">>> https listen on address: %q\n", address)
		err = server.ListenAndServeTLS(certFile, keyFile)
	}
	if err == http.ErrServerClosed {
		log.Println("exit program.")
	} else {
		log.Fatalf("%v\n", err)
	}
}

func Datetime(writer http.ResponseWriter, req *http.Request) {
	var (
		n      int
		nowStr string
		errStr string
		now    time.Time
		err    error
	)

	now = time.Now()
	nowStr = now.Format(TimeFormat)

	writer.Header().Add("StatusCode", strconv.Itoa(http.StatusOK))
	writer.Header().Add("Status", "ok")

	if strings.Contains(req.Header.Get("Accept"), "application/json") {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		bts, _ := json.Marshal(map[string]string{"now": nowStr})
		n, err = writer.Write(append(bts, '\n'))
	} else {
		n, err = writer.Write([]byte(nowStr + "\n"))
	}
	if err != nil {
		errStr = err.Error()
	}

	fmt.Printf(logFields,
		nowStr, ReqIP(req), "GET", "/",
		req.ContentLength, "Datetime", http.StatusOK, n,
		errStr, req.Header.Get("User-Agent"),
	)
	return
}

func ReqIP(req *http.Request) (ip string) {
	if ip = req.Header.Get("X-Real-IP"); ip != "" {
		return
	}
	if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		return
	}
	ip, _, _ = net.SplitHostPort(req.RemoteAddr)

	return
}
