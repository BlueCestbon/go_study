package main

import (
	//_ "study/base/init-import/lib1"  // 此为匿名导入，可以使用init方法
	"study/base/1-init-import/lib1"
	"study/base/1-init-import/lib2"
)

func main() {
	//fmt.Println("Hello, World!")
	lib1.Test("Hello, World!")
	lib2.Test("Hello, World!")
	//println("Hello, World!")
}
