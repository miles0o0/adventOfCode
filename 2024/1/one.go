package one

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func OneMain() int {
	left, right := intiilisation()
	// negative to positive
	// tally up difference
	var total int
	for i := 0; i < len(left); i++ {
		value := left[i] - right[i]
		if value < 0 {
			value = -value
		}
		total += value
	}
	return int(total)
}

// gets right list count
// multiply left value with right count
// return total product of the latter
func OneSec() int {
	left, right := intiilisation()

	dic := make(map[int]int)
	// maybe some optimisations because its sorted? cba
	for i := 0; i < len(right); i++ {
		_, exists := dic[right[i]]
		if exists {
			dic[right[i]] = dic[right[i]] + 1
		} else {
			dic[right[i]] = 1
		}
	}

	var total int
	for i := 0; i < len(left); i++ {
		_, exists := dic[left[i]]
		if exists {
			total += left[i] * dic[left[i]]
		}
	}
	return total
}

// both functions need the file into an array first needs sorted
func intiilisation() ([]int, []int) {
	file, _ := os.Open("2024/data/dayOne.txt")
	defer file.Close()
	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		value1, _ := strconv.Atoi(nums[0])
		value2, _ := strconv.Atoi(nums[1])
		left = append(left, value1)
		right = append(right, value2)
	}
	//sort :)
	slices.Sort(left)
	slices.Sort(right)
	return left, right
}
