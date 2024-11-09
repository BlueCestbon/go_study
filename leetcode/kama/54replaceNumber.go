package main

import (
	"fmt"
)

func main() {
	var strByte []byte
	//fmt.Scanln(&strByte)

	strByte = []byte("a1b2c3")

	oldSize := len(strByte)
	count := 0
	for _, str := range strByte {
		if str >= '0' && str <= '9' {
			count++
		}
	}

	// 扩容
	for i := 0; i < count; i++ {
		strByte = append(strByte, []byte("     ")...)
	}
	newSize := len(strByte)

	tmpBytes := []byte("number")
	left := oldSize - 1
	right := newSize - 1
	for left >= 0 {
		str := strByte[left]
		if str >= '0' && str <= '9' {
			for i, tmpByte := range tmpBytes {
				strByte[right-len(tmpBytes)+i+1] = tmpByte
			}
			right -= len(tmpBytes)
		} else {
			strByte[right] = strByte[left]
			right--
		}
		left--
	}
	fmt.Println(string(strByte))
}
