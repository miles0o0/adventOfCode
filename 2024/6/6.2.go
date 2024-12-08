package six

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PartTwo() int {
	var grid [][]string
	file, _ := os.Open("2024/data/daySix.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	count := simMovement2(grid)
	return count
}

func simMovement2(grid [][]string) int {
	dirs := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	placement := make(map[string]bool)

	var sX, sY int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "^" {
				sX, sY = x, y
				break
			}
		}
	}

	simulate := func(sX, sY int) bool {
		nx, ny := sX, sY
		dir := 0
		position := make(map[string]bool)

		for {
			state := fmt.Sprintf("%d,%d,%d", nx, ny, dir)
			if position[state] {
				return true
			}
			position[state] = true
			newX := nx + dirs[dir][0]
			newY := ny + dirs[dir][1]
			if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) {
				return false
			}
			if grid[newY][newX] == "#" {
				dir = (dir + 1) % 4
			} else {
				nx = newX
				ny = newY
			}
		}
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "." {
				grid[y][x] = "#"
				if simulate(sX, sY) {
					placement[fmt.Sprintf("%d,%d", x, y)] = true
				}
				grid[y][x] = "."
			}
		}
	}
	return len(placement)
}
