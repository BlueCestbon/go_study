package main

import "fmt"

import (
	"github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	//ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	algorithm.BubbleSort(numbers, comparator)

	fmt.Println(numbers)

	// Output:
	// [1 2 3 4 5 6]
}
