package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Eat() {
	fmt.Println("Person Eat")
}

func Walk() {
	fmt.Println("Person Walk")
}

type Student struct {
	Person
	Score float32
}

func (s *Student) Eat() {
	fmt.Println("Student Eat")
}

func main() {
	student := Student{}

	student.Name = "xiaowei"
	student.Age = 18
	student.Score = 99.9

	student.Eat()

}
