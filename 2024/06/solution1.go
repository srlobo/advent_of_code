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
	guard := Guard{}
	guardWalk := make(map[Coord]struct{})

	j := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff)
		for i, r := range buff {
			if r == '#' {
				board.elements[Coord{X: i, Y: j}] = r
			} else if r == '^' {
				guard.coord = Coord{X: i, Y: j}
				guard.direction = '^'
			}
		}
		j++
	}
	board.maxY = j - 1

	// board.drawMapWithCoords(guardWalk)

	for guard.coord.Y >= board.minY &&
		guard.coord.Y <= board.maxY &&
		guard.coord.X >= board.minX &&
		guard.coord.X <= board.maxX {

		guardWalk[guard.coord] = empty
		switch guard.direction {
		case '^':
			if _, ok := board.elements[Coord{X: guard.coord.X, Y: guard.coord.Y - 1}]; ok {
				guard.direction = '>'
			} else {
				guard.coord.Y--
			}
		case '>':
			if _, ok := board.elements[Coord{X: guard.coord.X + 1, Y: guard.coord.Y}]; ok {
				guard.direction = 'v'
			} else {
				guard.coord.X++
			}
		case '<':
			if _, ok := board.elements[Coord{X: guard.coord.X - 1, Y: guard.coord.Y}]; ok {
				guard.direction = '^'
			} else {
				guard.coord.X--
			}

		case 'v':
			if _, ok := board.elements[Coord{X: guard.coord.X, Y: guard.coord.Y + 1}]; ok {
				guard.direction = '<'
			} else {
				guard.coord.Y++
			}
		}
	}

	board.drawMapWithCoords(guardWalk)
	fmt.Println(len(guardWalk))
}

type Guard struct {
	coord     Coord
	direction rune
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
	minY = board.minY
	maxY = board.maxY

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
