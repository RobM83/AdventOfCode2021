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
	avgmod := 3
	windows := 4 // A, B, C, D
	averages := []int{0, 0, 0, 0}

	//var baseline int
	result := 0
	for i, _ := range input {
		//Make sure to stop when no window can be created anymore
		if i+avgmod > len(input) {
			break
		}
		//fmt.Printf("Window = %d & average = %d \n", i%windows, input[i]+input[i+1]+input[i+2])
		averages[i%windows] = input[i] + input[i+1] + input[i+2]

		if i > 0 { //At first there is nothing to compare
			//fmt.Printf("average[%d](%d) > average[%d](%d)\n", i%windows, averages[i%windows], (i-1)%windows, averages[(i-1)%windows])
			if averages[i%windows] > averages[(i-1)%windows] {
				result++
			}
		}
	}
	fmt.Println(result)
}
