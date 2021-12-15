package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Ocotopuse struct {
	Energy    int
	AsFlashed bool
}

func (o *Ocotopuse) Energyze() {
	if o.Energy < 10 && !o.AsFlashed {
		o.Energy++
	}
}

func (o *Ocotopuse) TryFlash() bool {
	if o.Energy == 10 && !o.AsFlashed {
		o.AsFlashed = true
		o.Energy = 0
		return true
	}
	return false
}

func (o *Ocotopuse) Reset() {
	o.AsFlashed = false
	o.Energy = 0
}

// TODO: Use a simple []*Ocotopuse to remove some nested for loops.
// As the matrix size is known, all the X and Y can be calculated.
type Matrix [][]*Ocotopuse

func (m Matrix) Get(x, y int) *Ocotopuse {
	if y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
		return nil
	}

	return m[y][x]
}

var flashMask = [][2]int{
	{-1, -1},
	{+0, -1},
	{+1, -1},
	{-1, +0},
	{+1, +0},
	{-1, +1},
	{+0, +1},
	{+1, +1},
}

func (m *Matrix) Flash(x, y int) {
	for _, maskPoint := range flashMask {
		if o := m.Get(x+maskPoint[0], y+maskPoint[1]); o != nil {
			o.Energyze()
		}
	}
}

func compute(matrix Matrix, findFullFlashStepMode bool) (flashes int) {
	steps := 100
	if findFullFlashStepMode {
		steps = math.MaxInt
	}
	for step := 1; step <= steps; step++ {
		for _, raw := range matrix {
			for _, octo := range raw {
				octo.Energyze()
			}
		}

		for {
			asFlashed := false

			for y, raw := range matrix {
				for x, octo := range raw {
					if flashed := octo.TryFlash(); flashed {
						asFlashed = true
						flashes++
						matrix.Flash(x, y)
					}
				}
			}

			if !asFlashed {
				break
			}
		}

		for _, raw := range matrix {
			for _, octo := range raw {
				if octo.AsFlashed {
					octo.Reset()
				}
			}
		}

		if findFullFlashStepMode {
			isFullOfZeros := true
			for _, raw := range matrix {
				for _, octo := range raw {
					if octo.Energy != 0 {
						isFullOfZeros = false
					}
				}
			}

			if isFullOfZeros {
				return step
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(file)
	}

	var matrix1 Matrix
	var matrix2 Matrix

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ocoto1 []*Ocotopuse
		var ocoto2 []*Ocotopuse
		for _, rawNum := range scanner.Text() {
			num, err := strconv.Atoi(string(rawNum))
			if err != nil {
				log.Fatal(err)
			}
			ocoto1 = append(ocoto1, &Ocotopuse{Energy: num})
			ocoto2 = append(ocoto2, &Ocotopuse{Energy: num})
		}
		matrix1 = append(matrix1, ocoto1)
		matrix2 = append(matrix2, ocoto2)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Part 1:", compute(matrix1, false))
	fmt.Println("Part 2:", compute(matrix2, true))
}
