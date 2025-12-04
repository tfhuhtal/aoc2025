package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []rune(lines[i])
	}

	var p1, p2 int64
	first := true

	start := time.Now()

	for {
		changed := false

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				nbr := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						rr := r + dr
						cc := c + dc
						if rr >= 0 && rr < rows && cc >= 0 && cc < cols && grid[rr][cc] == '@' {
							nbr++
						}
					}
				}
				if grid[r][c] == '@' && nbr < 5 {
					p1++
					changed = true
					if !first {
						p2++
						grid[r][c] = '.'
					}
				}
			}
		}
		if first {
			fmt.Println(p1)
			first = false
		}
		if !changed {
			break
		}
	}

	elapsed := time.Since(start)

	fmt.Println(p2)
	fmt.Println("elapsed: ", elapsed)
}
