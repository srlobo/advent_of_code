package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	X int
	Y int
}

type BoardElements map[Coord][]rune

type Board struct {
	elements BoardElements

	minX int
	minY int
	maxX int
	maxY int
}

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	row := 0
	var board Board
	board.elements = make(BoardElements)
	var initial_position, end_position Coord

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		for x, r := range buff {
			c := Coord{x, row}
			if r != '.' {
				board.elements[c] = []rune{r}
			}
		}
		row += 1
	}
	board.getBoundingBox()
	// initial_board := board.Copy()

	initial_position = Coord{1, 0}
	for i := 0; i < board.maxX; i++ {
		c := Coord{i, board.maxY}
		if _, ok := board.elements[c]; !ok {
			end_position = c
		}
	}
	fmt.Printf("Initial: %v; end: %v\n", initial_position, end_position)
	board.drawMap()

	empty := struct{}{}

	positions := make(map[Coord]struct{})
	positions[initial_position] = empty
	time := 0
	for {
		time += 1
		new_positions := make(map[Coord]struct{})
		board = board.iterateBlizard()

		for position := range positions {
			var c Coord

			c = position
			// Wait
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// UP
			c.Y = c.Y + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Down
			c.Y = c.Y - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Left
			c.X = c.X - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Right
			c.X = c.X + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}
		}

		// fmt.Printf("Time: %d; position: %v; posible possitions: %v\n", time, positions, new_positions)
		// board.drawMapWithPositions(positions)
		// fmt.Println()

		if _, ok := new_positions[end_position]; ok {
			break
		}
		positions = new_positions
	}

	fmt.Println(time)

	tmp := end_position
	end_position = initial_position
	initial_position = tmp
	// board = initial_board.Copy()

	positions = make(map[Coord]struct{})
	positions[initial_position] = empty
	for {
		time += 1
		new_positions := make(map[Coord]struct{})
		board = board.iterateBlizard()

		for position := range positions {
			var c Coord

			c = position
			// Wait
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// UP
			c.Y = c.Y + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Down
			c.Y = c.Y - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Left
			c.X = c.X - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Right
			c.X = c.X + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}
		}

		// fmt.Printf("Time: %d; position: %v; posible possitions: %v\n", time, positions, new_positions)
		// board.drawMapWithPositions(positions)
		// fmt.Println()

		if _, ok := new_positions[end_position]; ok {
			break
		}
		positions = new_positions
	}

	fmt.Println(time)

	tmp = end_position
	end_position = initial_position
	initial_position = tmp
	// board = initial_board.Copy()

	positions = make(map[Coord]struct{})
	positions[initial_position] = empty
	for {
		time += 1
		new_positions := make(map[Coord]struct{})
		board = board.iterateBlizard()

		for position := range positions {
			var c Coord

			c = position
			// Wait
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// UP
			c.Y = c.Y + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Down
			c.Y = c.Y - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Left
			c.X = c.X - 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}

			c = position
			// Right
			c.X = c.X + 1
			if _, ok := board.elements[c]; board.CheckInside(c) && !ok {
				new_positions[c] = empty
			}
		}

		// fmt.Printf("Time: %d; position: %v; posible possitions: %v\n", time, positions, new_positions)
		// board.drawMapWithPositions(positions)
		// fmt.Println()

		if _, ok := new_positions[end_position]; ok {
			break
		}
		positions = new_positions
	}

	fmt.Println(time)

}

func (board *Board) getBoundingBox() {
	if board.maxX != 0 {
		return
	}
	var minX, minY, maxX, maxY int
	for coord := range board.elements {
		// fmt.Printf("Comparing (%d, %d)\n", coord.X, coord.Y)
		if coord.X > maxX {
			maxX = coord.X
		}

		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	board.minX = minX
	board.minY = minY
	board.maxX = maxX
	board.maxY = maxY
}

func (board *Board) drawMap() {
	board.getBoundingBox()

	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			dot, ok := board.elements[Coord{X: i, Y: j}]
			if ok {
				if len(dot) == 1 {
					fmt.Print(string(dot))
				} else {
					fmt.Print(len(dot))
				}

			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (board *Board) drawMapWithPositions(positions map[Coord]struct{}) {
	board.getBoundingBox()

	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)

	for j := board.minY; j <= board.maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			if ok {
				if len(dot) == 1 {
					fmt.Print(string(dot))
				} else {
					fmt.Print(len(dot))
				}
			} else {
				if _, ok := positions[c]; ok {
					fmt.Print("E")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func (board Board) iterateBlizard() Board {
	var new_board Board
	new_board.elements = make(BoardElements)
	for coord, list := range board.elements {
		for _, dot := range list {
			new_coord := Coord{coord.X, coord.Y}
			switch dot {
			case '>':
				new_coord.X += 1
				if new_coord.X > board.maxX-1 {
					new_coord.X = 1
				}
			case '^':
				new_coord.Y = new_coord.Y - 1
				if new_coord.Y == 0 {
					new_coord.Y = board.maxY - 1
				}
			case 'v':
				new_coord.Y += 1
				if new_coord.Y > board.maxY-1 {
					new_coord.Y = 1
				}

			case '<':
				new_coord.X = new_coord.X - 1
				if new_coord.X == 0 {
					new_coord.X = board.maxX - 1
				}
			}

			new_board.elements[new_coord] = append(new_board.elements[new_coord], dot)
		}
	}

	new_board.getBoundingBox()

	return new_board
}

func (board Board) CheckInside(position Coord) bool {
	board.getBoundingBox()
	// fmt.Printf("maxX: %d, minX: %d, maxY: %d, minX: %d\n", board.maxX, board.minX, board.maxY, board.minY)
	// fmt.Printf("Checking position %v ", position)
	if position.X > board.maxX || position.X < board.minX || position.Y > board.maxY || position.Y < board.minY {
		// fmt.Println("false")
		return false
	} else {
		// fmt.Println("true")
		return true
	}
}

func (board Board) Copy() Board {
	new_board := board
	new_board.elements = make(BoardElements)
	for k, v := range board.elements {
		new_board.elements[k] = v
	}
	return new_board
}
