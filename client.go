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
	fmt.Println(status)
	return status
}

func waitForEnemy(port string) net.Conn {
	fmt.Println(port)
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
	fmt.Println(enemy)
	time.Sleep(1000 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:5423")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected!")
	fmt.Println(conn.LocalAddr().String())
	fmt.Println(conn.RemoteAddr().String())
	conn.Write([]byte("you what is up nigs\n"))
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
	return conn
}

func play(conn net.Conn, player Player) {
	conn.Close()
}

func beClient() {
	url := "127.0.0.1" + port
	enemy := requestMatch(url)
	var conn net.Conn
	if strings.Contains(enemy, "WAIT\n") {
		newPort := strings.Split(enemy, ":")
		ort := strings.Replace(newPort[len(newPort)-1], "WAIT\n", "", 1)
		ort = ":" + ort
		fmt.Printf("ok I'll wait a little bit, on port: %s\n", port)
		conn = waitForEnemy(port)
	} else {
		fmt.Printf("I'm going to connect with the enemy! port: %s\n", enemy)
		conn = connectTo(enemy)
	}

	player := Player{Name: "joe"}
	conn.Close()
	play(conn, player)
}
