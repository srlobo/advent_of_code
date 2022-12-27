package main

import (
	"bufio"
	"fmt"
	"math"
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

type Map map[Coord]struct{}

func main() {
	empty := struct{}{}

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	mapa := make(Map)

	y := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		for x := 0; x < len(buff); x++ {
			if buff[x] == '#' {
				c := Coord{x, y}
				mapa[c] = empty
			}
		}
		y += 1
	}
	fmt.Println()
	fmt.Println(mapa)
	fmt.Println(len(mapa))
	mapa.drawMap()

	first_direction := 0
	for round := 1; round <= 10; round++ {
		new_map := make(Map)
		proposed_directions := make(map[Coord]Coord)
		// First half
		for c := range mapa {
			fmt.Printf("Processing %v\n", c)
			if !mapa.HasSomethingAround(c) {
				fmt.Printf("element: %v, stay\n", c)
				proposed_directions[c] = c
			} else {

				fmt.Printf("We start from %d condition\n", first_direction)
				for i := 0; i < 4; i++ {
					direction := (first_direction + i) % 4
					moved := false
					switch direction {
					case 0:
						look1 := Coord{c.X, c.Y - 1}
						look2 := Coord{c.X + 1, c.Y - 1}
						look3 := Coord{c.X - 1, c.Y - 1}
						if !mapa.HasPoint(look1) && !mapa.HasPoint(look2) && !mapa.HasPoint(look3) {
							// Move north
							fmt.Printf("element: %v, move north\n", c)
							proposed_directions[c] = Coord{c.X, c.Y - 1}
							moved = true
						}
					case 1:
						look1 := Coord{c.X, c.Y + 1}
						look2 := Coord{c.X + 1, c.Y + 1}
						look3 := Coord{c.X - 1, c.Y + 1}
						if !mapa.HasPoint(look1) && !mapa.HasPoint(look2) && !mapa.HasPoint(look3) {
							// Move south
							fmt.Printf("element: %v, move south\n", c)
							proposed_directions[c] = Coord{c.X, c.Y + 1}
							moved = true
						}

					case 2:
						look1 := Coord{c.X - 1, c.Y - 1}
						look2 := Coord{c.X - 1, c.Y}
						look3 := Coord{c.X - 1, c.Y + 1}
						if !mapa.HasPoint(look1) && !mapa.HasPoint(look2) && !mapa.HasPoint(look3) {
							// Move West
							fmt.Printf("element: %v, move west\n", c)
							proposed_directions[c] = Coord{c.X - 1, c.Y}
							moved = true
						}

					case 3:
						look1 := Coord{c.X + 1, c.Y - 1}
						look2 := Coord{c.X + 1, c.Y}
						look3 := Coord{c.X + 1, c.Y + 1}
						if !mapa.HasPoint(look1) && !mapa.HasPoint(look2) && !mapa.HasPoint(look3) {
							// Move east
							fmt.Printf("element: %v, move east\n", c)
							proposed_directions[c] = Coord{c.X + 1, c.Y}
							moved = true
						}
					}
					if moved {
						break
					}
				}
				if _, ok := proposed_directions[c]; !ok {
					// Crowded area, stay
					fmt.Printf("element: %v, stay as it's crowded\n", c)
					proposed_directions[c] = c
				}

			}
		}
		first_direction = (first_direction + 1) % 4
		fmt.Println("Proposed:", proposed_directions)
		// Second half
		proposed_directions_inverse := make(map[Coord]Coord)
		for old, new := range proposed_directions {
			if _, ok := proposed_directions_inverse[new]; ok { // There's another elf to move there
				proposed_directions[old] = old
				other_elf_proposing_move := proposed_directions_inverse[new]
				proposed_directions[other_elf_proposing_move] = other_elf_proposing_move
			} else {
				proposed_directions_inverse[new] = old
			}
		}
		fmt.Println("Proposed directions", proposed_directions)
		fmt.Println("Proposed directions inverse", proposed_directions_inverse)

		for _, new := range proposed_directions {
			new_map[new] = empty
		}
		fmt.Println("Map", new_map)
		fmt.Println(len(new_map))

		fmt.Printf("Round: %d\n", round)
		new_map.drawMap()
		mapa = new_map
		fmt.Println()
	}
	fmt.Println(mapa.HowManyEmptyTiles())
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

	fmt.Print("    ")
	for i := minX; i <= maxX; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	for j := minY; j <= maxY; j++ {
		fmt.Printf("%2d: ", j)
		for i := minX; i <= maxX; i++ {
			_, ok := mapa[Coord{X: i, Y: j}]
			if ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (mapa Map) HasPoint(c Coord) bool {
	_, ok := mapa[c]
	return ok
}

func (mapa Map) HasSomethingAround(c Coord) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			l := Coord{c.X + i, c.Y + j}
			if mapa.HasPoint(l) {
				return true
			}
		}
	}
	return false
}

func (mapa Map) HowManyEmptyTiles() int {
	minX, minY, maxX, maxY := mapa.getBoundingBox()

	size := (maxX - minX + 1) * (maxY - minY + 1)
	fmt.Printf("maxX - minX: %d\n", maxX-minX+1)
	fmt.Printf("maxY - minY: %d\n", maxY-minY+1)

	return size - len(mapa)
}
