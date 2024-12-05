package d5p2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var beforeMap map[string][]string

type Day5Part2Solver struct{}

func (solver *Day5Part2Solver) Solve(lines []string) (string, error) {

	beforeMap = make(map[string][]string)
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

	invalidManuals := [][]string{}
	for _, manual := range manuals {
		foundInvalidPage := false
		manualStillInvalid := true
		for manualStillInvalid {
			manualStillInvalid = false
			prevPages := []string{}
			for i, page := range manual {
				pagesWhichMustComeFirst := beforeMap[page]
				for _, pageWhichMustComeFirst := range pagesWhichMustComeFirst {
					if slices.Contains(manual, pageWhichMustComeFirst) && !slices.Contains(prevPages, pageWhichMustComeFirst) {
						originalIndex := slices.Index(manual, pageWhichMustComeFirst)
						shiftPageLeft(manual, originalIndex, i)
						foundInvalidPage = true
						manualStillInvalid = true
						break
					}
				}
				if manualStillInvalid {
					break
				}
				prevPages = append(prevPages, page)
				if i == len(manual)-1 {
					manualStillInvalid = false
				}
			}
		}

		if foundInvalidPage {
			invalidManuals = append(invalidManuals, manual)
		}
	}

	sum := 0
	for _, invalidManual := range invalidManuals {
		middle := invalidManual[len(invalidManual)/2]
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

func shiftPageLeft(pages []string, originalIndex, newIndex int) {
	for i := originalIndex; i > newIndex; i-- {
		pages[i], pages[i-1] = pages[i-1], pages[i]
	}
}
