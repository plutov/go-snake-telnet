// Copyright (c) 2017 Alex Pliutau

package server

import (
	"bufio"
	"flag"
	"log"
	"net"
	"time"
)

var (
	serverRunning = false
	host          string
	port          string
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

	for serverRunning {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection")

			continue
		}

		log.Println("Accepted incoming connection from " + conn.RemoteAddr().String())
		go handleConnection(conn)
		go runTicker(time.Tick(1*time.Second), conn)
	}

	log.Println("Server is not running anymore.")
}

func runTicker(tick <-chan time.Time, conn net.Conn) {
	w := bufio.NewWriter(conn)

	for range tick {
		if !serverRunning || conn == nil {
			continue
		}

		// TODO: call ticker function
		w.Write([]byte("\rTick"))
		w.Flush()
	}
}

func handleConnection(conn net.Conn) {
	// TODO: Do something with client connection
}
