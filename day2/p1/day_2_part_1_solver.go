package d2p1

import (
	day2utils "advent-of-code-2024/day2"
	"fmt"
	"math"
	"strconv"
)

type Day2Part1Solver struct{}

func (solver *Day2Part1Solver) Solve(lines []string) (string, error) {

	safeReports := 0
	lineNums, err := day2utils.ConvertLinesToIntSlices(lines)
	if err != nil {
		fmt.Println("Error converting lines to int slices", err)
		panic(err)
	}

	for _, line := range lineNums {
		direction := getDirection(line)
		if isReportSafe(line, direction) {
			safeReports++
		}
	}

	answer := strconv.Itoa(safeReports)
	return answer, nil
}

func isReportSafe(line []int, direction int) bool {

	for i := 1; i < len(line); i++ {
		curr := line[i]
		prev := line[i-1]

		if (curr-prev)*direction <= 0 || math.Abs(float64(curr-prev)) > 3 {
			return false
		}
	}
	return true
}

func getDirection(line []int) int {
	prev := line[0]
	direction := 0
	for i := 1; i < len(line); i++ {
		if line[i] > prev {
			direction++
		} else if line[i] < prev {
			direction--
		}
		prev = line[i]
	}

	if direction > 0 {
		direction = 1
	} else if direction < 0 {
		direction = -1
	} else {
		direction = 0
	}

	return direction
}
