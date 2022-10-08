package fast_random

import (
	"math/bits"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/probably-not/pow"
)

type Source struct {
	orig   []rand.Source
	shards uint32
	mask   uint32

	atomic32 *uint32
	lock     []sync.Mutex
}

func NewSource(shards uint32, seedFn func() int64) *Source {
	closestPow := uint32(pow.ClosestPowerOfTwoBitwise(int64(shards)))
	factor := uint32(bits.TrailingZeros32(closestPow))

	sources := make([]rand.Source, 0, closestPow)
	for i := uint32(0); i < closestPow; i++ {
		sources = append(sources, rand.NewSource(seedFn()))
	}

	i := uint32(0)
	return &Source{
		orig:     sources,
		atomic32: &i,
		shards:   closestPow,
		mask:     (uint32(1) << uint32(factor)) - 1,
		lock:     make([]sync.Mutex, closestPow),
	}
}

func (s *Source) Int63() int64 {
	n := atomic.AddUint32(s.atomic32, 1)
	head := n & s.mask
	shard := (head + uint32(1)) & s.mask

	s.lock[shard].Lock()
	defer s.lock[shard].Unlock()
	return s.orig[shard].Int63()
}

func (s *Source) Seed(seed int64) {
	for _, o := range s.orig {
		o.Seed(seed)
	}
}
