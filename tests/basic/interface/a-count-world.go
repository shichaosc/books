package main

import (
	"bufio"
	"fmt"
	"strings"
)

type line interface {
	Count() int
}

func (c countWorld) Count() int {
	scanner := bufio.NewScanner(strings.NewReader(string(c)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func (c countLine) Count() int {
	scanner := bufio.NewScanner(strings.NewReader(string(c)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

type countWorld string
type countLine string

func count(l line) int {
	return l.Count()
}

func main() {
	input := "s a\n2"
	fmt.Println(count(countWorld(input)))
	fmt.Println(count(countLine(input)))
}
