package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//Represents a node
type node struct {
	x, y int
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

	//Input to array for easy graph parsing
	grid := getInputArray(input)

	//fmt.Println(grid)
	distance := map[node]int{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			distance[node{x, y}] = math.MaxInt
		}
	}
	distance[node{0, 0}] = 0 // Start node
	visited := map[node]bool{}
	currentNode := node{0, 0}

	for {
		//fmt.Println(currentNode)

		for _, n := range getNeighbours(currentNode, grid) {
			if visited[n] {
				continue
			}
			newDistance := distance[currentNode] + grid[n.y][n.x]
			if newDistance < distance[n] {
				distance[n] = newDistance
			}
		}
		visited[currentNode] = true //cuurent ?

		if visited[node{len(grid) - 1, len(grid) - 1}] { //End
			fmt.Println(distance[node{len(grid) - 1, len(grid) - 1}])
			break
		}

		minDist := math.MaxInt
		minNode := node{len(grid), len(grid[0])}

		for p, v := range distance {
			if !visited[p] && v < minDist {
				minDist = v
				minNode = p
			}
		}

		currentNode = minNode
	}

}

func getNeighbours(n node, grid [][]int) []node {
	output := []node{}

	if n.x+1 < len(grid[n.y]) { //right
		output = append(output, node{n.x + 1, n.y})
	}
	if n.x-1 > 0 { //Left
		output = append(output, node{n.x - 1, n.y})
	}
	if n.y+1 < len(grid) { //Bottom
		output = append(output, node{n.x, n.y + 1})
	}
	if n.y-1 > 0 { //Top
		output = append(output, node{n.x, n.y - 1})
	}

	return output
}

func getInputArray(input []string) [][]int { //Y, X !
	rows := len(input)
	cols := len(input[0])

	arr := make([][]int, rows) //Y
	for y := 0; y < rows; y++ {
		arr[y] = make([]int, cols) //X
	}

	for y, line := range input {
		for x, char := range line {
			arr[y][x], _ = strconv.Atoi(string(char))
		}
	}

	return arr
}
