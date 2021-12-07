package interview

import (
	"fmt"
	"sync"
	"time"
)

func Waittimeout1() {
	fmt.Println(">>> Waittimeout1:")

	job1 := func(wg *sync.WaitGroup) {
		fmt.Println("    >>>")
		time.Sleep(5 * time.Second)
		wg.Done()
		fmt.Println("    <<<")
		return
	}

	wg := new(sync.WaitGroup)
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go job1(wg)
	}
	go func() {
		wg.Wait()
		ch <- true
	}()

	select {
	case <-ch:
		fmt.Println("    all jobs done!")
	case <-time.After(20 * time.Second):
		fmt.Println("    timeout!")
	}
}
