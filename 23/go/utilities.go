package main

import (
	"fmt"
	"os"
	"strings"
)

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
