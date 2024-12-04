package main

import (
	"bufio"
	"fmt"
	"os"
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

	j := 0
	board := Board{elements: make(BoardElements)}
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff)
		for i := 0; i < len(buff); i++ {
			coord := Coord{X: i, Y: j}
			board.elements[coord] = rune(buff[i])
		}
		j++
	}
	board.minX = 0
	board.minY = 0
	board.maxY = j

	total := 0
	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			coord := Coord{X: i, Y: j}
			if findXmas(board, coord) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func findXmas(board Board, coord Coord) bool {
	newCoord := Coord{X: coord.X, Y: coord.Y}
	if board.elements[newCoord] != 'A' {
		return false
	}

	// up down left right diagonal
	upLeft := Coord{X: coord.X - 1, Y: coord.Y - 1}
	downRight := Coord{X: coord.X + 1, Y: coord.Y + 1}
	if !((board.elements[upLeft] == 'M' && board.elements[downRight] == 'S') ||
		(board.elements[upLeft] == 'S' && board.elements[downRight] == 'M')) {
		return false
	}

	// down up left right diagonal
	upRight := Coord{X: coord.X + 1, Y: coord.Y - 1}
	downLeft := Coord{X: coord.X - 1, Y: coord.Y + 1}
	if !((board.elements[upRight] == 'M' && board.elements[downLeft] == 'S') ||
		(board.elements[upRight] == 'S' && board.elements[downLeft] == 'M')) {
		return false
	}

	return true
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
