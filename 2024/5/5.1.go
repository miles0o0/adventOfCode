package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne() int {
	var grid []string
	file, _ := os.Open("2024/data/dayFive.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	rules, updates := rulesSort(grid)
	return valid(rules, updates)
}

func valid(rules map[string]bool, updates [][]string) int {
	total := 0

	ruleValidity := func(line []string) ([]string, bool) {
		for pos := range line {
			for i := pos + 1; i < len(line); i++ {
				key := fmt.Sprintf("%s,%s", line[pos], line[i])
				if _, exists := rules[key]; !exists {
					return line, false
				}
			}
		}
		return line, true
	}
	for _, line := range updates {
		update, valid := ruleValidity(line)
		if valid {
			mid, _ := strconv.Atoi(update[(len(update)+1)/2-1])
			total += mid
		}
	}
	return total
}

func rulesSort(grid []string) (map[string]bool, [][]string) {
	var rules = make(map[string]bool)
	var updates [][]string
	for _, line := range grid {
		rule := strings.Split(line, "|")
		if len(rule) == 2 {
			key := fmt.Sprintf("%s,%s", rule[0], rule[1])
			rules[key] = true
			continue
		}
		update := strings.Split(line, ",")
		if len(update) > 1 {
			updates = append(updates, update)
		}
	}
	return rules, updates
}
