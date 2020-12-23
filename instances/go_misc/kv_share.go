package rover

import (
	"errors"
	"fmt"
	"sync"
)

var (
	KVShareClosed   = errors.New("channel is closed")
	KVShareNotFound = errors.New("key not found")
)

type KVShare struct {
	data  map[string]interface{}
	mutex sync.Mutex
	c     chan struct{}
	n     int
}

func NewKVShare(ndata map[string]interface{}, n int) (kvs *KVShare, err error) {
	var i int
	var k string

	if n < 1 {
		err = fmt.Errorf("n must greater than 0")
		return
	}

	kvs = new(KVShare)
	kvs.mutex.Lock()
	kvs.n = n

	kvs.data = make(map[string]interface{}, len(ndata))
	for k = range ndata {
		kvs.data[k] = ndata[k]
	}

	kvs.c = make(chan struct{}, kvs.n)
	for i = 0; i < kvs.n; i++ {
		kvs.c <- struct{}{}
	}

	kvs.mutex.Unlock()

	return
}

func (kvs *KVShare) Get(key string) (value interface{}, err error) {
	var ok bool
	if _, ok = <-kvs.c; !ok {
		err = KVShareClosed
		return
	}

	if value, ok = kvs.data[key]; !ok {
		err = KVShareNotFound
		return
	}

	kvs.c <- struct{}{}

	return
}

func (kvs *KVShare) Update(ndata map[string]interface{}) {
	var i int
	var k string

	kvs.mutex.Lock()

	for i = 0; i < kvs.n; i++ {
		<-kvs.c
	}

	for k = range ndata {
		kvs.data[k] = ndata[k]
	}

	for i = 0; i < kvs.n; i++ {
		kvs.c <- struct{}{}
	}

	kvs.mutex.Unlock()
}

func (kvs *KVShare) Close() {
	var i int

	kvs.mutex.Lock()
	for i = 0; i < kvs.n; i++ {
		<-kvs.c
	}

	close(kvs.c)
	kvs.data = nil
	kvs.mutex.Unlock()
}
