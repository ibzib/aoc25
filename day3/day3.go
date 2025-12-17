package main

import (
	"fmt"
	"strconv"
)

func joltage1(bank []byte) int {
	firstDigit := byte('0')
	secondDigit := byte('0')
	firstDigitIdx := 0
	for i := 0; i < len(bank)-1; i++ {
		if bank[i] > firstDigit {
			firstDigit = bank[i]
			firstDigitIdx = i
		}
	}
	for i := firstDigitIdx + 1; i < len(bank); i++ {
		if bank[i] > secondDigit {
			secondDigit = bank[i]
		}
	}
	joltage, _ := strconv.Atoi(fmt.Sprintf("%c%c", firstDigit, secondDigit))
	return joltage
}

func joltage2(bank []byte) int {
	joltage := 0
	a := 0
	b := len(bank) - 12
	for range 12 {
		max := byte('0')
		maxIdx := 0
		for i := a; i <= b; i++ {
			if bank[i] > max {
				max = bank[i]
				maxIdx = i
			}
		}
		x, _ := strconv.Atoi(string(max))
		joltage = 10*joltage + x
		a = maxIdx + 1
		b++
	}
	return joltage
}

func main() {
	totalJoltage := 0
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		bank := []byte(line)
		// joltage := joltage1(bank)
		joltage := joltage2(bank)
		totalJoltage += joltage
	}
	fmt.Println(totalJoltage)
}
