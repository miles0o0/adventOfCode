package seven

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartTwo() int {
	var grid []string
	file, _ := os.Open("2024/data/daySeven.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return operations2(grid)
}

func operations2(grid []string) int {
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
		if search2(nums[1:], nums[0], result) {
			total += result
		}
	}
	return total
}

func search2(nums []int, current int, result int) bool {
	if len(nums) == 0 {
		return current == result
	}
	num := nums[0]
	newNums := nums[1:]
	if search2(newNums, current+num, result) {
		return true
	}
	if search2(newNums, current*num, result) {
		return true
	}
	newCurrent, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, num))
	if search2(newNums, newCurrent, result) {
		return true
	}
	return false
}
