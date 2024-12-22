package main

func foo6(a *int) {
	return
}

func main() {
	data := 10
	f := foo6
	f(&data)
	println(data)
}
