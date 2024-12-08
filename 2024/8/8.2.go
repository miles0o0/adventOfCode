package eight

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PartTwo() int {
	var grid [][]string
	file, _ := os.Open("2024/data/dayEight.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	return antiNodeCount2(grid)
}

func antiNodeCount2(grid [][]string) int {
	nodeLoc := make(map[string][][]int)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "." {
				continue
			} else {
				ant := grid[y][x]
				loc := []int{x, y}
				nodeLoc[ant] = append(nodeLoc[ant], loc)
			}
		}
	}

	locAnti := func(loc1 []int, loc2 []int) map[string]bool {
		dx := loc1[0] - loc2[0]
		dy := loc1[1] - loc2[1]
		antiMap := make(map[string]bool)
		addPoint := func(x, y int) {
			key := fmt.Sprintf("%d,%d", x, y)
			if !antiMap[key] {
				antiMap[key] = true
			}
		}
		x, y := loc1[0], loc1[1]
		for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
			addPoint(x, y)
			x += dx
			y += dy
		}
		x, y = loc2[0], loc2[1]
		for x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) {
			addPoint(x, y)
			x -= dx
			y -= dy
		}
		return antiMap
	}

	antNodeLoc := make(map[string]bool)
	for _, loc := range nodeLoc {
		for i := 0; i < len(loc); i++ {
			for j := i + 1; j < len(loc); j++ {
				results := locAnti(loc[i], loc[j])
				for key := range results {
					antNodeLoc[key] = true
				}
			}
		}
	}
	return len(antNodeLoc)
}
