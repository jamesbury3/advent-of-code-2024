package d4p1

import (
	"strconv"
)

type XmasIterator struct {
	xmas  []rune
	index int
}

func (xmasIterator *XmasIterator) next() rune {
	val := xmasIterator.xmas[xmasIterator.index]
	xmasIterator.index++
	if xmasIterator.index >= len(xmasIterator.xmas) {
		xmasIterator.index = 0
	}
	return val
}

func (xmasIterator *XmasIterator) reset() {
	xmasIterator.index = 0
}

var xmasIterator XmasIterator = XmasIterator{[]rune{'X', 'M', 'A', 'S'}, 0}

type Day4Part1Solver struct{}

func (solver *Day4Part1Solver) Solve(lines []string) (string, error) {

	letters := [][]rune{}
	for i, line := range lines {
		letters = append(letters, make([]rune, len(line)))
		for j, char := range line {
			letters[i][j] = rune(char)
		}
	}

	sum := 0

	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters[i]); j++ {
			if letters[i][j] != 'X' && letters[i][j] != 'S' {
				continue
			}

			if checkLeft(i, j, letters) {
				sum++
			}
			if checkUp(i, j, letters) {
				sum++
			}
			if checkLeftUpDiagonal(i, j, letters) {
				sum++
			}
			if checkUpRightDiagonal(i, j, letters) {
				sum++
			}
		}
	}

	answer := strconv.Itoa(sum)
	return answer, nil
}

func checkLeft(row, col int, letters [][]rune) bool {
	if col < 3 {
		return false
	}
	foundXmas := true
	xmasIterator.reset()
	for j := col; j > col-4; j-- {
		currLetter := letters[row][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	if foundXmas {
		return true
	}
	foundXmas = true
	xmasIterator.reset()
	for j := col - 3; j <= col; j++ {
		currLetter := letters[row][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	return foundXmas
}

func checkLeftUpDiagonal(row, col int, letters [][]rune) bool {
	if row < 3 || col < 3 {
		return false
	}
	foundXmas := true
	xmasIterator.reset()
	for i, j := row, col; i > row-4 && j > col-4; i, j = i-1, j-1 {
		currLetter := letters[i][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	if foundXmas {
		return true
	}
	foundXmas = true
	xmasIterator.reset()
	for i, j := row-3, col-3; i <= row && j <= col; i, j = i+1, j+1 {
		currLetter := letters[i][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	return foundXmas
}

func checkUp(row, col int, letters [][]rune) bool {
	if row < 3 {
		return false
	}
	foundXmas := true
	xmasIterator.reset()
	for i := row; i > row-4; i-- {
		currLetter := letters[i][col]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	if foundXmas {
		return true
	}
	foundXmas = true
	xmasIterator.reset()
	for i := row - 3; i <= row; i++ {
		currLetter := letters[i][col]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	return foundXmas
}

func checkUpRightDiagonal(row, col int, letters [][]rune) bool {
	if row < 3 || col >= len(letters[0])-3 {
		return false
	}
	foundXmas := true
	xmasIterator.reset()
	for i, j := row, col; i > row-4 && j < col+4; i, j = i-1, j+1 {
		currLetter := letters[i][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}
	if foundXmas {
		return true
	}
	foundXmas = true
	xmasIterator.reset()
	for i, j := row-3, col+3; i <= row && j >= col; i, j = i+1, j-1 {
		currLetter := letters[i][j]
		itNext := xmasIterator.next()
		if currLetter != itNext {
			foundXmas = false
			break
		}
	}

	return foundXmas
}
