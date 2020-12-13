package elasticwp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerPool(t *testing.T) {
	testNum := 10
	a := assert.New(t)
	wp, _ := New(3, context.TODO())
	resChan := make(chan int, testNum)
	expected := make([]int, 0)

	for i := range make([]struct{}, testNum) {
		tmp := i
		expected = append(expected, tmp)
		wp.EnqueueTask(func() error {
			resChan <- tmp
			return nil
		})
	}

	a.Nil(wp.Wait())
	close(resChan)

	res := make([]int, 0)
	for i := range resChan {
		res = append(res, i)
	}

	a.ElementsMatch(expected, res)
}
