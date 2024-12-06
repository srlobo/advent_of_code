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
	guardInitialPosition := Guard{}

	j := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff)
		for i, r := range buff {
			if r == '#' {
				board.elements[Coord{X: i, Y: j}] = r
			} else if r == '^' {
				guardInitialPosition.coord = Coord{X: i, Y: j}
				guardInitialPosition.direction = '^'
			}
		}
		j++
	}
	board.maxY = j - 1

	guard := guardInitialPosition
	guardWalk := make(map[Coord][]rune)
	for guard.coord.Y >= board.minY &&
		guard.coord.Y <= board.maxY &&
		guard.coord.X >= board.minX &&
		guard.coord.X <= board.maxX {

		guardWalk[guard.coord] = append(guardWalk[guard.coord], guard.direction)
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

	count := 0
	for possibleObstacle := range guardWalk {
		board.elements[possibleObstacle] = '#'
		if board.guardWalkIsLoop(guardInitialPosition) {
			count++
		}
		delete(board.elements, possibleObstacle)
	}

	fmt.Println(count)
}

func (board *Board) guardWalkIsLoop(initialGuardPosition Guard) bool {
	guard := initialGuardPosition
	guardWalk := make(map[Coord][]rune)
	for guard.coord.Y >= board.minY &&
		guard.coord.Y <= board.maxY &&
		guard.coord.X >= board.minX &&
		guard.coord.X <= board.maxX {

		for i := 0; i < len(guardWalk[guard.coord]); i++ {
			if guardWalk[guard.coord][i] == guard.direction {
				// board.drawMapWithCoords(guardWalk)
				return true
			}
		}
		guardWalk[guard.coord] = append(guardWalk[guard.coord], guard.direction)
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
	// board.drawMapWithCoords(guardWalk)
	return false
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

func (board *Board) drawMapWithCoords(coords map[Coord][]rune) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY := board.minY
	maxY := board.maxY

	for j := minY; j <= maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			s, ok2 := coords[c]
			if len(s) == 1 {
				s = append(s, ' ')
			} else if len(s) == 0 {
				s = append(s, ' ')
				s = append(s, ' ')
			}
			if ok && ok2 {
				fmt.Print(string(dot) + string(dot))
			} else if ok && !ok2 {
				fmt.Print(string(string(dot) + string(dot)))
			} else if !ok && ok2 {
				fmt.Print(string(s))
			} else if !ok && !ok2 {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
