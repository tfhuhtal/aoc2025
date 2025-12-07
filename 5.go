package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	var ranges [][2]int64

	for _, line := range lines {
		bounds := strings.Split(line, "-")
		start, _ := strconv.ParseInt(bounds[0], 10, 64)
		end, _ := strconv.ParseInt(bounds[1], 10, 64)
		ranges = append(ranges, [2]int64{start, end})
	}

	var p1, p2 int64

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var curr int64 = -1
	for _, r := range ranges {
		start, end := r[0], r[1]
		if curr >= start {
			start = curr + 1
		}
		if start <= end {
			p2 += end - start + 1
		}
		curr = max(int64(curr), int64(end))
	}

	for scanner.Scan() {
		ing, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fresh := false
		for _, r := range ranges {
			if ing >= r[0] && ing <= r[1] {
				fresh = true
			}
		}

		if fresh {
			p1++
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
