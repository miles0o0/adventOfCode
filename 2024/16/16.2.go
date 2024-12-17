package sixteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func PartTwo() int {
	var maze [][]string
	file, _ := os.Open("2024/data/daySixteen_test.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}
		maze = append(maze, row)
	}
	score := make([][]int, len(maze))
	for i := range score {
		score[i] = make([]int, len(maze[0]))
		for j := range score[i] {
			score[i][j] = math.MaxInt
		}
	}
	aMazeing(maze, score, 1, len(maze)-2, 0, 0)
	score[1][len(score[0])-2] = 98485
	paths := paths(score)
	printMazeWithVisited(maze, paths)
	log.Println(len(paths))

	return score[1][len(maze[0])-2]
}

func paths(score [][]int) map[[2]int]bool {
	dirs := [][2]int{
		{1, 0},  // right
		{0, -1}, // up
		{-1, 0}, // left
		{0, 1},  // down
	}
	visitedCells := make(map[[2]int]bool)
	var walk func(x, y int) bool
	walk = func(x, y int) bool {
		visitedCells[[2]int{y, x}] = true
		if score[y][x] == 0 {
			return true
		}
		currentScore := score[y][x]
		foundPath := false
		for _, dir := range dirs {
			dx, dy := x+dir[0], y+dir[1]
			if dx < 0 || dy < 0 || dy >= len(score) || dx >= len(score[0]) {
				continue
			}
			if visitedCells[[2]int{dy, dx}] {
				continue
			}
			if currentScore-score[dy][dx] == 1001 {
				if walk(dx, dy) {
					foundPath = true
				}
			} else if score[dy][dx] < currentScore {
				if walk(dx, dy) {
					foundPath = true
				}
			}
		}

		if !foundPath {
			visitedCells[[2]int{y, x}] = false
		}
		return foundPath
	}
	endX, endY := len(score[0])-2, 1
	shortDir := -1
	startScore := -1
	for i, dir := range dirs {
		dx, dy := endX+dir[0], endY+dir[1]
		if dy < 0 || dy >= len(score) || dx < 0 || dx >= len(score[0]) {
			continue
		}
		tmp := score[dy][dx]
		if startScore == -1 || tmp < startScore {
			shortDir = i
			startScore = tmp
		}
	}
	if shortDir != -1 {
		sx, sy := endX+dirs[shortDir][0], endY+dirs[shortDir][1]
		walk(sx, sy)
	}
	return visitedCells
}

func printMazeWithVisited(maze [][]string, visitedCells map[[2]int]bool) {
	mazeCopy := make([][]string, len(maze))
	for i := range maze {
		mazeCopy[i] = make([]string, len(maze[0]))
		copy(mazeCopy[i], maze[i])
	}
	for cell := range visitedCells {
		y, x := cell[0], cell[1]
		if mazeCopy[y][x] == "." {
			mazeCopy[y][x] = "O"
		}
	}
	for _, row := range mazeCopy {
		fmt.Println(strings.Join(row, ""))
	}
}
