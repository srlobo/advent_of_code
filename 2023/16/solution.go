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

type Ray struct {
	position  Coord
	direction Coord
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

	board := Board{elements: make(BoardElements), minX: 0, minY: 0}
	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		for i, c := range buff {
			coord := Coord{X: i, Y: j}
			if c != '.' {
				board.elements[coord] = c
			}
			board.maxX = len(buff)
		}
		j++
		board.maxY = j

	}

	fmt.Println("Board size: ", board.maxX, board.maxY)
	rays := make([]Ray, 0)
	initial_ray := Ray{Coord{-1, 0}, Coord{1, 0}}
	rays = append(rays, initial_ray)
	ray_board := make(BoardRayElements)

	for {
		if len(rays) == 0 {
			break
		}
	top:
		for i := 0; i < len(rays); i++ {
			new_coord := Coord{rays[i].position.X + rays[i].direction.X, rays[i].position.Y + rays[i].direction.Y}
			if new_coord.X > board.maxX-1 || new_coord.X < board.minX || new_coord.Y > board.maxY-1 || new_coord.Y < board.minY { // we are out of the board
				rays = append(rays[:i], rays[i+1:]...)
				continue top
			}
			if this_coord_rays, ok := ray_board[new_coord]; ok { // In this coord there's another ray
				for _, direction := range this_coord_rays {
					if direction == rays[i].direction { // We've been already in this coord
						rays = append(rays[:i], rays[i+1:]...)
						continue top
					}
				}
			} else {
				ray_board[new_coord] = make([]Coord, 0)
			}

			ray_board[new_coord] = append(ray_board[new_coord], rays[i].direction)
			rays[i].position = new_coord
			if c, ok := board.elements[new_coord]; ok {
				switch c {
				case '/':
					if rays[i].direction.X == 1 {
						rays[i].direction = Coord{0, -1}
					} else if rays[i].direction.X == -1 {
						rays[i].direction = Coord{0, 1}
					} else if rays[i].direction.Y == 1 {
						rays[i].direction = Coord{-1, 0}
					} else if rays[i].direction.Y == -1 {
						rays[i].direction = Coord{1, 0}
					}

				case '\\':
					if rays[i].direction.X == 1 {
						rays[i].direction = Coord{0, 1}
					} else if rays[i].direction.X == -1 {
						rays[i].direction = Coord{0, -1}
					} else if rays[i].direction.Y == 1 {
						rays[i].direction = Coord{1, 0}
					} else if rays[i].direction.Y == -1 {
						rays[i].direction = Coord{-1, 0}
					}

				case '-':
					if rays[i].direction.X == 0 {
						rays[i].direction = Coord{1, 0}
						new_direction := Coord{-1, 0}
						rays = append(rays, Ray{new_coord, new_direction})
						ray_board[new_coord] = append(ray_board[new_coord], new_direction)
					}

				case '|':
					if rays[i].direction.Y == 0 {
						rays[i].direction = Coord{0, 1}
						new_direction := Coord{0, -1}
						rays = append(rays, Ray{new_coord, new_direction})
						ray_board[new_coord] = append(ray_board[new_coord], new_direction)
					}
				}
			}

		}
	}
	// board.drawMap(ray_board)

	fmt.Println(len(ray_board))
}

type Coord struct {
	X int
	Y int
}

type (
	BoardElements    map[Coord]rune
	BoardRayElements map[Coord][]Coord
)

type Board struct {
	elements BoardElements

	minX int
	minY int
	maxX int
	maxY int
}
type BoardRay struct {
	elements BoardRayElements

	minX int
	minY int
	maxX int
	maxY int
}

func (board *Board) drawMap(ray_board BoardRayElements) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				fmt.Print(string(dot))
			} else {
				if dot, ok := ray_board[c]; ok {
					if len(dot) > 1 {
						fmt.Print(len(dot))
					} else {
						switch dot[0] {
						case Coord{1, 0}:
							fmt.Print(">")
						case Coord{0, -1}:
							fmt.Print("^")
						case Coord{-1, 0}:
							fmt.Print("<")
						case Coord{0, 1}:
							fmt.Print("v")
						}
					}
				} else {
					fmt.Print(" ")
				}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
