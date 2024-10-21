package main

import (
	"fmt"
)

func main() {
	var dayOfWeek int = 4

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

}

// 结果：
// Thursday
// Friday
// Saturday
