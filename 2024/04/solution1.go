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
			total = total + findXmas(board, coord)
		}
	}
	fmt.Println(total)
}

func findXmas(board Board, coord Coord) int {
	newCoord := Coord{X: coord.X, Y: coord.Y}
	if board.elements[newCoord] != 'X' {
		return 0
	}
	total := 0
	// Horizontal:
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X + 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X + 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X + 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}
	// Horizontal - reverse:
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X - 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X - 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X - 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}
	// Vertical
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.Y = coord.Y - 1
	if board.elements[newCoord] == 'M' {
		newCoord.Y = newCoord.Y - 1
		if board.elements[newCoord] == 'A' {
			newCoord.Y = newCoord.Y - 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	// Vertical - reverse
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.Y = coord.Y + 1
	if board.elements[newCoord] == 'M' {
		newCoord.Y = newCoord.Y + 1
		if board.elements[newCoord] == 'A' {
			newCoord.Y = newCoord.Y + 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	// Diagonal: down-right
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X + 1
	newCoord.Y = coord.Y + 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X + 1
		newCoord.Y = newCoord.Y + 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X + 1
			newCoord.Y = newCoord.Y + 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	// Diagonal: down-left
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X - 1
	newCoord.Y = coord.Y + 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X - 1
		newCoord.Y = newCoord.Y + 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X - 1
			newCoord.Y = newCoord.Y + 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	// Diagonal: up-right
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X + 1
	newCoord.Y = coord.Y - 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X + 1
		newCoord.Y = newCoord.Y - 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X + 1
			newCoord.Y = newCoord.Y - 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	// Diagonal: up-left
	newCoord = Coord{X: coord.X, Y: coord.Y}
	newCoord.X = coord.X - 1
	newCoord.Y = coord.Y - 1
	if board.elements[newCoord] == 'M' {
		newCoord.X = newCoord.X - 1
		newCoord.Y = newCoord.Y - 1
		if board.elements[newCoord] == 'A' {
			newCoord.X = newCoord.X - 1
			newCoord.Y = newCoord.Y - 1
			if board.elements[newCoord] == 'S' {
				total = total + 1
			}
		}
	}

	return total
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
