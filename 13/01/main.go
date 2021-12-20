package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FoldInstruction struct {
	axis string //Y or X
	line int    //Line to fold
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
	paper := getPaper(input)
	foldInstructions := getFoldInstructions(input)

	//printPaper(paper)
	//fmt.Println("")

	paper = foldPaper(paper, foldInstructions)
	//nrDots := countDots(paper)

}

func countDots(paper [][]string) int {
	var count int
	for _, row := range paper {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}
	return count
}

func foldPaper(paper [][]string, foldInstruction []FoldInstruction) [][]string {
	var foldedPaper = paper
	for fc, fi := range foldInstruction {
		fc++
		if fi.axis == "y" {
			foldedPaper = foldY(foldedPaper, fi)
		} else {
			foldedPaper = foldX(foldedPaper, fi)
		}
		//printPaper(foldedPaper)
		fmt.Printf("Foldcount %d: %d\n\n", fc, countDots(foldedPaper))
	}
	return foldedPaper
}

func foldY(paper [][]string, fi FoldInstruction) [][]string {
	newLength := fi.line - 1                                  // -1 == folding line
	foldedPaper := createEmptyPaper(len(paper[0]), newLength) //X blijft ongewijzigd

	count := 0
	newPosY := 0
	for row := range paper {
		if row > fi.line {
			count++
			newPosY = row - (2 * count)
		}
		for col := range paper[row] {
			if row == fi.line {
				continue //Foldline
			}
			if row < fi.line {
				foldedPaper[row][col] = paper[row][col]
			} else {
				if paper[row][col] == "#" {
					foldedPaper[newPosY][col] = paper[row][col]
				}
			}
		}
	}

	return foldedPaper
}

func foldX(paper [][]string, fi FoldInstruction) [][]string {
	newWidth := fi.line - 1                                 // -1 == folding line
	foldedPaper := createEmptyPaper(newWidth, len(paper)-1) //y blijft ongewijzigd

	newPosX := 0
	for row := range paper {
		count := 0
		for col := range paper[row] {
			if col > fi.line {
				count++
				newPosX = col - (2 * count)
			}
			if col == fi.line {
				continue //Foldline
			}
			if col < fi.line {
				foldedPaper[row][col] = paper[row][col]
			} else {
				if paper[row][col] == "#" {
					foldedPaper[row][newPosX] = paper[row][col]
				}
			}
		}
	}

	return foldedPaper
}

func printPaper(paper [][]string) {
	for _, row := range paper {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func getFoldInstructions(input []string) []FoldInstruction {
	var foldInstructions []FoldInstruction

	for _, line := range input {
		//Read till fold instruction
		if len(line) != 0 && line[0] == 'f' { //Fold instruction
			var fi FoldInstruction
			instruction := strings.Split(line, " ")
			instruction = strings.Split(instruction[2], "=")
			fi.axis = instruction[0]
			fi.line = strToInt(instruction[1])
			foldInstructions = append(foldInstructions, fi)
		}
	}

	return foldInstructions
}

func getPaper(input []string) [][]string { //Y, X !
	rows, cols := getPaperSize(input)
	paper := createEmptyPaper(rows, cols)

	for _, line := range input {
		//Read till fold instruction
		if len(line) == 0 {
			return paper
		}

		coords := strings.Split(line, ",")
		x := strToInt(coords[0])
		y := strToInt(coords[1])

		paper[y][x] = "#"
	}

	return paper
}

func createEmptyPaper(x, y int) [][]string {
	//Array starts with 0, so create one bigger
	paper := make([][]string, y+1) //rows
	for i := 0; i < y+1; i++ {
		paper[i] = make([]string, x+1) //cols
	}

	//Place dots (mainly debug)
	for rows := range paper {
		for cols := range paper[rows] {
			paper[rows][cols] = "."
		}
	}

	return paper
}

func getPaperSize(input []string) (int, int) {
	x := 0
	y := 1
	paperSizeX := 0
	paperSizeY := 0

	for _, line := range input {
		//Read till fold instruction
		if len(line) == 0 {
			return paperSizeX, paperSizeY
		}

		coords := strings.Split(line, ",")
		if strToInt(coords[x]) > paperSizeX {
			paperSizeX = strToInt(coords[x])
		}
		if strToInt(coords[y]) > paperSizeY {
			paperSizeY = strToInt(coords[y])
		}
	}

	return paperSizeX, paperSizeY
}

func strToInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}
