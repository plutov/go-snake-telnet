// Copyright (c) 2017 Alex Pliutau

package server

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	serverRunning = false
	host          string
	port          string
	conn          net.Conn
)

// Run prepars the telnet server and begins running it.
func Run() {
	serverRunning = true

	flag.StringVar(&host, "host", "localhost", "TCP Host")
	flag.StringVar(&port, "port", "50000", "TCP Port")
	flag.Parse()

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Failed to start TCP server.")
	}

	log.Println("TCP server started on " + host + ":" + port)

	runServer(listener)
}

func runServer(listener net.Listener) {
	defer listener.Close()
	go runTicker(time.Tick(1 * time.Second))

	var err error
	for serverRunning {
		conn, err = listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection")

			continue
		}

		log.Println("Accepted incoming connection from " + conn.RemoteAddr().String())
		go handleConnection()
	}
}

func runTicker(tick <-chan time.Time) {
	for range tick {
		if !serverRunning || conn == nil {
			return
		}

		// TODO: call ticker function
		conn.Write([]byte("Tick.\r\n"))
	}
}

func handleConnection() {
	conn.Write([]byte("Welcome to Telnet Snake.\r\n"))
}
