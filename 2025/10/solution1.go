package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var empty = struct{}{}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
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

	total := 0
	for fileScanner.Scan() {
		buff := strings.Split(fileScanner.Text(), " ")

		puzzle := make(buttonSet)
		var buttonsets []buttonSet
		var joltage []int

		// Puzzle
		for i := 1; i < len(buff[0])-1; i++ {
			if buff[0][i] == '#' {
				puzzle[i-1] = empty
			}
		}

		// Buttonsets and joltage
		for i := 1; i < len(buff); i++ {
			if buff[i][0] == '(' {
				buttonSetArr := strings.Split(buff[i][1:len(buff[i])-1], ",")
				buttonset := make(buttonSet)
				for _, bStr := range buttonSetArr {
					bInt, _ := strconv.Atoi(bStr)
					buttonset[bInt] = empty
				}
				buttonsets = append(buttonsets, buttonset)
			} else if buff[i][0] == '{' {
				for jStr := range strings.SplitSeq(buff[i][1:len(buff[i])-1], ",") {
					jInt, _ := strconv.Atoi(jStr)
					joltage = append(joltage, jInt)
				}
			}
		}

		fmt.Println("Target: ", puzzle)
		fmt.Println("Buttons: ", buttonsets)
		fmt.Println("Joltage: ", joltage)
		fmt.Println("Start calculations")

		cache := stateCache{cache: make(map[string]struct{})}
		solution := cache.solve(puzzle, buttonsets, []buttonSet{}, 0) + 1
		fmt.Println("Solution: ", solution)
		total += solution

	}
	fmt.Println("Total: ", total)
}

type buttonSet map[int]struct{}

func (b buttonSet) String() string {
	var resArr []string
	for b := range b {
		resArr = append(resArr, strconv.Itoa(b))
	}
	return strings.Join(resArr, ",")
}

func compose(a, b buttonSet) buttonSet {
	res := make(buttonSet)
	for button := range a {
		res[button] = empty
	}

	for button := range b {
		if _, ok := res[button]; ok {
			delete(res, button)
		} else {
			res[button] = empty
		}

	}

	return res
}

func areEqual(a, b buttonSet) bool {
	if len(compose(a, b)) == 0 {
		return true
	} else {
		return false
	}
}

type stateCache struct {
	cache map[string]struct{}
}

func (cache *stateCache) includes(b buttonSet) bool {
	_, ok := cache.cache[b.toStr()]
	return ok
}
func (cache *stateCache) append(b buttonSet) {
	cache.cache[b.toStr()] = empty
}

func (b buttonSet) toStr() string {
	var buttonArr []string
	for k := range b {
		buttonArr = append(buttonArr, strconv.Itoa(k))
	}
	sort.Slice(buttonArr, func(i, j int) bool {
		return i < j
	})

	return strings.Join(buttonArr, "")
}

func (cache *stateCache) solve(objective buttonSet, originalButtons, actualStates []buttonSet, depth int) int {
	depth++
	fmt.Println("Round ", depth, "actualStates ", len(actualStates))
	var newButtonSet []buttonSet

	if len(actualStates) == 0 {
		for _, currentButton := range originalButtons {
			newButton := currentButton
			if cache.includes(newButton) {
				// fmt.Println("Cache hit!", newButton)
				continue
			}
			if areEqual(newButton, objective) { // Found
				// fmt.Println("Found! (", depth, ") ", currentState, currentButton)
				return 0
			}
			newButtonSet = append(newButtonSet, newButton)
			cache.append(newButton)
		}
	} else {
		for _, currentState := range actualStates {
			for _, currentButton := range originalButtons {
				newButton := compose(currentState, currentButton)
				// fmt.Println(currentState, currentButton, "-> ", newButton)
				if cache.includes(newButton) {
					// fmt.Println("Cache hit!", newButton)
					continue
				}
				if areEqual(newButton, objective) { // Found
					// fmt.Println("Found! (", depth, ") ", currentState, currentButton)
					return 0
				}
				newButtonSet = append(newButtonSet, newButton)
				cache.append(newButton)

			}
		}
	}

	return cache.solve(objective, originalButtons, newButtonSet, depth) + 1
}
