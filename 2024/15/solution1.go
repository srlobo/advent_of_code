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
	var instructions string
	j := 0
	var botPosition Coord

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		if buff == "" {
			break
		}
		board.maxX = len(buff)
		for i := 0; i < len(buff); i++ {
			if buff[i] != '.' {
				coord := Coord{X: i, Y: j}
				if buff[i] == '@' {
					botPosition = coord
				} else {
					board.elements[coord] = rune(buff[i])
				}
			}
		}
		j++
	}
	board.maxY = j
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		instructions = instructions + buff
	}

	for _, nextMovement := range instructions {
		// fmt.Println("Move", string(nextMovement))
		// fmt.Println("Bot position: ", botPosition)
		nextPosition := getNextPos(botPosition, nextMovement)
		// fmt.Println("Next position: ", nextPosition)

		if _, ok := board.elements[nextPosition]; !ok {
			botPosition = nextPosition
		} else if board.elements[nextPosition] == 'O' {
			if board.moveRock(nextPosition, nextMovement) {
				botPosition = nextPosition
			}
		}

		// board.drawMapWithCoords(map[Coord]struct{}{botPosition: empty})
		// fmt.Println("---------------")

	}
	fmt.Println(board.calculateGPSSum())
}

func (board *Board) calculateGPSSum() int {
	ret := 0

	for c, r := range board.elements {
		if r != '#' {
			ret += c.X + (c.Y * 100)
		}
	}
	return ret
}

func (board *Board) moveRock(rockPos Coord, direction rune) bool {
	newRockPos := getNextPos(rockPos, direction)
	if board.elements[newRockPos] == '#' {
		return false
	} else if board.elements[newRockPos] == 'O' {
		if board.moveRock(newRockPos, direction) {
			board.elements[newRockPos] = 'O'
			delete(board.elements, rockPos)
			return true
		} else {
			return false
		}
	} else {
		board.elements[newRockPos] = 'O'
		delete(board.elements, rockPos)
		return true
	}
}

func getNextPos(coord Coord, movement rune) Coord {
	switch movement {
	case '^':
		return Coord{X: coord.X, Y: coord.Y - 1}
	case 'v':
		return Coord{X: coord.X, Y: coord.Y + 1}
	case '<':
		return Coord{X: coord.X - 1, Y: coord.Y}
	case '>':
		return Coord{X: coord.X + 1, Y: coord.Y}
	}
	return coord
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
	minY := board.minY
	maxY := board.maxY

	for j := minY; j < maxY; j++ {
		for i := board.minX; i < board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			_, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(string(dot)))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string("@"))
			} else if !ok && !ok2 {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
