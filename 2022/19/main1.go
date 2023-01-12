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

type RobotCosts [4]int

type RobotsSet [4]RobotCosts

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	robots := make(map[int]RobotsSet)
	fmt.Println(robots)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		bp_id, line_robots := ParseLine(buff)
		robots[bp_id] = line_robots
	}

	fmt.Println(robots)
	build_sequence := []int{1, 1, 1, 2, 1, 2, 3, 3}
	result := Simulate(robots[1], 1, 0, 0, 0, build_sequence)
	fmt.Println(result)

	build_sequence = []int{1, 1, 1, 2, 1, 2, 3, 3, 1, 1, 1, 1, 1, 1}
	result = Simulate(robots[1], 1, 0, 0, 0, build_sequence)
	fmt.Println(result)

	fmt.Print("")

	quality_levels := 0
	for rob_n := range robots {
		var max_vector []int
		max := 0
		for _, vector := range GetAllCombinations(15, []int{1, 2, 3}) {
			v := append([]int{1}, vector...)
			// v = append(v, 3)
			res := Simulate(robots[rob_n], 1, 0, 0, 0, v[:])
			// fmt.Printf("Trying robot: %d; %v, res: %d\n", rob_n, v[:], res)
			if res > max {
				max = res
				max_vector = v
			}
		}
		fmt.Printf("Robot %d, max: %d, max_vector: %v\n", rob_n, max, max_vector)
		quality_levels += max * rob_n
	}

	fmt.Println(quality_levels)

}

func ParseLine(buff string) (int, RobotsSet) {
	r := regexp.MustCompile(`Blueprint (?P<BPID>[\d]+): Each ore robot costs (?P<ORE_COST_IN_ORE>[\d]+) ore. Each clay robot costs (?P<CLAY_COST_IN_ORE>[\d]+) ore. Each obsidian robot costs (?P<OBSIDIAN_COST_IN_ORE>[\d]+) ore and (?P<OBSIDIAN_COST_IN_CLAY>[\d]+) clay. Each geode robot costs (?P<GEODE_COST_IN_ORE>[\d]+) ore and (?P<GEODE_COST_IN_OBSIDIAN>[\d]+) obsidian.`)
	expression := r.FindStringSubmatch(buff)

	bp_id, _ := strconv.Atoi(expression[1])
	// 1st robot
	robot1 := RobotCosts{}
	robot1[0], _ = strconv.Atoi(expression[2])

	// 2nd robot
	robot2 := RobotCosts{}
	robot2[0], _ = strconv.Atoi(expression[3])

	// 3rd robot
	robot3 := RobotCosts{}
	robot3[0], _ = strconv.Atoi(expression[4])
	robot3[1], _ = strconv.Atoi(expression[5])

	// 4rd robot
	robot4 := RobotCosts{}
	robot4[0], _ = strconv.Atoi(expression[6])
	robot4[2], _ = strconv.Atoi(expression[7])

	ret := RobotsSet{robot1, robot2, robot3, robot4}
	return bp_id, ret
}

func Simulate(robots RobotsSet, begin_ore, begin_clay, begin_obsidian, begin_geode int, build_sequence []int) int {
	// fmt.Printf("In simulate, build_sequence: %v\n", build_sequence)
	current_materials := [4]int{0, 0, 0, 0}
	current_bots := [4]int{begin_ore, begin_clay, begin_obsidian, begin_geode}
	sequence_pointer := 0
	for t := 1; t <= 24; t++ {
		new_bots := current_bots

		// We build bots
		for {
			if sequence_pointer >= len(build_sequence) {
				fmt.Printf("We could build more bots, but there's no more instructions\n")
				break
			}
			bot_type := build_sequence[sequence_pointer]
			current_sequence_bot := robots[bot_type]
			if current_sequence_bot.CanIBuild(current_materials) {
				//fmt.Printf("We can buy a type %d bot, let's do\n", bot_type)
				// Substract the cost of each material
				for material_n := 0; material_n < 4; material_n++ {
					current_materials[material_n] = current_materials[material_n] - current_sequence_bot[material_n]
				}
				// And then add the bot to the list
				new_bots[bot_type] += 1
				sequence_pointer += 1
			} else {
				break
			}
		}

		// We harvest
		for bot_n := 0; bot_n < 4; bot_n++ {
			current_materials[bot_n] += current_bots[bot_n]
		}
		current_bots = new_bots

		//fmt.Printf("Time: %d; materials: %v; bots: %v, sequence: %v, sequence_pointer: %d\n", t, current_materials, current_bots, build_sequence, sequence_pointer)
	}
	return current_materials[3]
}

func (rs RobotCosts) CanIBuild(current_materials [4]int) bool {
	// fmt.Printf("CanIBuild? %v with materials %v\n", rs, current_materials)
	for material_n := 0; material_n < 4; material_n++ {
		// fmt.Printf("CanIBuild? material: %d; need: %d, have: %d\n", material_n, rs[material_n], current_materials[material_n])
		if rs[material_n] > current_materials[material_n] {
			return false
		}
	}
	return true
}

func GetAllCombinations(n int, lst []int) [][]int {
	if n == 0 {
		return [][]int{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	var res [][]int
	for i := 0; i < len(lst); i++ {
		subres := []int{lst[i]}
		for _, comb := range GetAllCombinations(n-1, lst) {
			res = append(res, append(subres, comb...))
		}
	}

	return res
}
