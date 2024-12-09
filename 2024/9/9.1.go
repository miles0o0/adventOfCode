package nine

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// start file block then space block. alturnate
func PartOne() int {
	file, _ := os.Open("2024/data/dayNine.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	expanded := compDiskSpace(line)
	comp := compact(expanded)
	return checkSum(comp)
}

func compDiskSpace(line string) []string {
	strNums := strings.Split(line, "")
	fileBlock := true
	var expanded []string
	id := 0
	for i := 0; i < len(strNums); i++ {
		val, _ := strconv.Atoi(strNums[i])
		var str []string
		if fileBlock {
			for j := 0; j < val; j++ {
				str = append(str, strconv.Itoa(id))
			}
			id++
			fileBlock = false
		} else {
			for j := 0; j < val; j++ {
				str = append(str, ".")
			}
			fileBlock = true
		}
		expanded = append(expanded, str...)
	}
	return expanded
}

func compact(lineLetters []string) []string {
	i, j := 0, len(lineLetters)-1
	for ; i < j; i++ {
		if lineLetters[i] != "." {
			continue
		} else if lineLetters[j] == "." {
			i--
			j--
			continue
		}
		lineLetters[i] = lineLetters[j]
		lineLetters[j] = "."
		j--
	}
	return lineLetters
}

func checkSum(line []string) int {
	var sum int
	for i := 0; i < len(line); i++ {
		if string(line[i]) == "." {
			break
		}
		val, _ := strconv.Atoi(string(line[i]))
		sum += (i * val)
	}
	return sum
}
