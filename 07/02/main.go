package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	positions := getPositions(input)

	average := getAverage(positions)
	steps := calcSteps(positions, average)

	fmt.Println(steps)
}

func calcSteps(positions []int, destination int) int {
	var steps int
	for _, pos := range positions {
		s := pos - destination
		if s < 0 {
			s = -s //make it positive
		}
		for i := 1; i <= s; i++ {
			steps += i
		}
	}
	return steps
}

func getAverage(positions []int) int {
	var sum int
	for _, pos := range positions {
		sum += pos
	}

	avg := float64(sum) / float64(len(positions))

	//return int(math.Round(avg))
	return int(avg)
}

func getPositions(input []string) []int {
	var positions []int
	for _, line := range input {
		strPos := strings.Split(line, ",")
		for _, pos := range strPos {
			intPos, _ := strconv.Atoi(pos)
			positions = append(positions, intPos)
		}
	}
	return positions
}
