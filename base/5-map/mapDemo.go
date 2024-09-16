package main

import "fmt"

func baseMap() {
	map1 := map[string]string{}
	map1["name"] = "xiaowei"
	map1["age"] = "18"
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["name"] = "xiaowei"
	map2["age"] = "18"
	fmt.Println(map2)

	map3 := map[string]string{
		"name": "xiaowei",
		"age":  "18",
	}
	fmt.Println(map3)
}

func appMap() {
	// 增加
	map1 := make(map[string]string)
	map1["name"] = "xiaowei"
	map1["age"] = "18"

	// 遍历
	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// 修改
	map1["age"] = "23"
	fmt.Println(map1)

	// 删除
	delete(map1, "age")
	fmt.Println(map1)

}

func main() {
	//baseMap()
	appMap()
}
