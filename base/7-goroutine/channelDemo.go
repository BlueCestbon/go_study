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
		//time.Sleep(2 * time.Second)
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
			//close(ch)  // 会报错panic
		}
		//close(ch) // 如果不关闭，下面的读取就报错死锁，因为channel已经没有数据了，但是for循环会一直死循环
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("main start")
	for {
		// ok为true表示channel没有关闭，false表示channel已经关闭
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
	closeChannel()
}
