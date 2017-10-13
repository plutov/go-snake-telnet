// Copyright (c) 2017 Alex Pliutau

package server

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"strings"
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
		log.Fatal("Failed to start TCP server: " + err.Error())
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

		go s.handleConnection(conn)
	}
}

func (s *server) handleConnection(conn net.Conn) {
	game := snake.NewGame()

	// Clear screen and move to 0:0
	conn.Write([]byte(clearASCII + leftTopASCII))
	conn.Write([]byte(leftTopASCII))

	go s.read(conn, game)
	go game.Start()

	tick := time.Tick(250 * time.Millisecond)
	for range tick {
		// Move to 0:0 and render
		conn.Write([]byte(leftTopASCII + game.Render()))
		if game.IsOver {
			conn.Close()
		}
	}
}

// Accept input and send it to KeyboardEventsChan
func (s *server) read(conn net.Conn, game *snake.Game) {
	reader := bufio.NewReader(conn)

	for {
		if game.IsOver {
			return
		}

		data, _, err := reader.ReadLine()

		if err == nil {
			key := strings.ToLower(strings.TrimSpace(string(data)))
			if len(key) > 0 {
				game.KeyboardEventsChan <- snake.KeyboardEvent{
					Key: string(key[0]),
				}
			}
		} else {
			log.Println("Read error: " + err.Error())
			if err == io.EOF {
				game.IsOver = true
				conn.Close()
				return
			}
		}
	}
}
