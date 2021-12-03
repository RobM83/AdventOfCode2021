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
	//input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	input, _ := readLines("input.txt")

	inputLength := 12 // in example 5 (adjust slices)
	totals := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, in := range input {
		for i := 0; i < inputLength; i++ {
			value, _ := strconv.Atoi(string(in[i]))
			totals[i] += value
		}
	}

	gammaRate := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	epsilonRate := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	maxResult := len(input)
	for i, bit := range totals {
		if bit >= (maxResult / 2) {
			gammaRate[i] = 1
		} else {
			epsilonRate[i] = 1
		}
	}

	grDec := binarySliceToDec(gammaRate)
	erDec := binarySliceToDec(epsilonRate)
	fmt.Printf("Gammarate is %v(%d)\nEpsilonrate is %v(%d)\n", gammaRate, grDec, epsilonRate, erDec)
	fmt.Printf("Power consumption: %d\n", grDec*erDec)

}

func binarySliceToDec(input []int) int {
	result := 0
	binary := []int{2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1} //Example is only 5 bits
	for i, _ := range input {
		if input[i] == 1 {
			result += binary[i]
		}
	}
	return result
}
