package fast_random

import (
	"testing"
	"time"
)

var x int64

// BenchmarkModulo-10    	68138263	        17.57 ns/op	       0 B/op	       0 allocs/op
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

// BenchmarkBitShift-10    	73486071	        15.76 ns/op	       0 B/op	       0 allocs/op
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
