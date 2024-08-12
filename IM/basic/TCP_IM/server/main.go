package main

func main() {
	s := NewServer("127.0.0.1", 8888)
	s.Start()
}