package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func pt1() {
	// read input
	var s string
	nums := []int{}
	ops := []string{}
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			if err == io.EOF {
				break
			} else if err.Error() == "unexpected newline" {
				continue
			}
		}
		if s == "*" || s == "+" {
			ops = append(ops, s)
		} else {
			x, _ := strconv.Atoi(s)
			nums = append(nums, x)
		}
	}

	// compute answer
	total := 0
	for i, op := range ops {
		var col int
		if op == "+" {
			col = 0
		} else {
			col = 1
		}
		for j := i; j < len(nums); j += len(ops) {
			if op == "+" {
				col += nums[j]
			} else {
				col *= nums[j]
			}
		}
		total += col
	}
	fmt.Println(total)
}

func pt2() {
	// read input
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// find column boundaries
	ops := lines[len(lines)-1]
	colStarts := []int{}
	for i, c := range ops {
		if c == '+' || c == '*' {
			colStarts = append(colStarts, i)
		}
	}

	total := 0
	for i, colStart := range colStarts {
		// normalize columns to rows
		var colEnd int // exclusive
		if i < len(colStarts)-1 {
			colEnd = colStarts[i+1] - 1
		} else {
			colEnd = len(ops)
		}
		numStrs := make([]string, colEnd-colStart)
		for j := colStart; j < colEnd; j++ {
			for _, line := range lines[:len(lines)-1] {
				if line[j] != ' ' {
					numStrs[j-colStart] += string(line[j])
				}
			}
		}

		// compute subtotal
		var subtotal int
		if ops[colStart] == '+' {
			subtotal = 0
		} else {
			subtotal = 1
		}
		for _, s := range numStrs {
			x, _ := strconv.Atoi(s)
			if ops[colStart] == '+' {
				subtotal += x
			} else {
				subtotal *= x
			}
		}
		total += subtotal
	}
	fmt.Println(total)
}

func main() {
	// pt1()
	pt2()
}
