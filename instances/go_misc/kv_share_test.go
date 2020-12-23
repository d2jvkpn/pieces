package rover

import (
	"fmt"
	"testing"
	"time"
)

func TestKVShare_x1(t *testing.T) {
	data := make(map[string]interface{}, 1e6)
	for i := 0; i < 1e6; i++ {
		data[fmt.Sprintf("%d", i)] = i
	}

	kvs, _ := NewKVShare(data, 1e2)

	t1 := time.Now()
	for i := 0; i < 1e6; i++ {
		v, err := kvs.Get(fmt.Sprintf("%d", i))
		fmt.Println(i, v, err)
	}

	fmt.Println(time.Now().Sub(t1))

	kvs.Close()

	for i := 0; i < 10; i++ {
		v, err := kvs.Get(fmt.Sprintf("%d", i))
		fmt.Println(i, v, err)
	}
}

func TestKVShare_x2(t *testing.T) {
	data := make(map[string]interface{}, 1e6)
	for i := 0; i < 1e6; i++ {
		data[fmt.Sprintf("%d", i)] = i
	}

	t1 := time.Now()
	for i := 0; i < 1e6; i++ {
		fmt.Println(data[fmt.Sprintf("%d", i)])
	}

	fmt.Println(time.Now().Sub(t1))
}
