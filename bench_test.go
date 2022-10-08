package fast_random

import (
	"testing"
	"time"
)

var x int64

// BenchmarkSource-10    	77507067	        15.56 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSource(b *testing.B) {
	src := NewSource(64, func() int64 {
		return time.Now().UnixNano()
	})

	b.ResetTimer()
	b.ReportAllocs()
	var n int64
	for i := 0; i < b.N; i++ {
		n = src.Int63()
	}
	b.StopTimer()
	x = n
}
