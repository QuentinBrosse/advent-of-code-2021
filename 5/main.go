package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func biggerFirst(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

func computePoints(file io.ReadSeeker, withDiagonals bool) (int, error) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	gameMap := [][]int{}

	for scanner.Scan() {
		x1, y1, x2, y2 := 0, 0, 0, 0
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2); err != nil {
			return 0, err
		}

		xMax, xMin := biggerFirst(x1, x2)
		yMax, yMin := biggerFirst(y1, y2)

		// Build y axis.
		for len(gameMap) < yMax+1 {
			gameMap = append(gameMap, nil)
		}

		// Build x axis.
		for i := 0; i < len(gameMap); i++ {
			for len(gameMap[i]) < xMax+1 {
				gameMap[i] = append(gameMap[i], 0)
			}
		}

		if withDiagonals {
			for x1 != x2 || y1 != y2 {
				gameMap[y1][x1] += 1

				if x1 < x2 {
					x1 += 1
				} else if x1 > x2 {
					x1 -= 1
				}

				if y1 < y2 {
					y1 += 1
				} else if y1 > y2 {
					y1 -= 1
				}
			}
			gameMap[y1][x1] += 1
		} else {
			if y1 == y2 { // horizontal
				for x := xMin; x < xMax+1; x++ {
					gameMap[y1][x] += 1
				}
			}

			if x1 == x2 { // vertical
				for y := yMin; y < yMax+1; y++ {
					gameMap[y][x1] += 1
				}
			}
		}
	}

	// Count the points.
	points := 0
	for _, raw := range gameMap {
		for _, point := range raw {
			if point > 1 {
				points += 1
			}
		}
	}

	return points, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	if result, err := computePoints(file, false); err == nil {
		fmt.Println("Part 1:", result)
	} else {
		log.Fatal(err)
	}

	if result, err := computePoints(file, true); err == nil {
		fmt.Println("Part 2:", result)
	} else {
		log.Fatal(err)
	}
}
