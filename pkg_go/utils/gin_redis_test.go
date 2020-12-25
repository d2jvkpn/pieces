package utils

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func defaultRedisClient(t *testing.T) (client *redis.Client) {
	client = redis.NewClient(
		&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0},
	)

	statusCmd := client.Ping()
	if err := statusCmd.Err(); err != nil {
		t.Fatal(err)
	}

	return
}

func TestRedisClient(t *testing.T) {
	client := defaultRedisClient(t)

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
