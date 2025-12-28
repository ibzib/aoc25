package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	a int
	b int
}

type ByA []Interval

func (s ByA) Len() int           { return len(s) }
func (s ByA) Less(i, j int) bool { return s[i].a < s[j].a }
func (s ByA) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func pt1() {
	var intervals []Interval
	for {
		var iv Interval
		_, err := fmt.Scanf("%d-%d", &iv.a, &iv.b)
		if err != nil {
			break
		}
		intervals = append(intervals, iv)
	}
	ans := 0
	// O(N*M), not optimal
	for {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			break
		}
		for _, iv := range intervals {
			if x >= iv.a && x <= iv.b {
				ans++
				break
			}
		}
	}
	fmt.Println(ans)
}

func main() {
	var intervals []Interval
	for {
		var iv Interval
		_, err := fmt.Scanf("%d-%d", &iv.a, &iv.b)
		if err != nil {
			break
		}
		intervals = append(intervals, iv)
	}
	sort.Sort(ByA(intervals))
	start := 0
	end := -1
	ans := 0
	for _, iv := range intervals {
		if iv.a <= end {
			if iv.b >= end {
				// Merge intervals
				end = iv.b
			}
		} else {
			ans += end - start + 1
			start = iv.a
			end = iv.b
		}
	}
	ans += end - start + 1
	fmt.Println(ans)
}
