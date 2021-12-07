package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	median := getMedian(positions)
	steps := calcSteps(positions, median)

	fmt.Println(steps)
}

func calcSteps(positions []int, destination int) int {
	var steps int
	for _, pos := range positions {
		s := pos - destination
		if s < 0 {
			steps += -s //make it positive
		} else {
			steps += s
		}
	}
	return steps
}

func getMedian(positions []int) int {
	sort.Ints(positions)

	if middle := len(positions) / 2; middle%2 == 0 {
		return positions[len(positions)/2]
	} else {
		return (positions[middle-1] + positions[middle]) / 2
	}
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
