package d9p1

import (
	"strconv"
)

type Day9Part1Solver struct{}

func (solver *Day9Part1Solver) Solve(lines []string) (string, error) {

	line := lines[0]
	fileId := -1
	space := []string{}

	for i, r := range line {
		size := runeToInt(r)
		var charToAdd string
		if i % 2 == 0 {
			fileId += 1
			charToAdd = strconv.Itoa(fileId)
		} else {
			charToAdd = "."
		}
		for range size {
			space = append(space, charToAdd)
		}
	}

	idx := 0
	for i := len(space)-1; i >= 0; i-- {
		if i == idx {
			break
		}
		if space[i] == "." {
			continue
		}

		for j := idx; j < len(space); j++ {
			if space[j] == "." {
				space[j] = space[i]
				space[i] = "."
				idx = j+1
				break
			}
		}
	}

	sum := 0
	for i, v := range space {
		if v == "." {
			continue
		}
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		sum += i * vInt
	}

	return strconv.Itoa(sum), nil
}

func runeToInt(r rune) int {
	return int(r - '0')
}
