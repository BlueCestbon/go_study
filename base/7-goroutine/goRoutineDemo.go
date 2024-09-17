package main

import (
	"fmt"
	"time"
)

func myTask() {
	for {
		fmt.Println("myTask")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go myTask()

	//time.Sleep(1 * time.Second)

	for {
		fmt.Println("main")
		time.Sleep(1 * time.Second)
	}
}
