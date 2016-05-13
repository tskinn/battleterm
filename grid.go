package main

// each coordinate is an integer with each bit in the integer
// signifying different attributes of ship or whatever

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	bit1, mask1                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3, mask3                           // bit3 == 8, mask3 == 7
	
	littleShip = 1 << iota
	sub
	frigate
	battleship
	airCraftCarrier
	vertical
	horizontal
	middle
	rightEnd
	topEnd
	hit
	miss
	cursor
)

const (
	dimensions = 9

	// everything below is just used for drawing 
	// ships 
	// littleShip      = 1   //     0000:0000:0001
	// sub             = 2   //     0000:0000:0010
	// frigate         = 4   //     0000:0000:0100
	// battleShip      = 8   //     0000:0000:1000
	// airCraftCarrier = 16  //     0000:0001:0000
	// // position    
	// vertical        = 32  //     0000:0010:0000
	// horizontal      = 64  //     0000:0100:0000
	// middle          = 128 //     0000:1000:0000
	// rightEnd        = 256 //     0001:0000:0000
	// topEnd          = 512 //     0010:0000:0000
	// // state    
	// hit             = 1024//     0100:0000:0000
	// miss            = 2048//     1000:0000:0000
	// cursor          = 4096//0001:0000:0000:0000
	antiCursor      = ^cursor    // all ones except 4096
	up =    "up"
	down =  "down"
	left =  "left"
	right = "right"
)
// type Grid [][]int

func toGrid (x, y int) (int, int) {
	return x + 41, y + 3
}

func (grid *Grid) setPoint(x, y, shipType, pos, axis int) bool {
	//edge := 9
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
	case battleship:// battleship
		return battleShipLength
	case airCraftCarrier:// airCraftCarrier
		return airCraftCarrierLength
	}
	return 2
}

// func (grid *Grid)setShip(xy, start []int) bool {
// 	return false
// }

func (grid *Grid)setShip(xyEnd, xyStart []int, ship int) bool {
	length := getShipLen(ship)
	direction := getDirection(xyEnd, xyStart)
	var pos int

	if direction == up || direction == down { // y increases downward
		if direction == up { // so everything is going down 
			y = y - length
		} 
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
		if direction == left {// so everything is going right
			x = x - length
		} 
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

func getDirection(start, end []int) string {
	if start[0] < end[0] {
		return "right"
	} else if start[0] > end[0] {
		return "left"
	} else if start[1] > end[1] {
		return "down"
	}
	return "up"
}
