package d1p2

import (
	"advent-of-code-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

type D1P2 struct {
	lines []string
}

func (solver *D1P2) Solve() (string, error) {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	solver.lines = lines

	left := []int{}
	right := make(map[int]int)

	for _, line := range lines {
		leftVal, err := strconv.Atoi(strings.Split(line, "   ")[0])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return "", nil
		}
		left = append(left, leftVal)
		rightVal, err := strconv.Atoi(strings.Split(line, "   ")[1])
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
