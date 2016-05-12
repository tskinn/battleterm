package main

import (
	"fmt"
	"net"
	"time"
	"log"
	"bufio"
	"strings"
	"strconv"
	"github.com/nsf/termbox-go"
)

//
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

// wait for opponents to connect to you
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

//
func connectTo(enemy string) net.Conn {
	// probably not needed 
	time.Sleep(100 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:5423")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to %s", conn.RemoteAddr().String())
	return conn
}

// convert Board(string slice) to int slice
func convToSlice(boardArray []string) [][]int64 {
	var err error

	// initialize 2d array
	grid := make([][]int64, 10)
	for i := 0; i < 10; i++ {
		grid[i] = make([]int64, 10)
	}

	// transfer 1d string array to 2d int array
	for i, num := range boardArray {
		grid[i / 10][i % 10], err = strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	return grid
}

func updateGame(board string) {
	msg := strings.TrimRight(board, "\n")
	boardArry := strings.Split(msg, ",")

	convToSlice(boardArry)
	
	return
}

func (game * Game) moveCursor(ev termbox.Event) {
	
	switch ev.Key {
	case termbox.KeyArrowUp:
		if y > 3 {
			game.actualY -= 2
			game.ourCursorY -= 1
			// y -= 2
			// xY[1] -= 1
		}
	case termbox.KeyArrowDown:
		if y < 20 {
			game.actualY += 2
			game.ourCursorY += 1
			// y += 2
			// xY[1] += 1
		}
	case termbox.KeyArrowLeft:
		if x > 5 + offset {
			game.actualX -= 4
			game.ourCursorX -= 1
			// x -= 4
			// xY[0] -= 1
		}
	case termbox.KeyArrowRight:
		if x < 41 + offset {
			game.actualX += 4
			game.ourCursorX += 1
			// x += 4
			// xY[0] += 1
		}

}

func (game *Game)controller(conn net.Conn, player Player, goFirst bool) {

	game.Setup()
	theirTurn := goFirst

loop:
	for {
		if theirTurn {
			theirTurn = game.Listen()
			break
		}

		ev := termbox.PollEvent()
		if ev.Type == termbox.KeyArrowDown  ||
			ev.Type == termbox.KeyArrowLeft  ||
			ev.Type == termbox.KeyArrowRight ||
			ev.Type == termbox.KeyArrowUp {
			// move cursor
			game.theirBoard.moveCursor(ev)
		} else if ev.Type == termbox.KeySpace {
			// make move
		} else if ev.Tyep == termbox.KeyEsc {
			// quit
		}

		// update view
	}
}

func play(conn net.Conn, player Player, goFirst bool) {
	defer conn.Close()
	game := Game{}
	game.playerSetPieces()
	
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
		case "WAIT": // TODO delete maybe
			log.Println("Ok. I'll wait...")
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
