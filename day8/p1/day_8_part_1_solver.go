package d8p1

import (
	"fmt"
	"strconv"
	"strings"
)

type Day8Part1Solver struct{}

func (solver *Day8Part1Solver) Solve(lines []string) (string, error) {

	antennas := make(map[string][][]int)
	grid := [][]string{}
	width := len(lines)
	height := len(lines[0])

	for i := 0; i < len(lines); i++ {
		gridRow := []string{}
		for j := 0; j < len(lines[i]); j++ {
			curr := string(lines[i][j])
			gridRow = append(gridRow, curr)
			if curr == "." {
				continue
			}
			antennas[curr] = append(antennas[curr], []int{i, j})
		}
		grid = append(grid, gridRow)
	}

	printGrid(grid)
	fmt.Println(antennas)

	for _, locations := range antennas {

		for _, location := range locations {
			for _, otherLocation := range locations {
				if location[0] == otherLocation[0] && location[1] == otherLocation[1] {
					continue
				}
				xDist := otherLocation[0] - location[0]
				yDist := otherLocation[1] - location[1]

				destinationRow := location[0] + (2 * xDist)
				destinationCol := location[1] + (2 * yDist)

				if destinationRow >= 0 && destinationRow < height &&
					destinationCol >= 0 && destinationCol < width {

					if strings.Contains(grid[destinationRow][destinationCol], "#") {
						continue
					}
					grid[destinationRow][destinationCol] += "#"
				}
			}
		}
	}

	antinodes := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if strings.Contains(grid[i][j], "#") {
				antinodes++
			}
		}
	}
	printGrid(grid)

	answer := strconv.Itoa(antinodes)
	return answer, nil
}

func printGrid(grid [][]string) {
	for _, gridRow := range grid {
		fmt.Println(gridRow)
	}
}
