package main

import (
	fast_random "github.com/maxim-kuderko/fast-random"
	"math/rand"
	"sync"
	"time"
)

func main() {
	concurrency := 16
	source := fast_random.NewSource(concurrency, func() int64 {
		return time.Now().UnixNano()
	})
	r := rand.New(source)

	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000000; j++ {
				r.Int63()
			}
		}()
	}
	wg.Wait()
}
