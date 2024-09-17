package main

import (
	"fmt"
	"time"
)

func noCacheChannel() {
	ch := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine start")
		ch <- 1
	}()

	num, ok := <-ch
	fmt.Println("main start")
	if ok {
		fmt.Println(num)
	}
	fmt.Println("main end")
}

func cacheChannel() {
	ch := make(chan int, 3)
	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine start")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("goroutine:", i)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main start")
	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println("main:", num)
	}
	time.Sleep(1 * time.Second)
}

func closeChannel() {
	ch := make(chan int, 3)
	go func() {
		//defer close(ch)
		fmt.Println("goroutine start")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("goroutine:", i)
		}
		//close(ch)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("main start")
	for {
		if num, ok := <-ch; ok {
			fmt.Println("main:", num)
		} else {
			fmt.Println("channel is closed")
			break
		}
	}
}

func main() {
	//noCacheChannel()
	//cacheChannel()
	//closeChannel()
}
