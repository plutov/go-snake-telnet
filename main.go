package main

import (
	"github.com/plutov/go-snake-telnet/server"
)

func main() {
	s := server.New(":8080")
	s.Run()
}
