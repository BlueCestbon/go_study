package main

func main() {
	ch := make(chan []string)
	s := []string{"xw"}
	go func() {
		ch <- s
	}()
}
