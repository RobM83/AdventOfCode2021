package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Represents the graph
type graph struct {
	nodes map[node][]edge
}

//Represents a node
type node struct {
	x, y int
}

//Represents edges and cost
type edge struct {
	n    node
	cost int
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
	input, _ := readLines("test.txt")

	//Input to array for easy graph parsing
	arr := getInputArray(input)

	//Build a graph
	g := newGrap()
	for y := range arr {
		for x := range arr[y] {
			// Check top, bottom, left, right
			//top

			if x+1 < len(arr[y]) { //right
				g.addEdge(node{x, y}, node{x + 1, y}, arr[y][x+1])
			}
			if x-1 > 0 { //Left
				g.addEdge(node{x, y}, node{x - 1, y}, arr[y][x-1])
			}
			if y+1 < len(arr) { //Bottom
				g.addEdge(node{x, y}, node{x, y + 1}, arr[y+1][x])
			}
			if y-1 > 0 { //Top
				g.addEdge(node{x, y}, node{x, y - 1}, arr[y-1][x])
			}
		}
	}

	fmt.Println("breakkie breakkie")
	// S topleft
	//Go right/down/up/left

}

func newGrap() *graph {
	return &graph{
		nodes: make(map[node][]edge),
	}
}

func (g *graph) addEdge(n1, n2 node, cost int) {
	g.nodes[n1] = append(g.nodes[n1], edge{n2, cost})
	g.nodes[n2] = append(g.nodes[n2], edge{n1, cost})
}

func (g *graph) getEdges(n node) []edge {
	return g.nodes[n]
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
