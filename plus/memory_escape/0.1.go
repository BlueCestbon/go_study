package main

//go:noinline  // 禁止内联
func foo2() *int {
	var fooVal1 = 11
	var fooVal2 = 12
	var fooVal3 = 13
	var fooVal4 = 14
	var fooVal5 = 15

	for i := 0; i < 5; i++ {
		println(&fooVal1, &fooVal2, &fooVal3, &fooVal4, &fooVal5)
	}

	//返回foo_val3给main函数
	return &fooVal3
}

func main() {
	println(foo2())
}
