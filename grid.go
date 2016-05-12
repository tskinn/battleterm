package main

// each coordinate is an integer with each bit in the integer
// signifying different attributes of ship or whatever
const (
	dimensions = 9

	// everything below is just used for drawing 
	// ships 
	littleShip      = 1   // 0000:0000:0001
	sub             = 2   // 0000:0000:0010
	frigate         = 4   // 0000:0000:0100
	battleShip      = 8   // 0000:0000:1000
	airCraftCarrier = 16  // 0000:0001:0000
	// position
	vertical        = 32  // 0000:0010:0000
	horizontal      = 64  // 0000:0100:0000
	middle          = 128 // 0000:1000:0000
	rightEnd        = 256 // 0001:0000:0000
	topEnd          = 512 // 0010:0000:0000
	// state
	hit             = 1024// 0100:0000:0000
	miss            = 2048// 1000:0000:0000
)
type Grid [][]int

func toGrid (x, y int) (int, int) {
	return x + 41, y + 3
}

func (grid *Grid) setPoint(x, y, shipType, pos, axis int) bool {
	edge := 9
	if (*grid)[y][x] == 0 {
		(*grid)[y][x] = shipType | pos | axis
		return true
	}
	return false
}

func getShipLen(ship int) int {
	switch ship {
	case littleShip:// little ship
		return littleShipLength
	case sub:// sub
		return subLength
	case frigate:// frigate
		return frigateLength
	case battleShip:// battleship
		return battleShipLength
	case airCraftCarrier:// airCraftCarrier
		return airCraftCarrierLength
	}
	return 2
}

func (grid *Grid)setShip(x, y, ship int, direction string) bool {
	length := getShipLen(ship)
	var pos int
	vert := false
	if direction == "up" || direction == "down" {
		vert = true
	}

	if direction == "up" || direction == "down" {
		// y increases downward
		if direction == "up" {y = y - length} // so everything is going down 
		end := y + length
		for i := y; i <= end; i++ {
			if i == y {
				pos = topEnd
			} else if i == end {
				pos = 0
			} else {
				pos = middle
			}
			if !grid.setPoint(x, i, ship, pos, vertical) {
				return false
			}
		}
	} else {
		if direction == "left" {x = x - length} // so everything is going right
		end := x + length
		for i := x; i <= end; i++ {
			if i == x {
				pos = 0
			} else if i == end {
				pos = rightEnd
			} else {
				pos = middle
			}
			if !grid.setPoint(i, y, ship, pos, horizontal) {
				return false
			}
		}
	}
	
	return true // it worked
}
