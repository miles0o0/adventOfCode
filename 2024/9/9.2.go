package nine

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// start file block then space block. alturnate
func PartTwo() int {
	file, _ := os.Open("2024/data/dayNine.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	expanded := compDiskSpace2(line)
	comp := compact2(expanded)
	return checkSum2(comp)
}

func compDiskSpace2(line string) []string {
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

func compact2(lineLetters []string) []string {
	i, j := 0, len(lineLetters)-1
	var lengthGap, lengthSec int
	for {
		if j <= 0 || lineLetters[j] == "0" {
			break
		}
		if lineLetters[i] != "." {
			i++
			continue
		} else if lineLetters[j] == "." {
			j--
			continue
		}
		if i >= j {
			i = 0
			j -= lengthSec
		}
		lengthGap = 0
		lengthSec = 0
		for k := i; lineLetters[k] == "."; k++ {
			lengthGap++
		}
		for k := j; k >= 0; k-- {
			if lineLetters[k] != lineLetters[j] {
				break
			}
			lengthSec++
		}
		if lengthGap >= lengthSec {
			for l := 0; l < lengthSec; l++ {
				lineLetters[l+i] = lineLetters[j]
				lineLetters[j] = "."
				j--
			}
			i = 0
		} else {
			i++
		}
	}
	return lineLetters
}

func checkSum2(line []string) int {
	var sum int
	for i := 0; i < len(line); i++ {
		if string(line[i]) == "." {
			continue
		}
		val, _ := strconv.Atoi(string(line[i]))
		sum += (i * val)
	}
	return sum
}
