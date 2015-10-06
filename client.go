package main

import (
	"fmt"
	"net"
	"time"
//	"io/ioutil"
	"log"
	"bufio"
	"flag"
	"strings"
)

var port string

// set flags
func init() {
	flag.StringVar(&port, "port", "8888", "port to listen on")
	flag.StringVar(&port, "p", "8888", "port to listen on")
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

func waitForEnemy(port string) {
	fmt.Println(port)
	ln, err := net.Listen("tcp", ":8080")
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
}

func connectTo(enemy string) {
	fmt.Println(enemy)
	time.Sleep(1000 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
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
	
}

func main() {
	flag.Parse()
	port = ":" + port
	url := "127.0.0.1:8081"
	enemy := requestMatch(url)
	if strings.Contains(enemy, "WAIT\n") {
		newPort := strings.Split(enemy, ":")
		ort := strings.Replace(newPort[len(newPort)-1], "WAIT\n", "", 1)
		ort = ":" + ort
		fmt.Printf("ok I'll wait a little bit, on port: %s\n", ort)
		waitForEnemy(port)
	} else {
		fmt.Printf("I'm going to connect with the enemy! %s\n", enemy)
		connectTo(enemy)
	}
}
