// Copyright (c) 2017 Alex Pliutau

package server

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/plutov/go-snake-telnet/snake"
)

var (
	host string
	port string
)

const (
	leftTopASCII = "\033[0;0H"
	clearASCII   = "\033[2J"
)

type server struct {
	listener net.Listener
}

// Run prepars the telnet server and begins running it.
func Run() {
	flag.StringVar(&host, "host", "localhost", "TCP Host")
	flag.StringVar(&port, "port", "50000", "TCP Port")
	flag.Parse()

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Failed to start TCP server.")
	}

	log.Println("TCP server started on " + host + ":" + port)

	server := new(server)
	server.listener = listener
	server.runServer()
}

func (s *server) runServer() {
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err.Error())

			continue
		}

		log.Println("Accepted incoming connection from " + conn.RemoteAddr().String())

		game := snake.NewGame()
		go s.handleConnection(conn, game)
	}
}

func (s *server) handleConnection(conn net.Conn, game *snake.Game) {
	// Clear screen and move to 0:0
	conn.Write([]byte(clearASCII + leftTopASCII))
	conn.Write([]byte(leftTopASCII + game.Render()))

	tick := time.Tick(1 * time.Second)
	for range tick {
		// Move to 0:0 and render
		conn.Write([]byte(leftTopASCII + game.Render()))
	}
}
