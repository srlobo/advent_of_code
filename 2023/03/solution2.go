package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	board.drawMap()

	// collect numbers

	total := 0
	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok && dot == '*' {
				fmt.Println("found")
				coords_around := board.numbersAroundSymbol(c)
				fmt.Println(coords_around)
				if len(coords_around) == 2 {
					total += coords_around[0] * coords_around[1]
				}
			}
		}
	}
	fmt.Println(total)
}

func (board *Board) numbersAroundSymbol(symbol_coord Coord) []int {
	res := make([]int, 0)

	possible_coords := make([]Coord, 0)
	empty := struct{}{}
	done_coords := map[Coord]struct{}{}

	for j := symbol_coord.Y - 1; j <= symbol_coord.Y+1; j++ {
		for i := symbol_coord.X - 1; i <= symbol_coord.X+1; i++ {
			if j < board.minY || j > board.maxY || i < board.minX || i > board.maxX {
				continue
			}
			c := Coord{X: i, Y: j}
			if char, ok := board.elements[c]; ok {
				if char >= '0' && char <= '9' {
					possible_coords = append(possible_coords, c)
				}
			}
		}
	}
	// Let's grow those possible numbers
	for _, c := range possible_coords {
		fmt.Printf("testing c: %o; done_coords: %o;\n", c, done_coords)
		this_number_coords := make([]Coord, 0)
		if _, ok := done_coords[c]; ok {
			continue
		}
		done_coords[c] = empty
		this_number_coords = append(this_number_coords, c)
		testing_c := Coord{X: c.X - 1, Y: c.Y}
		for { // Let's go to the left
			if _, ok := done_coords[testing_c]; ok {
				break
			}
			if char, ok := board.elements[testing_c]; ok {
				if char >= '0' && char <= '9' {
					done_coords[testing_c] = empty
					this_number_coords = append(this_number_coords, testing_c)
					testing_c = Coord{X: testing_c.X - 1, Y: testing_c.Y}
				} else {
					break
				}
			} else {
				break
			}
		}
		testing_c = Coord{X: c.X + 1, Y: c.Y}
		for { // Let's go to the right
			if _, ok := done_coords[testing_c]; ok {
				break
			}
			if char, ok := board.elements[testing_c]; ok {
				if char >= '0' && char <= '9' {
					done_coords[testing_c] = empty
					this_number_coords = append(this_number_coords, testing_c)
					testing_c = Coord{X: testing_c.X + 1, Y: testing_c.Y}
				} else {
					break
				}
			} else {
				break
			}
		}
		sort.Slice(this_number_coords, func(i, j int) bool {
			return this_number_coords[i].X < this_number_coords[j].X
		})
		number := 0
		for _, c := range this_number_coords {
			number_part, _ := board.elements[c]
			number = number*10 + int(number_part-'0')
		}
		res = append(res, number)

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
