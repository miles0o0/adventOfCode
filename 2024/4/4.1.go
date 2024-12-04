package four

import (
	"bufio"
	"os"
)

func PartOne() int {
	var grid [][]rune
	file, _ := os.Open("2024/data/dayFour.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return wordSearch(grid)
}

func wordSearch(grid [][]rune) int {
	const word = "XMAS"
	total := 0
	length := len(word)
	dirs := [][2]int{
		{0, 1},   // up
		{1, 0},   // right
		{-1, 0},  // left
		{0, -1},  // down
		{1, 1},   // up right
		{-1, 1},  // up left
		{1, -1},  // down right
		{-1, -1}, // down left
	}
	// joke name don't take this all so seriously
	// searches for xmas starting with x
	omniDirectionalSearch := func(x, y, dx, dy int) bool {
		for i := 0; i < length; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(word[i]) {
				return false
			}
		}
		return true
	}
	for x, row := range grid {
		for y := range row {
			if string(grid[x][y]) != "X" {
				continue
			}
			for _, dir := range dirs {
				if omniDirectionalSearch(x, y, dir[0], dir[1]) {
					total += 1
				}
			}
		}
	}
	return total
}
