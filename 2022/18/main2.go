package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	X int
	Y int
	Z int
}

type Droplet map[Coord]struct{}

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

	droplet := make(Droplet)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		buff_splitted := strings.Split(buff, ",")
		x, _ := strconv.Atoi(buff_splitted[0])
		y, _ := strconv.Atoi(buff_splitted[1])
		z, _ := strconv.Atoi(buff_splitted[2])

		point := Coord{x, y, z}
		droplet[point] = empty
	}

	n_facets := 0
	for cube := range droplet {
		// For each cube we check the 6 faces
		var c Coord
		c = Coord{cube.X - 1, cube.Y, cube.Z}
		if !droplet.Has(c) {
			n_facets++
		}
		c = Coord{cube.X + 1, cube.Y, cube.Z}
		if !droplet.Has(c) {
			n_facets++
		}
		c = Coord{cube.X, cube.Y - 1, cube.Z}
		if !droplet.Has(c) {
			n_facets++
		}
		c = Coord{cube.X, cube.Y + 1, cube.Z}
		if !droplet.Has(c) {
			n_facets++
		}
		c = Coord{cube.X, cube.Y, cube.Z - 1}
		if !droplet.Has(c) {
			n_facets++
		}
		c = Coord{cube.X, cube.Y, cube.Z + 1}
		if !droplet.Has(c) {
			n_facets++
		}
	}
	fmt.Println(n_facets)

	fmt.Println(droplet.GetBoundingBox())
	maxX, minX, maxY, minY, maxZ, minZ := droplet.GetBoundingBox()

	// Let's search for an air cube
	var c Coord
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				c = Coord{x, y, z}
				if !droplet.Has(c) {
					break
				}

			}
		}
	}

	// And now extract every connected air cube
	visited_air_cubes := make(Droplet)
	tentative_air_cubes := make(Droplet)
	tentative_air_cubes[c] = empty

	for {
		if len(tentative_air_cubes) == 0 {
			break
		}
		for cube := range tentative_air_cubes {
			// For each cube we check the 6 faces
			var c Coord
			c = Coord{cube.X - 1, cube.Y, cube.Z}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			c = Coord{cube.X + 1, cube.Y, cube.Z}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			c = Coord{cube.X, cube.Y - 1, cube.Z}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			c = Coord{cube.X, cube.Y + 1, cube.Z}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			c = Coord{cube.X, cube.Y, cube.Z - 1}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			c = Coord{cube.X, cube.Y, cube.Z + 1}
			if !droplet.Has(c) && !visited_air_cubes.Has(c) && c.Inside(maxX, minX, maxY, minY, maxZ, minZ) {
				tentative_air_cubes[c] = empty
			}
			delete(tentative_air_cubes, cube)
			visited_air_cubes[cube] = empty
		}
	}

	bubbles := make(Droplet)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				c = Coord{x, y, z}
				if !visited_air_cubes.Has(c) && !droplet.Has(c) {
					bubbles[c] = empty
				}

			}
		}
	}

	// Now in bubbles we have the air cubes trapped inside the droplet. Let's count the faces they have not touching other bubble cubes
	for cube := range bubbles {
		// For each cube we check the 6 faces
		var c Coord
		c = Coord{cube.X - 1, cube.Y, cube.Z}
		if !bubbles.Has(c) {
			n_facets--
		}
		c = Coord{cube.X + 1, cube.Y, cube.Z}
		if !bubbles.Has(c) {
			n_facets--
		}
		c = Coord{cube.X, cube.Y - 1, cube.Z}
		if !bubbles.Has(c) {
			n_facets--
		}
		c = Coord{cube.X, cube.Y + 1, cube.Z}
		if !bubbles.Has(c) {
			n_facets--
		}
		c = Coord{cube.X, cube.Y, cube.Z - 1}
		if !bubbles.Has(c) {
			n_facets--
		}
		c = Coord{cube.X, cube.Y, cube.Z + 1}
		if !bubbles.Has(c) {
			n_facets--
		}
	}

	fmt.Println(n_facets)

}

func (d Droplet) Has(c Coord) bool {
	_, ok := d[c]
	return ok
}

func (droplet Droplet) GetBoundingBox() (int, int, int, int, int, int) {
	var maxX, minX, maxY, minY, maxZ, minZ int
	for cube := range droplet {
		if maxX < cube.X {
			maxX = cube.X
		}
		if maxY < cube.Y {
			maxY = cube.Y
		}
		if maxZ < cube.Z {
			maxZ = cube.Z
		}
		if minX > cube.X {
			minX = cube.X
		}
		if minY > cube.Y {
			minY = cube.Y
		}
		if minZ > cube.Z {
			minZ = cube.Z
		}
	}

	return maxX, minX, maxY, minY, maxZ, minZ
}

func (c Coord) Inside(maxX, minX, maxY, minY, maxZ, minZ int) bool {
	return c.X <= maxX && c.X >= minX && c.Y <= maxY && c.Y >= minY && c.Z <= maxZ && c.Z >= minZ
}
