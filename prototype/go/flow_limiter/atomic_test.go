package main

import (
	"fmt"
	"testing"
)

func BenchmarkCounter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
