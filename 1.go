package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	dial := 50
	p1 := 0
	p2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		d := line[0]
		amtStr := line[1:]
		amt, err := strconv.Atoi(amtStr)
		if err != nil {
			continue
		}

		for i := 0; i < amt; i++ {
			if d == 'L' {
				dial = (dial - 1 + 100) % 100
			} else {
				dial = (dial + 1) % 100
			}
			if dial == 0 {
				p2++
			}
		}
		if dial == 0 {
			p1++
		}
	}

	fmt.Printf("part 1 password: %d\n", p1)
	fmt.Printf("part 2 password: %d\n", p2)
}
