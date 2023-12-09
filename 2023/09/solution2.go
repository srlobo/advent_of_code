package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	inputs := make([][]int, 0)
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		line := make([]int, 0)
		for _, c := range strings.Split(buff, " ") {
			c_int, _ := strconv.Atoi(c)
			line = append(line, c_int)
		}
		slices.Reverse(line)
		inputs = append(inputs, line)
	}

	fmt.Println(inputs)

	total := 0
	for _, input := range inputs {
		differences := make([][]int, 0)
		this_line := input
		differences = append(differences, this_line)
		for {
			this_line = calculateDifferences(this_line)
			differences = append(differences, this_line)
			if checkDifferencesIsAllZero(this_line) {
				break
			}
		}
		fmt.Println(differences)
		for i := len(differences) - 2; i >= 0; i-- {
			new_number := differences[i][len(differences[i])-1] + differences[i+1][len(differences[i+1])-1]
			differences[i] = append(differences[i], new_number)
		}
		fmt.Println(differences)
		total += differences[0][len(differences[0])-1]
	}
	fmt.Println(total)
}

func checkDifferencesIsAllZero(differences []int) bool {
	for _, d := range differences {
		if d != 0 {
			return false
		}
	}
	return true
}

func calculateDifferences(input []int) []int {
	differences := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		difference := input[i+1] - input[i]
		differences = append(differences, difference)
	}
	return differences
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
