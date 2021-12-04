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

type Card struct {
	Numbers [][]int
	Hits    [][]int
}

func main() {
	input, _ := readLines("input.txt")

	numbersDrawn := strings.Split(input[0], ",")
	cards := getCards(input)
	var lastWinningCard *Card
	var lastWinningNumber int

	for _, number := range numbersDrawn {
		nr, _ := strconv.Atoi(number)
		drawNumber(nr, cards)
		winningCard := checkBingo(cards)
		for winningCard != nil {
			lastWinningNumber = nr
			lastWinningCard = winningCard
			cards = removeWinningCard(winningCard, cards)
			winningCard = checkBingo(cards) //Check again (could be multiple!)
		}
	}
	fmt.Printf("Solution: %d\n", getSumUnmarkedNumbers(lastWinningCard)*lastWinningNumber)
}

func removeWinningCard(card *Card, cards []*Card) []*Card {
	for i, c := range cards {
		if c == card {
			cards = append(cards[:i], cards[i+1:]...)
			break
		}
	}
	return cards
}

func getSumUnmarkedNumbers(card *Card) int {
	var unmarkedNumbers []int
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if card.Hits[row][col] == 0 {
				unmarkedNumbers = append(unmarkedNumbers, card.Numbers[row][col])
			}
		}
	}
	return sum(unmarkedNumbers)
}

func drawNumber(number int, cards []*Card) {
	//Mark hits on cards
	for _, card := range cards {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if card.Numbers[row][col] == number {
					card.Hits[row][col] = 1
				}
			}
		}
	}
}

//Check bingo return winning row/col and unmarked number of the winning card.
func checkBingo(cards []*Card) *Card {
	//Mark hits on cards
	for _, card := range cards {
		//Check rows
		for row := 0; row < 5; row++ {
			hits := 0
			for col := 0; col < 5; col++ {
				hits += card.Hits[row][col]
				if hits == 5 {
					return card
				}
			}
		}

		//Check cols
		for col := 0; col < 5; col++ {
			hits := 0
			for row := 0; row < 5; row++ {
				hits += card.Hits[row][col]
				if hits == 5 {
					return card
				}
			}
		}
	}

	return nil
}

func getCards(input []string) []*Card {
	var cards []*Card
	row := 0
	var c *Card

	for cnt, line := range input {
		if cnt == 0 { //First line contains the drawn numbers
			continue
		}

		newCard := false
		if len(line) == 0 {
			newCard = true

		}

		//Second line and onward contains bingo cards
		if newCard {
			if c != nil {
				cards = append(cards, c)
			}
			c = new(Card)
			c.Numbers = make([][]int, 5) //Assuming rows == 5
			c.Hits = make([][]int, 5)
			for i := 0; i < 5; i++ { //Initialize cols
				c.Numbers[i] = make([]int, 5) //Assuming cols == 5
				c.Hits[i] = make([]int, 5)
			}
			row = 0
			continue //Empty line, skip
		}

		for col := 0; col < 5; col++ {
			numbers := strings.Split(strings.Replace(strings.TrimLeft(line, " "), "  ", " ", -1), " ")
			c.Numbers[row][col], _ = strconv.Atoi(numbers[col])
			c.Hits[row][col] = 0
		}
		row++
	}

	//Add the last card as well range doesn't include last empty line.
	if c != nil {
		cards = append(cards, c)
	}

	return cards
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
