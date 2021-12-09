package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Point struct {
	row int
	col int
}

func main() {
	input, _ := readLines("input.txt")

	raster := getRaster(input)
	lowPoints := getLowPoints(raster)
	basins := []int{}

	for _, lowpoint := range lowPoints {
		basins = append(basins, getBasinSize(lowpoint, raster))
	}

	sort.Ints(basins)
	solution := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
	fmt.Println(solution)

}

func getBasinSize(p Point, raster [][]int) int {
	basinPoints := getBasinPoints(p, []Point{}, raster)
	//fmt.Println(basinPoints)

	return len(basinPoints)
}

func getBasinPoints(p Point, points []Point, raster [][]int) []Point {
	//Check if point is already in points, then return
	var exists func(Point) bool
	exists = func(p Point) bool {
		for _, point := range points {
			if point == p {
				return true
			}
		}
		return false
	}

	//Add point to points
	points = append(points, p)

	//New point, check direction
	if p.row-1 > -1 && raster[p.row-1][p.col] < 9 && !exists(Point{p.row - 1, p.col}) {
		points = getBasinPoints(Point{p.row - 1, p.col}, points, raster)
	}

	if p.col+1 < len(raster[p.row]) && raster[p.row][p.col+1] < 9 && !exists(Point{p.row, p.col + 1}) {
		points = getBasinPoints(Point{p.row, p.col + 1}, points, raster)
	}

	if p.col-1 > -1 && raster[p.row][p.col-1] < 9 && !exists(Point{p.row, p.col - 1}) {
		points = getBasinPoints(Point{p.row, p.col - 1}, points, raster)
	}

	if p.row+1 < len(raster) && raster[p.row+1][p.col] < 9 && !exists(Point{p.row + 1, p.col}) {
		points = getBasinPoints(Point{p.row + 1, p.col}, points, raster)
	}

	return points
}

func getLowPoints(raster [][]int) []Point {
	var points []Point
	for row, _ := range raster {
		for col, v := range raster[row] {
			point := Point{row, col}
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
				points = append(points, point)
			}
		}
	}

	return points
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
