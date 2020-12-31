package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	client, err := DefaultRedisClient()
	if err != nil {
		t.Fatal(err)
	}

	rd := NewRes(1, "something wrong")
	bts, _ := json.Marshal(rd.ResData)

	statusCmd := client.Set("--aaa", bts, 10*time.Second)
	if err := statusCmd.Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>>\n", client.Get("--aaa").String())
}

// $ go test -run=TestDemo_t1 -args ":8080"
func TestDemo_t1(t *testing.T) {
	var (
		addr string
		err  error
	)

	if flag.NArg() > 0 {
		addr = flag.Args()[0]
	} else {
		addr = ":8080"
	}

	if err = Demo(addr); err != nil {
		t.Fatal(err)
	}
}
