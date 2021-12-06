package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type boardNumber struct {
	Number  int
	Checked bool
}

type board [][]boardNumber

func sumUncheckedNumbers(board [][]boardNumber) int {
	sum := 0
	for _, raw := range board {
		for _, cell := range raw {
			if !cell.Checked {
				sum += cell.Number
			}
		}
	}

	return sum
}

func computeANumber(boards []board, number int, returnFirst bool) (int, []board) {
	lastResult := 0
	boardIndex := 0

	for _, board := range boards {
		boardComplete := false
		for i := 0; i < boardSize && !boardComplete; i++ {
			rawCompletion := 0
			colCompletion := 0
			for j := 0; j < boardSize; j++ {
				if board[i][j].Number == number {
					board[i][j].Checked = true
				}
				if board[i][j].Checked {
					rawCompletion++
				}

				if board[j][i].Number == number {
					board[j][i].Checked = true
				}
				if board[j][i].Checked {
					colCompletion++
				}
			}

			if rawCompletion == boardSize || colCompletion == boardSize {
				result := sumUncheckedNumbers(board) * number
				if returnFirst {
					return result, boards[:boardIndex]
				}
				lastResult = result
				boardComplete = true
			}
		}

		if !boardComplete {
			boards[boardIndex] = board
			boardIndex++
		}
	}

	return lastResult, boards[:boardIndex]
}

func part1(boards []board, numbers []int) (int, error) {
	for _, number := range numbers {
		if result, _ := computeANumber(boards, number, true); result != 0 {
			return result, nil
		}
	}

	return 0, fmt.Errorf("no completion found")
}

func part2(boards []board, numbers []int) (int, error) {
	lastResult := 0

	for _, number := range numbers {
		result, filteredBoards := computeANumber(boards, number, false)
		if result != 0 {
			lastResult = result
		}
		boards = filteredBoards
	}

	if lastResult == 0 {
		return 0, fmt.Errorf("no completion found")
	}

	return lastResult, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rawNumbers := strings.Split(scanner.Text(), ",")
	var numbers []int
	for _, rawNum := range rawNumbers {
		num, err := strconv.ParseInt(rawNum, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(num))
	}

	scanner.Scan() // skip empty line

	var boards []board
	var currentBoard board
	for i := 1; scanner.Scan(); i++ {
		rawNumbers := strings.Fields(scanner.Text())

		var numbers []boardNumber
		for _, rawNum := range rawNumbers {
			num, err := strconv.ParseInt(rawNum, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, boardNumber{Number: int(num)})
		}

		currentBoard = append(currentBoard, numbers)
		if i%boardSize == 0 {
			boards = append(boards, currentBoard)
			currentBoard = nil
			scanner.Scan() // skip empty line
			continue
		}
	}

	if result, err := part1(boards, numbers); err == nil {
		fmt.Println("Part 1:", result)
	} else {
		log.Fatal(err)
	}

	if result, err := part2(boards, numbers); err == nil {
		fmt.Println("Part 2:", result)
	} else {
		log.Fatal(err)
	}
}
