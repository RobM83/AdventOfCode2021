package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	T  = 0 //Top
	TR = 1 //Top Right
	BR = 2 //Bottom Right
	B  = 3 //Bottom
	BL = 4 //Bottom Left
	TL = 5 //Top Left
	M  = 6 //Middle
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

	var ssDisplay map[int][]string
	answer := 0

	for _, line := range input {
		strResult := ""
		outputLine := strings.Split(line, "|")
		ssDisplay = getWireMap(strings.Split(outputLine[0], " "))
		//printDisplay(ssDisplay)

		for _, o := range strings.Split(outputLine[1], " ") {
			if o == "" {
				continue
			}
			strResult += strconv.Itoa(getNumber(ssDisplay, o))
		}
		intResult, _ := strconv.Atoi(strResult)
		answer += intResult
		//fmt.Println(strResult)
	}
	fmt.Println("Answer: ", answer)
}

func getNumber(wireMap map[int][]string, input string) int {
	switch len(input) {
	case 2: // 1 (TR, BR)
		return 1
	case 3: // 7 (T, TL, TR)
		return 7
	case 4: // 4 (TR, BR, M, TL)
		return 4
	case 5: //
		return getWireNumber(wireMap, input)
	case 6: // 0, 6, 9
		return getWireNumber(wireMap, input)
	case 7: // 8
		return 8
	}

	return -1
}

func getWireNumber(wiremap map[int][]string, input string) int {
	result := []string{"0", "0", "0", "0", "0", "0", "0"}
	for i := 0; i < 7; i++ {
		for _, v := range strings.Split(input, "") {
			if v == wiremap[i][0] {
				result[i] = "1"
			}
		}
	}

	strResult := strings.Join(result, "")
	switch strResult {
	//T, TR, BR, B, BL, TL, M
	case "1111110":
		return 0
	case "1101101":
		return 2
	case "1111001":
		return 3
	case "1011011":
		return 5
	case "1011111":
		return 6
	case "1111011":
		return 9
	}

	return -1
}

func getWireMap(input []string) map[int][]string {
	l6 := []string{} //Buffer to hold 6pos (0,6,9) - for calculating TR, M, BL
	ssDisplay := initDisplay()

	for _, o := range input {
		ssDisplay = retractUnique(ssDisplay, o)
		if len(o) == 6 {
			l6 = addIfUnique(l6, o)
		}
	}

	// Find TR, M, BL -> values, retract from other
	possible := make(map[string]int)
	for _, l := range l6 {
		for _, c := range strings.Split(l, "") {
			possible[c]++
		}
	}

	//Retract from others
	for k := range possible { //don't care about order
		if possible[k] == 2 {
			for i := 0; i < 7; i++ {
				if !(i == TR || i == BL || i == M) { //Only[TR, BL, M]
					ssDisplay[i] = removeCharsFromInput(ssDisplay[i], strings.Split(k, ""))
					ssDisplay = checkForOneChar(ssDisplay)
				}
			}
		}
	}

	return ssDisplay
}

func addIfUnique(list []string, input string) []string {
	for _, v := range list {
		if v == input {
			return list
		}
	}
	return append(list, input)
}

//Retract posibilities 1,4,7
func retractUnique(ssDisplay map[int][]string, input string) map[int][]string {
	switch len(input) {
	case 2: // 1 (TR, BR)
		ssDisplay[TR] = []string{string(input[0]), string(input[1])}
		ssDisplay[BR] = []string{string(input[0]), string(input[1])}
		//Remove this from all other positions
		for i := 0; i < 7; i++ { // we can
			if i != TR && i != BR { //Skip[TR, BR]
				ssDisplay[i] = removeCharsFromInput(ssDisplay[i], strings.Split(input, ""))
			}
		}
	case 3: // 7 (T, TL, TR)
		//DAB - AB
		for i := 0; i < 7; i++ {
			if i == T || i == TR || i == BR { //Only[TR, BR]
				ssDisplay[i] = keepCharsFromInput(ssDisplay[i], strings.Split(input, ""))
			} else {
				ssDisplay[i] = removeCharsFromInput(ssDisplay[i], strings.Split(input, ""))
			}
		}
	case 4: // 4 (TR, BR, M, TL)
		for i := 0; i < 7; i++ {
			if i == TR || i == BR || i == M || i == TL { //Only[TR, BR, M, TL]
				ssDisplay[i] = keepCharsFromInput(ssDisplay[i], strings.Split(input, ""))
			} else {
				ssDisplay[i] = removeCharsFromInput(ssDisplay[i], strings.Split(input, ""))
			}
		}
	}

	//Always check if only one char is left, can be removed from other positions
	ssDisplay = checkForOneChar(ssDisplay)

	return ssDisplay
}

func checkForOneChar(ssDisplay map[int][]string) map[int][]string {
	for i := 0; i < 7; i++ {
		if len(ssDisplay[i]) == 1 {
			for j := 0; j < 7; j++ {
				if i == j {
					continue
				}
				ssDisplay[j] = removeCharsFromInput(ssDisplay[j], ssDisplay[i])
			}
		}
	}
	return ssDisplay
}

func keepCharsFromInput(input []string, chars []string) []string {
	result := []string{}
	for _, o := range input {
		for _, c := range chars {
			if o == c {
				result = append(result, c)
			}
		}
	}
	return result
}

func removeCharsFromInput(input []string, chars []string) []string {
	for i, o := range input {
		for _, c := range chars {
			if o == c {
				input = append(input[:i], input[i+1:]...)
				return removeCharsFromInput(input, chars)
			}
		}
	}
	return input
}

func initDisplay() map[int][]string {
	ssDisplay := make(map[int][]string)
	for i := 0; i < 7; i++ {
		ssDisplay[i] = []string{"a", "b", "c", "d", "e", "f", "g"}
	}
	return ssDisplay
}

func printDisplay(ssDisplay map[int][]string) {
	for i := 0; i < 7; i++ {
		fmt.Println(ssDisplay[i])
	}
}
