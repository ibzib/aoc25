package main

import "fmt"

type Vector struct {
	i int
	j int
}

var deltas = []Vector{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func isAccessible(a [][]byte, i int, j int) bool {
	adj := 0
	for _, d := range deltas {
		ii := i + d.i
		jj := j + d.j
		if ii >= 0 && ii < len(a) && jj >= 0 && jj < len(a[0]) && a[ii][jj] == '@' {
			adj++
		}
	}
	return adj < 4
}

func countRolls(a [][]byte) int {
	n := 0
	for i, s := range a {
		for j, c := range s {
			if c == '@' && isAccessible(a, i, j) {
				n++
			}
		}
	}
	return n
}

func removeRolls(a [][]byte) int {
	n := 0
	q := []Vector{}
	for i, s := range a {
		for j, c := range s {
			if c == '@' && isAccessible(a, i, j) {
				q = append(q, Vector{i, j})
			}
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if a[v.i][v.j] != '@' || !isAccessible(a, v.i, v.j) {
			continue
		}
		// Remove
		n++
		a[v.i][v.j] = 'x'
		// Remove neighbors
		for _, d := range deltas {
			ii := v.i + d.i
			jj := v.j + d.j
			if ii >= 0 && ii < len(a) && jj >= 0 && jj < len(a[0]) && a[ii][jj] == '@' {
				q = append(q, Vector{ii, jj})
			}
		}
	}
	for _, s := range a {
		fmt.Println(string(s))
	}
	return n
}

func main() {
	a := [][]byte{}
	var s string
	for {
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			break
		}
		a = append(a, []byte(s))
	}
	// n := countRolls(a)
	n := removeRolls(a)
	fmt.Println(n)
}
