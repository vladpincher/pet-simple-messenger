package main

import (
	"fmt"

	"my-pet-simple-messenger/api/http/server"
)

func main() {

	s := server.NewServer()
	defer s.Close()

	if err := s.Start(); err != nil {
		fmt.Println("server error:", err)
	}
}
