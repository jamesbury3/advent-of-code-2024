package d4p2

import "strconv"

type Day4Part2Solver struct{}

func (solver *Day4Part2Solver) Solve(lines []string) (string, error) {

	letters := [][]rune{}
	for i, line := range lines {
		letters = append(letters, make([]rune, len(line)))
		for j, char := range line {
			letters[i][j] = rune(char)
		}
	}

	sum := 0

	for i := 1; i < len(letters)-1; i++ {
		for j := 1; j < len(letters[i])-1; j++ {
			curr := letters[i][j]
			if curr != 'A' {
				continue
			}

			if letters[i-1][j-1] == 'M' && letters[i-1][j+1] == 'M' &&
				letters[i+1][j+1] == 'S' && letters[i+1][j-1] == 'S' {
				sum++
			} else if letters[i-1][j-1] == 'S' && letters[i-1][j+1] == 'S' &&
				letters[i+1][j+1] == 'M' && letters[i+1][j-1] == 'M' {
				sum++
			} else if letters[i-1][j-1] == 'M' && letters[i-1][j+1] == 'S' &&
				letters[i+1][j+1] == 'S' && letters[i+1][j-1] == 'M' {
				sum++
			} else if letters[i-1][j-1] == 'S' && letters[i-1][j+1] == 'M' &&
				letters[i+1][j+1] == 'M' && letters[i+1][j-1] == 'S' {
				sum++
			}
		}
	}

	answer := strconv.Itoa(sum)
	return answer, nil
}
