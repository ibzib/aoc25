package main

import (
	"bufio"
	"fmt"
	"os"
)

func pt1(lines []string) {
	// set of indexes to visit next
	q := map[int]struct{}{}

	// find starting position
	for i, c := range lines[0] {
		if c == 'S' {
			q[i] = struct{}{}
		}
	}

	// fire mah lazer
	nSplits := 0
	for _, line := range lines[1:] {
		nextQ := map[int]struct{}{}
		for i, _ := range q {
			if line[i] == '^' {
				// split
				nSplits++
				if i > 0 {
					nextQ[i-1] = struct{}{}
				}
				if i < len(line)-1 {
					nextQ[i+1] = struct{}{}
				}
			} else {
				// continue down
				nextQ[i] = struct{}{}
			}
		}
		q = nextQ
	}
	fmt.Println(nSplits)
}

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	pt1(lines)
}
