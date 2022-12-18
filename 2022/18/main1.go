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
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
		c = Coord{cube.X + 1, cube.Y, cube.Z}
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
		c = Coord{cube.X, cube.Y - 1, cube.Z}
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
		c = Coord{cube.X, cube.Y + 1, cube.Z}
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
		c = Coord{cube.X, cube.Y, cube.Z - 1}
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
		c = Coord{cube.X, cube.Y, cube.Z + 1}
		if _, ok := droplet[c]; !ok {
			n_facets++
		}
	}
	fmt.Println(n_facets)
}
