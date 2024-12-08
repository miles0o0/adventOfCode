package seven

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	var grid []string
	file, _ := os.Open("2024/data/daySeven.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return operations(grid)
}

func operations(grid []string) int {
	var total int
	for _, line := range grid {
		values := strings.Split(line, ":")
		result, _ := strconv.Atoi(values[0])
		values = strings.Split(values[1], " ")
		var nums []int
		for _, str := range values {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		if search(nums[1:], nums[0], result) {
			total += result
		}
	}
	return total
}

func search(nums []int, current int, result int) bool {
	if len(nums) == 0 {
		return current == result
	}
	num := nums[0]
	newNums := nums[1:]
	if search(newNums, current+num, result) {
		return true
	}
	if search(newNums, current*num, result) {
		return true
	}
	return false
}
