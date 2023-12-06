package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day5_1() {
	bytes, err := os.ReadFile("inputs/day5_input")
	if (err != nil) {
		fmt.Println("There was an error reading a file")
	}
	line_blocks := strings.Split(string(bytes), "\n\n")

	seeds := make([]int, 0)
	seeds_s := strings.Split(line_blocks[0][strings.Index(line_blocks[0], ":") + 2:], " ")
	for _, seed := range seeds_s {
		i, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error when converting string to number")
		}
		seeds = append(seeds, i)
	}

	for _, line_block := range line_blocks[1:] {
		lines := strings.Split(strings.TrimRight(line_block, "\n"), "\n")
		
		SEED:
		for seed_idx, seed := range seeds {
			for _, line := range lines[1:] {
				nums := strings.Split(line, " ")
				dest, _ := strconv.Atoi(nums[0])
				src, _ := strconv.Atoi(nums[1])
				rng, _ := strconv.Atoi(nums[2])
				if seed >= src && seed <= src + rng {
					seeds[seed_idx] = seed - src + dest
					continue SEED
				}
			}
		}
	}

	fmt.Println(slices.Min(seeds))
}
