package core

import (
	"advent-of-code-2024/utils"
	"fmt"
)

type Solver interface {
	Solve([]string) (string, error)
}

type DaySolver struct {
	DaySolverDelegate Solver
}

func (ds *DaySolver) CalculateAnswer() (string, error) {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Println("Error reading input lines:", err)
		return "", err
	}

	answer, err := ds.DaySolverDelegate.Solve(lines)
	if err != nil {
		fmt.Println("Error calculating answer:", err)
		return "", err
	}

	fmt.Println("Answer:", answer)
	return answer, nil
}
