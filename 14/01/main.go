package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	code := input[0] //Assuming always line 0
	codeMap := getCodeMap(input)
	charCount := initCharCount(code)

	steps := 10
	result := code
	for i := 0; i < steps; i++ {
		result = pairInsertion(result, codeMap, charCount, 0)
		//fmt.Printf("Step %d: %s\n", i+1, result)
	}

	for char, count := range charCount {
		fmt.Printf("%s: %d\n", char, count)
	}

	high, low := findHighLow(charCount)
	fmt.Printf("High: %d, Low: %d, Diff: %d\n", high, low, high-low)
}

func findHighLow(charCount map[string]int) (high int, low int) {
	high = 0
	low = 0
	for _, count := range charCount {
		if low == 0 {
			low = count
		}
		if count > high {
			high = count
		}
		if count < low {
			low = count
		}
	}
	return high, low
}

func initCharCount(code string) map[string]int {
	charCount := make(map[string]int)
	for _, char := range code {
		charCount = increaseCharCount(charCount, string(char))
	}
	return charCount
}

func increaseCharCount(charCount map[string]int, char string) map[string]int {
	if _, ok := charCount[char]; ok {
		charCount[char]++
	} else {
		charCount[char] = 1
	}
	return charCount
}

func pairInsertion(code string, codeMap map[string]string, charCount map[string]int, currentPos int) string {
	if currentPos+2 > len(code) {
		return code
	}

	sub := "x"
	pair := code[currentPos : currentPos+2]
	if sub, ok := codeMap[pair]; ok {
		code = code[:currentPos+1] + sub + code[currentPos+len(sub):]
		charCount = increaseCharCount(charCount, sub)
	}

	return pairInsertion(code, codeMap, charCount, currentPos+len(sub)+1)
}

func getCodeMap(input []string) map[string]string {
	codeMap := make(map[string]string)
	for _, line := range input {
		if strings.Contains(line, "->") {
			coderesult := strings.Split(line, "->")
			code := strings.TrimSpace(coderesult[0])
			result := strings.TrimSpace(coderesult[1])
			codeMap[code] = result
		}
	}
	return codeMap
}
