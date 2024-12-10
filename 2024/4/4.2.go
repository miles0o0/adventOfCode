package four

import (
	"bufio"
	"fmt"
	"os"
)

func PartTwo() int {
	var grid [][]rune
	file, _ := os.Open("2024/data/dayFour.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	// grid := [][]rune{
	// 	[]rune("MMMSXXMASM"), // Each string converted to a slice of runes
	// 	[]rune("MSAMXMSMSA"),
	// 	[]rune("AMXSXMAAMM"),
	// 	[]rune("MSAMASMSMX"),
	// 	[]rune("XMASAMXAMM"),
	// 	[]rune("XXAMMXXAMA"),
	// 	[]rune("SMSMSASXSS"),
	// 	[]rune("SAXAMASAAA"),
	// 	[]rune("MAMMMXMMMM"),
	// 	[]rune("MXMXAXMASX"),
	// }

	return wordSearch2(grid)
}

func wordSearch2(grid [][]rune) int {
	total := 0
	dirs := [][2]int{
		{-1, 1},  // Top-right
		{-1, -1}, // Top-left
		{1, -1},  // Bottom-left
		{1, 1},   // Bottom-right
	}
	checkXPattern := func(x, y int) string {
		var result []rune
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) {
				result = append(result, grid[nx][ny])
			} else {
				result = append(result, ' ')
			}
		}
		return string(result)
	}
	hasTwoMsAndTwoSs := func(letters []rune) bool {
		countM, countS := 0, 0
		for _, letter := range letters {
			if letter == 'M' {
				countM++
			} else if letter == 'S' {
				countS++
			}
		}
		return countM == 2 && countS == 2
	}
	for x, row := range grid {
		for y, cell := range row {
			if cell != 'A' {
				continue
			}
			letters := checkXPattern(x, y)
			fmt.Printf("At (%d, %d), letters: %s\n", x, y, letters)
			if letters == "MSMS" || letters == "SMSM" || !hasTwoMsAndTwoSs([]rune(letters)) {
				continue
			}
			total++
		}
	}
	return total
}
