package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	// in a continuous loop, reads a message from the server and display it.
	reader := bufio.NewReader(*conn)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Println(":: ", line)
	}
}

func write(conn *net.Conn) {
	// continually gets input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter text:")
		text, _ := reader.ReadString('\n')
		if text == "/quit\n" {
			break
		} else {
			fmt.Fprintf(*conn, text)
		}
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	// tries to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)

	// starts asynchronously reading and displaying messages
	go read(&conn)

	// starts getting and sending user messages.
	write(&conn)
}
