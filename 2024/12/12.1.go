package twelve

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartOne() int {
	file, _ := os.Open("2024/data/dayTwelve.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	dirs := [][2]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	region := segragate(lines, dirs)
	return fenceCost(region, dirs)
}

// create map / bounds of the different areas with the same letters
func segragate(lines [][]string, dirs [][2]int) [][]string {
	regons := make([][]string, len(lines))
	for i := range regons {
		regons[i] = make([]string, len(lines[0]))
	}
	cache := make(map[string]int)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if regons[y][x] == "" {
				cache[lines[y][x]] += 1
				walkRegion(lines, regons, x, y, dirs, cache)
			}

		}
	}
	return regons
}

// label the regions with the tag to mak out the segments
func walkRegion(lines, regions [][]string, x, y int, dirs [][2]int, cache map[string]int) {
	region := lines[y][x]
	nameRegion := fmt.Sprintf("%s-%d", region, cache[region])
	regions[y][x] = nameRegion

	for _, dir := range dirs {
		dx, dy := x+dir[0], y+dir[1]
		if dy < 0 || dx < 0 || dy >= len(lines) || dx >= len(lines[0]) {
			continue
		}
		if regions[dy][dx] == nameRegion {
			continue
		}
		if lines[dy][dx] == region {
			walkRegion(lines, regions, dx, dy, dirs, cache)
		}
	}
}

func fenceCost(regions [][]string, dirs [][2]int) int {
	var total int
	perimeter := make(map[string]int)
	area := make(map[string]int)

	// Calculate area and perimeter
	for y := 0; y < len(regions); y++ {
		for x := 0; x < len(regions[0]); x++ {
			region := regions[y][x]
			area[region] += 1

			for _, dir := range dirs {
				dx, dy := x+dir[0], y+dir[1]
				if dy < 0 || dy >= len(regions) || dx < 0 || dx >= len(regions[0]) || regions[dy][dx] != region {
					perimeter[region] += 1
				}
			}
		}
	}

	// Calculate total cost
	for key, p := range perimeter {
		log.Printf("Region %s: area = %d, Perimeter = %d", key, area[key], p)
		total += p * area[key]
	}
	return total
}
