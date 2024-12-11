package eleven

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	file, _ := os.Open("2024/data/dayEleven.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	stones := convertToIntList(line)
	blinks := 75
	stonesIndex := make(map[string]int)
	return simBlink(stones, blinks, stonesIndex)
}

func convertToIntList(stones string) []int {
	var intStones []int
	arrStones := strings.Split(stones, " ")
	for _, stone := range arrStones {
		intStone, _ := strconv.Atoi(stone)
		intStones = append(intStones, intStone)
	}
	return intStones
}

// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func simBlink(stones []int, depth int, stoneIndex map[string]int) int {
	if depth >= 45 {
		log.Printf("Depth: %d, Stones: %v\n", depth, stones)
	}

	if depth == 0 {
		return len(stones)
	}

	numStones := 0
	for _, stone := range stones {
		key := fmt.Sprintf("%d-%d", stone, depth)

		if value, exists := stoneIndex[key]; exists {
			numStones += value
			continue
		}
		var result int
		if stone == 0 {
			result = simBlink([]int{1}, depth-1, stoneIndex)
		} else {
			numDigits := int(math.Log10(float64(stone))) + 1

			if numDigits%2 == 0 {
				half := numDigits / 2
				divisor := int(math.Pow10(half))
				firstHalf := stone / divisor
				secondHalf := stone % divisor
				result = simBlink([]int{firstHalf, secondHalf}, depth-1, stoneIndex)
			} else {
				newStone := stone * 2024
				result = simBlink([]int{newStone}, depth-1, stoneIndex)
			}
		}
		stoneIndex[key] = result
		numStones += result
	}
	return numStones
}
