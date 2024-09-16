package main

import "fmt"

func myFunc(arg interface{}) {
	switch arg.(type) {
	case int:
		fmt.Println("this is int, value is ", arg)
	case string:
		fmt.Println("this is string, value is ", arg)
	default:
		fmt.Println("unknown type")
	}
}

func main() {

	myFunc(100)
	myFunc("hello")
	myFunc(nil)

}
