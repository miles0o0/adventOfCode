package one

import (
	"bufio"
	"log"
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
	var total float64 = 0
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

	var total int = 0
	dic := make(map[int]int)
	// maybe some optimisations because its sorted? cba
	for i := 0; i < len(right); i++ {
		_, exists := dic[int(right[i])]
		if exists {
			dic[int(right[i])] = dic[int(right[i])] + 1
		} else {
			dic[int(right[i])] = 1
		}
	}

	for i := 0; i < len(left); i++ {
		_, exists := dic[int(left[i])]
		if exists {
			total += int(left[i]) * dic[int(left[i])]
		}
	}

	return total
}

// both functions need the file into an array first needs sorted
func intiilisation() ([]float64, []float64) {
	file, err := os.Open("2024/data/dayOne.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []float64
	var right []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		// convert to float 64 because squaring and rooting to remove negatives
		value1, err1 := strconv.ParseFloat(nums[0], 64)
		value2, err2 := strconv.ParseFloat(nums[1], 64)
		if err1 != nil || err2 != nil {
			log.Fatal(err)
		}
		left = append(left, value1)
		right = append(right, value2)
	}

	//sort :)
	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
