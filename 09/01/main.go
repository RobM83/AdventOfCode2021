package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	raster := getRaster(input)
	lowPoints := getLowPoints(raster)

	sumRisk := 0
	for _, lp := range lowPoints {
		sumRisk += lp + 1
	}

	fmt.Println(sumRisk)
}

func getLowPoints(raster [][]int) []int {
	var lowPoints []int
	for row, _ := range raster {
		for col, v := range raster[row] {
			lowPoint := true

			//Check top
			if row-1 > -1 {
				if raster[row-1][col] <= v {
					lowPoint = false
					continue
				}
			}
			//Check bottom
			if row+1 < len(raster) {
				if raster[row+1][col] <= v {
					lowPoint = false
					continue
				}
			}
			//Check left
			if col-1 > -1 {
				if raster[row][col-1] <= v {
					lowPoint = false
					continue
				}
			}
			//Check right
			if col+1 < len(raster[row]) {
				if raster[row][col+1] <= v {
					lowPoint = false
					continue
				}
			}

			if lowPoint {
				//fmt.Println(v)
				lowPoints = append(lowPoints, v)
			}
		}
	}

	return lowPoints
}

func getRaster(input []string) [][]int {
	var raster [][]int

	for y, strY := range input {
		raster = append(raster, []int{})
		for _, strX := range strY {
			x, _ := strconv.Atoi(string(strX))
			raster[y] = append(raster[y], x)
		}
	}

	return raster
}
