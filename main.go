package main

import "fizz_buzz/server"

func main() {
	s := server.New("8080")
	s.Start()
}
