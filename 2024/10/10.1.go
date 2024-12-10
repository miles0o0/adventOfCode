package ten

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var trailTracks = make(map[string]bool)

func PartOne() int {
	// Read grid from file
	var grid []string
	file, _ := os.Open("2024/data/dayTen.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	numGrid := make([][]int, len(grid))
	for i, line := range grid {
		numRow := []int{}
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Printf("Error converting '%c' to int: %v\n", char, err)
				continue
			}
			numRow = append(numRow, num)
		}
		numGrid[i] = numRow
	}
	result := findTrails(numGrid)
	return result
}

func findTrails(grid [][]int) int {
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 0 {
				trailTracks = make(map[string]bool)
				_ = findTrailHeads(grid, x, y, 0, grid[y][x], true)
				total += len(trailTracks)
			}
		}
	}
	return total
}

func findTrailHeads(grid [][]int, x int, y int, heads int, start int, firstCall bool) int {
	dirs := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return heads
	}
	if grid[y][x]-1 != start && !firstCall {
		return heads
	}
	if grid[y][x] == 9 {
		key := fmt.Sprintf("%d,%d", x, y)
		trailTracks[key] = true
		return heads + 1
	}
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			continue
		}
		heads = findTrailHeads(grid, (x + dir[0]), (y + dir[1]), heads, grid[y][x], false)
	}
	return heads
}
