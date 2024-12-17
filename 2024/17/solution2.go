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
	state.A, _ = strconv.ParseInt(register[1], 10, 64)

	fileScanner.Scan()
	buff = fileScanner.Text()
	register = re.FindStringSubmatch(buff)
	state.B, _ = strconv.ParseInt(register[1], 10, 64)

	fileScanner.Scan()
	buff = fileScanner.Text()
	register = re.FindStringSubmatch(buff)
	state.C, _ = strconv.ParseInt(register[1], 10, 64)

	fileScanner.Scan()
	fileScanner.Scan()
	var instructions []int
	buff = fileScanner.Text()
	for _, strInstruction := range strings.Split(strings.Split(buff, " ")[1], ",") {
		intInstruction, _ := strconv.Atoi(strInstruction)
		instructions = append(instructions, intInstruction)
	}
	state.Program = instructions

	var i int64
	i = 0
	for {
		fmt.Println("Testing", strconv.FormatInt(int64(i), 8))

		currentState := state
		currentState.A = i
		for currentState.execNextInstruction() {
			// fmt.Println(currentState.String())
		}
		if equalSlices(currentState.Output, currentState.Program) {
			if len(currentState.Output) == len(currentState.Program) {
				fmt.Println(strconv.FormatInt(int64(i), 8), currentState.Output)
				break
			}
			fmt.Println(strconv.FormatInt(int64(i), 8), currentState.Output)
			i = i * 8
			continue
		}
		// fmt.Println(currentState.Output)
		i++
	}
	fmt.Println("Found", i)
}

func equalSlices(a, b []int) bool {
	// fmt.Print(a, b)
	for i := 0; i < len(a); i++ {
		if a[len(a)-i-1] != b[len(b)-i-1] {
			// fmt.Println("False")
			return false
		}
	}
	// fmt.Println("True")
	return true
}

type State struct {
	A       int64
	B       int64
	C       int64
	IP      int
	Program []int
	Output  []int
}

func (s *State) String() string {
	res := ""
	res += "Register A: "
	res += strconv.FormatInt(int64(s.A), 8)
	res += "\n"
	res += "Register B: "
	res += strconv.FormatInt(int64(s.B), 8)
	res += "\n"
	res += "Register C: "
	res += strconv.FormatInt(int64(s.C), 8)
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
	nextParam := int64(s.Program[s.IP+1])

	switch nextInstruction {
	case 0: // adv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int64(float64(numerator) / denominator)
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
			s.IP = int(nextParam)
		} else {
			s.IP += 2
		}
	case 4: // bxc
		s.B = s.B ^ s.C
		s.IP += 2
	case 5: // out
		output := s.Output
		output = append(output, int(s.getComboOperation(nextParam)%8))
		s.Output = output
		s.IP += 2
	case 6: // bdv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int64(float64(numerator) / denominator)
		s.B = result
		s.IP += 2
	case 7: // bdv
		numerator := s.A
		denominator := math.Pow(2, float64(s.getComboOperation(nextParam)))
		result := int64(float64(numerator) / denominator)
		s.C = result
		s.IP += 2
	}
	return true
}

func (s *State) getComboOperation(param int64) int64 {
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
