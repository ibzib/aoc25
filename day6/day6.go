package main

import (
	"fmt"
	"io"
	"strconv"
)

func pt1(nums []int, ops []string) {
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

func main() {
	// read input
	var s string
	tokens := []string{}
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			if err == io.EOF {
				break
			} else if err.Error() == "unexpected newline" {
				continue
			} else {
				panic(err)
			}
		}
		tokens = append(tokens, s)
	}

	// parse input
	nums := []int{}
	ops := []string{}
	for _, s := range tokens {
		if s == "*" || s == "+" {
			ops = append(ops, s)
		} else {
			x, err := strconv.Atoi(s)
			if err != nil {
				panic("bad token " + s)
			}
			nums = append(nums, x)
		}
	}

	pt1(nums, ops)
	// TODO pt2
}
