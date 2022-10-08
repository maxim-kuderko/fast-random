package fast_random

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

type Source struct {
	orig   []rand.Source
	shards uint32

	atomic32 *uint32
	lock     []sync.Mutex
}

func NewSource(shards uint32, seedFn func() int64) *Source {
	sources := make([]rand.Source, 0, shards)
	for i := uint32(0); i < shards; i++ {
		sources = append(sources, rand.NewSource(seedFn()))
	}
	i := uint32(0)
	return &Source{orig: sources, atomic32: &i, shards: shards, lock: make([]sync.Mutex, shards)}
}

func (s *Source) Int63() int64 {
	n := atomic.AddUint32(s.atomic32, 1)
	if n >= s.shards {
		atomic.CompareAndSwapUint32(s.atomic32, n, 0)
	}
	shard := n % s.shards
	s.lock[shard].Lock()
	defer s.lock[shard].Unlock()
	return s.orig[shard].Int63()
}

func (s *Source) Seed(seed int64) {
	for _, o := range s.orig {
		o.Seed(seed)
	}
}
