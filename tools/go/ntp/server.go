package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not server address provided")
		os.Exit(1)
	}

	addr := os.Args[1]
	ser, _ := NewNTPServer(addr, 10)
	fmt.Printf(">>> NTP server listening on: %q\n", addr)
	ser.Run()
}

// https://en.wikipedia.org/wiki/Network_Time_Protocol
// https://en.wikipedia.org/wiki/File:NTP-Algorithm.svg
type NTPServer struct {
	delay int64
	*http.Server
}

func (ser *NTPServer) Run() error {
	return ser.ListenAndServe()
}

func NewNTPServer(addr string, delay int64) (ser *NTPServer, err error) {
	if delay < 0 {
		return nil, fmt.Errorf("invalid delay")
	}
	if delay == 0 {
		delay = 10
	}
	ser = new(NTPServer)
	ser.delay = delay

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error

		t2 := time.Now()

		if delayStr := r.URL.Query().Get("delay"); delayStr != "" {
			if delay, err = strconv.ParseInt(delayStr, 10, 64); err != nil {
				data := map[string]int64{}
				w.WriteHeader(http.StatusBadRequest)

				json.NewEncoder(w).Encode(
					map[string]interface{}{"code": -1, "message": "bad request", "data": data},
				)
				return
			}
		}

		time.Sleep(time.Duration(ser.delay) * time.Millisecond)
		t3 := time.Now()

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		data := map[string]int64{"t2": t2.UnixMilli(), "t3": t3.UnixMilli()}

		json.NewEncoder(w).Encode(
			map[string]interface{}{"code": 0, "message": "ok", "data": data},
		)
	})

	ser.Server = &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 10, // 8M
	}

	return ser, nil
}
