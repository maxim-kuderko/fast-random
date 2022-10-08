package fast_random

import (
	"testing"
	"time"
)

var x int64

// BenchmarkModulo-10    	70365412	        17.30 ns/op	       0 B/op	       0 allocs/op
func BenchmarkModulo(b *testing.B) {
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

// BenchmarkBitShift-10    	75204264	        15.99 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBitShift(b *testing.B) {
	src := NewBitShiftSource(64, func() int64 {
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
