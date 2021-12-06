package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(file io.Reader) (int, error) {
	horizontal := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		switch direction {
		case "forward":
			horizontal += dist
		case "down":
			depth += dist
		case "up":
			depth -= dist
		}
	}

	return horizontal * depth, nil
}

func part2(file io.Reader) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		switch direction {
		case "forward":
			horizontal += dist
			depth += aim * dist
		case "down":
			aim += dist
		case "up":
			aim -= dist
		}
	}

	return horizontal * depth, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// display err (in all files)
	if value, err := part1(file); err == nil {
		fmt.Println("Part 1:", value)
	} else {
		log.Fatal(err)
	}

	file.Seek(0, 0)

	if value, err := part2(file); err == nil {
		fmt.Println("Part 2:", value)
	} else {
		log.Fatal(err)
	}
}
