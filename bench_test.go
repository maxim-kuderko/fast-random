package fast_random

import (
	"testing"
	"time"
)

var x int64

// BenchmarkModulo-10    	78919891	        15.11 ns/op	       0 B/op	       0 allocs/op
func BenchmarkModulo(b *testing.B) {
	src := NewSource(16, func() int64 {
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

// BenchmarkBitShift-10    	85562766	        14.01 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBitShift(b *testing.B) {
	src := NewBitShiftSource(16, func() int64 {
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
