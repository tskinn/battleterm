package main

import (
	"github.com/nsf/termbox-go"
	"strings"
	"strconv"
)
//  0123456789012345678901    
var ui = [...]string{"                    YOURS                                         THEIRS                  ",
	"   | A | B | C | D | E | F | G | H | I | J |     | A | B | C | D | E | F | G | H | I | J |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 0 |   |   |   |   |   |   |   |   |   |   |   0 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 1 |   |   |   |   |   |   |   |   |   |   |   1 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 2 |   |   |   |   |   |   |   |   |   |   |   2 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 3 |   |   |   |   |   |   |   |   |   |   |   3 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 4 |   |   |   |   |   |   |   |   |   |   |   4 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 5 |   |   |   |   |   |   |   |   |   |   |   5 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 6 |   |   |   |   |   |   |   |   |   |   |   6 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 7 |   |   |   |   |   |   |   |   |   |   |   7 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 8 |   |   |   |   |   |   |   |   |   |   |   8 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 9 |   |   |   |   |   |   |   |   |   |   |   9 |   |   |   |   |   |   |   |   |   |   |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	"  _________ __+=__ __|7_  _1_  ____             _________ __(>__ __|7_  _1_  ____         ",
	"   \\#####/  \\####/ \\###/ /###\\ \\##/              \\#####/  \\####/ \\###/ /###\\ \\##/         "}


func drawPlayerPieces(grid Grid, offset int, fg, color termbox.Attribute) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			uiX := (x * 5) + 5 + offset
			uiY := (y * 2) + 2

			triple := translatePositionChars(grid[x][y])
			i := -1
			for _, value := range(triple) {
				termbox.SetCell(uiX + i, uiY, value, color, fg)
			}
			
			// termbox.SetCell(x-1, y, triple[0], color, fg)
			// termbox.SetCell(x,   y, triple[1], color, fg)
			// termbox.SetCell(x+1, y, triple[2], color, fg)
			// set at realX,realY
		}
	}
}

func translatePositionChars(n int) []rune {
	val := make([]rune, 3)
	if (n & hit == hit) {
		val[1] = 'X'
	} else {
		val[1] = 'O'
	}
	if (n & vertical == vertical) {
		if (n & topEnd == topEnd) {
			val[0] = '/'
			val[2] = '\\'
			return val
		}
		if (n & middle == middle) {
			val[0] = '|'
			val[2] = '|'
			return val
		}
		val[0] = '\\'
		val[2] = '/'
		return val
	} 
	if (n & middle == middle) {
		return val
	}
	if (n & rightEnd == rightEnd) {
		val[2] = '>'
		return val
	}
	val[0] = '<'
	return val
}


const (
	msgPlaceBoats =    "            Place boats with space + arrowKey...              "
	msgWaitEnemyTurn = "            Enemy's turn.                           "
	msgYourTurn =      "            Your turn.                              "
	airCraftCarrierLength = 5
	battleShipLength = 4
	frigateLength = 3
	subLength = 3
	littleShipLength = 2
)


var height, width = 0, 0
var x, y = 0, 0
var offset = 0

func draw() {
	fg := termbox.ColorDefault
//	bg := termbox.ColorBlack
	blue := termbox.ColorBlue
	red := termbox.ColorRed
	for i, rx := range ui {
		for j, ry := range rx  {
			if j > 44 {
				termbox.SetCell(j, i, ry, red, fg)
			} else {
				termbox.SetCell(j, i, ry, blue, fg)
			}
		}
	}
	for i, r := range msgPlaceBoats {
		termbox.SetCell(i, 26, r, termbox.ColorWhite, fg)
	}
	termbox.Flush()
}

func moveCursor(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		if y > 3 {
			y -= 2
		}
	case termbox.KeyArrowDown:
		if y < 20 {
			y += 2
		}
	case termbox.KeyArrowLeft:
		if x > 5 + offset {
			x -= 4
		}
	case termbox.KeyArrowRight:
		if x < 41 + offset {
			x += 4
		}
	}
	termbox.SetCursor(x,y)
	termbox.Flush()
}

func view() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)
	width, height = termbox.Size() 

	draw()
	x = 5
//	x = 51
	y = 3
//	offset = 46
	termbox.SetCursor(x,y)
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			draw()
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			} else if ev.Key == termbox.KeyTab {
				if offset == 0 {
					offset = 46
					x = x + offset
				} else {
					x = x - offset
					offset = 0
				}
				termbox.SetCursor(x,y)
				termbox.Flush()
				break
			}
			moveCursor(ev)
		}
	}

}

// func (game *Game)getMove() ([]int, bool) {
	
// 	//	shipEndXY := make([]int, 2)
// 	shipStartXY := make([]int, 2)
// 	xY := make([]int, 2)
// 	xY[0], xY[1] = 0, 0
// 	x, y := 51, 3
// 	offset := 41
// 	termbox.SetCursor(x, y) // set cursor at 0,0 of right grid

