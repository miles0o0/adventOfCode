package five

import (
	"bufio"
	"os"
	"strings"
)

func PartOne() int {
	var grid [][]string
	file, _ := os.Open("2024/data/dayFive.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return -1
}
