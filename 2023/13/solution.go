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

	j := 0
	board := Board{elements: make(BoardElements), minX: 0, minY: 0}
	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		if buff == "" {
			fmt.Println("Line end")
			board.maxY += 1
			fmt.Println(board.elements)
			total += processBoardElements(board)
			board = Board{elements: make(BoardElements), minX: 0, minY: 0}
			j = 0
			continue
		}
		for i := 0; i < len(buff); i++ {
			if buff[i] == '#' {
				c := Coord{X: i * 2, Y: j * 2}
				board.elements[c] = rune(buff[i])
			}
		}
		board.maxX = (len(buff) * 2) - 1
		board.maxY = j * 2
		j++

	}
	board.maxY += 1
	fmt.Println("EOF")
	total += processBoardElements(board)

	fmt.Println(total)
}

func processBoardElements(board Board) int {
	total := 0
	var mirror_line_x, mirror_line_y int
mirrorx:
	for mirror_line_x = 1; mirror_line_x < board.maxX; mirror_line_x += 2 {
		fmt.Println("board.maxX:", board.maxX, "board.maxY:", board.maxY, "mirror_line_x:", mirror_line_x)
		for j := 0; j <= board.maxY; j += 2 {
			for i := 0; i <= board.maxX; i += 2 {
				c := Coord{X: i, Y: j}
				// fmt.Println("Testing c:", c)
				if _, ok := board.elements[c]; ok {
					mirror_c := mirrorCoordX(c, mirror_line_x)
					// fmt.Println("c:", c, "mirror_c:", mirror_c, "mirror_line_x:", mirror_line_x)
					if mirror_c.X > board.maxX || mirror_c.X < 0 {
						continue
					}
					if _, ok := board.elements[mirror_c]; !ok {
						// fmt.Println("Mirror no encontrado")
						continue mirrorx
					}

				}
			}
		}
		if mirror_line_x < board.maxX {
			if total > 0 {
				fmt.Println("Double")
			}
			fmt.Println("We have found a mirror line", mirror_line_x)
			board.drawMapWithMirrorX(mirror_line_x)
			total += (mirror_line_x + 1) / 2
			fmt.Println(total)
		}
		break
	}

mirrory:
	for mirror_line_y = 1; mirror_line_y < board.maxY; mirror_line_y += 2 {
		fmt.Println("board.maxX:", board.maxX, "board.maxY:", board.maxY, "mirror_line_y:", mirror_line_y)
		for i := 0; i < board.maxX; i += 2 {
			for j := 0; j < board.maxY; j += 2 {
				c := Coord{X: i, Y: j}
				fmt.Println("Testing c:", c)
				if _, ok := board.elements[c]; ok {
					mirror_c := mirrorCoordY(c, mirror_line_y)
					fmt.Println("c:", c, "mirror_c:", mirror_c, "mirror_line_y:", mirror_line_y)
					if mirror_c.Y > board.maxY || mirror_c.Y < 0 {
						continue
					}
					if _, ok := board.elements[mirror_c]; !ok {
						fmt.Println("Mirror no encontrado")
						continue mirrory
					}

				}
			}
		}
		if mirror_line_y < board.maxY && total == 0 {
			if total > 0 {
				fmt.Println("Double")
			}

			fmt.Println("We have found a mirror line", mirror_line_y)
			board.drawMapWithMirrorY(mirror_line_y)
			total += 100 * ((mirror_line_y + 1) / 2)
			fmt.Println(total)
		}
		break
	}

	if total == 0 {
		board.drawMap()
		panic("No mirror found")
	}
	return total
}

func mirrorCoordX(coord Coord, mirror_line_x int) Coord {
	return Coord{X: mirror_line_x - (coord.X - mirror_line_x), Y: coord.Y}
}

func mirrorCoordY(coord Coord, mirror_line_y int) Coord {
	return Coord{X: coord.X, Y: mirror_line_y - (coord.Y - mirror_line_y)}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func (board *Board) drawMapWithMirrorY(mirror int) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	fmt.Print("   ")
	for i := board.minX; i <= board.maxX; i++ {
		if i%2 == 0 {
			fmt.Printf("%1d ", (i/2+1)%10)
		}
	}
	fmt.Println()

	for j := board.minY; j <= board.maxY; j++ {
		if j%2 == 0 {
			fmt.Printf("%2d ", j/2+1)
		} else {
			fmt.Print("   ")
		}
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			if c.Y == mirror {
				fmt.Print("-")
				continue
			}
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

func (board *Board) drawMapWithMirrorX(mirror int) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	fmt.Print("   ")
	for i := board.minX; i <= board.maxX; i++ {
		if i%2 == 0 {
			fmt.Printf("%1d ", (i/2+1)%10)
		}
	}

	fmt.Println()
	for j := board.minY; j <= board.maxY; j++ {
		if j%2 == 0 {
			fmt.Printf("%2d ", j/2+1)
		} else {
			fmt.Print("   ")
		}

		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			if c.X == mirror {
				fmt.Print("|")
				continue
			}
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
