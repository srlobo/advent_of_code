package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func main() {
	// sequences := []string{"A>^^A", "A^>^A", "A^^>A"}
	// for _, sequence := range sequences {
	// 	fmt.Println(sequence)
	// 	for i := 1; i < len(sequence); i++ {
	// 		ret, ok := arrowMovement(rune(sequence[i-1]), rune(sequence[i]), 0)
	// 		fmt.Println("ret ", string(sequence[i-1]), string(sequence[i]), ret, ok)
	// 	}
	// }
	// return
	// sequences := []string{"17", "28"}
	// for _, sequence := range sequences {
	// 	fmt.Println(sequence)
	// 	for i := 1; i < len(sequence); i++ {
	// 		ret, ok := goFromNumberToNumber(rune(sequence[i-1]), rune(sequence[i]), 0)
	// 		fmt.Println("ret ", string(sequence[i-1]), string(sequence[i]), ret, ok)
	// 	}
	// }
	// return

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		defer readFile.Close()
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cache := make(instructionsCache)
	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		number, _ := strconv.Atoi(buff[:len(buff)-1])

		sum := 0
		buff = "A" + buff
		for i := 1; i < len(buff); i++ {
			ret, _ := goFromNumberToNumber(rune(buff[i-1]), rune(buff[i]), 0)
			minNum := MaxInt
			initialDepth := 25
			for j := 0; j < len(ret); j++ {
				// fmt.Println("ret", ret[j])
				var res int
				res = cache.extractInstructions("A"+ret[j], initialDepth)
				if res < minNum {
					minNum = res
				}
			}
			sum += minNum

			// In this point we have the movement from one key to the next

		}

		fmt.Println("sum", sum)
		total += sum * number
	}
	fmt.Println(total)
}

type instructionsCacheKey struct {
	instructions string
	depth        int
}

type instructionsCache map[instructionsCacheKey]int

func (cache *instructionsCache) extractInstructions(originalInstructions string, depth int) int {
	if val, ok := (*cache)[instructionsCacheKey{originalInstructions, depth}]; ok {
		return val
	}
	total := 0
	// fmt.Println("originalInstructions: ", originalInstructions, "depth: ", depth)
	if depth == 0 {
		total = len(originalInstructions) - 1 // because of the added A at the beginning
	} else {
		for i := 1; i < len(originalInstructions); i++ {
			ret, ok := arrowMovement(rune(originalInstructions[i-1]), rune(originalInstructions[i]), 0)
			// fmt.Println("Return", ret)
			if ok {
				ret = getSmallest(ret)
				minNum := MaxInt
				var num int
				for j := 0; j < len(ret); j++ {
					num = cache.extractInstructions("A"+ret[j], depth-1)
					if num < minNum {
						minNum = num
					}
				}
				total += minNum
			}
		}
	}
	// fmt.Println("depth:", depth, total)
	(*cache)[instructionsCacheKey{originalInstructions, depth}] = total
	return total
}

type RunePair [2]rune

