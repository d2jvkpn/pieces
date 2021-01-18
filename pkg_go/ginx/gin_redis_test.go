package ginx

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"testing"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func testRedisClient(t *testing.T) (client *redis.Client) {
	client, err := DefaultRedisClient()
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func TestRedisClient(t *testing.T) {
	client := testRedisClient(t)

	rd := NewRes(1, "something wrong")
	bts, _ := json.Marshal(rd.ResData)

	statusCmd := client.Set(context.TODO(), "--aaa", bts, 10*time.Second)
	if err := statusCmd.Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>>\n", client.Get(context.TODO(), "--aaa").String())
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

func TestRedisHash(t *testing.T) {
	client := testRedisClient(t)
	now := time.Now()
	key := "device:retention_callback:xxxxxxxx"

	intCmd := client.HSet(
		context.TODO(),
		key,
		"source", "kuaishou",
		"created_at", strconv.FormatInt(now.Unix(), 10),
		"append_tt", "ms",
		"2020-01-12", "https://site.example",
	)

	if err := intCmd.Err(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(intCmd.Val(), intCmd.String())

	ssm := client.HGetAll(context.TODO(), key)
	result, err := ssm.Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

	ssm = client.HGetAll(context.TODO(), "~)(JK)(j`op1mn`0po1jm`m1po`1iopjasoias")
	result, err = ssm.Result()
	fmt.Println(err)
	fmt.Println(result)
}
