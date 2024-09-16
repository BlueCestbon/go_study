package main

import "fmt"

func appendFunc() {
	s1 := make([]int, 3, 5) // len=3, cap=5
	fmt.Println("len:", len(s1), "cap:", cap(s1))
	fmt.Println(s1)
	s1 = append(s1, 1)
	fmt.Println("len:", len(s1), "cap:", cap(s1))
	fmt.Println(s1)
}

func sliceFunc() {
	s1 := []int{1, 2, 3}
	s2 := s1[0 : len(s1)-1]
	s1[0] = 100
	fmt.Println(s2)
}

func copyFunc() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	copy(s2, s1)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}

func main() {
	//appendFunc()
	//sliceFunc()
	//copyFunc()
}
