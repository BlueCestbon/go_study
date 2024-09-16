package lib1

import "fmt"

func init() {
	fmt.Println("lib1 init")
}

func Test(str string) {
	fmt.Println("Lib1Test: " + str)
}
