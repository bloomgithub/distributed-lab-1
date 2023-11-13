package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	// print out lots of messages
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		// send reply to client
		fmt.Fprintln(*conn, "OK")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8030")
	// persists and handles new connections
	for {
		conn, _ := ln.Accept()
		go handleConnection(&conn)
	}

}
