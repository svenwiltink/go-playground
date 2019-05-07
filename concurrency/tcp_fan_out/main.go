package main

import (
	"fmt"
	"net"
)

const maxConn = 2

func main() {

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		panic(err)
	}

	connChan := make(chan net.Conn)
	startWorkers(connChan)

	for {
		connection, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		select {
		case connChan <- connection:
		default:
			connection.Write([]byte("too many clients"))
			connection.Close()
		}
	}
}

func startWorkers(channel chan net.Conn) {
	for i := 0; i < maxConn; i++ {
		go startWorker(channel)
	}
}

func startWorker(channel chan net.Conn) {
	for connection := range channel {
		handleConn(connection)
	}
}

func handleConn(conn net.Conn) {
	defer fmt.Println("connection closed")
	defer conn.Close()
	conn.Write([]byte("HEY"))

	buff := make([]byte, 4096)
	for _, err := conn.Read(buff); err == nil; {

	}
}
