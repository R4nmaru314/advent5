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
	var resultString []string

	stackList := []Stack{
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

		stackList = calculateStack(howMany, fromWhere, toWhere, stackList)
	}

	for _, val := range stackList {
		resultString = append(resultString, val.stack[len(val.stack)-1])
	}
	log.Println("ResultString:", strings.Join(resultString, ""))
}

func calculateStack(howMany int, fromWhere int, toWhere int, stackList []Stack) []Stack {
	fromStack := stackList[fromWhere-1].stack
	toStack := stackList[toWhere-1].stack

	reverseStack(fromStack)
	toStack = append(toStack, fromStack[:howMany]...)
	fromStack = fromStack[howMany:]
	reverseStack(fromStack)

	stackList[fromWhere-1].stack = fromStack
	stackList[toWhere-1].stack = toStack

	return stackList
}

func reverseStack(stack []string) {
	for i := 0; i < len(stack)/2; i++ {
		stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
	}
}
