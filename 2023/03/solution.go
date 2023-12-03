package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	var board Board
	board.elements = make(BoardElements)
	line_n := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		board.maxX = len(buff)
		for char_n, c := range buff {
			if c != '.' {
				board.elements[Coord{char_n, line_n}] = c
			}
		}
		line_n += 1
	}
	board.maxY = line_n
	board.minY = 0
	board.minX = 0

	// board.drawMap()

	// collect numbers

	number_tmp := make([]rune, 0)
	coords_number_tmp := make([]Coord, 0)

	total := 0
	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				if dot >= '0' && dot <= '9' {
					number_tmp = append(number_tmp, dot)
					coords_number_tmp = append(coords_number_tmp, c)
				} else {
					if len(number_tmp) > 0 {
						number := 0
						for _, n := range number_tmp {
							number = number*10 + int(n-'0')
						}
						coords_around_number := board.coordsAroundNumber(coords_number_tmp)
						fmt.Println("========================================")
						fmt.Println(number)
						board.drawMapWithCoords(coords_around_number)
						if board.checkSymbolInCoords(coords_around_number) {
							total += number
							fmt.Printf("%d -> true\n", number)
						} else {
							fmt.Printf("%d -> false\n", number)
						}
						fmt.Println("========================================")

						number_tmp = make([]rune, 0)
						coords_number_tmp = make([]Coord, 0)
					}
				}
			} else {
				if len(number_tmp) > 0 {
					number := 0
					for _, n := range number_tmp {
						number = number*10 + int(n-'0')
					}
					coords_around_number := board.coordsAroundNumber(coords_number_tmp)
					fmt.Println("========================================")
					fmt.Println(number)
					board.drawMapWithCoords(coords_around_number)
					if board.checkSymbolInCoords(coords_around_number) {
						total += number
						fmt.Printf("%d -> true\n", number)
					} else {
						fmt.Printf("%d -> false\n", number)
					}
					fmt.Println("========================================")

					number_tmp = make([]rune, 0)
					coords_number_tmp = make([]Coord, 0)
				}
			}
		}
		if len(number_tmp) > 0 {
			number := 0
			for _, n := range number_tmp {
				number = number*10 + int(n-'0')
			}
			coords_around_number := board.coordsAroundNumber(coords_number_tmp)
			fmt.Println("========================================")
			fmt.Println(number)
			board.drawMapWithCoords(coords_around_number)
			if board.checkSymbolInCoords(coords_around_number) {
				total += number
				fmt.Printf("%d -> true\n", number)
			} else {
				fmt.Printf("%d -> false\n", number)
			}
			fmt.Println("========================================")

			number_tmp = make([]rune, 0)
			coords_number_tmp = make([]Coord, 0)
		}

	}
	fmt.Println(total)
}

func (board *Board) checkSymbolInCoords(coords map[Coord]struct{}) bool {
	for coord := range coords {
		if c, ok := board.elements[coord]; ok {
			if c < '0' || c > '9' { // c is symbol
				return true
			}
		}
	}
	return false
}

func (board *Board) coordsAroundNumber(number_coords []Coord) map[Coord]struct{} {
	res := make(map[Coord]struct{})
	empty := struct{}{}

	for _, c := range number_coords {
		for j := c.Y - 1; j <= c.Y+1; j++ {
			if j < board.minY || j > board.maxY {
				continue
			}
			for i := c.X - 1; i <= c.X+1; i++ {
				if i < board.minX || i > board.maxX {
					continue
				}
				res[Coord{X: i, Y: j}] = empty
			}
		}
	}

	return res
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
