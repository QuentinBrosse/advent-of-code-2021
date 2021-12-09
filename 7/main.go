package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func computeFuel(positions []int, min, max int, useConstantRate bool) int {
	minFuel := math.MaxInt64

	for point := min; point <= max; point++ {
		currentFuel := 0
		for _, position := range positions {
			cost := int(math.Abs(float64(position - point)))
			if useConstantRate {
				currentFuel += cost
			} else {
				currentFuel += cost * (cost + 1) / 2
			}
		}

		if currentFuel < minFuel {
			minFuel = currentFuel
		}
	}

	return minFuel
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rawPositions := strings.Split(string(content), ",")
	positions, min, max := []int{}, math.MaxInt64, -1

	for _, rawPositions := range rawPositions {
		position, err := strconv.Atoi(rawPositions)
		if err != nil {
			log.Fatal(err)
		}

		positions = append(positions, position)
		if position < min {
			min = position
		}
		if position > max {
			max = position
		}
	}

	fmt.Println("Part 1:", computeFuel(positions, min, max, true))
	fmt.Println("Part 2:", computeFuel(positions, min, max, false))
}
