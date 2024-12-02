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
	var total float64 = 0
	for i := 0; i < len(left); i++ {
		value1 := left[i] - right[i]
		var value2 float64 = float64(value1 * value1)
		value2 = math.Sqrt(value2)
		total += value2
	}

	return int(total)
}

func OneSec() int {
	left, right := intiilisation()

	type value struct {
		value   int
		numOcur int
	}
	var total int = 0
	dic := make(map[int]int)

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
		value1, err1 := strconv.ParseFloat(nums[0], 64)
		value2, err2 := strconv.ParseFloat(nums[1], 64)
		if err1 != nil || err2 != nil {
			log.Fatal(err)
		}
		left = append(left, value1)
		right = append(right, value2)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
