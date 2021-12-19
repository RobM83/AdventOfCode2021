package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	id      string
	nextHop []string
}

type Paths struct {
	p [][]string
}

var paths [][]string

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

	caves := getCavesAndPaths(input)

	p := new(Paths)
	p.findPath(caves["start"], caves, []string{})

	// for _, p := range p.p {
	// 	fmt.Println(p)
	// }

	fmt.Println(len(p.p))
}

func (p *Paths) findPath(c Cave, caves map[string]Cave, path []string) {

	if c.id == "end" {
		path = append(path, c.id)
		endpath := append(make([]string, 0, len(path)), path...) //Make a new one
		p.p = append(p.p, endpath)
		return
	}

	path = append(path, c.id)

	for _, nextHop := range c.nextHop {
		if isLower(nextHop) && twiceHitLower(path) && existInStringSlice(nextHop, path) { //Don't go back to small caves (lowercase)
			continue
		}
		p.findPath(caves[nextHop], caves, path)
	}

	return
}

func getCavesAndPaths(input []string) map[string]Cave {
	caves := make(map[string]Cave)
	for _, line := range input {
		c := strings.Split(line, "-") // A-B
		caves = addCavesAndHop(c[0], c[1], caves)
		caves = addCavesAndHop(c[1], c[0], caves)
	}
	return caves
}

func addCavesAndHop(c, h string, caves map[string]Cave) map[string]Cave {
	if h == "start" {
		return caves
	}
	if _, ok := caves[c]; ok {
		//Cave already exists, extend paths
		nh := caves[c].nextHop
		if !existInStringSlice(h, nh) {
			nh = append(nh, h)
			caves[c] = Cave{c, nh}
		}
	} else {
		caves[c] = Cave{c, []string{h}}
	}

	return caves
}

func twiceHitLower(path []string) bool {
	for _, c := range path {
		hit := 0

		if !isLower(c) || c == "start" || c == "end" { //Only lowercase
			continue
		}

		for _, c2 := range path {
			if c == c2 {
				hit++
			}
			if hit == 2 {
				return true
			}
		}
	}
	return false
}

func existInStringSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

func isLower(s string) bool {
	if s == strings.ToLower(s) {
		return true
	}
	return false
}
