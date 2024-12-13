package thirteen

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

// three tokens to push the a button and 1 token to push the b
func PartOne() int {
	file, _ := os.Open("2024/data/dayThirteen.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	instructions := convInstructions(lines)
	return clawCost(instructions)
}

func convInstructions(lines []string) [][3][2]int {
	var instructions [][3][2]int
	var temp [3][2]int

	for i := 0; i < len(lines); i += 4 {
		if lines[i] == "" {
			continue
		}

		var x1, y1, x2, y2, px, py int
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &x1, &y1)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &x2, &y2)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &px, &py)

		temp[0][0] = x1
		temp[0][1] = y1
		temp[1][0] = x2
		temp[1][1] = y2
		temp[2][0] = px
		temp[2][1] = py

		instructions = append(instructions, temp)
	}
	return instructions
}

func clawCost(instructions [][3][2]int) int {
	total := 0
	for _, set := range instructions {
		// A button translation
		// B button translation
		// P Prize location
		aButt, bButt, prize := set[0], set[1], set[2]

		queue := list.New()
		queue.PushBack([3]int{0, 0, 0}) // [x, y, cost]
		visited := make(map[[2]int]bool)
		visited[[2]int{0, 0}] = true

		for queue.Len() > 0 {
			current := queue.Remove(queue.Front()).([3]int)
			x, y, cost := current[0], current[1], current[2]

			if x == prize[0] && y == prize[1] {
				total += cost
				break
			}

			// Move using button B
			nBX, nBY := x+bButt[0], y+bButt[1]
			if nBX >= 0 && nBY >= 0 && nBX <= prize[0] && nBY <= prize[1] && !visited[[2]int{nBX, nBY}] {
				queue.PushBack([3]int{nBX, nBY, cost + 1})
				visited[[2]int{nBX, nBY}] = true
			}

			// Move using button A
			nAX, nAY := x+aButt[0], y+aButt[1]
			if nAX >= 0 && nAY >= 0 && nAX <= prize[0] && nAY <= prize[1] && !visited[[2]int{nAX, nAY}] {
				queue.PushBack([3]int{nAX, nAY, cost + 3})
				visited[[2]int{nAX, nAY}] = true
			}
		}
	}

	return total
}
