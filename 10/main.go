package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Stack []rune

func (s *Stack) Push(value rune) {
	*s = append(*s, value)
}

func (s *Stack) Pop() rune {
	value, newFilo := (*s)[len(*s)-1], (*s)[:len(*s)-1]
	*s = newFilo
	return value
}

func (s Stack) Len() int {
	return len(s)
}

func check(line string) (isIncorect bool, score int, stack Stack) {
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			stack.Push(char)
			continue
		}

		switch {
		case char == ')' && stack.Pop() != '(':
			return true, 3, nil
		case char == ']' && stack.Pop() != '[':
			return true, 57, nil
		case char == '}' && stack.Pop() != '{':
			return true, 1197, nil
		case char == '>' && stack.Pop() != '<':
			return true, 25137, nil
		}
	}

	return false, 0, stack
}

func part1(file io.ReadSeeker) (score int) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isIncorect, lineScore, _ := check(scanner.Text()); isIncorect {
			score += lineScore
		}
	}
	return
}

func part2(file io.ReadSeeker) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	var scores []int

	for scanner.Scan() {
		isIncorect, _, stack := check(scanner.Text())
		if isIncorect {
			continue // ignore incorrect line
		}

		score := 0
		for stack.Len() > 0 {
			switch stack.Pop() {
			case '(':
				score = score*5 + 1
			case '[':
				score = score*5 + 2
			case '{':
				score = score*5 + 3
			case '<':
				score = score*5 + 4
			}
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	return scores[len(scores)/2]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(file)
	}

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}
