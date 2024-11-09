package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	} // 解析参数
	fmt.Println(r.Form) // 输出到服务端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	name := r.Form["name"]
	fmt.Println(name)
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello %s , how are you", strings.Join(name, ", "))
}

func main() {
	http.HandleFunc("/", sayHelloName)       // 请求地址和方法绑定
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
