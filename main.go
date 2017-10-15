// Copyright (c) 2017 Alex Pliutau

package main

import (
	"flag"

	"github.com/plutov/go-snake-telnet/server"
)

func main() {
	var host, port string
	flag.StringVar(&host, "host", "localhost", "TCP Host")
	flag.StringVar(&port, "port", "8080", "TCP Port")
	flag.Parse()

	s := server.New(host + ":" + port)
	s.Run()
}
