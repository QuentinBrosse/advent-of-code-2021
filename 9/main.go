package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	Value   int
	Visited bool
}

type Matrix struct {
	matrix [][]*Point
}

type Coordinate struct {
	X int
	Y int
}

func (m *Matrix) Sizes() (int, int) {
	return len(m.matrix[0]), len(m.matrix)
}

func (m *Matrix) getPoint(x, y int) *Point {
	sizeX, sizeY := m.Sizes()
	if y < 0 || y >= sizeY || x < 0 || x >= sizeX {
		return nil
	}

	return m.matrix[y][x]
}

func (m *Matrix) Value(x, y int) *int {
	if p := m.getPoint(x, y); p != nil {
		return &p.Value
	}
	return nil
}

func (m *Matrix) Point(x, y int) *Point {
	return m.getPoint(x, y)
}

func (m *Matrix) SetAsVisited(x, y int) {
	if p := m.getPoint(x, y); p != nil {
		p.Visited = true
	}
}

func (m *Matrix) ResetVisited() {
	for _, raw := range m.matrix {
		for _, point := range raw {
			point.Visited = false
		}
	}
}

func (m *Matrix) DeepestCoordinates() (cords []*Coordinate) {
	sizeX, sizeY := m.Sizes()
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			height := *m.Value(x, y)
			candidates := []*int{
				m.Value(x-1, y-1),
				m.Value(x, y-1),
				m.Value(x+1, y-1),
				m.Value(x-1, y),
				m.Value(x+1, y),
				m.Value(x-1, y+1),
				m.Value(x, y+1),
				m.Value(x+1, y+1),
			}

			if isSmallest(height, candidates) {
				cords = append(cords, &Coordinate{x, y})
			}
		}
	}

	return cords
}

var digDirections = []*Coordinate{
	{0, -1},
	{+1, 0},
	{0, +1},
	{-1, 0},
}

func (m *Matrix) DigBasin(x, y int) int {
	current := *m.Value(x, y)
	m.SetAsVisited(x, y)

	basinSize := 1
	for _, direction := range digDirections {
		point := m.Point(x+direction.X, y+direction.Y)
		if point != nil && !point.Visited && point.Value > current && point.Value < 9 {
			basinSize += m.DigBasin(x+direction.X, y+direction.Y)
		}
	}
	return basinSize
}

func isSmallest(target int, candidates []*int) bool {
	for _, candidate := range candidates {
		if candidate != nil {
			if target > *candidate {
				return false
			}
		}
	}
	return true
}

func part1(matrix *Matrix) int {
	totalRiskLevel := 0
	for _, coordinate := range matrix.DeepestCoordinates() {
		totalRiskLevel += 1 + *matrix.Value(coordinate.X, coordinate.Y)
	}
	return totalRiskLevel
}

func part2(matrix *Matrix) int {
	var basinSizes []int
	for _, coordinate := range matrix.DeepestCoordinates() {
		matrix.ResetVisited()
		basinSizes = append(basinSizes, matrix.DigBasin(coordinate.X, coordinate.Y))
	}

	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	sum := 1
	for _, size := range basinSizes[:3] {
		sum *= size
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(file)
	}

	matrix := &Matrix{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []*Point
		for _, rawNum := range scanner.Text() {
			num, err := strconv.Atoi(string(rawNum))
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, &Point{Value: num})
		}
		matrix.matrix = append(matrix.matrix, line)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Part 1:", part1(matrix))
	fmt.Println("Part 2:", part2(matrix))
}
