package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type engine_num struct {
	value string
	row int
	start_idx int
	adjacent bool
}

func day3_1() {
	lines := read_lines("inputs/day3_input", "\n")
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
