package util

import (
	"bufio"
	"os"
)

func ReadFile(filepath string) *bufio.Scanner {
	file, _ := os.Open("2024/data/dayTwo.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return scanner
}
