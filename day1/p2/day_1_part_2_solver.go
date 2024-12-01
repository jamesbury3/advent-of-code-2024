package d1p2

import (
	"fmt"
	"strconv"
	"strings"
)

type Day1Part2Solver struct{}

func (solver *Day1Part2Solver) Solve(lines []string) (string, error) {

	left := []int{}
	right := make(map[int]int)

	for _, line := range lines {
		split_line := strings.Split(line, "   ")

		leftVal, err := strconv.Atoi(split_line[0])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return "", nil
		}
		left = append(left, leftVal)
		rightVal, err := strconv.Atoi(split_line[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return "", nil
		}

		right[rightVal]++
	}

	sum := 0
	for _, val := range left {
		sum += val * right[val]
	}
	answer := strconv.Itoa(sum)
	return answer, nil
}
