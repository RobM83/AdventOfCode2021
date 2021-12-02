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
	const (
		horizontal = 0
		depth      = 1
		aim        = 2
	)

	//input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	input, _ := readLines("input.txt")
	position := []int{0, 0, 0} // Horizontal, Depth

	for _, v := range input {
		movement := strings.Split(v, " ")
		switch movement[0] {
		case "forward":
			forward, _ := strconv.Atoi(movement[1])
			position[horizontal] += forward
			position[depth] += position[aim] * forward
		case "up":
			up, _ := strconv.Atoi(movement[1])
			position[aim] -= up
		case "down":
			down, _ := strconv.Atoi(movement[1])
			position[aim] += down
		}
	}
	fmt.Printf("Horizontal=%d, Depth=%d, Position=%d\n", position[horizontal], position[depth], position[horizontal]*position[depth])
}
