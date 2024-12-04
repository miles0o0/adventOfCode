package four

import (
	"bufio"
	"os"
)

func PartTwo() int {
	var grid [][]rune
	file, _ := os.Open("2024/data/dayFour.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return wordSearch2(grid)
}

func wordSearch2(grid [][]rune) int {
	const word = "MS"
	total := 0
	length := len(word)
	dirs := [][2]int{
		{1, 1},   // up right
		{-1, 1},  // up left
		{-1, -1}, // down left
		{1, -1},  // down right

	}
	// joke name don't take this all so seriously
	// searches for xmas starting with x
	omniDirectionalSearch := func(x, y, dx, dy int) rune {
		for i := 0; i < length; i++ {
			nx, ny := x+dx, y+dy
			if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(word[i]) {
				return rune(word[i])
			}
		}
		return rune('X')
	}
	for x, row := range grid {
		for y := range row {
			if string(grid[x][y]) != "A" {
				continue
			}
			var letters []rune
			for _, dir := range dirs {
				letter := omniDirectionalSearch(x, y, dir[0], dir[1])
				if string(letter) != "X" {
					letters = append(letters, letter)
				}
			}
			if !isEqual(letters, []rune{'M', 'S', 'M', 'S'}) &&
				!isEqual(letters, []rune{'S', 'M', 'S', 'M'}) &&
				countRune(letters, 'M') == 2 &&
				countRune(letters, 'S') == 2 &&
				len(letters) == 4 {
				total += 1
				// log.Print(string(letters))
			}
		}
	}
	return total
}

func countRune(letters []rune, r rune) int {
	count := 0
	for _, letter := range letters {
		if letter == r {
			count++
		}
	}
	return count
}

func isEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
