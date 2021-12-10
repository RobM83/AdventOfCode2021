package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	input, _ := readLines("input.txt")
	breakingChars := []string{}
	acScore := []int{}

	for _, line := range input {
		result, breakingChar := findPair(line, 0, []string{})
		if result == "corrupted" {
			breakingChars = append(breakingChars, breakingChar...)
		}
		if result == "incomplete" {
			stack := autoComplete(breakingChar)
			score := autoCompletionScore(stack)
			acScore = append(acScore, score)
		}
		//fmt.Printf("Line %d (%s) \t%s(%s)\n", i, line, result, breakingChar)
	}

	sort.Ints(acScore)

	fmt.Println("Breaking chars: ", breakingChars)
	fmt.Println("Score syntax error: ", scoringChars(breakingChars))
	fmt.Println("Score auto completion: ", acScore[len(acScore)/2])
}

func autoCompletionScore(stack []string) int {
	score := 0
	for _, char := range stack {
		score = score * 5
		switch char {
		case ")":
			score += 1
		case "]":
			score += 2
		case "}":
			score += 3
		case ">":
			score += 4
		}
	}
	return score
}

func autoComplete(input []string) []string {
	var stack []string
	for i := len(input) - 1; i >= 0; i-- {
		switch input[i] {
		case "(":
			stack = append(stack, ")")
		case "[":
			stack = append(stack, "]")
		case "{":
			stack = append(stack, "}")
		case "<":
			stack = append(stack, ">")
		}
	}
	return stack
}

func scoringChars(chars []string) int {
	score := 0
	for _, char := range chars {
		switch char {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}
	return score
}

func findPair(line string, pos int, stack []string) (string, []string) {
	if len(line) != pos { // if not at end of line
		if isOpening(string(line[pos])) { // Opening char
			stack = append(stack, string(line[pos]))
		} else { // Closing char - must form pair, otherwise corrupted
			if isPair(stack[len(stack)-1], string(line[pos])) {
				stack = stack[:len(stack)-1]
			} else {
				return "corrupted", []string{string(line[pos])}
			}
		}

		// End of line and empty stack
		if len(line) == pos && len(stack) == 0 {
			return "correct", []string{}
		}

		return findPair(line, pos+1, stack)
	}

	//Stack is not 0, but end of line reached
	return "incomplete", stack
}

func isOpening(char string) bool {
	if char == "(" || char == "[" || char == "{" || char == "<" {
		return true
	}
	return false
}

func isPair(open, close string) bool {
	pairs := make(map[string]string)
	pairs["("] = ")"
	pairs["["] = "]"
	pairs["{"] = "}"
	pairs["<"] = ">"

	if pairs[open] == close {
		return true
	}
	return false
}
