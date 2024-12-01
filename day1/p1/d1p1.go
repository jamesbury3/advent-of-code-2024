package d1p1

import (
	"advent-of-code-2025/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type D1P1 struct {
	lines []string
}

func (solver *D1P1) Solve() (string, error) {

	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	solver.lines = lines

	left := []int{}
	right := []int{}

	for _, line := range lines {
		// split each line and convert to int
		leftVal, err := strconv.Atoi(strings.Split(line, "   ")[0])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return "", nil
		}
		rightVal, err := strconv.Atoi(strings.Split(line, "   ")[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return "", nil
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(right[i] - left[i])))
	}

	answer := strconv.Itoa(sum)
	return answer, nil
}
