package d5p1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day5Part1Solver struct{}

func (solver *Day5Part1Solver) Solve(lines []string) (string, error) {

	beforeMap := make(map[string][]string)
	nextLoopStart := 0
	for i, line := range lines {
		if len(line) == 0 {
			nextLoopStart = i + 1
			break
		}
		splitLine := strings.Split(line, "|")
		first := splitLine[0]
		second := splitLine[1]
		if len(beforeMap[second]) == 0 {
			beforeMap[second] = make([]string, 0)
		}
		beforeMap[second] = append(beforeMap[second], first)
	}

	manuals := [][]string{}
	for i := nextLoopStart; i < len(lines); i++ {
		line := lines[i]
		splitLine := strings.Split(line, ",")
		manuals = append(manuals, splitLine)
	}

	validManuals := [][]string{}
	for _, manual := range manuals {
		prevPages := []string{}
		foundInvalidPage := false
		for _, page := range manual {
			pagesWhichMustComeFirst := beforeMap[page]
			for _, pageWhichMustComeFirst := range pagesWhichMustComeFirst {
				if slices.Contains(manual, pageWhichMustComeFirst) && !slices.Contains(prevPages, pageWhichMustComeFirst) {
					foundInvalidPage = true
					break
				}
			}
			if foundInvalidPage {
				break
			} else {
				prevPages = append(prevPages, page)
			}
		}
		if !foundInvalidPage {
			validManuals = append(validManuals, manual)
		}
	}

	sum := 0
	for _, validManual := range validManuals {
		middle := validManual[len(validManual)/2]
		middleInt, err := strconv.Atoi(middle)
		if err != nil {
			fmt.Println("Error converting middle to int")
			return "", err
		}
		sum += middleInt
	}
	answer := strconv.Itoa(sum)
	return answer, nil
}
