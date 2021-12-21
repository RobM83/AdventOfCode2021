package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	char  = 0
	pair1 = 1
	pair2 = 2
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
	pairMap := initPairMap(input)

	steps := 40

	charCount := initCharCount(code)
	pairMap = pairCalculation(pairMap, codeMap, steps, charCount)

	for char, count := range charCount {
		fmt.Printf("%s: %d\n", char, count)
	}

	high, low := findHighLow(charCount)
	fmt.Printf("High: %d, Low: %d, Diff: %d\n", high, low, high-low)
}

func initPairMap(input []string) map[string]int {
	pairMap := make(map[string]int)
	code := input[0]
	for i := 0; i < len(string(code))-1; i++ {
		pair := string(code)[i : i+2]
		pairMap[pair]++
	}
	return pairMap
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
		charCount = increaseCharCount(charCount, string(char), 1)
	}
	return charCount
}

func increaseCharCount(charCount map[string]int, char string, count int) map[string]int {
	if _, ok := charCount[char]; ok {
		charCount[char] += count
	} else {
		charCount[char] = 1
	}
	return charCount
}

func pairCalculation(pairMap map[string]int, codeMap map[string]string, steps int, charCount map[string]int) map[string]int {
	for step := 0; step < steps; step++ {
		fmt.Printf("Step: %d - %s\n", step+1, time.Now().String())

		updateMap := make(map[string]int)
		for pair, count := range pairMap {
			updateMap[string(pair[0])+codeMap[pair]] += count
			updateMap[codeMap[pair]+string(pair[1])] += count
			charCount = increaseCharCount(charCount, codeMap[pair], count)
		}

		pairMap = updateMap
	}

	return pairMap
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
