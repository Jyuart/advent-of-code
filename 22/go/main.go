package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve_day1_1() int {
	data, err := os.ReadFile("./day1_input")
	check(err)
	all_calories := string(data)
	per_elf_calories := strings.Split(all_calories, "\n\n")

	max := 0
	for i := 0; i < len(per_elf_calories); i++ {
		elf_calories := strings.Split(per_elf_calories[i], "\n")
		elf_calories_total := 0

		for j := 0; j < len(elf_calories); j++ {
			calories, _ := strconv.Atoi(elf_calories[j])
			elf_calories_total += calories
		}

		if elf_calories_total > max {
			max = elf_calories_total
		}
	}

	return max
}

func solve_day1_2() int {
	data, err := os.ReadFile("./day1_input")
	check(err)
	all_calories := string(data)
	per_elf_calories := strings.Split(all_calories, "\n\n")

	top3 := []int{ 0, 0, 0 }
	for i := 0; i < len(per_elf_calories); i++ {
		elf_calories := strings.Split(per_elf_calories[i], "\n")
		elf_calories_total := 0

		for j := 0; j < len(elf_calories); j++ {
			calories, _ := strconv.Atoi(elf_calories[j])
			elf_calories_total += calories
		}
		top4 := top3
		top4 = append(top4, elf_calories_total)
		slices.Sort(top4)
		top3 = top4[1:]
	}
	var result int
	for _, i := range top3 {
		result += i
	}
	return result
}

type game struct {
	first string
	second string
}

func solve_day2_1() int {
	data, err := os.ReadFile("./day2_input")
	check(err)
	all_games := string(data)
	games := strings.Split(all_games, "\n")
	games = games[:len(games)-1]
	choice_scores := make(map[string]int)
	result_scores := make(map[game]int)

	choice_scores["X"] = 1
	choice_scores["Y"] = 2
	choice_scores["Z"] = 3
	result_scores[game{first: "A", second: "X"}] = 3
	result_scores[game{first: "A", second: "Y"}] = 6
	result_scores[game{first: "A", second: "Z"}] = 0
	result_scores[game{first: "B", second: "X"}] = 0
	result_scores[game{first: "B", second: "Y"}] = 3
	result_scores[game{first: "B", second: "Z"}] = 6
	result_scores[game{first: "C", second: "X"}] = 6
	result_scores[game{first: "C", second: "Y"}] = 0
	result_scores[game{first: "C", second: "Z"}] = 3

	total_score := 0
	for _, g := range games {
		total_score += choice_scores[string(g[2])]
		theirs := string(g[0])
		mine := string(g[2])
		if theirs == mine {
			total_score += 3
			continue
		}
		total_score += result_scores[game{first: theirs, second: mine}]
	}
	
	return total_score
}

func main() {
	// day1_1_solution := solve_day1_1()
	// day1_2_solution := solve_day1_2()
	day2_1_solution := solve_day2_1()
	fmt.Println(day2_1_solution)
}
