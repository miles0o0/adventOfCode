package fourteen

import (
	"bufio"
	"fmt"
	"os"
)

func PartTwo() int {
	var robots [][2][2]int
	file, _ := os.Open("2024/data/dayFourteen.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var p, v [2]int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p[0], &p[1], &v[0], &v[1])
		robots = append(robots, [2][2]int{p, v})
	}

	for i := 0; i < 100000000; i++ {
		floor := simMovement2(robots, i)
		if checkRowForRobots(floor, 31) {
			printGrid2(floor)
			return i
		}
	}
	return -1
}

func printGrid2(grid [height][width]int) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func checkRowForRobots(grid [height][width]int, count int) bool {
	for y := 0; y < height; y++ {
		consecutive := 0
		for x := 0; x < width; x++ {
			if grid[y][x] != 0 {
				consecutive++
				if consecutive == count {
					return true
				}
			} else {
				consecutive = 0
			}
		}
	}
	return false
}

func simMovement2(robots [][2][2]int, time int) [height][width]int {
	var floor [height][width]int
	for _, robot := range robots {
		position := robot[0]
		vector := robot[1]

		finalX := (position[0] + vector[0]*time) % width
		finalY := (position[1] + vector[1]*time) % height

		if finalX < 0 {
			finalX += width
		}
		if finalY < 0 {
			finalY += height
		}

		floor[finalY][finalX]++
	}
	return floor
}
