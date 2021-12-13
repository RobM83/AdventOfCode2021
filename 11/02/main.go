package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Octopus struct {
	value   int
	flashed bool
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

	grid := buildGrid(input)

	gridsize := len(grid) * len(grid[0])

	for i := 0; i > -1; i++ {
		//fmt.Println("Step: ", i+1)
		grid = increaseByOne(grid)
		//fmt.Println(numberOfFlashes(grid))
		if numberOfFlashes(grid) == gridsize {
			fmt.Println("All lights ON step:", i+1)
			break
		}
		//printGrid(grid)
		grid = resetFlashed(grid)

	}
}

func resetFlashed(grid [][]Octopus) [][]Octopus {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col].flashed = false
		}
	}
	return grid
}

func withinBoundaries(grid [][]Octopus, row, col int) bool {
	if row < 0 || row > len(grid)-1 {
		return false
	}
	if col < 0 || col > len(grid[row])-1 {
		return false
	}
	return true
}

func flashImpact(grid [][]Octopus, row, col int) [][]Octopus {
	if withinBoundaries(grid, row, col) {
		if !grid[row][col].flashed {
			grid[row][col].value++
			if grid[row][col].value == 10 {
				grid[row][col].value = 0
				grid[row][col].flashed = true //Action!
				grid = findFlashImpactedOctopuses(grid, row, col)
			}
		}
	}
	return grid
}

func findFlashImpactedOctopuses(grid [][]Octopus, row, col int) [][]Octopus {
	grid = flashImpact(grid, row-1, col) //Top
	grid = flashImpact(grid, row+1, col) //Bottom
	grid = flashImpact(grid, row, col-1) //Left
	grid = flashImpact(grid, row, col+1) //Right

	grid = flashImpact(grid, row-1, col-1) //TopLeft
	grid = flashImpact(grid, row-1, col+1) //TopRight
	grid = flashImpact(grid, row+1, col-1) //DownLeft
	grid = flashImpact(grid, row+1, col+1) //DownRight

	return grid
}

func numberOfFlashes(grid [][]Octopus) int {
	var flashes int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].flashed {
				flashes++
			}
		}
	}
	return flashes
}

func increaseByOne(grid [][]Octopus) [][]Octopus {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if !grid[row][col].flashed {
				grid[row][col].value++
				if grid[row][col].value == 10 {
					grid[row][col].value = 0
					grid[row][col].flashed = true
					grid = findFlashImpactedOctopuses(grid, row, col)
				}
			}
		}
	}
	return grid
}

func buildGrid(input []string) [][]Octopus {
	var grid [][]Octopus
	for row, line := range input {
		//Every line is a row
		grid = append(grid, []Octopus{})
		for _, char := range line {
			//Every character is a octopus
			o := Octopus{
				value:   strToInt(string(char)),
				flashed: false,
			}
			grid[row] = append(grid[row], o)
		}

	}
	return grid
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func printGrid(grid [][]Octopus) {
	for row := 0; row < len(grid); row++ {
		fmt.Println(append(grid[row]))
	}
}
