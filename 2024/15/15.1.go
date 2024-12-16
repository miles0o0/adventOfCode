package fifteen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PartOne() int {
	var floor [][]string
	var instructions [][]string
	file, _ := os.Open("2024/data/dayFifteen.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") {
			var row []string
			for _, char := range line {

				row = append(row, string(char))
			}
			floor = append(floor, row)
		} else if strings.Contains(line, "^") || strings.Contains(line, "<") || strings.Contains(line, ">") || strings.Contains(line, "v") {
			var row []string
			for _, char := range line {
				row = append(row, string(char))
			}
			instructions = append(instructions, row)
		}
	}

	simulateBots(floor, instructions)
	return boxGPS(floor)
}

func boxGPS(floor [][]string) int {
	total := 0
	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[0]); x++ {
			cell := floor[y][x]
			if cell == "O" {
				total += (y * 100) + x
			}
		}
	}
	return total
}

func botDir(inst string) [2]int {
	var dir int
	dirs := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	switch inst {
	case "^":
		dir = 0
	case ">":
		dir = 1

	case "v":
		dir = 2
	case "<":
		dir = 3
	default:
		panic("direction not found")
	}
	return dirs[dir]
}

func findBot(floor [][]string) (x, y int) {
	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[0]); x++ {
			if floor[y][x] == "@" {

				return x, y
			}
		}
	}
	panic("robot not found")
}

func printGrid2(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

func simulateBots(floor, instructions [][]string) {
	x, y := findBot(floor)
	for _, inst := range instructions {
		for _, command := range inst {
			dir := botDir(command)
			dx, dy := x+dir[0], y+dir[1]
			if floor[dy][dx] == "O" {
				push := true
				for i := 2; ; i++ {
					tdx, tdy := dx+dir[0]*i, dy+dir[1]*i
					if tdx < 0 || tdy < 0 || tdx >= len(floor[0]) || tdy >= len(floor) || floor[tdy][tdx] == "#" {
						push = false
						break
					}
					if floor[tdy][tdx] == "." {
						for j := i; j > 0; j-- {
							floor[dy+dir[1]*j][dx+dir[0]*j] = "O"
							floor[dy+dir[1]*(j-1)][dx+dir[0]*(j-1)] = "."
						}
						break
					}
				}
				if !push {
					continue
				}
			}
			if floor[dy][dx] != "#" {
				floor[dy][dx] = "@"
				floor[y][x] = "."
				x, y = dx, dy
			}
		}
	}
}
