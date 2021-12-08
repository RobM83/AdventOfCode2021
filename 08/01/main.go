package main

import (
	"bufio"
	"fmt"
	"os"
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
	input, _ := readLines("input.txt")

	cnt := 0
	for _, line := range input {
		outputLine := strings.Split(line, "|")
		output := strings.Split(outputLine[1], " ")
		for _, o := range output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
