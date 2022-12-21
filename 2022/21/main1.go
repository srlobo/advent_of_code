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

type Line struct {
	name       string
	value      int
	left       string
	right      string
	operator   string
	calculated bool
}

type Lines map[string]*Line

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := make(Lines)
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		line := ParseLine(buff)
		lines[line.name] = &line
	}

	fmt.Println(lines.Calculate(lines["root"]))
}

func ParseLine(buff string) Line {
	r := regexp.MustCompile(`(?P<LVALUE>[^:]+): (?P<RVALUE1>[^ ]+) (?P<OPERATION>[-+/*]) (?P<RVALUE2>[^ ]+)`)
	expression := r.FindStringSubmatch(buff)
	if len(expression) > 0 {
		return Line{
			name:       expression[1],
			value:      0,
			left:       expression[2],
			operator:   expression[3],
			right:      expression[4],
			calculated: false,
		}
	} else {
		r := regexp.MustCompile(`(?P<LVALUE>[^:]+): (?P<VALUE>[\d]+)`)
		expression = r.FindStringSubmatch(buff)
		value, _ := strconv.Atoi(expression[2])
		return Line{name: expression[1], value: value, calculated: true}
	}
}

func (lines Lines) Calculate(node *Line) int {
	fmt.Printf("Entering Calculate, node: %v\n", *node)
	lvalue := lines[node.left]
	if !lvalue.calculated {
		lines.Calculate(lvalue)
	}

	rvalue := lines[node.right]
	if !rvalue.calculated {
		lines.Calculate(rvalue)
	}

	var ret int

	switch node.operator {
	case "+":
		ret = lvalue.value + rvalue.value
	case "-":
		ret = lvalue.value - rvalue.value
	case "*":
		ret = lvalue.value * rvalue.value
	case "/":
		ret = lvalue.value / rvalue.value
	}

	node.value = ret
	node.calculated = true
	fmt.Printf("Ending Calculate, node: %v; ret: %d\n", *node, ret)

	return ret
}
