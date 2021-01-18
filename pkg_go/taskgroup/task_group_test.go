package taskgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestInst1(t *testing.T) {
	tg := NewTaskGroup(context.TODO(), "node1", true)

	tg.Go(func() error {
		time.Sleep(5 * time.Second)
		return fmt.Errorf("unexpected 1")
	}, nil, "t1")

	tg.Go(func() error {
		time.Sleep(4 * time.Second)
		return nil
	}, nil)

	do3 := func() error {
		time.Sleep(3 * time.Second)
		return fmt.Errorf("unexpected 3")
	}

	tg.Go(do3, nil)
	tg.Wait()
	tgs := tg.GetTasks()
	for i := range tgs {
		tgs[i].PrintError()
	}

	fmt.Println(tg.GetStage())
}

func TestInst2(t *testing.T) {
	tg := NewTaskGroup(context.TODO(), "node2", true)

	tg.Go(func() error {
		time.Sleep(5 * time.Second)
		return fmt.Errorf("task1 unexpected")
	}, nil, "t1")

	tg.GoWithContext(func(ctx context.Context) error {
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
	tg.Cancel()
	tgs := tg.GetTasks()
	for i := range tgs {
		// tgs[i].PrintError()
		fmt.Println(tgs[i].String())
	}

	fmt.Println(tg.GetStage())
}

// $ timeout 5 go test -run TestInst3
func TestInst3(t *testing.T) { // test with accepting os interrupt signal
	tg := NewTaskGroup(context.TODO(), "node3", true)

	tg.Go(func() error {
		time.Sleep(4 * time.Second)
		return fmt.Errorf("task1 unexpected")
	}, nil, "t1")

	tg.GoWithContext(func(ctx context.Context) error {
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

	tg.ListenOSIntr()
	tgs := tg.GetTasks()
	for i := range tgs {
		// tgs[i].PrintError()
		fmt.Println(tgs[i].String())
	}

	fmt.Println(tg.GetStage())
}

func TestInst4(t *testing.T) {
	tg := NewTaskGroup(context.TODO(), "node2", true)

	tg.Go(func() error {
		time.Sleep(2 * time.Second)
		v := 0
		fmt.Println(1 / v)
		return nil
	}, nil, "t1")

	tg.Wait()
	tgs := tg.GetTasks()
	for i := range tgs {
		fmt.Println(tgs[i].String())
	}

	fmt.Println(tg.GetStage())
}
