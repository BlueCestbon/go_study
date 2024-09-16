package main

import "fmt"

// 这里和方法一样，大写表示public，小写表示private

type TStudent struct {
	Name  string
	Age   int
	Score float32
}

func (s *TStudent) String() string {
	return fmt.Sprintf("Name=%v, Age=%v, Score=%v", s.Name, s.Age, s.Score)
}

func (s *TStudent) getName() string {
	return s.Name
}

func (s *TStudent) setAge(age int) {
	s.Age = age
}

func main() {
	student := TStudent{"xiaowei", 18, 99.9}
	//fmt.Println(student)

	fmt.Println(student.String())
	fmt.Println(student.getName())
	student.setAge(23)
	fmt.Println(student.String())

}
