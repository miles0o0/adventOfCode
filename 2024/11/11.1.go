package eleven

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	file, _ := os.Open("2024/data/dayEleven_test.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	stones := convertToIntList(line)
	blinks := 75
	for i := 0; i < blinks; i++ {
		log.Println(i)
		stones = simBlink(stones)
	}
	return len(stones)
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
func simBlink(stones []int) []int {
	result := []int{}
	for _, stone := range stones {
		if stone == 0 {
			result = append(result, 1)
		} else if (len(strconv.Itoa(stone)) % 2) == 0 {
			str := strconv.Itoa(stone)
			mid := len(str) / 2
			firstHalf, _ := strconv.Atoi(str[:mid])
			secondHalf, _ := strconv.Atoi(str[mid:])
			result = append(result, firstHalf, secondHalf)
		} else {
			result = append(result, stone*2024)
		}
	}
	return result
}
