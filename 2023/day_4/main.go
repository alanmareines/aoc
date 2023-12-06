package day_4

import (
	"aoc/utils"
	"regexp"
	"strings"
)

func Puzzle1() int {
	var total int
	fileString := utils.InputToString("./day_4/input.txt")

	lines := strings.Split(fileString, "\n")

	for _, l := range lines {
		cards := strings.Split(l, ": ")[1]

		winnerString := strings.Split(cards, " | ")[0]
		playedString := strings.Split(cards, " | ")[1]

		winners := regexp.MustCompile(`\d+`).FindAllString(winnerString, -1)
		played := regexp.MustCompile(`\d+`).FindAllString(playedString, -1)

		var winningNums []string
		for _, w := range winners {
			for _, p := range played {
				if w == p {
					winningNums = append(winningNums, w)
				}
			}
		}

		if len(winningNums) > 0 {
			sum := 1

			for i := 1; i < len(winningNums); i++ {
				sum = sum * 2
			}

			total += sum
		}

	}

	return total
}

func Puzzle2() int {
	var total int
	fileString := utils.InputToString("./day_4/input.txt")

	lines := strings.Split(fileString, "\n")

	cardSheets := make(map[int]int)
	for i := range lines {
		cardSheets[i] = 1
	}

	for i, l := range lines {
		cards := strings.Split(l, ": ")[1]

		winnerString := strings.Split(cards, " | ")[0]
		playedString := strings.Split(cards, " | ")[1]

		winners := regexp.MustCompile(`\d+`).FindAllString(winnerString, -1)
		played := regexp.MustCompile(`\d+`).FindAllString(playedString, -1)

		var winningNums []string
		for _, w := range winners {
			for _, p := range played {
				if w == p {
					winningNums = append(winningNums, w)
				}
			}
		}

		for j := 0; j < cardSheets[i]; j++ {
			for k := range winningNums {
				cardSheets[i+k+1] = cardSheets[i+k+1] + 1
			}
		}
	}

	for _, num := range cardSheets {
		total += num
	}

	return total
}
