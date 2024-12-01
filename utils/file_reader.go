package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadLines() ([]string, error) {
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return []string{}, err
	}

	filePath := filepath.Join(absPath, "input.txt")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
