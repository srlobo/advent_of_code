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

	// lines.Print()

	for k, l := range lines {
		if l.left == "humn" {
			new_humn_line := ExpressLineAsLeft(*l)
			delete(lines, k)
			lines["humn"] = &new_humn_line
			break
		}
	}

	solutions := make(Lines)
	root_right_str := lines["root"].right
	root_left_str := lines["root"].left
	converted_nodes := make(map[string]struct{})
	converted_nodes["humn"] = struct{}{}
	delete(lines, "root")

	root_right_result := lines.Calculate(lines[root_right_str], converted_nodes, solutions)
	solutions[root_right_str] = lines[root_right_str]

	lines.Print()
	lines.DetectMissingNode(root_left_str, converted_nodes)

	solutions[root_left_str] = &Line{name: root_left_str, value: root_right_result, calculated: true}
	fmt.Println("=========================")
	lines.Print()
	fmt.Printf("Converted nodes %v\n", converted_nodes)

	fmt.Println(lines.Calculate(lines["humn"], converted_nodes, solutions))
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

func (lines Lines) Calculate(node *Line, converted_nodes map[string]struct{}, solutions Lines) int {
	fmt.Printf("Entering Calculate, node: %v\n", *node)
	//lines.Print()
	lvalue, ok := solutions[node.left]
	if !ok {
		lvalue, ok = lines[node.left]
		if !ok {
			lines.DetectMissingNode(node.left, converted_nodes)
			lvalue, ok = lines[node.left]
			if !ok {
				fmt.Printf("lines[%s] not found\n", node.left)
			}
		}
	}
	fmt.Printf("lvalue: %s\n", node.left)
	fmt.Println(lvalue)
	if !lvalue.calculated {
		lines.Calculate(lvalue, converted_nodes, solutions)
	}

	rvalue, ok := solutions[node.right]
	if !ok {
		rvalue, ok = lines[node.right]
		if !ok {
			lines.DetectMissingNode(node.right, converted_nodes)
			rvalue, ok = lines[node.right]
			if !ok {
				fmt.Printf("lines[%s] not found\n", node.left)
			}
		}
	}

	if !rvalue.calculated {
		lines.Calculate(rvalue, converted_nodes, solutions)
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
		if lvalue.value%rvalue.value != 0 {
			fmt.Printf("Problem with this division: %d / %d\n", lvalue.value, rvalue.value)
		}

	}

	node.value = ret
	solutions[node.name] = node
	node.calculated = true
	fmt.Printf("Ending Calculate, node: %v; ret: %d\n", *node, ret)

	return ret
}

func (lines Lines) Print() {
	for k, v := range lines {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}
	fmt.Println()
}

func ExpressLineAsLeft(line Line) Line {
	var ret Line

	ret.name = line.left
	ret.calculated = false
	switch line.operator {
	case "+": // name = left + right -> left = name - right
		ret.operator = "-"
		ret.left = line.name
		ret.right = line.right
	case "-": // name = left - right -> left = name + right
		ret.operator = "+"
		ret.left = line.name
		ret.right = line.right

	case "*": // name = left * right -> left = name / right
		ret.operator = "/"
		ret.left = line.name
		ret.right = line.right
	case "/": // name = left / right -> left = name * right
		ret.operator = "*"
		ret.left = line.name
		ret.right = line.right
	}
	fmt.Printf("Convert %v in %v\n", line, ret)
	return ret
}

func ExpressLineAsRight(line Line) Line {
	var ret Line

	ret.name = line.right
	ret.calculated = false
	switch line.operator {
	case "+": // name = left + right -> right = name - left
		ret.operator = "-"
		ret.left = line.name
		ret.right = line.left
	case "-": // name = left - right -> right = left - name
		ret.operator = "-"
		ret.left = line.left
		ret.right = line.name

	case "*": // name = left * right -> right = name / left
		ret.operator = "/"
		ret.left = line.name
		ret.right = line.left
	case "/": // name = left / right -> right = left / name
		ret.operator = "/"
		ret.left = line.left
		ret.right = line.name
	}
	fmt.Printf("Convert %v in %v\n", line, ret)
	return ret
}

func (lines Lines) DetectMissingNode(node_str string, converted_nodes map[string]struct{}) {
	for k, v := range lines {
		if v.left == node_str {
			if _, ok := converted_nodes[v.right]; ok {
				continue
			}
			if _, ok := converted_nodes[v.name]; ok {
				continue
			}
			new_v := ExpressLineAsLeft(*v)
			delete(lines, k)
			lines[new_v.name] = &new_v
			converted_nodes[new_v.name] = struct{}{}
			break
		} else if v.right == node_str {
			if _, ok := converted_nodes[v.left]; ok {
				continue
			}
			if _, ok := converted_nodes[v.name]; ok {
				continue
			}

			new_v := ExpressLineAsRight(*v)
			delete(lines, k)
			lines[new_v.name] = &new_v
			converted_nodes[new_v.name] = struct{}{}
			break
		}
	}
}
