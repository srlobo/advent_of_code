package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var megacache map[string]int

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
		fmt.Println(buff)
		megacache = make(map[string]int)
		checks := make([]int, 0)
		for _, check := range strings.Split(buff[1], ",") {
			check_int, _ := strconv.Atoi(check)
			checks = append(checks, check_int)
		}
		full_line := buff[0]
		full_checks := make([]int, 0)
		full_checks = append(full_checks, checks...)
		for i := 0; i < 4; i++ {
			full_line += "?" + buff[0]
			full_checks = append(full_checks, checks...)
		}

		fmt.Println("checking ", full_line, full_checks)
		tmp := tryNextPossibility(full_line, full_checks)
		total += tmp
		fmt.Println(tmp)
		fmt.Println()
	}
	fmt.Println(total)
}

func calculateCacheKey(condition string, check []int, current_char_count int) string {
	return condition + " " + fmt.Sprint(check) + strconv.Itoa(current_char_count)
}

func tryNextPossibility(condition string, check []int) int {
	// check if condition is possible with the checks
	var i int

	group := 0
	group_len := 0
	for i = 0; i < len(condition); i++ {
		if group >= len(check) && group_len > 0 {
			// fmt.Println("The combination is not valid (group>=len(check))", condition)
			return 0
		}

		if condition[i] == '#' {
			group_len++
		} else if condition[i] == '.' && group_len > 0 {
			if check[group] != group_len {
				// fmt.Println("The combination is not valid (check[group] != group_len)", condition)
				return 0
			} else {
				group_len = 0
				group++
			}
		} else if condition[i] == '?' {
			break
		}
	}
	// fmt.Println("len(condition): ", len(condition), "i: ", i, "group: ", group, "group_len: ", group_len, "len(check): ", len(check))
	if i == len(condition) && group == len(check) && group_len == 0 { // We end there when we have finished and there's some . after the last group
		// fmt.Println("The combination is valid", condition)
		return 1
	}
	if i == len(condition) && group != len(check)-1 { // The combination is not valid
		// fmt.Println("The combination is not valid (len(condition) && group != len(check)-1)", condition)
		return 0
	} else if i == len(condition) && group == len(check)-1 { // The combination is valid
		if check[group] != group_len {
			// fmt.Println("The combination is not valid (check[group] != group_len)", condition)
			return 0
		} else {
			// fmt.Println("The combination is valid", condition)
			return 1
		}
	}

	//	fmt.Println(megacache)
	var my_condition_1, my_condition_2 string
	total := 0
	for i = 0; i < len(condition); i++ {
		if condition[i] == '?' {
			break
		}
		my_condition_1 += string(condition[i])
		my_condition_2 += string(condition[i])
	}

	my_condition_1_cache_key := calculateCacheKey("."+condition[i+1:], check[group:], group_len)
	my_condition_1 += "." + condition[i+1:]
	// fmt.Println("Trying ", my_condition_1)
	if val, ok := megacache[my_condition_1_cache_key]; ok {
		// fmt.Println("Cache hit:", my_condition_1_cache_key, val)
		total += val
	} else {
		tmp := tryNextPossibility(my_condition_1, check)
		megacache[my_condition_1_cache_key] = tmp
		total += tmp
	}

	my_condition_2_cache_key := calculateCacheKey("#"+condition[i+1:], check[group:], group_len)
	my_condition_2 += "#" + condition[i+1:]
	// fmt.Println("Trying ", my_condition_2)
	if val, ok := megacache[my_condition_2_cache_key]; ok {
		// fmt.Println("Cache hit:", my_condition_2_cache_key, val)
		total += val
	} else {
		tmp := tryNextPossibility(my_condition_2, check)
		megacache[my_condition_2_cache_key] = tmp
		total += tmp
	}

	// fmt.Println("Keep on searching (examined ", my_condition_1, my_condition_2, ") total:", total)
	return total
}
