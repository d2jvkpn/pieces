package TaskLoader

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTaskLoader_t1(t *testing.T) {
	tl := NewTaskLoader(context.TODO(), "node1", true)

	tl.Go(func() error {
		time.Sleep(5 * time.Second)
		return fmt.Errorf("unexpected 1")
	}, nil, "t1")

	tl.Go(func() error {
		time.Sleep(4 * time.Second)
		return nil
	}, nil)

	do3 := func() error {
		time.Sleep(3 * time.Second)
		return fmt.Errorf("unexpected 3")
	}

	tl.Go(do3, nil)
	tl.Wait()
	tls := tl.GetTasks()
	for i := range tls {
		tls[i].PrintError()
	}

	fmt.Println(tl.GetStage())
}

func TestTaskLoader_t2(t *testing.T) {
	tl := NewTaskLoader(context.TODO(), "node2", true)

	tl.Go(func() error {
		time.Sleep(5 * time.Second)
		return fmt.Errorf("task1 unexpected")
	}, nil, "t1")

	tl.GoWithContext(func(ctx context.Context) error {
		ch := make(chan error)

		go func() {
			<-ctx.Done()
			ch <- fmt.Errorf("task2 was cancelled")
		}()

		go func() {
			time.Sleep(9 * time.Second)
			fmt.Println("task2 is done")
			ch <- nil
		}()

		err := <-ch
		return err
	}, "t2")

	time.Sleep(7 * time.Second)
	tl.Cancel()
	tls := tl.GetTasks()
	for i := range tls {
		// tls[i].PrintError()
		fmt.Println(tls[i].String())
	}

	fmt.Println(tl.GetStage())
}

// $ timeout 5 go test -run TestInst3
func TestTaskLoader_t3(t *testing.T) { // test with accepting os interrupt signal
	tl := NewTaskLoader(context.TODO(), "node3", true)

	tl.Go(func() error {
		time.Sleep(4 * time.Second)
		return fmt.Errorf("task1 unexpected")
	}, nil, "t1")

	tl.GoWithContext(func(ctx context.Context) error {
		ch := make(chan error)

		go func() {
			<-ctx.Done()
			ch <- fmt.Errorf("task2 was cancelled")
		}()

		go func() {
			time.Sleep(9 * time.Second)
			fmt.Println("task2 is done")
			ch <- nil
		}()

		err := <-ch
		return err
	}, "t2")

	tl.ListenOSIntr()
	tls := tl.GetTasks()
	for i := range tls {
		// tls[i].PrintError()
		fmt.Println(tls[i].String())
	}

	fmt.Println(tl.GetStage())
}

func TestTaskLoader_t4(t *testing.T) {
	tl := NewTaskLoader(context.TODO(), "node2", true)

	tl.Go(func() error {
		time.Sleep(2 * time.Second)
		v := 0
		fmt.Println(1 / v)
		return nil
	}, nil, "t1")

	tl.Wait()
	tls := tl.GetTasks()
	for i := range tls {
		fmt.Println(tls[i].String())
	}

	fmt.Println(tl.GetStage())
}
