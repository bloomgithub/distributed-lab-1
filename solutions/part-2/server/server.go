package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// deals with an error event.
	if err != nil {
		// print error
		fmt.Println("error")
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// continuously accepts network connections from the Listener
	// add connections to the channel for handling connections.
	for {
		conn, err := ln.Accept()
		handleError(err)
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// so long as connection is alive:
	// - reads in new messages as delimited by '\n's
	// - tidies up each message and add it to the messages channel,
	// - records which client it came from.
	reader := bufio.NewReader(client)
	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		message := Message{
			sender:  clientid,
			message: msg,
		}

		msgs <- message
	}
}

func main() {
	// reads in the network port we should listen on, from the commandline argument.
	// - default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	// creates a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", *portPtr)
	handleError(err)

	// creates a channel for connections
	conns := make(chan net.Conn)
	// creates a channel for messages
	msgs := make(chan Message)
	// creates a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	n := 0

	// starts accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			// deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
			client := conn
			clients[n] = conn
			go handleClient(client, n, msgs)
			n++
		case msg := <-msgs:
			// deals with a new message
			// sends the message to all clients that aren't the sender
			for i, client := range clients {
				if msg.sender != i {
					fmt.Fprintf(client, msg.message)
				}
			}
		}
	}
}
