package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func invalid(x int, p2 bool) bool {
	s := strconv.Itoa(x)
	n := len(s)

	var kStart, kEnd int
	kStart = 2
	if p2 {
		kEnd = n + 1
	} else {
		kEnd = 3
	}

	for k := kStart; k < kEnd; k++ {
		if n%k == 0 {
			ok := true
			sz := n / k
			i := 0
			for i < n {
				if s[i:i+sz] != s[:sz] {
					ok = false
				}
				i += sz
			}
			if ok {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	p1 := 0
	p2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			r = strings.TrimSpace(r)
			if r == "" {
				continue
			}
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				continue
			}
			first, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			last, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 != nil || err2 != nil {
				continue
			}
			for x := first; x <= last; x++ {
				if invalid(x, false) {
					p1 += x
				}
				if invalid(x, true) {
					p2 += x
				}
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
