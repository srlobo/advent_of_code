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

	var board Board
	board.elements = make(BoardElements)
	board.minX = 0
	board.minY = 0
	var start_pos Coord
	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		// fmt.Println(buff)
		board.maxX = len(buff)
		for i, c := range buff {
			if c == '.' {
				continue
			}
			coord := Coord{X: i, Y: j}
			board.elements[coord] = c
			if c == 'S' {
				start_pos = coord
			}
		}
		j++
		board.maxY = j
	}

	// fmt.Println()
	// board.drawMap()
	// fmt.Println(start_pos)

	path_tiles := make(map[Coord]Coord)
	path_tiles[start_pos] = start_pos
	var actual_pos, prev_pos Coord
	prev_pos = start_pos
	actual_pos = board.getInitialPipePair(start_pos)[0]

	for {
		if actual_pos == start_pos {
			break
		}

		new_actual_pos := board.getNextStep(actual_pos, prev_pos)

		path_tiles[actual_pos] = prev_pos
		prev_pos = actual_pos
		actual_pos = new_actual_pos
	}

	// fmt.Println(path_tiles)

	inside_out_tiles := make(map[Coord]rune)
	for i := 0; i < board.maxX; i++ {
		for j := 0; j < board.maxY; j++ {
			coord := Coord{X: i, Y: j}
			if _, ok := path_tiles[coord]; ok {
				continue
			} else if _, ok := inside_out_tiles[coord]; ok {
				continue
			} else {
				if inside_out, ok := inferInsideOutTile(coord, path_tiles, board); ok {
					for k := range growTileSet(coord, path_tiles, board) {
						inside_out_tiles[k] = inside_out
					}
				}
			}
		}
	}

	// Now in reverse
	prev_pos = start_pos
	actual_pos = board.getInitialPipePair(start_pos)[1]

	for {
		if actual_pos == start_pos {
			break
		}

		new_actual_pos := board.getNextStep(actual_pos, prev_pos)

		path_tiles[actual_pos] = prev_pos
		prev_pos = actual_pos
		actual_pos = new_actual_pos
	}

	// fmt.Println(path_tiles)

	for i := 0; i < board.maxX; i++ {
		for j := 0; j < board.maxY; j++ {
			coord := Coord{X: i, Y: j}
			if _, ok := path_tiles[coord]; ok {
				continue
			} else if _, ok := inside_out_tiles[coord]; ok {
				continue
			} else {
				if inside_out, ok := inferInsideOutTile(coord, path_tiles, board); ok {
					for k := range growTileSet(coord, path_tiles, board) {
						if inside_out == 'I' {
							inside_out_tiles[k] = 'O'
						} else if inside_out == 'O' {
							inside_out_tiles[k] = 'I'
						}
					}
				}
			}
		}
	}

	for i := 0; i < board.maxX; i++ {
		for j := 0; j < board.maxY; j++ {
			coord := Coord{X: i, Y: j}
			if _, ok := path_tiles[coord]; ok {
				continue
			} else if _, ok := inside_out_tiles[coord]; ok {
				continue
			} else {
				fmt.Println(coord)
				inside_out_tiles[coord] = '?'
			}
		}
	}
	board.drawMapWithCoords(inside_out_tiles)
	total1 := 0
	total2 := 0
	for _, v := range inside_out_tiles {
		if v == 'I' {
			total1++
		} else if v == 'O' {
			total2++
		}
	}
	fmt.Println("Total tiles on board", (board.maxX+1)*(board.maxY+1))
	fmt.Println("Total path_tiles", len(path_tiles))
	fmt.Println("Total inside", total1)
	fmt.Println("Total outside", total2)
	fmt.Println("diff:", (board.maxX+1)*(board.maxY+1)-len(path_tiles)-total1-total2)
}

func inferInsideOutTile(tile Coord, path_tiles map[Coord]Coord, board Board) (rune, bool) {
	// Return the runes I or O if it's inside or outside the path. The bool part returns true if the function has been successful with the infer, or false if it doesn't know.
	debug := false
	if tile.Y == 27 {
		if tile.X == 19 {
			debug = true
		}
	}
	for _, t := range getAroundTiles(tile) {
		//		if t.X < board.minX || t.X > board.maxX || t.Y < board.minY || t.Y > board.maxY { // If we are in the border, we are outside
		//			return 'O', true
		//		}
		if prev, ok := path_tiles[t]; ok {
			// Let's calculate the vector between prev and t:
			pipe_vector := Coord{X: t.X - prev.X, Y: t.Y - prev.Y} // pipe_vector will go (0,1) or (0, -1) or (1,0) or (-1, 0)
			// And now the vector between t and our tile
			calc_vector := Coord{X: tile.X - t.X, Y: tile.Y - t.Y}
			if debug {
				fmt.Println("Checking tile: ", tile, "path_tile: ", t, "prev: ", prev, "; pipe_vector", pipe_vector, "calc_vector", calc_vector)
			}

			if pipe_vector.Y == 1 && calc_vector.X == 1 {
				return 'O', true
			} else if pipe_vector.Y == 1 && calc_vector.X == -1 {
				return 'I', true
			} else if pipe_vector.Y == -1 && calc_vector.X == 1 {
				return 'I', true
			} else if pipe_vector.Y == -1 && calc_vector.X == -1 {
				return 'O', true
			} else if pipe_vector.X == 1 && calc_vector.Y == 1 {
				return 'I', true
			} else if pipe_vector.X == 1 && calc_vector.Y == -1 {
				return 'O', true
			} else if pipe_vector.X == -1 && calc_vector.Y == 1 {
				return 'O', true
			} else if pipe_vector.X == -1 && calc_vector.Y == -1 {
				return 'I', true
			}
		}
	}
	if debug {
		fmt.Println("not found")
	}

	return ' ', false
}

