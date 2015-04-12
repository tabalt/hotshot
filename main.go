package main

import (
	"./server"
	"fmt"
)

func main() {
	fmt.Println("start server")

	server := server.NewServer("4732")
	server.Start()

	fmt.Println("end server")
}
