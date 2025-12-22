package main

import (
	"bufio"
	"fmt"
	"os"
)

var empty = struct{}{}

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

	board := Board{minX: 0, minY: 0, elements: make(BoardElements)}
	var x, y int
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		x = 0
		for _, c := range buff {
			if c == '@' {
				board.elements[Coord{x, y}] = '@'
			}
			x++
		}
		board.maxX = x
		y++
	}
	board.maxY = y

	board.drawMap()

	total := 0
	for {
		partial := board.findAndRemoveRolls()
		if partial == 0 {
			break
		}
		total += partial
	}

	fmt.Println(total)

}

func (board *Board) findAndRemoveRolls() int {
	totalAccessibleRolls := make(map[Coord]struct{})

	for coord, _ := range board.elements {
		rollsAround := 0
		var testCoord Coord

		testCoord = Coord{coord.X - 1, coord.Y}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}
		testCoord = Coord{coord.X - 1, coord.Y - 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}
		testCoord = Coord{coord.X - 1, coord.Y + 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}

		testCoord = Coord{coord.X, coord.Y - 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}
		testCoord = Coord{coord.X, coord.Y + 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}

		testCoord = Coord{coord.X + 1, coord.Y}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}
		testCoord = Coord{coord.X + 1, coord.Y - 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}
		testCoord = Coord{coord.X + 1, coord.Y + 1}
		if board.elements[testCoord] == '@' {
			rollsAround++
		}

		if rollsAround < 4 {
			totalAccessibleRolls[coord] = empty
		}
	}

	for coord, _ := range totalAccessibleRolls {
		delete(board.elements, coord)
	}

	return len(totalAccessibleRolls)
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
