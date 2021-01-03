package slidingwindow_homework

import (
	"sync"
	"testing"
	"time"
)

func TestNewSlidingWindow(t *testing.T) {
	var wg sync.WaitGroup
	sw := NewSlidingWindow(10)
	wg.Add(1)
	go ticker(1000, sw)

	for i := 1; i < 100; i++ {
		bucket := NewBucket()
		bucket.Success.Store(i)
		sw.cb <- bucket
	}
	wg.Wait()
}

func ticker(timeInterval time.Duration, sw *SlidingWindow) {
	tt := time.NewTicker(timeInterval * time.Millisecond)
	for {
		select {
		case <-tt.C:
			sw.updateSlidingWindow(sw.cb)
		}
	}
}