// 	startSet := false
	
// loop:
// 	for {
// 		ev := termbox.PollEvent();
// 		switch ev.Type {
// 		case termbox.EventResize:
// 			//
// 		case termbox.EventKey:
// 			switch ev.Key {
// 			case termbox.KeyArrowUp:
// 				if y > 3 {
// 					y -= 2
// 					xY[1] -= 1
// 				}
// 			case termbox.KeyArrowDown:
// 				if y < 20 {
// 					y += 2
// 					xY[1] += 1
// 				}
// 			case termbox.KeyArrowLeft:
// 				if x > 5 + offset {
// 					x -= 4
// 					xY[0] -= 1
// 				}
// 			case termbox.KeyArrowRight:
// 				if x < 41 + offset {
// 					x += 4
// 					xY[0] += 1
// 				}
// 			case termbox.KeySpace:
// 				return xY, true
// 			}
// 			termbox.SetCursor(x,y)
// 			termbox.Flush()
			
// 		}
// 		if ev.Key == termbox.KeyCtrlQ {
// 			return xY, false
// 			break loop
// 		} 
// 	}
// 	return xY, false
// }

func (game *Game)setPieceShip(ship int) {
	logger.Printf("Setting ship number: %d", ship)
	//	shipEndXY := make([]int, 2)
	shipStartXY := make([]int, 2)
	xY := make([]int, 2)
	xY[0], xY[1] = 0, 0
	x, y := 51, 3
	offset := 0
	termbox.SetCursor(x, y) // set cursor at 0,0 of right grid
	startSet := false

	x = 5
	y = 3
	draw()
	termbox.SetCursor(x,y)
	termbox.Flush()
loop:
	for {
		ev := termbox.PollEvent();
		switch ev.Type {
		case termbox.EventResize:
			//
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if y > 3 {
					y -= 2
					xY[1] -= 1
				}
			case termbox.KeyArrowDown:
				if y < 20 {
					y += 2
					xY[1] += 1
				}
			case termbox.KeyArrowLeft:
				if x > 5 + offset {
					x -= 4
					xY[0] -= 1
				}
			case termbox.KeyArrowRight:
				if x < 41 + offset {
					x += 4
					xY[0] += 1
				}
			case termbox.KeySpace:
				logger.Println("Set command pushed")
				if startSet {
					logger.Println("Start set")
					if game.openSquare(xY){
						logger.Println("Square is open")
						worked := game.ourBoard.setShip(xY, shipStartXY, ship)
						if worked {
							return
						}
						startSet = false
					}
				} else {
					logger.Println("Start NOT set")
					if game.openSquare(xY) {
						shipStartXY = xY
						startSet = true
					}
				}
			}
			termbox.SetCursor(x,y)
			termbox.Flush()
			
		}
		if ev.Key == termbox.KeyCtrlQ {
			break loop
		}
		draw()
		game.draw()
		termbox.Flush()
		logger.Print(shipStartXY)
		logger.Printf("Position: %d,%d", xY[0], xY[1])
		printGrid(game.ourBoard)
	}

	return
}



func (game Game)playerSetPieces() {
	// setPieceShip(airCraftCarrier)
	// setPieceShip(battleShip)
	// setPieceShip(frigate)
	// setPieceShip(sub)
	// setPieceShip(littleShip)
	
	return
}

func drawEnemyCursor(x, y int) {
	// dont
	termbox.SetCursor(x, y)
}

// is square empty or not
func (game Game)openSquare(xY []int) bool {
	
	if game.ourBoard[xY[0]][xY[1]] == 0 {
		return true
	}
	return false
}

func (game Game)draw() {
	fg := termbox.ColorDefault
	//	bg := termbox.ColorBlack
	blue := termbox.ColorBlue
	red := termbox.ColorRed
	// print generic boards
	for i, rx := range ui {
		for j, ry := range rx  {
			if j > 44 {
				termbox.SetCell(j, i, ry, red, fg)
			} else {
				termbox.SetCell(j, i, ry, blue, fg)
			}
		}
	}
	// print messages at bottom of boards
	for i, r := range msgPlaceBoats {
		termbox.SetCell(i, 26, r, termbox.ColorWhite, fg)
	}
	drawPlayerPieces(game.ourBoard, 0, fg, blue)
	drawPlayerPieces(game.theirBoard, 4, fg, red)
	termbox.Flush()
}

func printGrid(grid Grid) {
	for i := 0; i < len(grid); i++ {
		logger.Println(grid[i])
	}
}

func updateEnemyCursor(msg string) {
	cPos := strings.TrimRight(msg, "\n")
	xy := strings.Split(cPos, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	logger.Printf("Enemy moved to x:%d, y:%d", x, y)
	return
}
