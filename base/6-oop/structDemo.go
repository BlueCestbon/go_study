package main

import "fmt"

type TPerson struct {
	Name string
	Age  int
}

func changePerson1(p TPerson) {
	p.Age = 23
}

func changePerson2(p *TPerson) {
	p.Age = 23
}

func main() {
	person := TPerson{}
	fmt.Println(person)

	person.Name = "xiaowei"
	person.Age = 18
	fmt.Println(person)
	fmt.Println("这个changePerson1函数是值传递，所以不会改变person的值")
	changePerson1(person)
	fmt.Println(person)
	fmt.Println("这个changePerson2函数是指针传递，所以会改变person的值")
	changePerson2(&person)
	fmt.Println(person)
}
