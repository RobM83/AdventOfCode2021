package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	DaysLeft int //daysLeft for new fish
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
	nrDays := 80
	input, _ := readLines("input.txt")
	fish := getFish(input)
	for d := 1; d <= nrDays; d++ {
		fish = dayPassed(fish)
		fmt.Printf("Day %d: ", d)
		for _, f := range fish {
			fmt.Printf("%d ", f.DaysLeft)
		}
		fmt.Printf(" (%d) \n", len(fish))
	}
}

func dayPassed(fish []*Fish) []*Fish {
	newFish := []*Fish{}
	for _, f := range fish {
		f.DaysLeft--
		if f.DaysLeft == -1 {
			f.DaysLeft = 6
			var nf = Fish{DaysLeft: 8}
			newFish = append(newFish, &nf)
		}
	}
	fish = append(fish, newFish...)
	return fish
}

func getFish(input []string) []*Fish {
	var fish []*Fish
	for _, line := range input {
		fishes := strings.Split(line, ",")
		for _, v := range fishes {
			f := new(Fish)
			f.DaysLeft, _ = strconv.Atoi(v)
			fish = append(fish, f)
		}
	}
	return fish
}
