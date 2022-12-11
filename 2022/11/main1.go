package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Monkey struct {
	items     []int
	operation struct {
		operator string
		operand  string
	}
	test   int
	action map[bool]int
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

	monkeys := make(map[int]Monkey)

	for {
		var monkey Monkey
		// Monkey line
		fileScanner.Scan()
		buff := fileScanner.Text()
		// Remove the :
		buff = buff[:len(buff)-1]
		monkey_n, _ := strconv.Atoi(strings.Split(buff, " ")[1])
		fmt.Printf("Monkey number %d\n", monkey_n)

		// Starting
		fileScanner.Scan()
		buff = fileScanner.Text()
		items := strings.Split(strings.ReplaceAll(buff, " ", ""), ":")[1]
		for _, s := range strings.Split(items, ",") {
			val, _ := strconv.Atoi(s)
			monkey.items = append(monkey.items, val)
		}

		// Operation
		fileScanner.Scan()
		buff = fileScanner.Text()

		operation := strings.Split(buff, " ")
		operator := operation[len(operation)-2]
		operand := operation[len(operation)-1]
		monkey.operation.operator = operator
		monkey.operation.operand = operand

		// Test
		fileScanner.Scan()
		buff = fileScanner.Text()
		test := strings.Split(buff, " ")
		monkey.test, _ = strconv.Atoi(test[len(test)-1])

		monkey.action = make(map[bool]int)
		// If true
		fileScanner.Scan()
		buff = fileScanner.Text()
		action_a := strings.Split(buff, " ")
		cond_val, _ := strconv.ParseBool(action_a[5][:len(action_a[5])-1])
		val, _ := strconv.Atoi(action_a[len(action_a)-1])
		monkey.action[cond_val] = val

		// If false
		fileScanner.Scan()
		buff = fileScanner.Text()
		action_a = strings.Split(buff, " ")
		cond_val, _ = strconv.ParseBool(action_a[5][:len(action_a[5])-1])
		val, _ = strconv.Atoi(action_a[len(action_a)-1])
		monkey.action[cond_val] = val

		fmt.Println(monkey)
		monkeys[monkey_n] = monkey

		// Empty line
		if !fileScanner.Scan() {
			fmt.Println("Found EOF")
			break
		}
	}

	turn_n := 0
	monkey_inspections := make(map[int]int)
	for {
		turn_n++
		for monkey_n := 0; monkey_n < len(monkeys); monkey_n++ {
			monkey := monkeys[monkey_n]
			fmt.Printf("Monkey %d:\n", monkey_n)
			fmt.Println(monkey)
			// For each item
			for i := 0; i < len(monkey.items); i++ {
				// Increment the number of item inspections
				monkey_inspections[monkey_n] += 1

				item := monkey.items[i]
				fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
				worryLevel := monkey.CalculateWorryLevel(item)
				// worryLevel = int(math.Round(float64(worryLevel) / float64(3)))
				worryLevel = worryLevel / 3
				fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", worryLevel)

				test := worryLevel%monkey.test == 0
				if test {
					fmt.Printf("    Current worry level is divisible by %d.\n", monkey.test)
				} else {
					fmt.Printf("    Current worry level is not divisible by %d.\n", monkey.test)
				}
				dst_monkey_n := monkey.action[test]
				items := append(monkeys[dst_monkey_n].items, worryLevel)
				dst_monkey, _ := monkeys[dst_monkey_n]
				dst_monkey.items = items
				monkeys[dst_monkey_n] = dst_monkey
				fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", worryLevel, dst_monkey_n)
			}
			// Clean monkey items
			monkey.items = []int{}
			monkeys[monkey_n] = monkey
		}
		fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", turn_n)
		for monkey_n := 0; monkey_n < len(monkeys); monkey_n++ {
			var items_s []string
			for i := 0; i < len(monkeys[monkey_n].items); i++ {
				items_s = append(items_s, fmt.Sprintf("%d", monkeys[monkey_n].items[i]))
			}
			items := strings.Join(items_s, ", ")
			fmt.Printf("Monkey %d: %s\n", monkey_n, items)
		}
		fmt.Println()
		if turn_n == 20 {
			break
		}
	}
	fmt.Println(monkey_inspections)
	var inspections []int
	for _, i := range monkey_inspections {
		inspections = append(inspections, i)
	}
	sort.Ints(inspections)
	fmt.Println(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}

func (monkey Monkey) CalculateWorryLevel(item int) int {
	var operand int
	var ret int

	if monkey.operation.operand == "old" {
		operand = item
	} else {
		operand, _ = strconv.Atoi(monkey.operation.operand)
	}

	switch monkey.operation.operator {
	case "+":
		ret = item + operand
		fmt.Printf("    Worry level increases by %d to %d.\n", operand, ret)
	case "*":
		ret = item * operand
		fmt.Printf("    Worry level is multiplied by %d to %d.\n", operand, ret)
	}
	return ret
}
