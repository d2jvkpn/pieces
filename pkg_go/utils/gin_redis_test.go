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

	rd := NewResData(1, "something wrong")
	bts, _ := json.Marshal(rd)

	statusCmd := client.Set("--aaa", bts, 10*time.Second)
	if err := statusCmd.Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>>\n", client.Get("--aaa").String())
}

// treat ResData as err, and mashal the error
func TestResData(t *testing.T) {
	newError := func() error {
		return NewResData(1, "something wrong")
	}

	err := newError()
	fmt.Println("!!!", err)
	bts, _ := json.Marshal(err)
	fmt.Printf(">>> %s\n", bts)
}

func TestDemo_t1(t *testing.T) {
	if err := Demo(":8080"); err != nil {
		t.Fatal(err)
	}
}

// $ go test -run=TestDemo_t2 -args ":8080"
func TestDemo_t2(t *testing.T) {
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
