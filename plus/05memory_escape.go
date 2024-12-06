package main

//go:noinline  // 禁止内联
func foo(argVal int) *int {

	var fooVal1 = 11
	var fooVal2 = 12
	var fooVal3 = 13
	var fooVal4 = 14
	var fooVal5 = 15

	//此处循环是防止go编译器将foo优化成inline(内联函数)
	//如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量，即使foo_val3发生逃逸，地址与其他也是连续的
	for i := 0; i < 5; i++ {
		println(&argVal, &fooVal1, &fooVal2, &fooVal3, &fooVal4, &fooVal5)
	}

	//返回foo_val3给main函数
	return &fooVal3
}

func main() {
	mainVal := foo(666)

	println(*mainVal, mainVal)
}
