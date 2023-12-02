package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func day1_1() {
	bytes, err := os.ReadFile("day1_input")
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	lines := strings.Split(string(bytes), "\n")
	// To remove the last empty item generated after splitting by \n
	lines = lines[:len(lines) - 1]
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

	bytes, err := os.ReadFile("day1_input")
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	lines := strings.Split(string(bytes), "\n")
	// To remove the last empty item generated after splitting by \n
	lines = lines[:len(lines) - 1]
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

func day2_1() {
	bytes, err := os.ReadFile("day2_input")
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	lines := strings.Split(string(bytes), "\n")
	// To remove the last empty item generated after splitting by \n
	lines = lines[:len(lines) - 1]

	max_red := 12
	max_green := 13
	max_blue := 14

	sum := 0

	LINE:
	for _, line := range lines {
		game := line[strings.Index(line, " ")+1 : strings.Index(line, ":")]
		game_num, _ := strconv.Atoi(game)

		reveals_s := line[strings.Index(line, ":")+1:]
		reveals := strings.Split(reveals_s, ";")

		for _, reveal := range reveals {

			balls := strings.Split(reveal, ",")

			for _, ball := range balls {
				ball = strings.Trim(ball, " ")
				num_s := ball[:strings.Index(ball, " ")]
				color := ball[strings.Index(ball, " ")+1:]

				num, _ := strconv.Atoi(num_s)
				if color == "red" && num > max_red {
					continue LINE
				}
				if color == "green" && num > max_green {
					continue LINE
				}
				if color == "blue" && num > max_blue {
					continue LINE
				}
			}
		}

		
		sum += game_num
	}

	fmt.Println(sum)
}

func day2_2() {
	bytes, err := os.ReadFile("day2_input")
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	lines := strings.Split(string(bytes), "\n")
	// To remove the last empty item generated after splitting by \n
	lines = lines[:len(lines) - 1]

	sum := 0

	for _, line := range lines {
		reveals_s := line[strings.Index(line, ":")+1:]
		reveals := strings.Split(reveals_s, ";")

		mins := map[string]int {
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		for _, reveal := range reveals {
			balls := strings.Split(reveal, ",")

			for _, ball := range balls {
				ball = strings.Trim(ball, " ")
				num_s := ball[:strings.Index(ball, " ")]
				color := ball[strings.Index(ball, " ")+1:]

				num, _ := strconv.Atoi(num_s)
				mins[color] = Max(mins[color], num)
			}
		}

		power := mins["red"] * mins["green"] * mins["blue"]
		sum += power
	}

	fmt.Println(sum)
}

func main() {
	day2_2()
}

// Utility functions
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
