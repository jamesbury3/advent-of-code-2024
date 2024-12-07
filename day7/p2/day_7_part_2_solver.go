package d7p2

import (
	"math"
	"strconv"
	"strings"
)

type Day7Part2Solver struct{}

func (solver *Day7Part2Solver) Solve(lines []string) (string, error) {

	equations := make(map[float64][]float64)
	for _, line := range lines {
		solution := strings.Split(line, ": ")[0]
		inputsString := strings.Split(line, ": ")[1]
		inputsStringSplit := strings.Split(inputsString, " ")
		inputs := []float64{}
		for _, inputString := range inputsStringSplit {
			val, err := strconv.Atoi(inputString)
			if err != nil {
				return "", err
			}
			inputs = append(inputs, float64(val))
		}
		solutionInt, err := strconv.Atoi(solution)
		if err != nil {
			return "", err
		}
		equations[float64(solutionInt)] = inputs
	}

	var sum float64 = 0
	for solution, inputs := range equations {
		if solver.applyOperators(1, inputs[0], solution, inputs) {
			sum += solution
		}
	}

	answer := strconv.FormatFloat(sum, 'f', -1, 64)
	return answer, nil
}

func (solver *Day7Part2Solver) applyOperators(currIdx int, solution, target float64, inputs []float64) bool {
	if solution == target && currIdx == len(inputs) {
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
	} else if solver.applyOperators(currIdx+1, combineFloat64s(solution, curr), target, inputs) {
		return true
	}
	return false
}

func combineFloat64s(a, b float64) float64 {
	numDigits := 1
	temp := b
	for temp >= 10 {
		temp /= 10
		numDigits++
	}
	multiplier := math.Pow(10, float64(numDigits))
	return ((a) * multiplier) + float64(b)
}
