package main

import "fmt"

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func main() {
	pw := 0
	pos := 50
	var dir string // "L" or "R"
	var x int
	for {
		_, err := fmt.Scanf("%1s%d\n", &dir, &x)
		if err != nil {
			break
		}
		zero := pos == 0
		if dir == "L" {
			pos -= x
		} else {
			pos += x
		}
		pw += abs(pos / 100)
		if !zero && pos <= 0 {
			pw++
		}
		pos %= 100
		if pos < 0 {
			pos += 100
		}
	}
	fmt.Println(pw)
}
