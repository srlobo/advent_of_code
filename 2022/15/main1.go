package main

import (
	"bufio"
	"fmt"
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

type SensorBeacon struct {
	sensor   Coord
	beacon   Coord
	distance int
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

	var sensorbeacon []SensorBeacon
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		sensor, beacon := parseLine(buff)
		sensorbeacon = append(sensorbeacon, SensorBeacon{sensor: sensor, beacon: beacon, distance: GetManhattanDistance(sensor, beacon)})
	}

	examineRow := 2000000
	emptyPositions := make(map[Coord]int)

	for _, sensor := range sensorbeacon {
		firstPoint := Coord{X: sensor.sensor.X, Y: examineRow}
		fmt.Printf("Sensor %v, firstPoint: %v\n", sensor, firstPoint)
		if GetManhattanDistance(firstPoint, sensor.sensor) > sensor.distance { // We are outside the zone where no beacon can exist
			continue
		}
		if firstPoint != sensor.beacon {
			emptyPositions[firstPoint] = 1
		}

		var j int
		j = 1
		for {
			tentativePoint := Coord{X: firstPoint.X + j, Y: examineRow}
			// fmt.Printf("tentativePoint: %v ", tentativePoint)
			if GetManhattanDistance(tentativePoint, sensor.sensor) <= sensor.distance {
				// fmt.Println("Inside!")
				if tentativePoint != sensor.beacon {
					emptyPositions[tentativePoint] = 1
				}
			} else {
				// fmt.Println("Outside!")
				break
			}
			j += 1
		}

		j = 1
		for {
			tentativePoint := Coord{X: firstPoint.X - j, Y: examineRow}
			// fmt.Printf("tentativePoint: %v ", tentativePoint)
			if GetManhattanDistance(tentativePoint, sensor.sensor) <= sensor.distance {
				// fmt.Println("Inside!")
				if tentativePoint != sensor.beacon {
					emptyPositions[tentativePoint] = 1
				}
			} else {
				// fmt.Println("Outside!")
				break
			}
			j += 1
		}
	}

	fmt.Println(sensorbeacon)
	// fmt.Println(emptyPositions)
	fmt.Println(len(emptyPositions))
}

func parseLine(buff string) (Coord, Coord) {
	r := regexp.MustCompile(`x=(?P<SX>-?[\d]+), y=(?P<SY>-?[\d]+).*beacon.*x=(?P<BX>-?[\d]+), y=(?P<BY>-?[\d]+)`)
	// var sensor, beacon Coord
	ret := r.FindStringSubmatch(buff)
	SX, _ := strconv.Atoi(ret[1])
	SY, _ := strconv.Atoi(ret[2])
	BX, _ := strconv.Atoi(ret[3])
	BY, _ := strconv.Atoi(ret[4])
	sCoord := Coord{X: SX, Y: SY}
	bCoord := Coord{X: BX, Y: BY}

	return sCoord, bCoord
}

func GetManhattanDistance(a Coord, b Coord) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
