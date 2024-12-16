package sixteen

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func PartOne() int {
	var maze [][]string
	file, _ := os.Open("2024/data/daySixteen.txt")
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
	score[0][0] = 1
	printScore(score, maze)
	log.Println(score[1][len(maze[0])-3])
	log.Println(score[2][len(maze[0])-2])
	return score[1][len(maze[0])-2]
}

// the print out needs +1 to the total as it does not step onto the final dot

func printScore(score [][]int, maze [][]string) {
	for y, row := range score {
		for x, cell := range row {
			mCell := maze[y][x]
			if mCell == "#" {
				fmt.Printf("%s  ", mCell)
			} else if cell == math.MaxInt {
				fmt.Print(".  ")
			} else {
				fmt.Printf("%d", cell)
			}
		}
		fmt.Println()
	}
}

// forward == 1 point
// 90 degree turns == 1000
func aMazeing(maze [][]string, score [][]int, x, y, dir, cur int) {
	dirs := [][2]int{
		{1, 0},  // right
		{0, -1}, // up
		{-1, 0}, // left
		{0, 1},  // down
	}
	if score[y][x] <= cur {
		return
	}
	score[y][x] = cur
	for i, d := range dirs {
		dx, dy := x+d[0], y+d[1]
		cell := maze[dy][dx]
		if cell == "#" {
			tCur := cur + 1000
			aMazeing(maze, score, x, y, i, tCur)
			continue
		}
		if cell == "." {
			tCur := cur + 1
			if i != dir {
				tCur += 1000
			}
			aMazeing(maze, score, dx, dy, i, tCur)
		}
	}
}
