package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Cameron-Xie/golang-code/workerpool"
	"time"
)

func main() {
	p, ctx := workerpool.New(3, context.TODO())

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			p.EnqueueTask(echoNum(i))
		}
	}

	fmt.Println(p.Wait())
}

func echoNum(i int) workerpool.Task {
	return func() error {
		if i == 5 {
			return errors.New("error: found five")
		}

		time.Sleep(1 * time.Second)

		fmt.Printf("number: %d\n", i)
		return nil
	}
}
