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
	//var mySlice []int // mySlice is nil  // 这是建议写法，因为这样字不会初始化，只有真正赋值的时候才会进行内存分配，可以节省内存
	mySlice := make([]int, 5) // mySlice is not nil
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	if mySlice == nil {
		fmt.Println("mySlice is nil")
	} else {
		fmt.Println("mySlice is not nil")
	}
}

// 数组是值传递，切片是引用传递
func changeArrayAndSlice(myArr [3]int, mySlice []int) {
	myArr[0] = 99
	mySlice[0] = 99
}

func invokeArrayAndSlice() {
	myArray := [3]int{1, 2, 3}
	mySlice := []int{1, 2, 3}
	changeArrayAndSlice(myArray, mySlice)
	fmt.Println(myArray)
	fmt.Println(mySlice)
}

func main() {
	//sliceDemo()
	//sliceInit()
	invokeArrayAndSlice()
}
