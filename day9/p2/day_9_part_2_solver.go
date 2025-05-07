package d9p2

import (
	"slices"
	"strconv"
)

type Day9Part2Solver struct{}

func (solver *Day9Part2Solver) Solve(lines []string) (string, error) {

	line := lines[0]
	space := cleanInput(line)

	prev := ""
	count := 0
	visited := []string{}
	for i := len(space) - 1; i >= 0; i-- {
		curr := space[i]
		if prev == "" && curr != "." {
			prev = curr
			count = 0
		}
		if curr == prev {
			count += 1
			continue
		}
		if slices.Contains(visited, prev) {
			curr, prev, count = updateCounters(curr)
			continue
		}
		startingIndex := findStartingIndex(space, count)
		if startingIndex == -1 || startingIndex >= i {
			visited = append(visited, prev)
			curr, prev, count = updateCounters(curr)
			continue
		}
		for j := range count {
			space[startingIndex+j] = prev
			space[i+1+(j)] = "."
		}
		visited = append(visited, prev)
		curr, prev, count = updateCounters(curr)
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

func findStartingIndex(space []string, length int) int {
	count := 0
	for i, v := range space {
		if v == "." {
			count += 1
		} else {
			count = 0
		}
		if count == length {
			return i + 1 - length
		}
	}
	return -1
}

func cleanInput(line string) []string {
	fileId := -1
	space := []string{}

	for i, r := range line {
		size := runeToInt(r)
		var charToAdd string
		if i%2 == 0 {
			fileId += 1
			charToAdd = strconv.Itoa(fileId)
		} else {
			charToAdd = "."
		}
		for range size {
			space = append(space, charToAdd)
		}
	}

	return space
}

func runeToInt(r rune) int {
	return int(r - '0')
}

// updateCounters returns curr, prev, count based on whether curr is "."
func updateCounters(curr string) (string, string, int) {
	var prev string
	var count int
	if curr == "." {
		prev = ""
		count = 0
	} else {
		prev = curr
		count = 1
	}
	return curr, prev, count
}
