package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

type Data struct {
	Id    int64
	Name  string
	Tags  []uint8
	At    time.Time
	mutex *sync.RWMutex
}

func main() {
	var m1, m2 runtime.MemStats

	runtime.ReadMemStats(&m1)

	d := Data{
		Id:    1,
		Name:  "你好,",
		Tags:  []uint8{'A', 'B', 'C'},
		At:    time.Now(),
		mutex: new(sync.RWMutex),
	}

	fmt.Printf(
		"d: %d, d.Id: %d, d.Name: %d, d.Tags: %d, d.At: %d, d.mutex: %d, *d.mutex: %d\n",
		unsafe.Sizeof(d),
		unsafe.Sizeof(d.Id),
		unsafe.Sizeof(d.Name),
		unsafe.Sizeof(d.Tags),
		unsafe.Sizeof(d.At),
		unsafe.Sizeof(d.mutex),
		unsafe.Sizeof(*d.mutex),
	)

	runtime.ReadMemStats(&m2)
	fmt.Println(m1.Alloc, m2.Alloc, m2.Alloc-m1.Alloc)
}
