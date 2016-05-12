package main

import (
	"net"
	"bufio"
)

type Grid [][]int
type Ships [][]int


type Game struct {
	ourBoard   Grid
	theirBoard Grid
	connection net.Conn
	ourShips   Ships
	messageOne string
	messageTwo string
	ourCursorX int // todo combine these. don't need two. might as well just use
	ourCursorY int // the grid values themselves to keep track of cursor or something like that
	actualX    int
	actualY    int
}

func (game *Game) Listen() bool {
	// get message
	msg, err := bufio.NewReader(game.connection).ReadString('\n')
	if err != nil {
		log.Println(err) // TODO change this to handle it better
	}
	msg = strings.TrimRight(msg, "\n")
	switch msg {
	case "CURSORUPDATE":
		cPos, _ := bufio.NewReader(conn).ReadString('\n')
		coordinates := strings.TrimRight(coordinates, "\n")
		game.ourBoard.updateCursor(coordinates)
		return true		
	case "TURN":
		board, _ := bufio.NewReader(conn).ReadString('\n')
		updateGame(board)
		return false
	}
	return true
}

// remove the cursor from the grid
func (grid * Grid) removeCursor() {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] & cursor == cursor {
				grid[i][j] = grid[i][j] & antiCursor
			}
		}
	}
}

func (grid *Grid)updateCursor(coordinates string) {
	xy := strings.Split(coordinates, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	game.ourBoard.removeCursor()
	game.ourBoard[x][y] |= cursor

}
