package utils

import (
	"fmt"
)

type Solver interface {
	Solve([]string) (string, error)
}

type DaySolver struct {
	DaySolverDelegate Solver
}

func (ds *DaySolver) CalculateAnswer() (string, error) {
	lines, err := readLines()
	if err != nil {
		fmt.Println("Error reading input lines:", err)
		return "", err
	}

	return ds.DaySolverDelegate.Solve(lines)
}