func goFromNumberToNumber(origin, destination rune, depth int) ([]string, bool) {
	// fmt.Println("Enter: origin: ", string(origin), "destination: ", string(destination), "depth", depth)
	result := []string{}

	if depth > 6 {
		return result, false
	}

	numberRoute := map[RunePair]string{
		{'A', '0'}: "<",
		{'A', '3'}: "^",

		{'0', '2'}: "^",

		{'1', '2'}: ">",
		{'1', '4'}: "^",

		{'2', '3'}: ">",
		{'2', '5'}: "^",

		{'3', '6'}: "^",

		{'4', '5'}: ">",
		{'4', '7'}: "^",

		{'5', '6'}: ">",
		{'5', '8'}: "^",

		{'6', '9'}: "^",

		{'7', '8'}: ">",

		{'8', '9'}: ">",
	}

	if origin == destination {
		result = append(result, "")
	} else if res, ok := numberRoute[RunePair{origin, destination}]; ok {
		result = append(result, res)
	} else if res, ok := numberRoute[RunePair{destination, origin}]; ok {
		result = append(result, reverse(res))
	} else {
		// Let's dig on the different ways to reach the number
		for r1, d := range numberRoute {
			if origin == r1[0] {
				// fmt.Println("Trying", r1.String())
				r2, ok := goFromNumberToNumber(r1[1], destination, depth+1)
				if ok {
					for i := 0; i < len(r2); i++ {
						result = append(result, d+r2[i])
						// fmt.Println("Posible result: ", result)
					}
				}

			} else if origin == r1[1] {
				// fmt.Println("Trying (reverse)", r1.String())
				r2, ok := goFromNumberToNumber(r1[0], destination, depth+1)
				if ok {
					for i := 0; i < len(r2); i++ {
						result = append(result, reverse(d)+r2[i])
						// fmt.Println("Posible result: ", result)
					}
				}

			}
		}
	}

	// fmt.Println("result: ", result)
	minLength := MaxInt
	for i := 0; i < len(result); i++ {
		if len(result[i]) < minLength {
			minLength = len(result[i])
		}
	}

	refinedResult := []string{}
	for i := 0; i < len(result); i++ {
		if len(result[i]) > minLength {
			continue
		}
		if depth == 0 {
			refinedResult = append(refinedResult, result[i]+"A")
		} else {
			refinedResult = append(refinedResult, result[i])
		}
	}
	// fmt.Println("origin: ", string(origin), "destination: ", string(destination), "depth", depth, "refinedResult: ", refinedResult)
	return refinedResult, true
}

func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '>':
			result = result + "<"
		case '<':
			result = result + ">"
		case '^':
			result = result + "v"
		case 'v':
			result = result + "^"
		}
	}
	return result
}

func arrowMovement(origin, destination rune, depth int) ([]string, bool) {
	result := []string{}
	if depth > 4 {
		return result, false
	}
	arrowRoute := map[RunePair]string{
		{'^', 'A'}: ">",
		{'^', 'v'}: "v",

		{'A', '>'}: "v",

		{'<', 'v'}: ">",

		{'v', '>'}: ">",
	}
	if origin == destination {
		result = append(result, "")
	} else if res, ok := arrowRoute[RunePair{origin, destination}]; ok {
		result = append(result, res)
	} else if res, ok := arrowRoute[RunePair{destination, origin}]; ok {
		result = append(result, reverse(res))
	} else {
		// Let's dig on the different ways to reach the other key
		for r1, d := range arrowRoute {
			if origin == r1[0] {
				// fmt.Println("Trying", r1.String())
				r2, ok := arrowMovement(r1[1], destination, depth+1)
				if ok {
					for i := 0; i < len(r2); i++ {
						result = append(result, d+r2[i])
						// fmt.Println("Posible result: ", result)
					}
				}

			} else if origin == r1[1] {
				// fmt.Println("Trying (reverse)", r1.String())
				r2, ok := arrowMovement(r1[0], destination, depth+1)
				if ok {
					for i := 0; i < len(r2); i++ {
						result = append(result, reverse(d)+r2[i])
						// fmt.Println("Posible result: ", result)
					}
				}

			}
		}
	}

	result = getSmallest(result)

	for i := 0; i < len(result); i++ {
		if depth == 0 {
			result[i] = result[i] + "A"
		}
	}
	// fmt.Println("origin: ", string(origin), "destination: ", string(destination), "depth", depth, "refinedResult: ", refinedResult)

	return result, true
}

func (r RunePair) String() string {
	return fmt.Sprintf("%c -> %c", r[0], r[1])
}

func getSmallest(array []string) []string {
	min := MaxInt
	for i := 0; i < len(array); i++ {
		if len(array[i]) < min {
			min = len(array[i])
		}
	}
	ret := []string{}
	for i := 0; i < len(array); i++ {
		if len(array[i]) == min {
			ret = append(ret, array[i])
		}
	}
	return ret
}
