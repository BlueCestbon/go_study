package main

func main() {
	// 因为我需要在wsl里访问win的机器，写的地址就不能是127.0.0.1
	server := NewServer("0.0.0.0", 8888)
	server.Start()
}
