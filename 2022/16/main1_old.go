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

	path := []string{"AA"}

	type pathEstimation struct {
		path       []string
		estimation int
	}

	var paths_to_try [][]string
	paths_to_try = append(paths_to_try, path)
	var path_estimations []*pathEstimation
	path_estimation := pathEstimation{path: path, estimation: 0}
	path_estimations = append(path_estimations, &path_estimation)

	for {
		news := false
		var new_path_estimations []*pathEstimation
		for i := 0; i < len(path_estimations) && i <= 3; i++ {
			new_path_estimations = path_estimations
			pe := path_estimations[i]
			fmt.Printf("Trying path %v, estimation: %d\n", pe.path, pe.estimation)
			path := pe.path
			fmt.Println(path)
			news = false
			for i := 0; i < len(valves_with_flow); i++ {
				v := valves_with_flow[i]
				if !IsIn(path, v.name) {
					np := append(path, v.name)
					estimation := GetPressureFromPath(np, valves)
					path_estimation := pathEstimation{path: np, estimation: estimation}
					new_path_estimations = append(new_path_estimations, &path_estimation)
					fmt.Printf("Trying path (inside): %v, estimation: %d\n", path_estimation.path, path_estimation.estimation)
					// fmt.Println(new_path_estimations)
					news = true
				}
				if !news {
					break
				}
			}
		}

		sort.SliceStable(new_path_estimations, func(i, j int) bool {
			return new_path_estimations[i].estimation > new_path_estimations[j].estimation
		})

		path_estimations = nil
		path_set := make(map[string]bool)
		for i := 0; i < len(new_path_estimations); i++ {
			p := new_path_estimations[i]
			if _, ok := path_set[Path2String(p.path)]; ok {
				continue
			}
			path_set[Path2String(p.path)] = true
			fmt.Printf("{path: %v; estimation: %d} ", p.path, p.estimation)
			path_estimations = append(path_estimations, p)
		}
		fmt.Println()

		if !news {
			break
		}

	}
	fmt.Println(*path_estimations[1])
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
