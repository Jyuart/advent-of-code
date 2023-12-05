package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type seed_map struct {
	source_from int
	source_to int
	dest_from int
	dest_to int
}

func day5_1() {
	lines_blocks := read_lines("inputs/test5", "\n\n")

	seeds := make([]int, 0)
	seeds_s := strings.Split(lines_blocks[0][strings.Index(lines_blocks[0], ":") + 2:], " ")
	for _, seed := range seeds_s {
		i, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error when converting string to number")
		}
		seeds = append(seeds, i)
	}

	BLOCK:
	for _, line_block := range lines_blocks[1:] {
		lines := strings.Split(line_block, "\n")

		for _, line := range lines[1:] {
			nums := strings.Split(line, " ")
			dest, _ := strconv.Atoi(nums[0])
			src, _ := strconv.Atoi(nums[1])
			rng, _ := strconv.Atoi(nums[2])

			for seed_idx, seed := range seeds {
				if seed >= src && seed < src + rng {
					seeds[seed_idx] = seed - src + dest
					continue BLOCK
				}
			}
		}
	}

	fmt.Println(slices.Min(seeds))
}
