package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Display struct {
	OutputValues  []string
	PaternsPerLen map[int][]string
}

type Segments struct {
	Top         string
	TopLeft     string
	TopRight    string
	Middle      string
	BottomLeft  string
	BottomRight string
	Bottom      string
}

var digitsSegmentsRef = map[string]rune{
	"abcefg":  '0',
	"cf":      '1',
	"acdeg":   '2',
	"acdfg":   '3',
	"bcdf":    '4',
	"abdfg":   '5',
	"abdefg":  '6',
	"acf":     '7',
	"abcdefg": '8',
	"abcdfg":  '9',
}

func ifXTimes(x int, value []string) (uniq string) {
	ref := make(map[rune]int)
	for _, v := range value {
		for _, char := range v {
			ref[char] += 1
		}
	}
	for char, count := range ref {
		if count == x {
			uniq += string(char)
		}
	}
	return uniq
}

func sortStr(value string) string {
	sorted := []rune(value)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return string(sorted)
}

func convertDigitSegments(signature string, segments *Segments) string {
	converted := ""
	for _, seg := range signature {
		switch seg {
		case 'a':
			converted += segments.Top
		case 'b':
			converted += segments.TopLeft
		case 'c':
			converted += segments.TopRight
		case 'd':
			converted += segments.Middle
		case 'e':
			converted += segments.BottomLeft
		case 'f':
			converted += segments.BottomRight
		case 'g':
			converted += segments.Bottom
		}
	}
	return sortStr(converted)
}

func part1(displays []*Display) int {
	count := 0
	for _, display := range displays {
		for _, outputValue := range display.OutputValues {
			switch len(outputValue) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return count
}

func part2(displays []*Display) (int, error) {
	sum := 0

	for _, display := range displays {
		paterns := display.PaternsPerLen

		s := &Segments{}
		s.Top = ifXTimes(1, append(paterns[2], paterns[3][0]))
		s.BottomLeft = ifXTimes(1, append(paterns[5], paterns[4][0]))
		s.Middle = ifXTimes(4, append(paterns[5], paterns[4][0]))
		s.TopLeft = ifXTimes(2, append(paterns[5], paterns[4][0]))
		s.TopRight = ifXTimes(2, append(paterns[6], s.Middle+s.BottomLeft))
		s.BottomRight = ifXTimes(2, append(paterns[2], paterns[3][0], s.TopRight))
		s.Bottom = ifXTimes(1, append(paterns[7], s.Top+s.TopRight+s.TopLeft+s.Middle+s.BottomLeft+s.BottomRight))

		digitsSegments := make(map[string]rune)
		for signatureRef, digit := range digitsSegmentsRef {
			convertedSignature := convertDigitSegments(signatureRef, s)
			digitsSegments[convertedSignature] = digit
		}

		var digits []rune
		for _, outputValue := range display.OutputValues {
			digits = append(digits, digitsSegments[sortStr(outputValue)])
		}

		number, err := strconv.Atoi(string(digits))
		if err != nil {
			return 0, err
		}
		sum += number
	}

	return sum, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(file)
	}

	scanner := bufio.NewScanner(file)
	var dislays []*Display
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		signalPaterns, outputValues := strings.Fields(parts[0]), strings.Fields(parts[1])

		paternsPerLen := make(map[int][]string)
		for _, signalPatern := range signalPaterns {
			l := len(signalPatern)
			paternsPerLen[l] = append(paternsPerLen[l], signalPatern)
		}

		dislays = append(dislays, &Display{
			OutputValues:  outputValues,
			PaternsPerLen: paternsPerLen,
		})

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Part 1:", part1(dislays))
	if result, err := part2(dislays); err == nil {
		fmt.Println("Part 2:", result)
	} else {
		log.Fatal(err)
	}
}
