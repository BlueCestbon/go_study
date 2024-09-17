package main

import (
	"fmt"
)

func rangeChannel() {
	ch := make(chan int, 3)
	go func() {
		fmt.Println("goroutine start")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("goroutine:", i)
		}
		close(ch)
	}()

	fmt.Println("main start")
	for num := range ch {
		fmt.Println("main:", num)
	}
	fmt.Println("main end")

}

func fibinacci(ch chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		// 如果ch可以写入数据，则写入数据
		case ch <- x:
			x, y = y, x+y
		// 如果quit可以读取数据，则关闭quit
		case <-quit:
			fmt.Println("quit")
			// 退出当前死循环
			return
		}
	}
}
func selectChannel() {
	ch := make(chan int)
	quit := make(chan bool)

	// 1. 通过select实现多路复用
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		// 2. 通过quit关闭goroutine
		quit <- false
	}()

	fmt.Println("main start")
	fibinacci(ch, quit)
	fmt.Println("main end")
}

func main() {
	//rangeChannel()
	selectChannel()
}
