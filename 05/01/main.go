package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	x         = 0
	y         = 1
	fieldSize = 1000 //For test (10x10)
)

type inputLine struct {
	p1 [2]int //x,y p1
	p2 [2]int //x,y p2
}

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

	inputLines := readInput(input)
	field := drawField(inputLines)
	crossing := getOverlappingLinePoints(field)

	//Debug
	// for _, line := range field {
	// 	fmt.Println(line)
	// }

	fmt.Println(crossing)

}

func getOverlappingLinePoints(field [fieldSize][fieldSize]int) int {
	crossing := 0
	for y := 0; y < fieldSize; y++ {
		for x := 0; x < fieldSize; x++ {
			if field[y][x] >= 2 {
				crossing++
			}
		}
	}
	return crossing
}

func drawField(inputLines []inputLine) [fieldSize][fieldSize]int {
	var field [fieldSize][fieldSize]int //Y, X

	for _, inputLine := range inputLines {
		var small, big int
		//Only draw horizontal and vertical lines.
		if inputLine.p1[y] == inputLine.p2[y] { //Horizontal
			if inputLine.p1[x] < inputLine.p2[x] {
				small = inputLine.p1[x]
				big = inputLine.p2[x]
			} else {
				small = inputLine.p2[x]
				big = inputLine.p1[x]
			}
			for i := small; i <= big; i++ {
				field[inputLine.p1[y]][i] += 1
			}
		}
		if inputLine.p1[x] == inputLine.p2[x] { //Vertical
			if inputLine.p1[y] < inputLine.p2[y] {
				small = inputLine.p1[y]
				big = inputLine.p2[y]
			} else {
				small = inputLine.p2[y]
				big = inputLine.p1[y]
			}
			for i := small; i <= big; i++ {
				field[i][inputLine.p1[x]] += 1
			}
		}
	}
	return field
}

func readInput(lines []string) []inputLine {
	var result []inputLine
	for _, in := range lines {

		inputLine := new(inputLine)
		line := strings.Split(in, " ")

		coords := strings.Split(line[0], ",")
		inputLine.p1[x] = strToInt(coords[x])
		inputLine.p1[y] = strToInt(coords[y])
		coords = strings.Split(line[2], ",")
		inputLine.p2[x] = strToInt(coords[x])
		inputLine.p2[y] = strToInt(coords[y])

		result = append(result, *inputLine)
	}
	return result
}

func strToInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
