package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day4_1() {
	lines := read_lines("inputs/day4_input", "\n")
	sum := 0

	for _, line := range lines {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winning := strings.Fields(numbers[0])
		owned := strings.Fields(numbers[1])

		set := make([]int, 0)
		hash := make(map[int]int)
		for _, num := range winning {
			i, _ := strconv.Atoi(num)
			hash[i] = 0
		}

		for _, num := range owned {
			i, _ := strconv.Atoi(num)
			if _, ok := hash[i]; ok {
				set = append(set, i)
			}
		}

		if len(set) == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(len(set) - 1)))
	}

	fmt.Println(sum)
}

func day4_2() {
	lines := read_lines("inputs/day4_input", "\n")
	sum := 0

	for _, line := range lines {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winning := strings.Fields(numbers[0])
		owned := strings.Fields(numbers[1])

		set := make([]int, 0)
		hash := make(map[int]int)
		for _, num := range winning {
			i, _ := strconv.Atoi(num)
			hash[i] = 0
		}

		for _, num := range owned {
			i, _ := strconv.Atoi(num)
			if _, ok := hash[i]; ok {
				set = append(set, i)
			}
		}

	}
	fmt.Println(sum)
}
