package day2utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertLinesToIntSlices(lines []string) ([][]int, error) {
	lineNums := [][]int{}
	for _, line := range lines {
		lineNum := []int{}
		splitLine := strings.Split(line, " ")
		for _, num := range splitLine {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error converting %s to int\n", num)
				return nil, err
			}
			lineNum = append(lineNum, val)
		}
		lineNums = append(lineNums, lineNum)
	}
	return lineNums, nil
}
