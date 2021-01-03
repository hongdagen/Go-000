package slidingwindow_homework

import (
	"container/list"
	"fmt"
)

type SlidingWindow struct {
	Buckets           *list.List
	MaximumBucketsNum int
	cb                chan *Bucket
}

func NewSlidingWindow(maximumBucketsNum int) *SlidingWindow {
	if maximumBucketsNum < 0 {
		panic("")
	}
	sw := &SlidingWindow{Buckets: list.New(), MaximumBucketsNum: maximumBucketsNum, cb: make(chan *Bucket, maximumBucketsNum)}
	return sw
}

func (this *SlidingWindow) updateSlidingWindow(cb chan *Bucket) {
	if this.Buckets.Len() < this.MaximumBucketsNum {
		this.Buckets.PushFront(<-cb)
	} else {
		this.Buckets.Remove(this.Buckets.Back())
		this.Buckets.PushFront(<-cb)
	}
	fmt.Printf("新加入bucket：%s\n", this.Buckets.Front().Value.(*Bucket).name)
	fmt.Printf("sw长度：%d\n", this.Buckets.Len())
	fmt.Printf("bucket success cnt：%d\n", this.Buckets.Front().Value.(*Bucket).GetSuccessIndicator())
}

func (this *SlidingWindow) GetSlidingWindowData() *list.List {
	return this.Buckets
}


