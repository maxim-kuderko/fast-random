package main

import (
	fast_random "github.com/maxim-kuderko/fast-random"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	source := fast_random.NewSource(8, func() int64 {
		return time.Now().UnixNano()
	})
	r := rand.New(source)

	wg := sync.WaitGroup{}
	wg.Add(runtime.GOMAXPROCS(0))
	for i := 0; i < 16; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000000; j++ {
				r.Int63()
			}
		}()
	}
	wg.Wait()
}
