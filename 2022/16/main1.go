package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Valve struct {
	name             string
	flow             int
	next_valves      []*Valve
	next_valve_names []string
}

type pathEstimation struct {
	path       Path
	estimation int
}

type pathEstimations []*pathEstimation

type Path []string

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	valves := make(map[string]*Valve)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		valve := parseLine(buff)
		valves[valve.name] = &valve
	}

	var valves_with_flow []*Valve
	for _, v := range valves {
		if v.flow > 0 {
			valves_with_flow = append(valves_with_flow, v)
			fmt.Printf("%d -> %s\n", v.flow, v.name)
		}
		for _, subvalve := range v.next_valve_names {
			v.next_valves = append(v.next_valves, valves[subvalve])
		}
	}
	sort.SliceStable(valves_with_flow, func(i, j int) bool {
		return valves_with_flow[i].flow > valves_with_flow[j].flow
	})

	path := Path{"AA", "DD", "BB", "JJ", "HH", "EE", "CC"}
	fmt.Printf("Winner is %v, %d\n", path, GetPressureFromPath(path, valves))

	path_estimations := pathEstimations{{[]string{"AA"}, 0}}
	max_path_estimations := 40

	for {
		fmt.Println("Considering paths:")
		original_path_estimations := make(pathEstimations, len(path_estimations))
		copy(original_path_estimations, path_estimations)
		path_estimations.print()
		for _, pe := range path_estimations {
			base_path := pe.path
			for _, v := range valves_with_flow {
				if IsIn(base_path, v.name) { // The component is in path
					continue
				}
				path := append(base_path, v.name)
				estimation := GetPressureFromPath(path, valves)
				//fmt.Printf("%v -> %d\n", path, estimation)
				if !path_estimations.Contains(path) {
					path_estimations = append(path_estimations, &pathEstimation{path, estimation})
				}
			}
		}
		fmt.Println("Once sorted:")
		path_estimations.sort()
		path_estimations.print()

		len_path_estimations := len(path_estimations)
		if len_path_estimations > max_path_estimations {
			len_path_estimations = max_path_estimations
		}
		path_estimations = path_estimations[:len_path_estimations]
		if original_path_estimations.IsEqual(path_estimations) {
			break
		}
		fmt.Println("================================")
	}
	path_estimations.print()
}

func IsIn(set []string, el string) bool {
	for _, item := range set {
		if el == item {
			return true
		}
	}
	return false
}

func parseLine(buff string) Valve {
	r := regexp.MustCompile(`Valve (?P<valve_name>[A-Z]{2}) has flow rate=(?P<flow>-?[\d]+); tunnels? leads? to valves? (?P<valves>.*)`)
	ret := r.FindStringSubmatch(buff)
	valve_name := ret[1]
	flow_rate, _ := strconv.Atoi(ret[2])
	next_valves := strings.Split(ret[3], ", ")

	valve := Valve{name: valve_name, flow: flow_rate, next_valve_names: next_valves}
	return valve
}

func GetShortestPath(valves map[string]*Valve, start_node *Valve, target_node *Valve) []string {
	// Init distances
	distances := make(map[string]int)

	for k := range valves {
		distances[k] = math.MaxInt32
	}
	complete_nodes := make(map[string]int)
	distances[start_node.name] = 0

	parents := make(map[string][]string)

	current_node := start_node
	parents[start_node.name] = append(parents[start_node.name], current_node.name)
	for {
		if current_node == target_node {
			break
		}
		current_node_d := distances[current_node.name]
		for _, v := range current_node.next_valves {
			if _, ok := complete_nodes[v.name]; ok { // We have already processed this node
				continue
			}
			if distances[v.name] > current_node_d+1 {
				distances[v.name] = current_node_d + 1
				parents[v.name] = append(parents[current_node.name], v.name)
			}
		}

		complete_nodes[current_node.name] = current_node_d
		delete(distances, current_node.name)

		min := math.MaxInt32
		var min_node *Valve
		for k, d := range distances {
			// fmt.Printf("Checking %s", k)
			if d < min {
				// fmt.Print(" bingo")
				min = d
				min_node = valves[k]
			}
			// fmt.Println()
		}

		current_node = min_node
	}
	return parents[target_node.name]
}

func GetPressureFromPath(path []string, valves map[string]*Valve) int {
	total_presure := 0
	current_presure := 0
	total_time := 0
	for i := 1; i < len(path); i++ {
		current_node := valves[path[i-1]]
		next_node := valves[path[i]]
		sub_path := GetShortestPath(valves, current_node, next_node)
		//fmt.Printf("Calculating %v -> %v; %v\n", current_node, next_node, sub_path)
		for t := 0; t < len(sub_path); t++ {
			total_presure += current_presure
			total_time += 1
			// fmt.Printf("Total time: %d; current_presure: %d; total_presure: %d\n", total_time, current_presure, total_presure)
			if total_time == 30 {
				return total_presure
			}
		}
		current_presure += next_node.flow
	}

	for {
		total_time += 1
		total_presure += current_presure
		// fmt.Printf("Total time: %d; current_presure: %d; total_presure: %d\n", total_time, current_presure, total_presure)
		if total_time == 30 {
			return total_presure
		}
	}
}

func GetPressureFromPathNotTill30(path []string, valves map[string]*Valve) (int, int) {
	total_presure := 0
	current_presure := 0
	total_time := 0
	for i := 1; i < len(path); i++ {
		current_node := valves[path[i-1]]
		next_node := valves[path[i]]
		sub_path := GetShortestPath(valves, current_node, next_node)
		//fmt.Printf("Calculating %v -> %v; %v\n", current_node, next_node, sub_path)
		for t := 0; t < len(sub_path); t++ {
			total_presure += current_presure
			total_time += 1
			// fmt.Printf("Total time: %d; current_presure: %d; total_presure: %d\n", total_time, current_presure, total_presure)
		}
		current_presure += next_node.flow
	}
	return total_presure, current_presure
}

func RemoveElementFromPaths(paths [][]string, el []string) [][]string {
	var ret [][]string
	for _, p := range paths {
		ret = append(ret, p)
	}
	return ret
}

func Path2String(path []string) string {
	return strings.Join(path, "/")
}

func (path_estimations pathEstimations) sort() {
	sort.SliceStable(path_estimations, func(i, j int) bool {
		return path_estimations[i].estimation > path_estimations[j].estimation
	})
}

func (path_estimations pathEstimations) print() {
	for i := 0; i < len(path_estimations); i++ {
		path := path_estimations[i].path
		estimation := path_estimations[i].estimation
		fmt.Printf("{ %v %d }", path, estimation)
	}
	fmt.Println()
}

func (p1 Path) IsEqual(p2 Path) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := 0; i < len(p1); i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}
func (path_estimations pathEstimations) Contains(path Path) bool {
	for i := 0; i < len(path_estimations); i++ {
		if path.IsEqual(path_estimations[i].path) {
			return true
		}
	}
	return false
}

func (path_estimations1 pathEstimations) IsEqual(path_estimations2 pathEstimations) bool {
	for i := 0; i < len(path_estimations1); i++ {
		found := false
		for j := 0; j < len(path_estimations2); j++ {
			if path_estimations1[i].path.IsEqual(path_estimations2[j].path) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
