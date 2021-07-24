package errorx

import (
	"bytes"
	"fmt"
	"testing"
)

var randString = "gSIBPBje2O6cOccRlr9WkLjchKeSsqvUKinHluAGi3WXDJJZdZKsq44kCeTPkUe0RlmZKSFYcaXd2krG8UWYwXOuLoC7MqqbhpNZjVOy9m7izGqueEw9bM6WXxD3gJf3JODz1MEaNdL869qSTRD1GRMd5PCl5i9kkBonnE2PsRhjVs7ze2yV14fLnOWVwYRWUunWAF8q9Ra4cg2mctLyEIxRVRQsRMlBGeuFpOn9q3CZ5XKuD2spDEJzB4L9JEhS"

func TestLogger_t1(t *testing.T) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>> 1", lg.Output(2, "wow"))
	fmt.Println(">>> 2", lg.Output(2, "wow"))
	fmt.Println(">>> 3", lg.Output(2, "wow"))

	if err := lg.Close(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>> 4", lg.Output(2, "wow"))
}

func TestLogger_t2(t *testing.T) {
	lg, err := NewLogger2("wk_logs/abc", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>> 1", lg.Output(2, "wow"))
	fmt.Println(">>> 2", lg.Output(2, "wow"))
	fmt.Println(">>> 3", lg.Output(2, "wow"))

	if err := lg.Close(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>> 4", lg.Output(2, "wow"))
}

func TestLogger_t3(t *testing.T) {
	lg, err := NewLogger2("wk_logs/abc", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	lg.SetStdout()

	fmt.Println(">>> 1", lg.Output(2, "wow"))
	fmt.Println(">>> 2", lg.Output(2, "wow"))
	fmt.Println(">>> 3", lg.Output(2, "wow"))

	if err := lg.Close(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(">>> 4", lg.Output(2, "wow"))
}

func TestLogger_t4(t *testing.T) {
	lg, err := NewLogger2("wk_logs/abc", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	lg.SetStdout()

	str := "world"
	lg.Info("hello, %s\n", str)
}

func TestLogger_t5(t *testing.T) {
	lg, err := NewLogger2("wk_logs/abc", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	lg.SetStdout()

	str := "world"
	lg.Printf("hello, %s\n", str)
}

// go test -bench=Logger_b1 -run=^BenchmarkLogger_b1$ -benchmem -count 10 -v
func BenchmarkLogger_b1(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err = lg.Output(2, randString); err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Logger_x1 -run=^BenchmarkLogger_x1$ -benchmem -count 10 -v
func BenchmarkLogger_x1(b *testing.B) {
	lg, err := NewLogger2("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err = lg.Output(2, randString); err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Logger_x2 -run=^BenchmarkLogger_x2$ -benchmem -count 10 -v
func BenchmarkLogger_x2(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		lg.Info(randString)
	}
}

// go test -bench=Logger_x3 -run=^BenchmarkLogger_x3$ -benchmem -count 10 -v
func BenchmarkLogger_x3(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		lg.Write([]byte(randString + "\n"))
	}
}

// go test -bench=Logger_x4 -run=^BenchmarkLogger_x4$ -benchmem -count 10 -v
func BenchmarkLogger_x4(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	bts := []byte(randString + "\n")

	for i := 0; i < b.N; i++ {
		lg.Write(bts)
	}
}

// go test -bench=Logger_b2 -run=_b2$ -benchmem -count 10 -v
func BenchmarkLogger_b2(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02")
	if err != nil {
		b.Fatal(err)
	}

	ch := make(chan bool, 1000)
	for i := 0; i < b.N; i++ {
		ch <- true
		go func() {
			lg.Output(2, randString)
			<-ch
		}()
	}
}

// go test -bench=Logger_b3 -run=_b3$ -benchmem -count 10 -v
func BenchmarkLogger_b3(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02T15-04-05")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err = lg.Output(2, randString); err != nil {
			b.Fatal(err)
		}
	}
}

// go test -bench=Logger_b4 -run=_b4$ -benchmem -count 10 -v
func BenchmarkLogger_b4(b *testing.B) {
	lg, err := NewLogger("wk_logs/abc", "2006-01-02T15-04-05")
	if err != nil {
		b.Fatal(err)
	}

	ch := make(chan bool, 1000)
	for i := 0; i < b.N; i++ {
		ch <- true
		go func() {
			lg.Output(2, randString)
			<-ch
		}()
	}
}

func TestBuffer(t *testing.T) {
	bts := []byte{'x', 'y', 'z'}
	buf := bytes.NewBuffer(bts)
	buf.WriteString("abcde")
	fmt.Println(buf.String(), string(bts))
}
