package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func main() {
	//input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	input, _ := readLines("input.txt")

	var baseline int
	result := 0
	for i, v := range input {
		if i == 0 {
			baseline = v
			continue
		} else {
			if v > baseline {

				result++
			}
		}
		baseline = v
	}
	fmt.Println(result)
}
