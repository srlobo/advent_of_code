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
	for {
		new_piece := make(Piece, len(pieces[rock_count%5]))
		copy(new_piece, pieces[rock_count%5])
		rock_count += 1
		new_piece.MoveUp(max_height + 4)
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

		if rock_count == 2022 {
			// if rock_count == 20 {
			break
		}
	}
	cave.PrintCave()
	fmt.Println(cave.GetMaxY() + 1)
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

	for y := maxY; y >= 0; y-- {
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
