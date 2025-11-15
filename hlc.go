package hlc

import (
	"sync/atomic"
	"time"
)

const lower16BitsMask int64 = 0xFFFF

type HLC struct {
	lastTS int64
}

func NewHLC() *HLC {
	return &HLC{}
}

func (h *HLC) Next() int64 {
	now := time.Now().UnixNano()
	now = now &^ lower16BitsMask // Clear lower 16 bits for counter

	old := atomic.LoadInt64(&h.lastTS)

	for now > old {
		if atomic.CompareAndSwapInt64(&h.lastTS, old, now) {
			return now
		}
		old = atomic.LoadInt64(&h.lastTS)
	}

	return atomic.AddInt64(&h.lastTS, 1)
}
