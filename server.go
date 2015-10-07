package main

import (
	"time"
	"net"
	"bufio"
	"log"
)

// connect players
func connectQueue(chqueue <-chan net.Conn) {
	for {
		p1 := <-chqueue
		p2 := <-chqueue
		log.Printf("Connecting players from\n\t%s\n\t%s\n", p1.RemoteAddr().String(), p2.RemoteAddr().String())
		p1.Write([]byte(p2.RemoteAddr().String() + "\n"))
		p1.Close()

		p2.Write([]byte(p2.RemoteAddr().String() + "WAIT\n"))
		p2.Close()
		log.Println("Players succesfully connected")
		time.Sleep(10000 * time.Millisecond)
	}
}

func addToQueue(chqueue chan<- net.Conn, conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	log.Printf("Message Received: %s \tfrom %s", string(message), conn.RemoteAddr().String())
	chqueue <- conn
}

func beServer() {
	chqueue := make(chan net.Conn, 20)
	go connectQueue(chqueue)

	log.Printf("Launching server on port %s...\n", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go addToQueue(chqueue, conn)		
	}
}
