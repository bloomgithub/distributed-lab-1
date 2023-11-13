package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	// type lots of messages
	for {
		// send user defined messages
		fmt.Println("Enter text:")
		text, _ := stdin.ReadString('\n')
		// send text down connection
		fmt.Fprintln(conn, text)
		read(&conn)
	}
}
