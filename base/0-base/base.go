package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(str)
	strArray := strings.Split(str, "")
	sort.Strings(strArray)
	str = strings.Join(strArray, "")
	fmt.Println(str)

}
