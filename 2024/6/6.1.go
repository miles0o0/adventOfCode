package six

import (
	"bufio"
	"os"
	"strings"
)

func PartOne() int {
	var grid [][]string
	file, _ := os.Open("2024/data/daySix.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	grid = simMovement(grid)
	count, _ := countX(grid)
	return count
}

func simMovement(grid [][]string) [][]string {
	dirs := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	nx, ny := countX(grid)
	dir := 0

	changeDir := func(grid [][]string) {
		newX := nx + dirs[dir][0]
		newY := ny + dirs[dir][1]
		grid[ny][nx] = "X"
		if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[newY]) {
			if grid[newY][newX] != "#" {
				nx = newX
				ny = newY
			} else {
				dir = (dir + 1) % 4
			}
		} else {
			nx = newX
			ny = newY
		}
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			grid[ny][nx] = "^"
		}
	}

	for nx >= 0 && ny >= 0 && ny < len(grid) && nx < len(grid[ny]) {
		changeDir(grid)
	}
	return grid
}

func countX(grid [][]string) (int, int) {
	var count int
	nx := 0
	ny := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "X" {
				count++
			} else if grid[y][x] == "^" {
				nx, ny = x, y
			}
		}
	}
	if count == 0 {
		count = nx
	}
	return count, ny
}
