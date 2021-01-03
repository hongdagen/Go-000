package slidingwindow_homework

import (
	"fmt"
	"testing"
)

func TestBucket_AddSuccessIndicator(t *testing.T) {
	bucket:= NewBucket()
	bucket.AddSuccessIndicator()
	fmt.Println(bucket.Success.Load())
}
