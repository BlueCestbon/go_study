package main

import "fmt"

func foo() *int {
	val := 100
	return &val
}

func main() {
	fmt.Println(*foo())
}
