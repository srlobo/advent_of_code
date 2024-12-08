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

	board := Board{elements: make(BoardElements)}
	antennas := make(map[rune][]Coord)

	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff) - 1
		for i, r := range buff {
			if r == '.' {
				continue
			}
			coords := antennas[r]
			coords = append(coords, Coord{X: i, Y: j})
			antennas[r] = coords
		}

		j++
	}
	board.maxY = j - 1

	for r, coords := range antennas {
		fmt.Println("rune:", string(r))
		board.combineAntennas(coords)
	}

	for r, coords := range antennas {
		for _, coord := range coords {
			board.elements[coord] = r
		}
	}

	board.drawMap()
	fmt.Println("Antennas:", len(board.elements))
}

func (board *Board) combineAntennas(coords []Coord) {
	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords); j++ {
			if i == j {
				continue
			}
			board.getMirrorCoord(coords[i], coords[j])
		}
	}
}

func (board *Board) getMirrorCoord(a, b Coord) {
	i := 1
	for true {
		newCoord := Coord{X: b.X + i*(b.X-a.X), Y: b.Y + i*(b.Y-a.Y)}
		if board.isOutside(newCoord) {
			return
		}
		board.elements[newCoord] = '#'
		i++
	}
}

func (board *Board) isOutside(a Coord) bool {
	fmt.Print("coord:", a, "outside:", board.minX, board.minY, board.maxX, board.maxY)
	if a.X < board.minX || a.X > board.maxX || a.Y < board.minY || a.Y > board.maxY {
		fmt.Println(" true")
		return true
	}

	fmt.Println(" false")
	return false
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
				fmt.Print(".")
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
