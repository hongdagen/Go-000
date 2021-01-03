package slidingwindow_homework

import (
	"sync/atomic"
	"time"
)

type Bucket struct {
	name      string
	Success   atomic.Value
	Failure   atomic.Value
	Timeout   atomic.Value
	Rejection atomic.Value
}

func NewBucket() *Bucket {
	return &Bucket{name: time.Now().Format(time.RFC3339Nano)}
}

func (b *Bucket) AddSuccessIndicator() {
	b.Success.Store(1)
}

func (b *Bucket) AddFailureIndicator() {
	b.Failure.Store(1)
}
func (b *Bucket) AddTimeoutIndicator() {
	b.Timeout.Store(1)
}
func (b *Bucket) AddRejectionIndicator() {
	b.Rejection.Store(1)
}
func (b *Bucket) GetSuccessIndicator() int {
	return b.Success.Load().(int)
}
func (b *Bucket) GetFailureIndicator() int {
	return b.Failure.Load().(int)
}
func (b *Bucket) GetTimeoutIndicator() int {
	return b.Timeout.Load().(int)
}
func (b *Bucket) GetRejectionIndicator() int {
	return b.Rejection.Load().(int)
}
