package d2p2

import (
	day2utils "advent-of-code-2024/day2"
	"fmt"
	"math"
	"slices"
	"strconv"
)

type Day2Part2Solver struct{}

func (solver *Day2Part2Solver) Solve(lines []string) (string, error) {
	safeReports := 0
	lineNums, err := day2utils.ConvertLinesToIntSlices(lines)
	if err != nil {
		fmt.Println("Error converting lines to int slices", err)
		panic(err)
	}

	for _, line := range lineNums {
		if isReportSafeInAnyOrder(line) {
			safeReports++
		}
	}

	answer := strconv.Itoa(safeReports)
	return answer, nil
}

func isReportSafeInAnyOrder(line []int) bool {
	if isReportSafeWithOneRemoval(line, 1, 1, nil, false) {
		return true
	} else if isReportSafeWithOneRemoval(line, 1, -1, nil, false) {
		return true
	}
	slices.Reverse(line)
	if isReportSafeWithOneRemoval(line, 1, 1, nil, false) {
		return true
	}
	return isReportSafeWithOneRemoval(line, 1, -1, nil, false)
}

func isReportSafeWithOneRemoval(line []int, idx, direction int,
	backupPrev *int, removalHasOccurred bool) bool {

	if idx >= len(line) {
		return true
	}

	curr := line[idx]
	var prev *int
	if backupPrev == nil {
		prev = &line[idx-1]
	} else {
		prev = backupPrev
	}

	if (curr-*prev)*direction <= 0 || math.Abs(float64(curr-*prev)) > 3 {
		if removalHasOccurred {
			return false
		} else {
			return isReportSafeWithOneRemoval(line, idx+1, direction, prev, true)
		}
	}
	return isReportSafeWithOneRemoval(line, idx+1, direction, nil, removalHasOccurred)
}
