package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

	var tokenizedArray []string
	for {
		// First line
		fileScanner.Scan()
		packet1 := fileScanner.Text()
		tokenizedArray = append(tokenizedArray, packet1)

		// Second line
		fileScanner.Scan()
		packet2 := fileScanner.Text()
		tokenizedArray = append(tokenizedArray, packet2)

		fmt.Println()
		if !fileScanner.Scan() {
			fmt.Println("Found EOF")
			break
		}
	}
	tokenizedArray = append(tokenizedArray, "[[2]]")
	tokenizedArray = append(tokenizedArray, "[[6]]")

	// Now we need to sort the array with the compareTokens function
	sort.SliceStable(tokenizedArray, func(i, j int) bool {
		return 1 == compareTokens(tokenize(tokenizedArray[i]), tokenize(tokenizedArray[j]))
	})

	fmt.Println()
	res := 1
	for i, tok := range tokenizedArray {
		if tok == "[[2]]" || tok == "[[6]]" {
			res *= (i + 1)
		}
		fmt.Println(tok)
	}
	fmt.Println(res)
}

func tokenize(buff string) []string {
	level := 0
	buff = buff[1 : len(buff)-1]
	var result []string
	if len(buff) == 0 {
		return nil
	}
	var item []byte
	j := 0
	for {
		if buff[j] == '[' {
			level += 1
			item = append(item, buff[j])
		} else if buff[j] == ']' && level != 0 {
			level = level - 1
			item = append(item, buff[j])
		} else if buff[j] == ',' && level == 0 { // finished this token
			result = append(result, string(item))
			item = nil
		} else {
			item = append(item, buff[j])
		}

		j++
		if j >= len(buff) {
			result = append(result, string(item))
			break
		}
	}

	// fmt.Printf("Tokenizado: %s\n", strings.Join(result, "|"))
	return result
}

func compareTokens(token1 []string, token2 []string) int {
	fmt.Printf("Compare %s vs %s\n", token1, token2)
	// Return values:
	// 1: packet in order
	// -1: packet not in order
	// 0: Don't know
	i := 0
	j := 0
	for {

		if i >= len(token1) && j < len(token2) {
			// First list exhausted before
			fmt.Print("Left side ran out of items, so inputs are in the right order")
			return 1
		} else if j >= len(token2) && i < len(token1) {
			// Second list exhausted before
			fmt.Printf("Right side ran out of items, so inputs are not in the right order")
			return -1
		} else if i >= len(token1) && j >= len(token2) {
			return 0
		}

		left := token1[i]
		right := token2[j]
		fmt.Printf("Compare %s vs %s\n", left, right)

		if isInteger(left) && isInteger(right) {
			intLeft, _ := strconv.Atoi(left)
			intRight, _ := strconv.Atoi(right)
			if intLeft < intRight {
				fmt.Printf("Left side is smaller, so inputs are in the right order")
				return 1
			} else if intRight < intLeft {
				fmt.Printf("Right side is smaller, so inputs are not in the right order")
				return -1
			} else {
				i++
				j++
			}
		} else if isList(left) && isList(right) {
			ordered := compareTokens(tokenize(left), tokenize(right))
			if ordered != 0 {
				return ordered
			} else {
				i++
				j++
			}
		} else if isList(left) && isInteger(right) {
			// Convert right into a list
			fmt.Printf("Mixed types; convert right to [%s] and retry comparison\n", right)
			newpacket := fmt.Sprintf("[%s]", right)
			ordered := compareTokens(tokenize(left), tokenize(newpacket))
			if ordered != 0 {
				return ordered
			} else {
				i++
				j++
			}
		} else if isInteger(left) && isList(right) {
			fmt.Printf("Mixed types; convert left to [%s] and retry comparison\n", left)
			// Convert right into a list
			newpacket := fmt.Sprintf("[%s]", left)
			ordered := compareTokens(tokenize(newpacket), tokenize(right))
			if ordered != 0 {
				return ordered
			} else {
				i++
				j++
			}
		}

	}
}

func isInteger(s string) bool {
	_, ok := strconv.Atoi(s)
	return ok == nil
}

func isList(s string) bool {
	return s[0] == '['
}
