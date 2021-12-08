package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func run(timers [9]uint, days uint) uint {
	for dayNumber := uint(0); dayNumber < days; dayNumber++ {
		zeros := timers[0]

		for i := 0; i < len(timers)-1; i++ {
			timers[i] = timers[i+1]
		}

		timers[8] = zeros
		timers[6] += zeros
	}

	sum := uint(0)
	for _, timer := range timers {
		sum += timer
	}
	return sum
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rawTimers := strings.Split(string(content), ",")
	timers := [9]uint{}
	for _, rawTimer := range rawTimers {
		timer, err := strconv.Atoi(rawTimer)
		if err != nil {
			log.Fatal(err)
		}
		timers[timer] += 1
	}

	fmt.Println("Part 1:", run(timers, 80))
	fmt.Println("Part 2:", run(timers, 256))
}