func growTileSet(tile Coord, path_tiles map[Coord]Coord, board Board) map[Coord]struct{} {
	ret := make(map[Coord]struct{})

	tiles_to_process := make(map[Coord]struct{}, 0)
	tiles_to_process[tile] = empty

	for {
		if len(tiles_to_process) == 0 {
			break
		}
		var t Coord
		for t = range tiles_to_process { // Choose one random element from tiles_to_process
			break
		}
		delete(tiles_to_process, t)

		ret[t] = empty
		for _, tile := range getAroundTiles(t) {
			if _, ok := ret[tile]; ok { // Already processed
				continue
			}
			if _, ok := path_tiles[tile]; ok { // Reached a path side
				continue
			}
			if tile.X > board.maxX || tile.X < board.minX || tile.Y > board.maxY || tile.Y < board.minY { // Reached a board side
				continue
			}
			tiles_to_process[tile] = empty

		}
	}
	return ret
}

func getAroundTiles(tile Coord) []Coord {
	ret := make([]Coord, 0)
	ret = append(ret, Coord{X: tile.X - 1, Y: tile.Y})
	ret = append(ret, Coord{X: tile.X + 1, Y: tile.Y})
	ret = append(ret, Coord{X: tile.X, Y: tile.Y - 1})
	ret = append(ret, Coord{X: tile.X, Y: tile.Y + 1})
	return ret
}

func (board *Board) getInitialPipePair(initial_pos Coord) [2]Coord {
	ret := make([]Coord, 0)

	var coord Coord
	var c rune
	var ok bool

	coord = Coord{X: initial_pos.X, Y: initial_pos.Y - 1}
	c, ok = board.elements[coord]
	if ok && (c == '|' || c == '7' || c == 'F') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X, Y: initial_pos.Y + 1}
	c, ok = board.elements[coord]
	if ok && (c == '|' || c == 'J' || c == 'L') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X + 1, Y: initial_pos.Y}
	c, ok = board.elements[coord]
	if ok && (c == '-' || c == '7' || c == 'J') {
		ret = append(ret, coord)
	}

	coord = Coord{X: initial_pos.X - 1, Y: initial_pos.Y}
	c, ok = board.elements[coord]
	if ok && (c == '-' || c == 'L' || c == 'F') {
		ret = append(ret, coord)
	}

	if len(ret) != 2 {
		panic("Invalid initial pipe pair")
	}

	return [2]Coord{ret[0], ret[1]}
}

func (board *Board) getNextStep(actual_pos, prev_pos Coord) Coord {
	var new_pos Coord
	actual_pipe := board.elements[actual_pos]
	switch actual_pipe {
	case '|':
		if actual_pos.Y-1 == prev_pos.Y {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		}
	case '-':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case 'L':
		if actual_pos.X+1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		} else {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		}
	case 'J':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y - 1}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case '7':
		if actual_pos.X-1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X - 1, Y: actual_pos.Y}
		}
	case 'F':
		if actual_pos.X+1 == prev_pos.X {
			new_pos = Coord{X: actual_pos.X, Y: actual_pos.Y + 1}
		} else {
			new_pos = Coord{X: actual_pos.X + 1, Y: actual_pos.Y}
		}
	default:
		panic("Unknown pipe")
	}
	return new_pos
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

func (board *Board) drawMapWithCoords(coords map[Coord]rune) {
	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", board.minX, board.minY, board.maxX, board.maxY)
	minY := 0
	maxY := board.maxY

	for j := minY; j <= maxY; j++ {
		for i := board.minX; i <= board.maxX; i++ {
			c := Coord{X: i, Y: j}
			dot, ok := board.elements[c]
			coord_payload, ok2 := coords[c]
			if ok && ok2 {
				fmt.Print(string(coord_payload))
			} else if ok && !ok2 {
				fmt.Print(string(dot))
			} else if !ok && ok2 {
				fmt.Print(string(coord_payload))
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
