package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

/**
闭包是由函数和与其相关的引用环境组合而成的实体。
简单来说，闭包就是一个引用了作用域之外的变量的函数（Func），该函数的存在时间可以超过创建他的作用域。
*/
// 闭包
func closure() {
	// 当count初始化时，外部函数执行并返回匿名函数，这个匿名函数持有对i的引用，即使外部函数结束，i仍然保存在内存中
	// i在逃逸分析后被保留，没有随着函数的执行结束而结束
	//左边的func()表示一个方法，右边的func() int表示返回值类型是一个函数，函数的值是int类型
	count := func() func() int {
		i := 0 // 初始化函数内变量
		return func() int {
			i++ // 函数内变量加 1
			return i
		}
	}()

	fmt.Println(count())
	fmt.Println(count())
}

// 闭包应用场景
// 返回了一个 http.HandlerFunc, 这个函数里面调用了 fn, 这样的话我们就可以实现链式操作，既执行了中间件代码，又可以继续执行函数，非常方便。
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validPath := regexp.MustCompile(`a(x*)b(y|z)c`)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // 如果没问题则继续执行 fn
	}
}

// 闭包原理的迭代器函数，状态共享
func testIter() {
	num := []int{1, 2, 3, 4}

	// 第一个func声明这个方法，入参是一个数组arr，返回值是一个函数，具体函数的返回值是int, bool
	myIterator := func(arr []int) func([]int) (int, bool) {
		i := -1
		return func(arr []int) (int, bool) {
			i++
			if i < len(arr) {
				return arr[i], true
			}
			return 0, false
		}
	}
	myIter := myIterator(num)

	for {
		value, ok := myIter(num)
		if !ok {
			return
		}
		fmt.Println(value)
	}
}

// 回调函数
func getData(data int, callback func(int)) {
	go func() {
		result := data + 2
		callback(result)
	}()
}
func invokeGetData() {
	callback := func(data int) {
		//fmt.Println(data)
		fmt.Println(strconv.Itoa(data) + "xw")
	}
	getData(1, callback)
	time.Sleep(1 * time.Second)
}

// 函数工厂
func calculationFactory(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "subtract":
		return func(a, b int) int {
			return a - b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	case "divide":
		return func(a, b int) int {
			if b != 0 {
				return a / b
			}
			return 0
		}
	default:
		return nil
	}
}
func invokeCalculationFactory() {
	myCalculate := calculationFactory("add")
	fmt.Println(myCalculate(1, 1))
}
func main() {
	//invokeGetData()
	invokeCalculationFactory()
}
