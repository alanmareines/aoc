package main

import (
	"aoc/day_1"
	"aoc/day_2"
	"aoc/day_3"
	"aoc/day_4"
	"fmt"
)

func main() {
	fmt.Println("\t----Day 1----\t")
	fmt.Println("First Puzzle:\t", day_1.Puzzle1())
	fmt.Println("Second Puzzle:\t", day_1.Puzzle2())

	fmt.Println("\t----Day 2----\t")
	fmt.Println("First Puzzle:\t", day_2.Puzzle1())
	fmt.Println("Second Puzzle:\t", day_2.Puzzle2())

	fmt.Println("\t----Day 3----\t")
	fmt.Println("First Puzzle:\t", day_3.Puzzle1())
	fmt.Println("Second Puzzle:\t", day_3.Puzzle2())

	fmt.Println("\t----Day 4----\t")
	fmt.Println("First Puzzle:\t", day_4.Puzzle1())
	fmt.Println("Second Puzzle:\t", day_4.Puzzle2())
}
