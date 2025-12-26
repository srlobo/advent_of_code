package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

	var coords []Coord3D
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		// fmt.Println(buff)
		arrBuff := strings.Split(buff, ",")
		x, _ := strconv.Atoi(arrBuff[0])
		y, _ := strconv.Atoi(arrBuff[1])
		z, _ := strconv.Atoi(arrBuff[2])
		coords = append(coords, Coord3D{x, y, z})
	}

	var distances []Coord3DDistance
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			d := distance(coords[i], coords[j])

			distances = append(distances, Coord3DDistance{coords[i], coords[j], d})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].d < distances[j].d
	})

	coordsLabels := make(map[Coord3D]int)
	nextLabel := 1
	count := 0
	for _, d := range distances {
		if count > 1000 {
			break
		}
		// fmt.Println("Round ", count)

		var label int
		if coordsLabels[d.A] != 0 && coordsLabels[d.B] != 0 {
			// fmt.Println("Already labeled: ", coordsLabels[d.A], coordsLabels[d.B], "Joining all in ", coordsLabels[d.A])
			destLabel := coordsLabels[d.A]
			srcLabel := coordsLabels[d.B]
			for coord, l := range coordsLabels {
				if l == srcLabel {
					coordsLabels[coord] = destLabel
				}
			}
		} else if coordsLabels[d.A] != 0 {
			label = coordsLabels[d.A]
			coordsLabels[d.B] = label
		} else if coordsLabels[d.B] != 0 {
			label = coordsLabels[d.B]
			coordsLabels[d.A] = label
		} else {
			label = nextLabel
			nextLabel++
			coordsLabels[d.A] = label
			coordsLabels[d.B] = label
		}
		// fmt.Println("Processing ", d, "label", label)
		count++
		// fmt.Println(coordsLabels)
		// fmt.Println(count)

		if count == 1000 {
			labelsCoords := make(map[int][]Coord3D)
			var sizeLists []int

			for coord, label := range coordsLabels {
				labelsCoords[label] = append(labelsCoords[label], coord)
			}

			for label, coords := range labelsCoords {
				fmt.Println(label, len(coords), coords)
				sizeLists = append(sizeLists, len(coords))
			}

			sort.Slice(sizeLists, func(i, j int) bool {
				return sizeLists[i] > sizeLists[j]
			})

			fmt.Println(sizeLists)
			t := 0
			for i := 0; i < len(sizeLists); i++ {
				t += sizeLists[i]
			}
			fmt.Println("Total: ", t)

			if len(sizeLists) > 3 {
				fmt.Println(sizeLists[0] * sizeLists[1] * sizeLists[2])
			} else {
				fmt.Println(0)
			}
		}

	}

}

type Coord3DDistance struct {
	A Coord3D
	B Coord3D
	d float64
}

type Coord3D struct {
	X int
	Y int
	Z int
}

func (a Coord3D) module() float64 {
	return math.Sqrt(float64((a.X * a.X) + (a.Y * a.Y) + (a.Z * a.Z)))
}

func distance(a, b Coord3D) float64 {
	dist := Coord3D{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
	return dist.module()
}
