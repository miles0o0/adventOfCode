package three

import (
	"bufio"
	"os"
)

func PartTwo() int {
	var total int = 0
	file, _ := os.Open("2024/data/dayThree.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	do := true
	var value int
	for scanner.Scan() {
		value, do = searchMul2(scanner.Text(), do)
		total += value
	}
	return total
}

// adding buffer because of the 13 length in the edge case of minimum function size at the end of the line
func searchMul2(line string, do bool) (int, bool) {
	gLine := append([]rune(line), []rune("ooooo")...)
	total := 0
	for pos, char := range gLine {
		if char == 'd' && pos+6 <= len(gLine) {
			substring := gLine[pos : pos+6]
			do = flag(substring, do)
		}
		if char == 'm' && pos+13 <= len(gLine) && do {
			substring := gLine[pos : pos+13]
			total += extract(substring)
		}
	}
	return total, do
}

func flag(command []rune, state bool) bool {
	if string(command[0:4]) == "do()" {
		return true
	} else if string(command[0:7]) == "don't()" {
		return false
	}
	return state
}
