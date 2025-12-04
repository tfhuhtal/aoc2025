package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	p1 := 0
	p2 := 0

	for scanner.Scan() {
		line := scanner.Text()
	}

	fmt.Printf("part 1 password: %d\n", p1)
	fmt.Printf("part 2 password: %d\n", p2)
}
