package main

import "fmt"

func foo7(a []string) {
	return
}

func main() {
	s := []string{"xw"}
	foo7(s)
	println(s)
	fmt.Println(s)
}
