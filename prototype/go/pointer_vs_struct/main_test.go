package main

import (
	"fmt"
	"io"
	"testing"
)

// go test -bench=BenchmarkNewD1_t1 -run=^BenchmarkNewD1_t1$ -benchmem -count 10 -v
func BenchmarkNewD1_t1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewD1()
	}
}

// go test -bench=BenchmarkNewD1_t2 -run=^BenchmarkNewD1_t2$ -benchmem -count 10 -v
func BenchmarkNewD1_t2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 := NewD1()
		fmt.Fprintf(io.Discard, "%v", d1)
	}
}

// go test -bench=BenchmarkNewD2_t1 -run=^BenchmarkNewD2_t1$ -benchmem -count 10 -v
func BenchmarkNewD2_t1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewD2()
	}
}

// go test -bench=BenchmarkNewD2_t2 -run=^BenchmarkNewD2_t2$ -benchmem -count 10 -v
func BenchmarkNewD2_t2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 := NewD2()
		fmt.Fprintf(io.Discard, "%v", d1)
	}
}

// go test -bench=BenchmarkNewD2_t3 -run=^BenchmarkNewD2_t3$ -benchmem -count 10 -v
func BenchmarkNewD2_t3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 := NewD2()
		fmt.Fprintf(io.Discard, "%v", &d1)
	}
}
