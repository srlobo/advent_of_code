package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type Map [][]string

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var mapa Map

	var end Coord
	var startPositions []Coord

	j := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		var line []string
		for i := 0; i < len(buff); i++ {
			line = append(line, string(buff[i]))
			switch string(buff[i]) {
			case "S":
				startPositions = append(startPositions, Coord{X: i, Y: j})
			case "a":
				startPositions = append(startPositions, Coord{X: i, Y: j})
			case "E":
				end = Coord{X: i, Y: j}
			}
		}
		mapa = append(mapa, line)
		j += 1
	}

	fmt.Printf("Number of startPositions: %d", len(startPositions))
	var steps []int
	for _, start := range startPositions {
		unvisited_set := make(map[Coord]int)
		visited_set := make(map[Coord]int)
		for j := 0; j < len(mapa); j++ {
			for i := 0; i < len(mapa[j]); i++ {
				unvisited_set[Coord{X: i, Y: j}] = math.MaxInt32
			}
		}

		unvisited_set[start] = 0

		fmt.Println(start)
		fmt.Println(end)

		current_node := start
		continueWithNextCandidate := false
		for {
			current_node_distance := unvisited_set[current_node]
			fmt.Printf("Current distance: %d\n", current_node_distance)
			possibleNeighbourgs := mapa.getUnvisitedNeighbours(current_node, unvisited_set)
			fmt.Print("possibleNeighbourgs ")
			fmt.Println(possibleNeighbourgs)

			for _, node := range possibleNeighbourgs {
				if unvisited_set[node] > current_node_distance+1 {
					unvisited_set[node] = current_node_distance + 1
				}
			}

			delete(unvisited_set, current_node)
			visited_set[current_node] = current_node_distance
			if _, ok := visited_set[end]; ok {
				break
			}

			next_node, ok := getNextNode(unvisited_set)
			if !ok {
				// printVisitedSet(visited_set, mapa)
				// panic("No next node found")
				fmt.Println("No next node found")
				continueWithNextCandidate = true
				break
			}

			fmt.Printf("Jumping to (%d, %d)\n", next_node.X, next_node.Y)

			current_node = next_node
			fmt.Printf("Visited_nodes: %d, unvisited_nodes: %d\n", len(visited_set), len(unvisited_set))
			fmt.Println()
		}
		if !continueWithNextCandidate {
			fmt.Println(visited_set[end])
			steps = append(steps, visited_set[end])

			fmt.Println(end)
		} else {
			fmt.Println("You are on an unconnected island")
		}
	}

	sort.Ints(steps)
	fmt.Print(steps)

}

func (mapa Map) getPossibleNeighbours(node Coord) []Coord {
	current_node_value := mapa[node.Y][node.X]
	if current_node_value == "S" {
		current_node_value = "a"
	} else if current_node_value == "E" {
		current_node_value = "z"
	}

	var ret []Coord

	for j := node.Y - 1; j <= node.Y+1; j++ {
		if j < 0 || j >= len(mapa) || (Coord{X: node.X, Y: j} == node) {
			continue
		}
		fmt.Printf("Testing (%d, %d) (max: %d, %d); value: %s, current_node_value: %s, diff: %d", node.X, j, len(mapa[0]), len(mapa), mapa[j][node.X], current_node_value, Diff(mapa[j][node.X][0], current_node_value[0]))
		if Diff(mapa[j][node.X][0], current_node_value[0]) <= 1 {
			ret = append(ret, Coord{X: node.X, Y: j})
			fmt.Print(", appending")
		}
		fmt.Println()
	}

	for i := node.X - 1; i <= node.X+1; i++ {
		if i < 0 || i >= len(mapa[0]) || (Coord{X: i, Y: node.Y} == node) {
			continue
		}
		fmt.Printf("Testing (%d, %d) (max: %d, %d); value: %s, current_node_value: %s, diff: %d", i, node.Y, len(mapa[0]), len(mapa), mapa[node.Y][i], current_node_value, Diff(mapa[node.Y][i][0], current_node_value[0]))
		if Diff(mapa[node.Y][i][0], current_node_value[0]) <= 1 {
			ret = append(ret, Coord{X: i, Y: node.Y})
			fmt.Print(", appending")
		}
		fmt.Println()
	}

	return ret
}

func (mapa Map) getUnvisitedNeighbours(node Coord, unvisitedSet map[Coord]int) []Coord {
	var ret []Coord
	possibleNeighbourgs := mapa.getPossibleNeighbours(node)
	for i := 0; i < len(possibleNeighbourgs); i++ {
		neighbour := possibleNeighbourgs[i]
		if _, ok := unvisitedSet[neighbour]; ok {
			ret = append(ret, neighbour)
		}
	}
	return ret
}

func Abs(a byte) int {
	if a < 0 {
		return int(-a)
	}
	return int(a)
}

func Diff(a byte, b byte) int {
	if a == 'E' {
		a = 'z'
	}
	if b == 'E' {
		b = 'z'
	}

	return int(a) - int(b)
}

func getNextNode(unvisited_set map[Coord]int) (Coord, bool) {
	min := math.MaxInt32
	ok := false
	var ret Coord
	for node, distance := range unvisited_set {
		if distance < math.MaxInt32 {
			//fmt.Printf("getNextNode, trying node (%d, %d), distance: %d\n", node.X, node.Y, distance)
		}
		if distance < min {
			min = distance
			ok = true
			ret = node
		}
	}
	return ret, ok
}

func printVisitedSet(visited_set map[Coord]int, mapa Map) {
	for j := 0; j < len(mapa); j++ {
		for i := 0; i < len(mapa[j]); i++ {
			node := Coord{X: i, Y: j}
			if distance, ok := visited_set[node]; ok {
				fmt.Printf("%3d ", distance)
			} else {
				fmt.Print("... ")
			}
		}
		fmt.Println()
	}
}
