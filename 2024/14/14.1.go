package fourteen

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const width = 101
const height = 103

// test values
// const width = 11
// const height = 7

const time = 100

func PartOne() int {
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
	floor := simMovement(robots)
	return calQuod(floor)
}

func calQuod(floor [height][width]int) int {
	var quad [4]int
	midY := len(floor) / 2
	midX := len(floor[0]) / 2

	for y := 0; y < len(floor); y++ {
		for x := 0; x < len(floor[0]); x++ {
			if y < midY {
				if x < midX {
					quad[0] += floor[y][x]
				} else if x > midX {
					quad[1] += floor[y][x]
				}
			} else if y > midY {
				if x < midX {
					quad[2] += floor[y][x]
				} else if x > midX {
					quad[3] += floor[y][x]
				}
			}
		}
	}
	log.Println(quad[0], quad[1], quad[2], quad[3])
	return quad[0] * quad[1] * quad[2] * quad[3]
}

func simMovement(robots [][2][2]int) [height][width]int {
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
