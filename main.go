package main

import (
	"fmt"
	"os"
)

func main() {
	pattern := os.Args[1]

	result := match(pattern)
	// fmt.Fprintln(os.Stdout, result)
	fmt.Println(result)
}

func match(pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	count := 0
	for _, v := range pattern {
		if string(v) == " " {
			count++
		}
	}
	if count >= 1 || len(pattern) >= 1 {
		count++
	}
	return count
}
