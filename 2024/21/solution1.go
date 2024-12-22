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

	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		number, _ := strconv.Atoi(buff[:len(buff)-1])

		sum := 0
		buff = "A" + buff
		full := []string{}
		for i := 1; i < len(buff); i++ {
			ret, _ := goFromNumberToNumber(rune(buff[i-1]), rune(buff[i]), 0)
			fmt.Println("ret", ret)

			kbd1Instructions := []string{}
			for j := 0; j < len(ret); j++ {
				kbd1Instructions = append(kbd1Instructions, extractInstructions("A"+ret[j])...)
			}
			kbd1Instructions = getSmallest(kbd1Instructions)
			fmt.Println("kbd1Instructions", kbd1Instructions, len(kbd1Instructions[0]))

			kbd2Instructions := []string{}
			for j := 0; j < len(kbd1Instructions); j++ {
				kbd2Instructions = append(kbd2Instructions, extractInstructions("A"+kbd1Instructions[j])...)
			}
			kbd2Instructions = getSmallest(kbd2Instructions)
			fmt.Println("kbd2Instructions", kbd2Instructions, len(kbd2Instructions[0]))
			if len(full) == 0 {
				full = kbd2Instructions
			} else {
				newFull := []string{}
				for j := 0; j < len(full); j++ {
					for h := 0; h < len(kbd2Instructions); h++ {
						newFull = append(newFull, full[j]+kbd2Instructions[h])
					}
				}
				full = newFull
			}

			sum += len(kbd2Instructions[0])

			// In this point we have the movement from one key to the next

		}

		fmt.Println("sum", sum)
		total += sum * number
	}
	fmt.Println(total)
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

func extractInstructions(originalInstructions string) []string {
	retInstructions := []string{}
	for i := 1; i < len(originalInstructions); i++ {
		ret, ok := arrowMovement(rune(originalInstructions[i-1]), rune(originalInstructions[i]), 0)
		if ok {
			if len(retInstructions) == 0 {
				retInstructions = append(retInstructions, ret...)
			} else {
				newInstructions := []string{}
				for h := 0; h < len(retInstructions); h++ {
					for j := 0; j < len(ret); j++ {
						newInstructions = append(newInstructions, retInstructions[h]+ret[j])
					}
				}
				retInstructions = newInstructions
			}
		}
	}
	return retInstructions
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

type ArrowCacheKey struct {
	origin      rune
	destination rune
	depth       int
}

type ArrowCache map[ArrowCacheKey][]string

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
