package main

import (
	"bufio"
	"strings"
	"strconv"
	"github.com/nsf/termbox-go"
)

type Grid [][]int
type Ships [][]int
type Connection interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
}

type Game struct {
	ourBoard   Grid
	theirBoard Grid
	connection Connection
	ourShips   Ships
	messageOne string
	messageTwo string
	ourCursorX int // todo combine these. don't need two. might as well just use
	ourCursorY int // the grid values themselves to keep track of cursor or something like that
	actualX    int
	actualY    int
}

func (game *Game) Listen() (bool, bool) {
	// get message
	msg, err := bufio.NewReader(game.connection).ReadString('\n')
	if err != nil {
		logger.Println(err) // TODO change this to handle it better
	}
	msg = strings.TrimRight(msg, "\n")
	switch msg {
	case "CURSORUPDATE":
		coordString, _ := bufio.NewReader(game.connection).ReadString('\n')
		coordinates := strings.TrimRight(coordString, "\n")
		game.updateCursor(coordinates)
		return true, false
	case "TURN":
		board, _ := bufio.NewReader(game.connection).ReadString('\n')
		updateGame(board)
		return false, false
	case "QUIT":
		return false, true
	}
	return true, false
}

// remove the cursor from the grid
func (grid * Grid) removeCursor() {
	for i := range (*grid) {
		for j := range (*grid)[i] {
			if (*grid)[i][j] & cursor == cursor {
				(*grid)[i][j] = (*grid)[i][j] & antiCursor
			}
		}
	}
}

func (game *Game)updateCursor(coordinates string) {
	xy := strings.Split(coordinates, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	game.ourBoard.removeCursor()
	game.ourBoard[x][y] |= cursor

}

func (grid *Grid)init() {
	*grid = make([][]int, 10)
	for i, _ := range (*grid) {
		(*grid)[i] = make([]int, 10)
	}
}

func (game *Game)Setup() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputAlt)
	
	width, height = termbox.Size() 

	draw()
	x = 5
	//	x = 51
	y = 3
	//	offset = 46
	termbox.SetCursor(x,y)
	termbox.Flush()
	
	game.ourBoard.init()
	game.theirBoard.init()
	game.setPieceShip(littleShip)
	game.draw()
	game.setPieceShip(sub)
	game.draw()
	game.setPieceShip(frigate)
	game.draw()
	game.setPieceShip(battleship)
	game.draw()
	game.setPieceShip(airCraftCarrier)
	game.draw()
}

// func (game *Game)setShip(xy, start []int) {

// }
