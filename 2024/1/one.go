package one

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func OneMain() int {

	left, right := intiilisation()

	// square and root to remove negative
	// tally up difference
	var total float64
	for i := 0; i < len(left); i++ {
		value1 := left[i] - right[i]
		var value2 float64 = float64(value1 * value1)
		value2 = math.Sqrt(value2)
		total += value2
	}

	return int(total)
}

// gets right list count
// multiply left value with right count
// return total product of the latter
func OneSec() int {
	left, right := intiilisation()

	dic := make(map[float64]float64)
	// maybe some optimisations because its sorted? cba
	for i := 0; i < len(right); i++ {
		_, exists := dic[right[i]]
		if exists {
			dic[right[i]] = dic[right[i]] + 1
		} else {
			dic[right[i]] = 1
		}
	}

	var total float64
	for i := 0; i < len(left); i++ {
		_, exists := dic[left[i]]
		if exists {
			total += left[i] * dic[left[i]]
		}
	}

	return int(total)
}

// both functions need the file into an array first needs sorted
func intiilisation() ([]float64, []float64) {

	file, _ := os.Open("2024/data/dayOne.txt")
	defer file.Close()

	var left []float64
	var right []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		// convert to float 64 because squaring and rooting to remove negatives
		value1, _ := strconv.ParseFloat(nums[0], 64)
		value2, _ := strconv.ParseFloat(nums[1], 64)
		left = append(left, value1)
		right = append(right, value2)
	}

	//sort :)
	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
