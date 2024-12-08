package eight

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PartOne() int {
	var grid [][]string
	file, _ := os.Open("2024/data/dayEight.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}
	return antiNodeCount(grid)
}

func antiNodeCount(grid [][]string) int {
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
	locAnti := func(loc1 []int, loc2 []int) [][]int {
		dx := loc1[0] - loc2[0]
		dy := loc1[1] - loc2[1]
		nx1 := loc1[0] + dx
		ny1 := loc1[1] + dy
		nx2 := loc2[0] - dx
		ny2 := loc2[1] - dy
		return [][]int{{nx1, ny1}, {nx2, ny2}}
	}

	antNodeLoc := make(map[string]bool)
	for _, loc := range nodeLoc {
		for i := 0; i < len(loc); i++ {
			for j := i + 1; j < len(loc); j++ {
				for _, result := range locAnti(loc[i], loc[j]) {
					if result[0] >= len(grid[0]) || result[1] >= len(grid) || result[0] < 0 || result[1] < 0 {
						continue
					}
					loc := fmt.Sprintf("%d,%d", result[0], result[1])
					antNodeLoc[loc] = true
				}
			}
		}
	}
	return len(antNodeLoc)
}
