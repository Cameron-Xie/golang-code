package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Cameron-Xie/golang-code/workerpool/elasticwp"
	"github.com/Cameron-Xie/golang-code/workerpool/staticwp"
)

func main() {
	fmt.Println("static worker pool")

	sp, ctx := staticwp.New(3, context.TODO())
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			sp.EnqueueTask(echoNum(i))
		}
	}

	fmt.Println(sp.Wait())

	fmt.Println("elastic worker pool")

	ep, ctx := elasticwp.New(1000, context.TODO())

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			ep.EnqueueTask(echoNumElastic(i))
		}
	}

	fmt.Println(ep.Wait())
}

func echoNum(i int) staticwp.Task {
	return func() error {
		if i == 5 {
			return errors.New("error: found five")
		}

		time.Sleep(1 * time.Second)

		fmt.Printf("number: %d\n", i)
		return nil
	}
}

func echoNumElastic(i int) elasticwp.Task {
	return func() error {
		if i == 5 {
			return errors.New("error: found five")
		}

		time.Sleep(1 * time.Second)

		fmt.Printf("number: %d\n", i)
		return nil
	}
}
