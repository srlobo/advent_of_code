package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MinUint = 0
	MinInt  = -MaxInt - 1
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

	board := Board{elements: make(BoardElements)}
	current := Coord{X: 0, Y: 0}
	board.elements[current] = '#'
	for fileScanner.Scan() {
		buff := strings.Split(fileScanner.Text(), " ")
		hexa_number_str := buff[2][2 : len(buff[2])-1]
		direction := hexa_number_str[len(hexa_number_str)-1]
		amount_, ok := strconv.ParseUint("0x"+hexa_number_str[0:len(hexa_number_str)-1], 0, 64)
		if ok != nil { // Not a number
			fmt.Println(ok)
		}
		amount := int(amount_)
		fmt.Println("hexa_number_str: ", "0x"+hexa_number_str[0:len(hexa_number_str)-1], "direction: ", direction, " amount: ", amount)
		var new_coord Coord
		for i := 0; i <= amount; i++ {
			switch direction {
			case '3':
				new_coord = Coord{X: current.X, Y: current.Y - i}
			case '1':
				new_coord = Coord{X: current.X, Y: current.Y + i}
			case '0':
				new_coord = Coord{X: current.X + i, Y: current.Y}
			case '2':
				new_coord = Coord{X: current.X - i, Y: current.Y}
			}
			board.elements[new_coord] = '#'
		}
		current = new_coord
	}
	board.calculateBoundingBox()
	// Let's find a point inside of the path
	var c Coord
first:
	for i := 0; i < board.maxX; i++ {
		for j := 0; j < board.maxY; j++ {
			c = Coord{X: i, Y: j}
			if _, ok := board.elements[c]; ok {
				continue
			}
			count := 0
			for l := c.X; l <= board.maxX; l++ {
				counting_pos := Coord{X: l, Y: c.Y}
				if _, ok := board.elements[counting_pos]; ok {
					count++
				}
			}
			// fmt.Println("Trying c: ", c, " count: ", count)
			if count == 1 {
				break first
			}
		}
	}
	// Now in c we have a point inside of the path
	for tile := range growTileSet(c, board.elements, board) {
		board.elements[tile] = '.'
	}
	board.drawMap()
	fmt.Println(len(board.elements))
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

func (board *Board) drawMap() {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				fmt.Print(string(dot))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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

func (board *Board) calculateBoundingBox() {
	board.minX = MaxInt
	board.minY = MaxInt
	board.maxX = MinInt
	board.maxY = MinInt

	for c := range board.elements {
		if c.X < board.minX {
			board.minX = c.X
		}
		if c.X > board.maxX {
			board.maxX = c.X
		}
		if c.Y < board.minY {
			board.minY = c.Y
		}
		if c.Y > board.maxY {
			board.maxY = c.Y
		}
	}
}

func growTileSet(tile Coord, path_tiles BoardElements, board Board) map[Coord]struct{} {
	ret := make(map[Coord]struct{})

	tiles_to_process := make(map[Coord]struct{}, 0)
	tiles_to_process[tile] = empty

	for {
		if len(tiles_to_process) == 0 {
			break
		}
		var t Coord
		for t = range tiles_to_process { // Choose one random element from tiles_to_process
			break
		}
		delete(tiles_to_process, t)

		ret[t] = empty
		for _, tile := range getAroundTiles(t) {
			if _, ok := ret[tile]; ok { // Already processed
				continue
			}
			if _, ok := path_tiles[tile]; ok { // Reached a path side
				continue
			}
			if tile.X > board.maxX || tile.X < board.minX || tile.Y > board.maxY || tile.Y < board.minY { // Reached a board side
				continue
			}
			tiles_to_process[tile] = empty

		}
	}
	return ret
}

func getAroundTiles(tile Coord) []Coord {
	ret := make([]Coord, 0)
	ret = append(ret, Coord{X: tile.X - 1, Y: tile.Y})
	ret = append(ret, Coord{X: tile.X + 1, Y: tile.Y})
	ret = append(ret, Coord{X: tile.X, Y: tile.Y - 1})
	ret = append(ret, Coord{X: tile.X, Y: tile.Y + 1})
	return ret
}
