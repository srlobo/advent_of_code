package main

import (
	"bufio"
	"fmt"
	"os"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		defer readFile.Close()
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	board := Board{elements: make(BoardElements), minX: 0, minY: 0}
	j := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		board.maxX = len(buff)
		for i, c := range buff {
			if c != '.' {
				coord := Coord{X: i, Y: j}
				board.elements[coord] = c
			}
		}
		j++
		board.maxY = j
	}
	// Code for extracting the cycle count: 9 is similar to 2
	var init, i int
	board_hashes := make(map[string]int)
	board_hashes[board.strMap()] = 0
	for i = 1; i < 100000; i++ {
		board.cycleRocks(1)
		board_str := board.strMap()
		var ok bool
		if init, ok = board_hashes[board_str]; ok {
			fmt.Println("Cycle found, i:", i, "init:", init)
			break
		} else {
			board_hashes[board_str] = i
		}
	}

	total_number := 1000000000
	cycle := i - init
	total_mod := (total_number - init) % cycle
	fmt.Println(total_mod)

	// board.drawMap()
	board.cycleRocks(total_mod)
	fmt.Println()
	// Calculate the weights
	total := 0
	for j := 0; j < board.maxY; j++ {
		for i := 0; i < board.maxX; i++ {
			coord := Coord{X: i, Y: j}
			if c, ok := board.elements[coord]; ok {
				if c == 'O' {
					total += board.maxY - j
				}
			}
		}
	}
	fmt.Println(total)
}

func (board *Board) cycleRocks(how_much int) {
	for i := 0; i < how_much; i++ {
		board.moveRocksNorth()
		board.moveRocksWest()
		board.moveRocksSouth()
		board.moveRocksEast()
	}
}

func (board *Board) moveRocksNorth() {
	for j := 0; j < board.maxY; j++ {
		for i := 0; i < board.maxX; i++ {
			orig_coord := Coord{X: i, Y: j}
			if board.elements[orig_coord] == 'O' {
				previous_dest_coord := orig_coord
				var j2 int
				for j2 = j - 1; j2 >= board.minY; j2-- {
					dest_coord := Coord{X: i, Y: j2}
					// fmt.Println("Rock at ", orig_coord, "Trying ", dest_coord)
					if _, ok := board.elements[dest_coord]; !ok {
						// fmt.Println(dest_coord, "empty")
						previous_dest_coord = dest_coord
					} else {
						// fmt.Println(dest_coord, "full, moving rock to ", previous_dest_coord)
						delete(board.elements, orig_coord)
						board.elements[previous_dest_coord] = 'O'
						break
					}
				}
				if j2 == board.minY-1 {
					// fmt.Println("limit reached, moving rock to ", previous_dest_coord)
					delete(board.elements, orig_coord)
					board.elements[previous_dest_coord] = 'O'
				}
			}
		}
	}
}

func (board *Board) moveRocksSouth() {
	for j := board.maxY - 1; j >= board.minY; j-- {
		for i := 0; i < board.maxX; i++ {
			orig_coord := Coord{X: i, Y: j}
			if board.elements[orig_coord] == 'O' {
				previous_dest_coord := orig_coord
				var j2 int
				for j2 = j + 1; j2 < board.maxY; j2++ {
					dest_coord := Coord{X: i, Y: j2}
					// fmt.Println("Rock at ", orig_coord, "Trying ", dest_coord)
					if _, ok := board.elements[dest_coord]; !ok {
						// fmt.Println(dest_coord, "empty")
						previous_dest_coord = dest_coord
					} else {
						// fmt.Println(dest_coord, "full, moving rock to ", previous_dest_coord)
						delete(board.elements, orig_coord)
						board.elements[previous_dest_coord] = 'O'
						break
					}
				}
				if j2 == board.maxY {
					// fmt.Println("limit reached, moving rock to ", previous_dest_coord)
					delete(board.elements, orig_coord)
					board.elements[previous_dest_coord] = 'O'
				}
			}
		}
	}
}

func (board *Board) moveRocksEast() {
	for i := board.maxX - 1; i >= board.minX; i-- {
		for j := 0; j < board.maxY; j++ {
			orig_coord := Coord{X: i, Y: j}
			if board.elements[orig_coord] == 'O' {
				previous_dest_coord := orig_coord
				var i2 int
				for i2 = i + 1; i2 < board.maxX; i2++ {
					dest_coord := Coord{X: i2, Y: j}
					// fmt.Println("Rock at ", orig_coord, "Trying ", dest_coord)
					if _, ok := board.elements[dest_coord]; !ok {
						// fmt.Println(dest_coord, "empty")
						previous_dest_coord = dest_coord
					} else {
						// fmt.Println(dest_coord, "full, moving rock to ", previous_dest_coord)
						delete(board.elements, orig_coord)
						board.elements[previous_dest_coord] = 'O'
						break
					}
				}
				if i2 == board.maxX {
					// fmt.Println("limit reached, moving rock to ", previous_dest_coord)
					delete(board.elements, orig_coord)
					board.elements[previous_dest_coord] = 'O'
				}
			}
		}
	}
}

func (board *Board) moveRocksWest() {
	for i := 0; i < board.maxX; i++ {
		for j := 0; j < board.maxY; j++ {
			orig_coord := Coord{X: i, Y: j}
			if board.elements[orig_coord] == 'O' {
				previous_dest_coord := orig_coord
				var i2 int
				for i2 = i - 1; i2 >= board.minY; i2-- {
					dest_coord := Coord{X: i2, Y: j}
					// fmt.Println("Rock at ", orig_coord, "Trying ", dest_coord)
					if _, ok := board.elements[dest_coord]; !ok {
						// fmt.Println(dest_coord, "empty")
						previous_dest_coord = dest_coord
					} else {
						// fmt.Println(dest_coord, "full, moving rock to ", previous_dest_coord)
						delete(board.elements, orig_coord)
						board.elements[previous_dest_coord] = 'O'
						break
					}
				}
				if i2 == board.minY-1 {
					// fmt.Println("limit reached, moving rock to ", previous_dest_coord)
					delete(board.elements, orig_coord)
					board.elements[previous_dest_coord] = 'O'
				}
			}
		}
	}
}

type Coord struct {
	X int
	Y int
}

type BoardElements map[Coord]rune

type Board struct {
	elements BoardElements

	minX int
	minY int
	maxX int
	maxY int
}

func (board *Board) strMap() string {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	ret := ""

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				ret += string(dot)
			} else {
				ret += " "
			}
		}
		ret += "\n"
	}

	return ret
}

func (board *Board) drawMap() {
	fmt.Print(board.strMap())
}

func (board *Board) drawMapWithCoords(coords map[Coord]struct{}) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY := MaxInt
	maxY := 0

	for c := range coords {
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	for j := minY; j <= maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			_, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(dot))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string("."))
			} else if !ok && !ok2 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
