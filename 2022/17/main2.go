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

type Piece []Coord

type Cave map[Coord]struct{}

func main() {

	pieces := []Piece{
		{{2, 0}, {3, 0}, {4, 0}, {5, 0}},
		{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},
		{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
		{{2, 0}, {2, 1}, {3, 0}, {3, 1}}}

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	buff := fileScanner.Text()
	fmt.Println(buff)

	cave := make(Cave)

	rock_count := 0
	max_height := -1
	time := 0
	super_time := 0
	last_height := 0
	last_rocks := 1
	finish := false
	for {
		new_piece := make(Piece, len(pieces[rock_count%5]))
		copy(new_piece, pieces[rock_count%5])
		rock_count += 1
		new_piece.MoveUp(max_height + 4)
		for {
			// cave.PrintCaveWithFallinPiece(new_piece)
			time = time % len(buff)
			if time == 0 {
				if super_time%5 == 0 {
					if super_time/5 == 4 { // Enough
						finish = true
						break
					}
					cave.PrintCaveWithFallinPiece(new_piece)
					height := cave.GetMaxY()
					fmt.Printf("super_time: %d, piece: %d, num_pieces: %d (diff: %d), height: %d; diff: %d\n", super_time, rock_count%5, rock_count, rock_count-last_rocks, height, height-last_height)
					last_height = height
					last_rocks = rock_count
				}
				super_time += 1
			}
			switch buff[time] {
			case '>':
				new_piece = MoveRight(cave, new_piece)
			case '<':
				new_piece = MoveLeft(cave, new_piece)
			}
			time += 1

			var stopped bool
			new_piece, stopped = MoveDown(cave, new_piece)

			if stopped { // Piece stopped
				cave.AddPiece(new_piece)
				max_height = cave.GetMaxY()
				break
			}

		}
		if rock_count%1000 == 0 {
			fmt.Printf("rock number %d, cave size %d\n", rock_count, len(cave))
		}
		if rock_count == 1000000000000 {
			// if rock_count == 2022 {
			break
		}
		if finish {
			break
		}

	}
	//cave.PrintCave()

	// In this point we have 2 iterations and the deltas for rocks and height. We should calculate how many rocks we can simulate
	height := cave.GetMaxY()
	fmt.Printf("super_time: %d, piece: %d, num_pieces: %d (diff: %d), height: %d; diff: %d\n", super_time, rock_count%5, rock_count, rock_count-last_rocks, height, height-last_height)

	objetive := 1000000000000
	iterations := (objetive - rock_count) / (rock_count - last_rocks)
	fmt.Printf("We must do %d iterations. The result will be -> height: %d, pieces: %d\n", iterations, iterations*(height-last_height), iterations*(rock_count-last_rocks))

	rock_count += iterations*(rock_count-last_rocks) - 1
	// We must put a piece on height
	new_piece := make(Piece, len(pieces[0]))
	copy(new_piece, pieces[0])
	new_piece.MoveUp(height + iterations*(height-last_height))
	cave.PrintCaveWithFallinPiece(new_piece)
	cave.AddPiece(new_piece)
	max_height = cave.GetMaxY()
	time = 0

	for {
		new_piece := make(Piece, len(pieces[rock_count%5]))
		copy(new_piece, pieces[rock_count%5])
		rock_count += 1
		new_piece.MoveUp(max_height + 4)
		// cave.PrintCaveWithFallinPiece(new_piece)
		for {
			// cave.PrintCaveWithFallinPiece(new_piece)
			time = time % len(buff)
			switch buff[time] {
			case '>':
				new_piece = MoveRight(cave, new_piece)
			case '<':
				new_piece = MoveLeft(cave, new_piece)
			}
			time += 1

			var stopped bool
			new_piece, stopped = MoveDown(cave, new_piece)

			if stopped { // Piece stopped
				cave.AddPiece(new_piece)
				max_height = cave.GetMaxY()
				break
			}

		}
		if rock_count%1000 == 0 {
			fmt.Printf("rock number %d, cave size %d\n", rock_count, len(cave))
		}
		if rock_count == 1000000000000 {
			// if rock_count == 2022 {
			break
		}
	}

	height = cave.GetMaxY()
	fmt.Printf("super_time: %d, piece: %d, num_pieces: %d (diff: %d), height: %d; diff: %d\n", super_time, rock_count%5, rock_count, rock_count-last_rocks, height, height-last_height)

	// Not sure why this is one more the input1 result, and one less than the input2 result
	fmt.Println(cave.GetMaxY())
}

func (piece Piece) MoveUp(amount int) {
	for i := 0; i < len(piece); i++ {
		piece[i].Y += amount
	}
}

func MoveDown(cave Cave, piece Piece) (Piece, bool) {
	new_piece := make(Piece, len(piece))
	copy(new_piece, piece)
	for i := 0; i < len(piece); i++ {
		new_piece[i].Y = new_piece[i].Y - 1
		if new_piece[i].Y < 0 { // We hit the bottom
			return piece, true
		}
	}

	for _, p := range new_piece {
		if _, ok := cave[p]; ok { // Hit another piece on the cave
			return piece, true
		}
	}

	return new_piece, false
}

func MoveLeft(cave Cave, piece Piece) Piece {
	new_piece := make(Piece, len(piece))
	copy(new_piece, piece)
	for i := 0; i < len(piece); i++ {
		new_piece[i].X = new_piece[i].X - 1
		if new_piece[i].X < 0 || new_piece[i].X >= 7 { // We hit the borders
			return piece
		}
	}

	for _, p := range new_piece {
		if _, ok := cave[p]; ok { // Hit another piece on the cave
			return piece
		}
	}

	return new_piece
}

func MoveRight(cave Cave, piece Piece) Piece {
	new_piece := make(Piece, len(piece))
	copy(new_piece, piece)
	for i := 0; i < len(piece); i++ {
		new_piece[i].X += 1
		if new_piece[i].X < 0 || new_piece[i].X >= 7 { // We hit the borders
			return piece
		}
	}

	for _, p := range new_piece {
		if _, ok := cave[p]; ok { // Hit another piece on the cave
			return piece
		}
	}

	return new_piece
}

func (cave Cave) PrintCave() {
	const minX = 0
	const maxX = 7
	const minY = 0
	maxY := cave.GetMaxY()

	for y := maxY; y >= 0; y-- {
		fmt.Print("|")
		for x := minX; x < maxX; x++ {
			if _, ok := cave[Coord{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("+-------+\n")
}
func (cave Cave) PrintCaveWithFallinPiece(piece Piece) {
	const minX = 0
	const maxX = 7
	const minY = 0
	maxY := cave.GetMaxY()

	piece_map := make(map[Coord]struct{})
	for _, p := range piece {
		piece_map[p] = struct{}{}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for y := maxY; y >= maxY-50; y-- {
		fmt.Print("|")
		for x := minX; x < maxX; x++ {
			if _, ok := cave[Coord{x, y}]; ok {
				fmt.Print("#")
			} else {
				if _, ok := piece_map[Coord{x, y}]; ok {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("+-------+\n")
}

func (cave Cave) AddPiece(piece Piece) {
	for _, c := range piece {
		// fmt.Printf("Adding %v\n", c)
		cave[c] = struct{}{}
	}
}

func (cave Cave) GetMaxY() int {
	maxY := 0
	for p := range cave {
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxY
}

func Clean(cave Cave) Cave {
	new_cave := make(Cave)

	xfound := make(map[int]struct{})
	maxY := cave.GetMaxY()
	for y := maxY; y > 0; y-- {
		for x := 0; x < 7; x++ {
			if _, ok := xfound[x]; ok {
				continue
			}
			c := Coord{x, y}
			if _, ok := cave[c]; ok {
				xfound[x] = struct{}{}
				new_cave[c] = struct{}{}
			}
		}
	}
	return new_cave
}

func (cave Cave) CheckIfIsFlat() bool {
	maxY := cave.GetMaxY()
	for x := 0; x < 7; x++ {
		c := Coord{x, maxY}
		if _, ok := cave[c]; !ok {
			return false
		}
	}
	return true
}
