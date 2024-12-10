package five

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartTwo() int {
	var grid []string
	file, _ := os.Open("2024/data/dayFive.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	rules, updates := rulesSort2(grid)
	return valid2(rules, updates)
}

func valid2(rules map[string]bool, updates [][]string) int {
	total := 0
	for _, line := range updates {
		update, valid := ruleValidity(line, rules)
		if valid {
			mid, _ := strconv.Atoi(update[(len(update)+1)/2-1])
			total += mid
		}
	}
	return total
}

func ruleValidity(line []string, rules map[string]bool) ([]string, bool) {
	reordered := false
	var lordSaveMeForMySins func([]string) ([]string, bool)
	lordSaveMeForMySins = func(line []string) ([]string, bool) {
		for pos := range line {
			for i := pos + 1; i < len(line); i++ {
				key := fmt.Sprintf("%s,%s", line[pos], line[i])
				if _, exists := rules[key]; !exists {
					reverseKey := fmt.Sprintf("%s,%s", line[i], line[pos])
					if _, reverseExists := rules[reverseKey]; reverseExists {
						element := line[i]
						line = append(line[:i], line[i+1:]...)
						line = append(line[:pos], append([]string{element}, line[pos:]...)...)
						reordered = true
						return lordSaveMeForMySins(line)
					} else {
						return line, false
					}
				}
			}
		}
		return line, true
	}
	line, valid := lordSaveMeForMySins(line)
	return line, valid && reordered
}

func rulesSort2(grid []string) (map[string]bool, [][]string) {
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
