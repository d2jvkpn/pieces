# taskloader

Package taskloader provides synchronization, error propagation, context cancelation, record for
groups of goroutines(tasks) working on subtasks of a common task.


#### 1. Support status
- task status
  - running
  - cancelled
  - done/failed/unexpected
  
- task group status
  - starting
  - running
  - canceling
  - done/exit
  
#### 2. Demo

- create taskloader
```go
tl := NewTaskLoader(context.TODO(), "node1")
// tl := NewTaskLoader(context.TODO(), "node1", true) // debug mode 
```

- add task
```go
tl.Go(func() error {
	time.Sleep(4 * time.Second)
	return fmt.Errorf("task1 unexpected")
}, nil, "t1")
```

- add task accept context
```go
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
```

- waiting all tasks exit
```go
tl.Wait()
```

- canceling all tasks
```go
tl.Cancel()
```

- listen os interrupt signal and cancel all task
```go
tl.ListenOSIntr()
```
