package main

import (
	"github.com/nsf/termbox-go"
)

var ui = [...]string{"                    YOURS                                         THEIRS                  ",
	"   | A | B | C | D | E | F | G | H | I | J |     | A | B | C | D | E | F | G | H | I | J |",
	"---+---+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 0 |/#\|   |   |   |   |   |   |   |   |   |   0 |   |   |   |   |   |   |   |   |   |   |",
	"---+|-|+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 1 ||#||   |   |   |   |   |   |   |   |   |   1 |   |   |   |   |   |   |   |   |   |   |",
	"---+|-|+---+---+---+---+---+---+---+---+---+  ---+---+---+---+---+---+---+---+---+---+---+",
	" 2 |\#/|   |   |   |   |   |   |   |   |   |   2 |   |   |   |   |   |   |   |   |   |   |",
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
	airCraftCarrier = 5
	battleShip = 4
	frigate = 3
	sub = 3
	littleShip = 2

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

func (game *Game)setShip(ship int) {

//	shipEndXY := make([]int, 2)
	shipStartXY := make([]int 2)
	xY := make([]int, 2)
	xY[0], xY[1] := 0, 0
	x, y := 51, 3
	offset := 41
	termbox.SetCursor(x, y) // set cursor at 0,0 of right grid

	startSet := false
	
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			//
		case termbox.EventKey:
			case ev.Key {
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
					if openSquare(game.MyGrid, xY){
						game.MyGrid, worked := setShip(game.MyGrid, xY, shipStartXY)
						
					}
				} else {
					if openSquare(game.MyGrid, xY) {
						shipStartXY = xY
					}
				}
			}
			termbox.SetCursor(x,y)
			termbox.Flush()
			
			}
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}
			else if ev.Key == termbox.KeyTab {
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
		}		
	}
	return
}

func playerSetPieces() {
	setShip(airCraftCarrier)
	setShip(battleShip)
	setShip(frigate)
	setShip(sub)
	setShip(littleShip)
	
	return
}

func drawEnemyCursor(x, y int) {
	// dont
	termbox.SetCursor(x, y)
}
