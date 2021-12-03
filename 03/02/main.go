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

	var oxygen string
	var co2scrub string
	moreOnes := moreOneBitsOnPosition(input, 0)
	oxygen = findLastBinary(input, moreOnes, 0, true)[0]
	co2scrub = findLastBinary(input, moreOnes, 0, false)[0]

	oxygenDecimal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2scrubDecimal, _ := strconv.ParseInt(co2scrub, 2, 64)

	fmt.Println("Oxygen:", oxygenDecimal)
	fmt.Println("Co2scrub:", co2scrubDecimal)

	fmt.Println("Life Support rating:", oxygenDecimal*co2scrubDecimal)
}

//Place input and looking for 1s or 0s
func findLastBinary(input []string, ones bool, bitnumber int, mostOccuring bool) []string {
	var result []string
	if !mostOccuring {
		ones = !ones
	}
	for _, in := range input {
		value, _ := strconv.Atoi(string(in[bitnumber]))
		if value == 1 && ones {
			result = append(result, in)
		}
		if value == 0 && !ones {
			result = append(result, in)
		}
	}

	bitnumber++

	if len(result) == 1 {
		return result
	}

	return findLastBinary(result, moreOneBitsOnPosition(result, bitnumber), bitnumber, mostOccuring)
}

//Returns true if there are more or equeal ones on the given bit position
func moreOneBitsOnPosition(input []string, bitnumber int) bool {
	total := 0
	for _, in := range input {
		value, _ := strconv.Atoi(string(in[bitnumber]))
		total += value
	}

	return len(input)-total <= total
}
