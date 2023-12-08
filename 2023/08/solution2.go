package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

type Node struct {
	left  string
	right string
}

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

	fileScanner.Scan()
	instructions := fileScanner.Text()
	nodes := make(map[string]Node)

	fileScanner.Scan()
	for fileScanner.Scan() {
		buff := strings.Split(fileScanner.Text(), " = ")
		lr := strings.Split(buff[1], ", ")
		n := Node{left: lr[0][1:], right: lr[1][:len(lr[1])-1]}
		nodes[buff[0]] = n
	}
	fmt.Println(instructions)
	fmt.Println(nodes)
	counts := make(map[string]int, 0)

	for k := range nodes {
		if k[len(k)-1] == 'A' {
			counts[k] = traverseMap(k, instructions, nodes)
		}
	}

	fmt.Println(counts)
	c := 1
	for _, v := range counts {
		c = LCM(c, v)
	}
	fmt.Println(c)
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

func traverseMap(origin string, instructions string, nodes map[string]Node) int {
	actual_pos := origin

	count := 0
	lr_pointer := 0
	for {
		fmt.Println("actual_pos: ", actual_pos, "actual instruction: ", string(instructions[lr_pointer]))

		if actual_pos[len(actual_pos)-1] == 'Z' {
			break
		}

		if instructions[lr_pointer] == 'R' {
			actual_pos = nodes[actual_pos].right
		} else if instructions[lr_pointer] == 'L' {
			actual_pos = nodes[actual_pos].left
		}

		if lr_pointer == len(instructions)-1 {
			lr_pointer = 0
		} else {
			lr_pointer++
		}
		count += 1
	}

	return count
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
