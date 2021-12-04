package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func part1(file io.Reader) (int, error) {
	lastValue := math.MaxInt
	increasedCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if value > lastValue {
			increasedCounter++
		}

		lastValue = value
	}

	return increasedCounter, nil
}

const WindowSize = 3

func part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	increasedCounter := 0
	previousSum := math.MaxInt
	window := []int{}

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if len(window) == WindowSize {
			_, window = window[0], window[1:]
		}

		if len(window) < WindowSize {
			window = append(window, value)
		}

		if len(window) == WindowSize {
			newSum := 0
			for _, v := range window {
				newSum += v
			}

			if newSum > previousSum {
				increasedCounter++
			}
			previousSum = newSum
		}
	}

	return increasedCounter, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if value, err := part1(file); err == nil {
		fmt.Println("Part 1:", value)
	}
	file.Seek(0, 0)
	if value, err := part2(file); err == nil {
		fmt.Println("Part 2:", value)
	}
}
