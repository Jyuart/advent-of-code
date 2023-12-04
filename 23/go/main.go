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
	lines := read_lines("day1_input")
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

	lines := read_lines("day1_input")
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
	lines := read_lines("day2_input")

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
	lines := read_lines("day2_input")
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
				mins[color] = max(mins[color], num)
			}
		}

		power := mins["red"] * mins["green"] * mins["blue"]
		sum += power
	}

	fmt.Println(sum)
}

type engine_num struct {
	value string
	row int
	start_idx int
	adjacent bool
}

func day3_1() {
	lines := read_lines("day3_input")
	nums := []engine_num{}
	line_len := len(lines[0])

	// looping over the lines in the file
	for line_idx, line := range lines {
		var num engine_num
		num.start_idx = -1

		// looping over the characters in the line
		for idx, char := range line {
			
			// if the current char is a digit
			if unicode.IsDigit(char) {

				// if it's a start of a number
				if num.start_idx == -1 {
					num.row = line_idx
					num.start_idx = idx

					// if the char before the start is a symbol
					if idx != 0 && string(line[idx-1]) != "." {
						num.adjacent = true
					}
				}

				num.value += string(char)

				// covering cases if the number ends at the end of the line
				if idx + 1 == line_len {
					nums = append(nums, num)

					num.start_idx = -1
					num.value = ""
					num.adjacent = false
				}

				continue
			}

			// if it's symbol after the last in the prev number
			if idx != 0 && unicode.IsDigit(rune(line[idx-1])) {
				// if the char after the end is a symbol
				if idx < len(line) && string(char) != "." {
					num.adjacent = true
				}
				nums = append(nums, num)

				num.start_idx = -1
				num.value = ""
				num.adjacent = false
			}
		}
	}

	lines_len := len(lines)
	sum := 0

	for _, num := range nums {
		for i := max(num.start_idx - 1, 0); i < min(line_len, num.start_idx + len(num.value) + 1); i++ {
			if num.row != 0 {
				if string(lines[num.row - 1][i]) != "." {
					num.adjacent = true
					break
				}
			}

			if num.row + 1 != lines_len {
				if string(lines[num.row + 1][i]) != "." {
					num.adjacent = true
					break
				}
			}
		}

		if num.adjacent {
			fmt.Println(num.value)
			i, err := strconv.Atoi(num.value)
			if err != nil {
				fmt.Println("Error when converting value to int")
			}	

			sum += i
		}
	}
	
	for _, num := range nums {
		if num.adjacent {

		}
	}

	fmt.Println(sum)
}

func day4_1() {
	lines := read_lines("d4_test")
	sum := 0

	numbers_count := len(lines[0]) - 1
	for _, line := range lines {
		numbers := strings.Split(line[strings.Index(line, ":")+1:], " ")

		fmt.Println(numbers_count)

		// -1 because we have a "|" there
		set := make(map[int]int)
		for _, num := range numbers {
			i, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			set[i] = 0
		}

		wins_count := numbers_count - len(set)
		fmt.Println(numbers_count)
		// fmt.Println(2^max((wins_count-1), 0))
		sum += 2^max((wins_count-1), 0)
	}

	fmt.Println(sum)
}

func main() {
	day4_1()
}

// Utility functions
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func read_lines(file_name string) []string {
	bytes, err := os.ReadFile(file_name)
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	lines := strings.Split(string(bytes), "\n")
	// To remove the last empty item generated after splitting by \n
	return lines[:len(lines) - 1]
}
