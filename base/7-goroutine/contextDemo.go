package main

import (
	"context"
	"fmt"
	"time"
)

func testWithTimeOut() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	//context.WithDeadline(context.Background(), time.Now().Add(3 * time.Second))

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("time out")
				return
			default:
				fmt.Println("running ... ")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
}

func testWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("time out")
				return
			default:
				fmt.Println("running ... ")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func main() {
	//testWithTimeOut()
	testWithCancel()
}
