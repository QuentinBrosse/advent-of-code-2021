package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// For simplicity, let's load the entire input in memory
// (x2 in the part 2) and assuming that it will not be too big.

func countOccurences(values []string, bitPosition int) (int, int) {
	zeros, ones := 0, 0

	for _, line := range values {
		switch c := line[bitPosition]; c {
		case '0':
			zeros++
		case '1':
			ones++
		default:
			// Assume the lines contains only 0 and 1.
		}
	}

	return zeros, ones
}

func part1(values []string) (int, error) {
	gammaRate, episilonRate := 0, 0
	lineLength := len(values[0])

	for n := 0; n < lineLength; n++ {
		zeros, ones := countOccurences(values, n)

		if ones > zeros {
			gammaRate += 1
		} else {
			episilonRate += 1
		}

		if n != lineLength-1 {
			gammaRate = gammaRate << 1
			episilonRate = episilonRate << 1
		}
	}

	return gammaRate * episilonRate, nil
}

func computeRating(values []string, bitPriotiry [2]rune) (int, error) {
	ratingValues := make([]string, len(values))
	copy(ratingValues, values)

	for n := 0; n < len(values[0]) && len(ratingValues) > 1; n++ {
		zeros, ones := countOccurences(ratingValues, n)

		bitCriteria := bitPriotiry[0]
		if ones >= zeros {
			bitCriteria = bitPriotiry[1]
		}

		i := 0
		for _, line := range ratingValues {
			if rune(line[n]) == bitCriteria {
				ratingValues[i] = line
				i++
			}
		}
		ratingValues = ratingValues[:i]
	}

	rating, err := strconv.ParseInt(ratingValues[0], 2, 64)
	return int(rating), err
}

func part2(values []string) (int, error) {
	oxyGenRating, err := computeRating(values, [2]rune{'0', '1'})
	if err != nil {
		return 0, err
	}

	co2ScrubRating, err := computeRating(values, [2]rune{'1', '0'})
	if err != nil {
		return 0, err
	}

	return oxyGenRating * co2ScrubRating, nil
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	values := strings.Split(string(content), "\n")

	if result, err := part1(values); err == nil {
		fmt.Println("Part 1:", result)
	} else {
		log.Fatal(err)
	}

	if result, err := part2(values); err == nil {
		fmt.Println("Part 2:", result)
	} else {
		log.Fatal(err)
	}
}
