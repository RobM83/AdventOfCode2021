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
	nrDays := 256
	input, _ := readLines("input.txt")

	fishPerDay := make(map[int]int)
	newFishDays := 8
	//Create the days and set start counter
	for i := 0; i == newFishDays; i++ {
		fishPerDay[i] = 0
	}

	//Prefill map
	for _, line := range input {
		fishes := strings.Split(line, ",")
		for _, v := range fishes {
			v, _ := strconv.Atoi(v)
			fishPerDay[v]++
		}
	}

	//Loop over days
	fmt.Printf("D\t0\t1\t2\t3\t4\t5\t6\t7\t8\t\n")
	for d := 0; d <= nrDays; d++ {

		//DEBUG
		fmt.Printf("%d ", d)
		for i := 0; i <= newFishDays; i++ {
			fmt.Printf("\t%d", fishPerDay[i])
		}
		fmt.Printf("\n")

		zeroDayChange := fishPerDay[0]
		for i := 0; i <= newFishDays; i++ {
			fishPerDay[i] = fishPerDay[i+1]
		}
		//Births
		fishPerDay[newFishDays] = zeroDayChange
		fishPerDay[6] += zeroDayChange
	}

	//Count and print the total fish
	cnt := 0
	for i := 0; i < newFishDays; i++ {
		cnt += fishPerDay[i]
	}
	fmt.Println("Total fish: ", cnt)

}
