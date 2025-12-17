package main

import (
	"fmt"
	"strconv"
)

func invalidIds1() map[int]struct{} {
	// Generate set of invalid IDs
	invalidIds := make(map[int]struct{})
	for i := 1; i <= 99999; i++ {
		id, _ := strconv.Atoi(fmt.Sprintf("%d%d", i, i))
		invalidIds[id] = struct{}{}
	}
	return invalidIds
}

func invalidIds2() map[int]struct{} {
	// Generate set of invalid IDs
	invalidIds := make(map[int]struct{})
	for i := 1; i <= 99999; i++ {
		for idStr := fmt.Sprintf("%d%d", i, i); len(idStr) <= 10; idStr += strconv.Itoa(i) {
			id, _ := strconv.Atoi(idStr)
			invalidIds[id] = struct{}{}
		}
	}
	return invalidIds
}

func main() {
	// invalidIds := invalidIds1()
	invalidIds := invalidIds2()

	// Example input (line-wrapped):
	// 11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
	// 1698522-1698528,446443-446449,38593856-38593862,565653-565659,
	// 824824821-824824827,2121212118-2121212124
	ans := 0
	var a, b int
	for {
		_, err := fmt.Scanf("%d-%d", &a, &b)
		if err != nil {
			break
		}
		for i := a; i <= b; i++ {
			if _, invalid := invalidIds[i]; invalid {
				ans += i
			}
		}
	}
	fmt.Println(ans)
}
