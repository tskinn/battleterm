package main

import (
	"fmt"
	"net"
	"time"
	"log"
	"bufio"
	"strings"
	"strconv"
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
	ln, err := net.Listen("tcp", ":5423")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to ", conn.RemoteAddr().String())
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

func setPieces() {
	return
}

func convToSlice(boardArray []string) (grid [][]int64) {
	var err error
	for i, num := range boardArray {
		grid[i / 10][i % 10], err = strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func updateGame(board string) {
	msg := strings.TrimRight(board, "\n")
	boardArry := strings.Split(msg, ",")

	convToSlice(boardArry)
	
	return
}

func updateEnemyCursor(msg string) {
	cPos := strings.TrimRight(msg, "\n")
	xy := strings.Split(cPos, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	log.Printf("Enemy moved to x:%d, y:%d", x, y)
	return
}

func play(conn net.Conn, player Player, goFirst bool) {
	defer conn.Close()

	playerSetPieces()
	
	// sync messaging between players aka who goes first
	if goFirst {
		
	}
	
	for {
		// get message
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err) // TODO change this to handle it better
		}
		msg = strings.TrimRight(msg, "\n")
		switch msg {
		case "CURSOR":
			cPos, _ := bufio.NewReader(conn).ReadString('\n')
			updateEnemyCursor(cPos)
		case "TURN":
			board, _ := bufio.NewReader(conn).ReadString('\n')
			updateGame(board)
		}
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
