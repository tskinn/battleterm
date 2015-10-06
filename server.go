package main

import (
	"fmt"
//	"net/http"
	"time"
	"net"
//	"strings"
	"bufio"
	"flag"
	"log"
)

var port string
var serverMode boolean

// set flags
func init() {
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.BoolVar(&serverMode, "s", false, "run in server mode to serve matches. this is a crappy description")
	flag.BoolVar(&serverMode, "server", false, "run in server mode... this is also a crappy description")
}

// connect players
func printQueue(chqueue <-chan net.Conn) {
	for {
		p1 := <-chqueue
		p2 := <-chqueue
		fmt.Printf("Local\nPlayer 1: %s\nPlayer 2: %s\n", p1.LocalAddr().String(), p2.LocalAddr().String())

				fmt.Printf("Remote\nPlayer 1: %s\nPlayer 2: %s\n", p1.RemoteAddr().String(), p2.RemoteAddr().String())

		
		p1.Write([]byte(p2.RemoteAddr().String() + "\n"))
		p1.Close()

		p2.Write([]byte(p2.RemoteAddr().String() + "WAIT\n"))
		p2.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

func addToQueue(chqueue chan<- net.Conn, conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received: ", string(message))
	chqueue <- conn
}

// listen on specified port for players and add connections to channel
// i guess the channel isn't necessary any more because this isn't
// concurrent
func main() {
	flag.Parse()
	if port == "" {
		return
	}
	port = ":" + port

	chqueue := make(chan net.Conn, 10)
	go printQueue(chqueue)

	fmt.Printf("Launching server on port %s...\n", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := ln.Accept()

		go addToQueue(chqueue, conn)
		
// 		message, _ := bufio.NewReader(conn).ReadString('\n')
// 		fmt.Print("Message Received: ", string(message))
// //		newmessage := strings.ToUpper(message)
// //		conn.Write([]byte(newmessage + "\n"))
// 		chqueue <- conn
	}
}
