package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const exitMessage = "exit"

func main() {
	cancel := make(chan int)

	go exitHandler(&cancel)
	go newServer()

	select {
	case <-cancel:
		fmt.Println("Server is shutting down...")
		return
	}
}

func newServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		handleConn(conn)
	}
}

func exitHandler(ch *chan int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == exitMessage {
			fmt.Println("Exit command received")
			*ch <- 1
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
