package d7p1

import (
	"strconv"
	"strings"
)

type Day7Part1Solver struct{}

func (solver *Day7Part1Solver) Solve(lines []string) (string, error) {

	equations := make(map[int][]int)
	for _, line := range lines {
		solution := strings.Split(line, ": ")[0]
		inputsString := strings.Split(line, ": ")[1]
		inputsStringSplit := strings.Split(inputsString, " ")
		inputs := []int{}
		for _, inputString := range inputsStringSplit {
			val, err := strconv.Atoi(inputString)
			if err != nil {
				return "", err
			}
			inputs = append(inputs, val)
		}
		solutionInt, err := strconv.Atoi(solution)
		if err != nil {
			return "", err
		}
		equations[solutionInt] = inputs
	}

	sum := 0
	for solution, inputs := range equations {
		if solver.applyOperators(1, inputs[0], solution, inputs) {
			sum += solution
		}
	}

	answer := strconv.Itoa(sum)
	return answer, nil
}

func (solver *Day7Part1Solver) applyOperators(currIdx, solution, target int, inputs []int) bool {
	if solution == target {
		return true
	}
	if currIdx >= len(inputs) {
		return false
	}
	curr := inputs[currIdx]
	if solver.applyOperators(currIdx+1, solution+curr, target, inputs) {
		return true
	} else if solver.applyOperators(currIdx+1, solution*curr, target, inputs) {
		return true
	}
	return false
}
