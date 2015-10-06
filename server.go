package main

import (
	"fmt"
//	"net/http"
	"time"
	"net"
	"strings"
	"bufio"
//	"flag"
	"log"
)

// struct Player {
// 	Name string
// }

// var port string
// var serverMode bool

// // set flags
// func init() {
// 	flag.StringVar(&port, "port", "8080", "port to listen on")
// 	flag.StringVar(&port, "p", "8080", "port to listen on")
// 	flag.BoolVar(&serverMode, "s", false, "run in server mode to serve matches. this is a crappy description")
// 	flag.BoolVar(&serverMode, "server", false, "run in server mode... this is also a crappy description")
// }

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

func beServer() {
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
	}
}

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
//	fmt.Fprintf(conn, "you what is up nigs\n")
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
		fmt.Printf("ok I'll wait a little bit, on port: %s\n", ort)
		conn = waitForEnemy(port)
	} else {
		fmt.Printf("I'm going to connect with the enemy! port: %s\n", enemy)
		conn = connectTo(enemy)
	}

	player := Player{Name: "joe"}
	conn.Close()
	play(conn, player)
}

// listen on specified port for players and add connections to channel
// i guess the channel isn't necessary any more because this isn't
// concurrent
// func main() {
// 	flag.Parse()
// 	port = ":" + port
// 	if serverMode {
// 		beServer()
// 	} else {
// 		beClient()
// 	}
	
// }
