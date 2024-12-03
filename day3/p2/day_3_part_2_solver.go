package d3p2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day3Part2Solver struct {
	enabled bool
}

func (solver *Day3Part2Solver) Solve(lines []string) (string, error) {
	sum := 0
	solver.enabled = true
	for _, line := range lines {
		sum = sum + solver.getSumOfProductsFromLine(line)
	}
	answer := strconv.Itoa(sum)
	return answer, nil
}

func (solver *Day3Part2Solver) getSumOfProductsFromLine(line string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(line, -1)
	lineSum := 0
	product := 0
	for _, match := range matches {
		if match == "do()" {
			solver.enabled = true
			continue
		} else if match == "don't()" {
			solver.enabled = false
			continue
		}
		if !solver.enabled {
			continue
		}
		firstNumberString := strings.Split(strings.Split(match, "(")[1], ",")[0]
		secondNumberString := strings.Split(strings.Split(match, ",")[1], ")")[0]
		firstNumber, err := strconv.Atoi(firstNumberString)
		if err != nil {
			fmt.Println("error converting first number")
			panic(err)
		}
		secondNumber, err := strconv.Atoi(secondNumberString)
		if err != nil {
			fmt.Println("error converting second number")
			panic(err)
		}
		product = firstNumber * secondNumber
		lineSum += product
	}

	return lineSum
}
