package main

import (
	rand "github.com/maxim-kuderko/fast-random"
	"sync"
)

func main() {
	concurrency := 16
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000000; j++ {
				rand.Int63()
			}
		}()
	}
	wg.Wait()
}
