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

func addPaths(q map[int]int, i int, nPaths int) {
	if _, ok := q[i]; !ok {
		q[i] = 0
	}
	q[i] += nPaths
}

func pt2(lines []string) {
	// map of indexes -> number of unique paths to that index
	q := map[int]int{}

	// find starting position
	for i, c := range lines[0] {
		if c == 'S' {
			q[i] = 1
		}
	}

	// fire mah lazer
	for _, line := range lines[1:] {
		nextQ := map[int]int{}
		for i, nPaths := range q {
			if line[i] == '^' {
				// split
				if i > 0 {
					addPaths(nextQ, i-1, nPaths)
				}
				if i < len(line)-1 {
					addPaths(nextQ, i+1, nPaths)
				}
			} else {
				// continue down
				addPaths(nextQ, i, nPaths)
			}
		}
		q = nextQ
	}
	total := 0
	for _, nPaths := range q {
		total += nPaths
	}
	fmt.Println(total)
}

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// pt1(lines)
	pt2(lines)
}
