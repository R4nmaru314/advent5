package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const file = "inputs.txt"

type Stack struct {
	stack []string
}

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	var resultStringPart1 []string
	var resultStringPart2 []string

	stackList1 := []Stack{
		{[]string{"W", "B", "D", "N", "C", "F", "J"}},
		{[]string{"P", "Z", "V", "Q", "L", "S", "T"}},
		{[]string{"P", "Z", "B", "G", "J", "T"}},
		{[]string{"D", "T", "L", "J", "Z", "B", "H", "C"}},
		{[]string{"G", "V", "B", "J", "S"}},
		{[]string{"P", "S", "Q"}},
		{[]string{"B", "V", "D", "F", "L", "M", "P", "N"}},
		{[]string{"P", "S", "M", "F", "B", "D", "L", "R"}},
		{[]string{"V", "D", "T", "R"}},
	}

	stackList2 := []Stack{
		{[]string{"W", "B", "D", "N", "C", "F", "J"}},
		{[]string{"P", "Z", "V", "Q", "L", "S", "T"}},
		{[]string{"P", "Z", "B", "G", "J", "T"}},
		{[]string{"D", "T", "L", "J", "Z", "B", "H", "C"}},
		{[]string{"G", "V", "B", "J", "S"}},
		{[]string{"P", "S", "Q"}},
		{[]string{"B", "V", "D", "F", "L", "M", "P", "N"}},
		{[]string{"P", "S", "M", "F", "B", "D", "L", "R"}},
		{[]string{"V", "D", "T", "R"}},
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		howMany, _ := strconv.Atoi(lineSplit[1])
		fromWhere, _ := strconv.Atoi(lineSplit[3])
		toWhere, _ := strconv.Atoi(lineSplit[5])

		calculateStackPart1(howMany, fromWhere, toWhere, stackList1)
		calculateStackPart2(howMany, fromWhere, toWhere, stackList2)
	}

	for _, val := range stackList1 {
		resultStringPart1 = append(resultStringPart1, val.stack[len(val.stack)-1])
	}
	for _, val := range stackList2 {
		resultStringPart2 = append(resultStringPart2, val.stack[len(val.stack)-1])
	}
	log.Println("ResultStringPart1:", strings.Join(resultStringPart1, ""))
	log.Println("ResultStringPart2:", strings.Join(resultStringPart2, ""))
}

func calculateStackPart1(howMany int, fromWhere int, toWhere int, stackList1 []Stack) {
	fromStack := stackList1[fromWhere-1].stack
	toStack := stackList1[toWhere-1].stack

	reverseStack(fromStack)
	toStack = append(toStack, fromStack[:howMany]...)
	fromStack = fromStack[howMany:]
	reverseStack(fromStack)

	stackList1[fromWhere-1].stack = fromStack
	stackList1[toWhere-1].stack = toStack
}

func calculateStackPart2(howMany int, fromWhere int, toWhere int, stackList2 []Stack) {
	fromStack := stackList2[fromWhere-1].stack
	toStack := stackList2[toWhere-1].stack

	reverseStack(fromStack)
	var tempStack []string
	tempStack = append(tempStack, fromStack[:howMany]...)
	reverseStack(tempStack)
	toStack = append(toStack, tempStack...)
	fromStack = fromStack[howMany:]
	reverseStack(fromStack)

	stackList2[fromWhere-1].stack = fromStack
	stackList2[toWhere-1].stack = toStack
}

func reverseStack(stack []string) {
	for i := 0; i < len(stack)/2; i++ {
		stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
	}
}
