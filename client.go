package main

import (
	"fmt"
	"net"
	"time"
	"log"
	"bufio"
	"strings"
)

func requestMatch(url string) string {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "hello\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)
	return status
}

func waitForEnemy(port string) net.Conn {
	log.Println(port)
	ln, err := net.Listen("tcp", ":5423")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Message Received: ", string(message))
	return conn
}

func connectTo(enemy string) net.Conn {
	// probably not needed 
	time.Sleep(100 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:5423")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to %s", conn.RemoteAddr().String())
//	conn.Write([]byte("you what is up nigs\n"))
//	status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(status)
	return conn
}

func play(conn net.Conn, player Player, goFirst bool) {
	defer conn.Close()

	// sync messaging between players aka who goes first
	if goFirst {
		
	}
	
	for {
		// get message
	}
	
}

func beClient() {
	serverAddress := serverAddress + port
	response := requestMatch(serverAddress)
	goFirst := true
	var conn net.Conn
	if strings.Contains(response, "WAIT\n") {
		fmt.Printf("ok I'll wait a little bit for enemy to engage, on port: %s\n", port)
		conn = waitForEnemy(port)
		goFirst = false
	} else {
		fmt.Printf("I'm going to connect with the enemy! port: %s\n", response)
		conn = connectTo(response)
	}

	player := Player{Name: name}
//	conn.Close()
	play(conn, player, goFirst)
}
