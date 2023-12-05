package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2_1() {
	lines := read_lines("inputs/day2_input", "\n")

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
	lines := read_lines("inputs/day2_input", "\n")
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
