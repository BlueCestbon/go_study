package main

import "fmt"

func processSlice(slice []int) {
	slice[0] = 100
}

func printSlice(slice []int) {
	for index, value := range slice {
		//fmt.Println("Index:", index, "Value:", value)
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

func sliceDemo() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	printSlice(mySlice)
	processSlice(mySlice)
	printSlice(mySlice)
}

func sliceInit() {
	//mySlice := []int{}  // mySlice is not nil
	//var mySlice []int // mySlice is nil
	mySlice := make([]int, 5) // mySlice is not nil
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	if mySlice == nil {
		fmt.Println("mySlice is nil")
	} else {
		fmt.Println("mySlice is not nil")
	}
}

func main() {
	//sliceDemo()
	sliceInit()
}
