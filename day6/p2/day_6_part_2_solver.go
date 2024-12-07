package d6p2

import (
	"fmt"
	"strconv"
)

type Day6Part2Solver struct{}

func (solver *Day6Part2Solver) Solve(lines []string) (string, error) {

	grid := [][]string{}
	originalGrid := [][]string{}
	row, col := 0, 0
	startRow, startCol := 0, 0
	for i := 0; i < len(lines); i++ {
		gridRow := []string{}
		originalGridRow := []string{}
		for j := 0; j < len(lines[i]); j++ {
			gridRow = append(gridRow, string(lines[i][j]))
			originalGridRow = append(originalGridRow, string(lines[i][j]))
			if string(lines[i][j]) == "^" {
				row, col = i, j
				startRow, startCol = i, j
			}
		}
		grid = append(grid, gridRow)
		originalGrid = append(originalGrid, originalGridRow)
	}

	newObstacles := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == startRow && j == startCol {
				continue
			}
			totalSteps := 0
			grid[i][j] = "#"
			row, col = startRow, startCol
			for !reachedEnd(row, col, grid) {
				row, col = move(row, col, grid)
				totalSteps++
				if totalSteps == len(grid)*len(grid[0]) {
					newObstacles++
					break
				}
			}
			resetGrid(grid, originalGrid)
		}
	}

	answer := strconv.Itoa(newObstacles)
	return answer, nil
}

func move(row, col int, grid [][]string) (int, int) {
	if grid[row][col] == "^" {
		return moveUpUntilStopped(row, col, grid)
	} else if grid[row][col] == ">" {
		return moveRightUntilStopped(row, col, grid)
	} else if grid[row][col] == "v" {
		return moveDownUntilStopped(row, col, grid)
	} else if grid[row][col] == "<" {
		return moveLeftUntilStopped(row, col, grid)
	}
	panic("Invalid direction position")
}

func moveUpUntilStopped(row, col int, grid [][]string) (int, int) {
	for row > 0 && grid[row-1][col] != "#" {
		grid[row][col], grid[row-1][col] = "X", grid[row][col]
		row--
	}
	grid[row][col] = ">"
	return row, col
}

func moveRightUntilStopped(row, col int, grid [][]string) (int, int) {
	for col < len(grid[0])-1 && grid[row][col+1] != "#" {
		grid[row][col], grid[row][col+1] = "X", grid[row][col]
		col++
	}
	grid[row][col] = "v"
	return row, col
}

func moveDownUntilStopped(row, col int, grid [][]string) (int, int) {
	for row < len(grid)-1 && grid[row+1][col] != "#" {
		grid[row][col], grid[row+1][col] = "X", grid[row][col]
		row++
	}
	grid[row][col] = "<"
	return row, col
}

func moveLeftUntilStopped(row, col int, grid [][]string) (int, int) {
	for col > 0 && grid[row][col-1] != "#" {
		grid[row][col], grid[row][col-1] = "X", grid[row][col]
		col--
	}
	grid[row][col] = "^"
	return row, col
}

func reachedEnd(row, col int, grid [][]string) bool {
	return row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
	fmt.Println()
}

func resetGrid(grid, originalGrid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = originalGrid[i][j]
		}
	}
}
