package main

import (
	"bufio"
	"fmt"
	"os"
)

type Key struct {
	i    int
	used int
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func pow10(n int) int64 {
	res := int64(1)
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}

func F(line string, i int, used int, dp map[Key]int64) int64 {
	if i == len(line) && used == 12 {
		return 0
	}
	if i == len(line) {
		return -1 << 60
	}
	key := Key{i: i, used: used}
	if v, ok := dp[key]; ok {
		return v
	}
	ans := F(line, i+1, used, dp)

	if used < 12 {
		d := int64(line[i] - '0')
		weight := pow10(11 - used)
		ans = max(ans, weight*d+F(line, i+1, used+1, dp))
	}

	dp[key] = ans
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var p1 int64 = 0
	var p2 int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		dp := make(map[Key]int64)
		p2 += F(line, 0, 0, dp)

		var best int64 = -1
		for i := 0; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				continue
			}
			di := int64(line[i] - '0')
			for j := i + 1; j < len(line); j++ {
				if line[j] < '0' || line[j] > '9' {
					continue
				}
				dj := int64(line[j] - '0')
				score := di*10 + dj
				if best < 0 || score > best {
					best = score
				}
			}
		}
		if best >= 0 {
			p1 += best
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
