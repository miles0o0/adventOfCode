package three

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	var total int = 0
	file, _ := os.Open("2024/data/dayThree.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += searchMul(scanner.Text())
	}
	return total
}

// adding buffer because of the 13 length in the edge case of minimum function size at the end of the line
func searchMul(line string) int {
	gLine := append([]rune(line), []rune("ooooo")...)
	total := 0
	for pos, char := range gLine {
		if char == 'm' && pos+13 <= len(gLine) {
			substring := gLine[pos : pos+13]
			total += extract(substring)
		}
	}
	return total
}

func extract(fun []rune) int {
	if len(fun) < 13 || string(fun[0:4]) != "mul(" {
		return 0
	}
	function := strings.Split(string(fun[4:13]), ")")
	if len(function) < 2 {
		return 0
	}
	numbers := strings.Split(function[0], ",")
	if !(len(numbers) == 2) {
		return 0
	}
	n1, err1 := strconv.Atoi(numbers[0])
	n2, err2 := strconv.Atoi(numbers[1])
	if err1 != nil || err2 != nil {
		return 0
	}
	return n1 * n2
}
