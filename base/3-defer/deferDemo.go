package main

import "fmt"

/**
 * defer 语句会在函数执行到最后的时候再执行
 * defer 语句的执行顺序是先进后出【defer的函数是入栈】
 */

func returnFunc() int {
	fmt.Println("returnFunc")
	return 0
}

func deferFunc1() {
	fmt.Println("deferFunc1")
}

func deferFunc2() {
	fmt.Println("deferFunc2")
}

func returnAndDefer() int {
	defer deferFunc1()
	defer deferFunc2()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
