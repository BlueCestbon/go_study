package main

import "fmt"

func swap(pa *int, pb *int) {
	*pa, *pb = *pb, *pa
}

func swapOld(pa *int, pb *int) {
	tmp := *pa
	*pa = *pb
	*pb = tmp
}

func main() {
	a := 1
	b := 2
	fmt.Println(a, b)
	//swap(&a, &b)
	swapOld(&a, &b)
	fmt.Println(a, b)

	p := &a
	fmt.Println(p)

	pp := &p // 二级指针
	fmt.Println(pp)

}
