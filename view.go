package main

import (
	"github.com/nsf/termbox-go"
	"strings"
	"strconv"
	"log"
)

var ui = [...]string{"                    YOURS                                         THEIRS                  ",
	"   | A | B | C | D | E | F | G | H | I | J |     | A | B | C | D | E | F | G | H | I | J |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 0 |/#\\|   |   |   |   |   |   |   |   |   |   0 |   |   |   |   |   |   |   |   |   |   |",
	"---+|-|+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 1 ||#||   |   |   |   |   |   |   |   |   |   1 |   |   |   |   |   |   |   |   |   |   |",
	"---+|-|+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 2 |\\#/|   |   |   |   |   |   |   |   |   |   2 |   |   |   |   |   |   |   |   |   |   |",
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

const (

	msg_place_boats = "            Place boats with space + arrowKey...              "
	msg_wait_enemy_turn = "            Enemy's turn.                           "
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
	for i, r := range msg_place_boats {
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

	//	shipEndXY := make([]int, 2)
	shipStartXY := make([]int, 2)
	xY := make([]int, 2)
	xY[0], xY[1] = 0, 0
	x, y := 51, 3
	offset := 41
	termbox.SetCursor(x, y) // set cursor at 0,0 of right grid

	startSet := false
	
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
				if startSet {
					if game.openSquare(xY){
						worked := game.MyGrid.setShip(xY, shipStartXY)
						if worked {
							log.Print("cool")
						}
					}
				} else {
					if game.openSquare(xY) {
						shipStartXY = xY
					}
				}
			}
			termbox.SetCursor(x,y)
			termbox.Flush()
			
		}
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
		log.Print(shipStartXY)
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
	
	if game.MyGrid[xY[0]][xY[1]] == 0 {
		return true
	}
	return false
}

func updateEnemyCursor(msg string) {
	cPos := strings.TrimRight(msg, "\n")
	xy := strings.Split(cPos, ",")
	x, _ := strconv.ParseInt(xy[0], 10, 64)
	y, _ := strconv.ParseInt(xy[1], 10, 64)
	log.Printf("Enemy moved to x:%d, y:%d", x, y)
	return
}
