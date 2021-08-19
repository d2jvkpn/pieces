package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	var (
		delay  int64
		addr   string
		err    error
		result *NTResult
	)

	flag.StringVar(&addr, "addr", "", "request addres")
	flag.Int64Var(&delay, "delay", 10, "delay in millsec")
	flag.Parse()

	if addr == "" {
		log.Fatalf("invalid addr: %q\n", addr)
	}
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}

	if result, err = GetNT(addr, delay); err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}

type NTResult struct {
	T1    time.Time `json:"t1"`
	T2    time.Time `json:"t2"`
	T3    time.Time `json:"t3"`
	T4    time.Time `json:"t4"`
	Sigma int64     `json:"sigma"`
	Delta int64     `json:"delta"`
}

func (result NTResult) String() string {
	bts, _ := json.MarshalIndent(result, "", "  ")
	return string(bts)
}

func GetNT(addr string, delay int64) (result *NTResult, err error) {
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}

	address := fmt.Sprintf("%s?delay=%d", addr, delay)
	result = new(NTResult)
	result.T1 = time.Now()
	var res *http.Response
	if res, err = http.Get(address); err != nil {
		return nil, fmt.Errorf("http.Get(%q): %w", address, err)
	}
	defer res.Body.Close()
	result.T4 = time.Now()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response.StatusCode: %d", res.StatusCode)
	}

	data := struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    struct {
			T2 int64 `json:"t2"` // unix millisec
			T3 int64 `json:"t3"` // unix millisec
		} `json:"data"`
	}{}

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decode response.Body: %w", err)

	}

	result.T2, result.T3 = time.UnixMilli(data.Data.T2), time.UnixMilli(data.Data.T3)
	// fmt.Printf("t1 = %s\nt2 = %s\nt3 = %s\nt4 = %s\n", t1, t2, t3, t4)

	d1 := result.T4.UnixMilli() - result.T1.UnixMilli()
	d2 := result.T3.UnixMilli() - result.T2.UnixMilli()
	result.Sigma = (d1 - d2) / 2

	result.Delta = result.T2.UnixMilli() - result.T1.UnixMilli() - result.Sigma
	return result, nil
}
