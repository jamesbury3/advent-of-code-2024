package d3p1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day3Part1Solver struct{}

func (solver *Day3Part1Solver) Solve(lines []string) (string, error) {
	sum := 0
	for _, line := range lines {
		sum = sum + getSumOfProductsFromLine(line)
	}
	answer := strconv.Itoa(sum)
	return answer, nil
}

func getSumOfProductsFromLine(line string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(line, -1)
	lineSum := 0
	product := 0
	for _, match := range matches {
		firstNumberString := strings.Split(strings.Split(match, "(")[1], ",")[0]
		secondNumberString := strings.Split(strings.Split(match, ",")[1], ")")[0]
		firstNumber, err := strconv.Atoi(firstNumberString)
		if err != nil {
			fmt.Println("error converting first number to an int")
			panic(err)
		}
		secondNumber, err := strconv.Atoi(secondNumberString)
		if err != nil {
			fmt.Println("error converting second number to an int")
			panic(err)
		}
		product = firstNumber * secondNumber
		lineSum += product
	}

	return lineSum
}
