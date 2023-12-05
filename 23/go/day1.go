package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func day1_1() {
	lines := read_lines("inputs/day1_input", "\n")
	sum := 0
	for line_idx, line := range lines {
		var number string
		for _, char := range line {
			if (unicode.IsDigit(rune(char))) {
				number += string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if (unicode.IsDigit(rune(line[i]))) {
				number += string(line[i])
				break
			}
		}
		i, err := strconv.Atoi(number)
		if (err != nil) {
			fmt.Println("An error during converting to numbers")
		}
		if line_idx < 10 {
			fmt.Println(i)
		}
		sum += i
	}

	fmt.Println(sum)
}

func day1_2() {
	digits_s := []string{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" }
	digits := []string{ "0", "1", "2", "3", "4", "5", "6", "7", "8", "9" }

	lines := read_lines("inputs/day1_input", "\n")
	sum := 0
	for _, line := range lines {
		first_idx := math.MaxInt
		last_idx := math.MinInt
		var first string
		var last string

		for _, digit := range digits_s {
			idx := strings.Index(line, digit)
			if idx == -1 {
				continue
			}
			if idx < first_idx {
				first = digit
				first_idx = idx
			}
			idx = strings.LastIndex(line, digit)
			if idx > last_idx {
				last = digit
				last_idx = idx
			}
		}

		for _, digit := range digits {
			idx := strings.Index(line, digit)
			if idx == -1 {
				continue
			}
			if idx < first_idx {
				first = digit
				first_idx = idx
			}
			idx = strings.LastIndex(line, digit)
			if idx > last_idx {
				last = digit
				last_idx = idx
			}
		}

		var number string
		if len(first) == 1 {
			number += first
		} else {
			number += digits[slices.Index(digits_s, first) + 1]
		}
		if len(last) == 1 {
			number += last
		} else {
			number += digits[slices.Index(digits_s, last) + 1]
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Failed parsing a number")
		}
		sum += n
	}
	fmt.Println(sum)
}
