package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
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

type Map map[Coord]rune

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	mapa := make(Map)
	var start_position Coord
	facing := '>'

	y := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			break
		}

		for x := 0; x < len(buff); x++ {
			if buff[x] == '#' || buff[x] == '.' {
				c := Coord{x, y}
				mapa[c] = rune(buff[x])
			}
		}
		y += 1
	}

	fileScanner.Scan()
	instructions := fileScanner.Text()
	minX, minY, maxX, _ := mapa.getBoundingBox()
	for i := minX; i <= maxX; i++ {
		c := Coord{i, minY}
		if v, ok := mapa[c]; ok {
			if v == '.' {
				start_position = c
				break
			}
		}
	}

	mapa.drawMap()
	fmt.Println(instructions)
	fmt.Println(start_position)
	fmt.Println(facing)

	movements := make(Map)
	instruction_pointer := 0
	current_position := start_position
	for instruction_pointer < len(instructions) {
		current_instruction := GetNextInstruction(instructions[instruction_pointer:])
		instruction_pointer += len(current_instruction)

		// fmt.Printf("Current instruction: %s; pos: %v; facing: %c\n", current_instruction, current_position, facing)
		// mapa.drawMapWithMovements(movements)
		// fmt.Println()

		if current_instruction == "R" {
			switch facing {
			case '>':
				facing = 'v'
			case 'v':
				facing = '<'
			case '<':
				facing = '^'
			case '^':
				facing = '>'
			}
		} else if current_instruction == "L" {
			switch facing {
			case '>':
				facing = '^'
			case 'v':
				facing = '>'
			case '<':
				facing = 'v'
			case '^':
				facing = '<'
			}
		} else {
			steps, _ := strconv.Atoi(current_instruction)
			// Let's see the next step
			for i := 0; i < steps; i++ {
				new_c := current_position
				switch facing {
				case '>':
					new_c.X += 1
				case '<':
					new_c.X = new_c.X - 1
				case 'v':
					new_c.Y += 1
				case '^':
					new_c.Y = new_c.Y - 1
				}
				// Now we check for walls
				object, ok := mapa[new_c]
				if ok { // There's something
					if object == '#' { // We hit a wall, exit the loop with the previous position
						break
					} else if object == '.' { // No problem with the path, continue
						current_position = new_c
						movements[current_position] = facing
					}
				} else { // There's no path, we must wrap
					// fmt.Println("There's no path, looking for the wrap")
					current_position, facing = mapa.GetExtreme(new_c, facing)
					// fmt.Printf("Obtained %v\n", current_position)
					movements[current_position] = facing
				}
			}
		}
	}
	//mapa.drawMapWithMovements(movements)
	// fmt.Println(current_position)
	// fmt.Println(string(facing))

	res := 1000*(current_position.Y+1) + 4*(current_position.X+1)
	switch facing {
	case 'v':
		res += 1
	case '<':
		res += 2
	case '^':
		res += 3
	}

	fmt.Println(res)

}

func (mapa Map) getBoundingBox() (int, int, int, int) {
	var minX, minY, maxX, maxY int
	minX = math.MaxInt
	minY = math.MaxInt

	maxX = math.MinInt
	maxY = math.MinInt

	for coord := range mapa {
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	return minX, minY, maxX, maxY

}

func (mapa Map) drawMap() {
	minX, minY, maxX, maxY := mapa.getBoundingBox()

	fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", minX, minY, maxX, maxY)

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			v, ok := mapa[Coord{X: i, Y: j}]
			if ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (mapa Map) drawMapWithMovements(movements Map) {
	minX, minY, maxX, maxY := mapa.getBoundingBox()

	fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", minX, minY, maxX, maxY)

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			if v, ok := movements[Coord{X: i, Y: j}]; ok {
				fmt.Printf("%c", v)
				continue
			}
			v, ok := mapa[Coord{X: i, Y: j}]
			if ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (mapa Map) HasPoint(c Coord) bool {
	_, ok := mapa[c]
	return ok
}

func GetNextInstruction(instructions string) string {
	r := regexp.MustCompile(`^(?P<Instruction>R|L|[\d]+)`)
	ret := r.FindStringSubmatch(instructions)

	return ret[1]
}

func (mapa Map) GetExtreme(pos Coord, face rune) (Coord, rune) {
	minX, minY, maxX, maxY := mapa.getBoundingBox()
	switch face {
	case '>':
		for i := minX; i <= maxX; i++ {
			c := Coord{i, pos.Y}
			// fmt.Printf("Trying %v ", c)
			if v, ok := mapa[c]; ok {
				// fmt.Printf(";found %c ", v)
				if v == '.' { // Path at the other side
					// fmt.Printf("returning %v\n", c)
					return c, face
				} else { // There's a rock at the other side
					return Coord{pos.X - 1, pos.Y}, face
				}
			}
			// fmt.Println()
		}
	case '<':
		for i := maxX; i >= minX; i-- {
			c := Coord{i, pos.Y}
			// fmt.Printf("Trying %v ", c)
			if v, ok := mapa[c]; ok {
				if v == '.' { // Path at the other side
					return c, face
				} else { // There's a rock at the other side
					return Coord{pos.X + 1, pos.Y}, face
				}
			}
		}
	case 'v':
		for i := minY; i <= maxY; i++ {
			c := Coord{pos.X, i}
			// fmt.Printf("Trying %v ", c)
			if v, ok := mapa[c]; ok {
				// fmt.Printf(";found %c ", v)
				if v == '.' { // Path at the other side
					// fmt.Printf("returning %v\n", c)
					return c, face
				} else { // There's a rock at the other side
					// fmt.Printf("There's a rock, returning %v\n", pos)
					return Coord{pos.X, pos.Y - 1}, face
				}
			}
		}
	case '^':
		for i := maxY; i >= minY; i-- {
			c := Coord{pos.X, i}
			if v, ok := mapa[c]; ok {
				if v == '.' { // Path at the other side
					return c, face
				} else { // There's a rock at the other side
					return Coord{pos.X, pos.Y + 1}, face
				}
			}
		}
	}

	// This should not happen
	return Coord{}, face
}
