package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	re := regexp.MustCompile(`Register .: (\d+)`)

	state := State{0, 0, 0, 0, []int{}, []int{}}
	var buff string
	var register []string

	fileScanner.Scan()
	buff = fileScanner.Text()
	register = re.FindStringSubmatch(buff)
	state.A, _ = strconv.Atoi(register[1])

	fileScanner.Scan()
	buff = fileScanner.Text()
	register = re.FindStringSubmatch(buff)
	state.B, _ = strconv.Atoi(register[1])

	fileScanner.Scan()
	buff = fileScanner.Text()
	register = re.FindStringSubmatch(buff)
	state.B, _ = strconv.Atoi(register[1])

	fileScanner.Scan()
	fileScanner.Scan()
	var instructions []int
	buff = fileScanner.Text()
	for _, strInstruction := range strings.Split(strings.Split(buff, " ")[1], ",") {
		intInstruction, _ := strconv.Atoi(strInstruction)
		instructions = append(instructions, intInstruction)
	}
	state.Program = instructions
	fmt.Println(state.String())
	for state.execNextInstruction() {
		fmt.Println(state.String())
	}
	fmt.Println(state.String())
}

type State struct {
	A       int
	B       int
	C       int
	IP      int
	Program []int
	Output  []int
}

func (s *State) String() string {
	res := ""
	res += "Register A: "
	res += strconv.Itoa(s.A)
	res += "\n"
	res += "Register B: "
	res += strconv.Itoa(s.B)
	res += "\n"
	res += "Register C: "
	res += strconv.Itoa(s.C)
	res += "\n"
	res += "IP: "
	res += strconv.Itoa(s.IP)
	res += "\n"

	res += "Program: "
	strProgram := make([]string, len(s.Program))
	for i := 0; i < len(s.Program); i++ {
		strProgram[i] = strconv.Itoa(s.Program[i])
	}
	res += strings.Join(strProgram, ",")
	res += "\n"
	res += "         "
	for i := 0; i < s.IP; i++ {
		for j := 0; j < len(strconv.Itoa(s.Program[i])); j++ {
			res += " "
		}
		res += " "
	}
	res += "^\n"

	res += "Output: "
	strOutput := make([]string, len(s.Output))
	for i := 0; i < len(s.Output); i++ {
		strOutput[i] = strconv.Itoa(s.Output[i])
	}
	res += strings.Join(strOutput, ",")
	res += "\n"

	return res
}

func (s *State) execNextInstruction() bool {
	if s.IP >= len(s.Program) {
		return false
	}
	nextInstruction := s.Program[s.IP]
	nextParam := s.Program[s.IP+1]

	switch nextInstruction {
	case 0: // adv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int(float64(numerator) / denominator)
		s.A = result
		s.IP += 2
	case 1: // bxl
		s.B = s.B ^ nextParam
		s.IP += 2
	case 2: // bst
		s.B = s.getComboOperation(nextParam) % 8
		s.IP += 2
	case 3: // jnz
		if s.A != 0 {
			s.IP = nextParam
		} else {
			s.IP += 2
		}
	case 4: // bxc
		s.B = s.B ^ s.C
		s.IP += 2
	case 5: // out
		output := s.Output
		output = append(output, s.getComboOperation(nextParam)%8)
		s.Output = output
		s.IP += 2
	case 6: // bdv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int(float64(numerator) / denominator)
		s.B = result
		s.IP += 2
	case 7: // bdv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int(float64(numerator) / denominator)
		s.C = result
		s.IP += 2
	}
	return true
}

func (s *State) getComboOperation(param int) int {
	switch param {
	case 0:
		return param
	case 1:
		return param
	case 2:
		return param
	case 3:
		return param
	case 4:
		return s.A
	case 5:
		return s.B
	case 6:
		return s.C
	case 7:
		panic("Combo 7 not implemented")
	}
	return 0
}
