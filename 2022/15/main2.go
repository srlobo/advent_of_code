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
	fmt.Println(len(sensorbeacon))
	for _, sb := range sensorbeacon {
		fmt.Println(sb)
	}

	limit := 4000000
	// limit = 20

	megamap := make(map[Coord]string)
	//	for y := 0; y <= limit; y++ {
	//		for x := 0; x <= limit; x++ {
	//			for c, sb := range sensorbeacon {
	//				sensor := sb.sensor
	//				point := Coord{X: x, Y: y}
	//				if GetManhattanDistance(point, sensor) <= sb.distance {
	//					//fmt.Printf("Adding %v\n", point)
	//					if point == sensor {
	//						megamap[point] = fmt.Sprintf("%s", "S ")
	//					} else if point == sb.beacon {
	//						megamap[point] = fmt.Sprintf("%s", "B ")
	//					} else {
	//						megamap[point] = fmt.Sprintf("%2d", c)
	//					}
	//				}
	//			}
	//		}
	//	}
	//	drawMap(megamap)

	megamap = make(map[Coord]string)
	for _, sb := range sensorbeacon {
		sensor := sb.sensor
		point := Coord{X: sensor.X + sb.distance + 1, Y: sensor.Y}
		if CheckPoint(point, 0, limit) {
			megamap[point] = " b"
		}
		for i := 0; i <= sb.distance+1; i++ {
			point = Coord{X: sensor.X + sb.distance + 1 - i, Y: sensor.Y + i}
			if CheckPoint(point, 0, limit) {
				megamap[point] = " b"
			}
			point = Coord{X: sensor.X + sb.distance + 1 - i, Y: sensor.Y - i}
			if CheckPoint(point, 0, limit) {
				megamap[point] = " b"
			}
		}
		point = Coord{X: sensor.X - sb.distance - 1, Y: sensor.Y}
		if CheckPoint(point, 0, limit) {
			megamap[point] = "b"
		}
		for i := 0; i <= sb.distance+1; i++ {
			point = Coord{X: sensor.X - sb.distance - 1 + i, Y: sensor.Y + i}
			if CheckPoint(point, 0, limit) {
				megamap[point] = " b"
			}
			point = Coord{X: sensor.X - sb.distance - 1 + i, Y: sensor.Y - i}
			if CheckPoint(point, 0, limit) {
				megamap[point] = " b"
			}
		}

	}
	// drawMap(megamap)
	fmt.Println(len(megamap))

	for point := range megamap {
		found := true
		for _, sb := range sensorbeacon {
			sensor := sb.sensor
			// fmt.Printf("Comparing %v vs %v; d: %d <=? %d; t:%v\n", point, sensor, GetManhattanDistance(point, sensor), sb.distance, found)
			if GetManhattanDistance(point, sensor) <= sb.distance {
				found = false
			}
		}
		if found {
			fmt.Printf("Found, it's %v\n", point)
			fmt.Println(4000000*point.X + point.Y)
			break
		}
	}
}

func drawMap(mapa map[Coord]string) {
	minX, minY, maxX, maxY := 0, 0, 20, 20

	// fmt.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d\n", minX, minY, maxX, maxY)

	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			dot, ok := mapa[Coord{X: i, Y: j}]
			if ok {
				fmt.Printf("%s ", dot)
			} else {
				fmt.Print(".. ")
			}
		}
		fmt.Println()
	}
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

func CheckPoint(p Coord, min int, max int) bool {
	if p.X < min || p.X > max || p.Y < min || p.Y > max {
		return false
	}
	return true
}
