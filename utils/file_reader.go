package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadLinesFromInputExample() ([]string, error) {
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return []string{}, err
	}

	filePath := filepath.Join(absPath, "input-example.txt")
	return readLines(filePath)
}

func ReadLinesFromInput() ([]string, error) {
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return []string{}, err
	}

	filePath := filepath.Join(absPath, "input.txt")
	return readLines(filePath)
}

func readLines(filePath string) ([]string, error) {
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
