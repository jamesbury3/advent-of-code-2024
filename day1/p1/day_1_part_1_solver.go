package d1p1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1Part1Solver struct{}

func (solver *Day1Part1Solver) Solve(lines []string) (string, error) {

	left := []int{}
	right := []int{}

	for _, line := range lines {
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
