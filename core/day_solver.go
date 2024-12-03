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

func (ds *DaySolver) CalculateAnswerFromInputExample() (string, error) {
	lines, err := utils.ReadLinesFromInputExample()
	if err != nil {
		fmt.Println("Error reading input lines:", err)
		return "", err
	}

	answer, err := ds.DaySolverDelegate.Solve(lines)
	if err != nil {
		fmt.Println("Error calculating answer:", err)
		return "", err
	}

	fmt.Println("Answer for Example:", answer)
	return answer, nil
}

func (ds *DaySolver) CalculateAnswerFromInput() (string, error) {
	lines, err := utils.ReadLinesFromInput()
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
