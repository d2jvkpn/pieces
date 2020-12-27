package utils

import (
	"encoding/json"
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

func TestDemo1(t *testing.T) {
	err := Demo1(":8080")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDemo2(t *testing.T) {
	err := Demo2(":8080")
	if err != nil {
		t.Fatal(err)
	}
}
