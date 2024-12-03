package two

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func TwoMain() int {
	reports := intiilisation()
	var safeCount int = 0
	for i := 0; i < len(reports); i++ {
		if isSafe(reports[i]) {
			safeCount += 1
		}
	}
	return safeCount
}

func TwoSec() int {
	reports := intiilisation()
	var safeCount int = 0
	for i := 0; i < len(reports); i++ {
		if isSafe2(reports[i]) {
			safeCount += 1
		}
	}
	return safeCount
}

func isSafe2(report []int) bool {
	// check asending or decending
	// could revers arrays but not sure whats faster
	acc := isAscending(report)
	return smartRep(report, acc, true)
}

func smartRep(report []int, acc bool, first bool) bool {
	check := func(report []int) bool {
		for i := 0; i < len(report)-1; i++ {
			diff := report[i] - report[i+1]
			if acc {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				return false
			}
		}
		return true
	}
	if check(report) {
		return true
	}
	if !first {
		return false
	}
	for i := 0; i < len(report); i++ {
		modified := append(append([]int{}, report[:i]...), report[i+1:]...)
		if smartRep(modified, acc, false) {
			return true
		}
	}
	return false
}

func isSafe(report []int) bool {
	var acc bool

	// check asending or decending
	// could revers arrays but not sure whats faster
	if report[0] > report[1] {
		acc = false
	} else {
		acc = true
	}

	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if acc {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isAscending(report []int) bool {
	var acc, dec int = 0, 0
	for i := 0; i < len(report)-1; i++ {
		if report[i] > report[i+1] {
			dec += 1
		} else {
			acc += 1
		}
	}
	return acc > dec
}

func intiilisation() [][]int {
	file, _ := os.Open("2024/data/dayTwo.txt")
	defer file.Close()
	var reports [][]int

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sNums := strings.Split(scanner.Text(), " ")
		var nums []int
		for i := 0; i < len(sNums); i++ {
			j, _ := strconv.Atoi(sNums[i])
			nums = append(nums, j)
		}
		reports = append(reports, nums)
		lineCount += 1
	}
	return reports
}
