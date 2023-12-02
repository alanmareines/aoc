package day_2

import (
	"aoc/utils"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	red, blue, green int
}

func findColorNumber(roundString string, color string) int {
	regexpString := `\d+ ` + color
	matches := regexp.MustCompile(regexpString).FindAllString(roundString, -1)

	if len(matches) == 0 {
		return 0
	}

	stringInt := strings.Split(matches[0], " "+color)[0]
	result, err := strconv.ParseInt(stringInt, 0, 0)

	utils.Check(err)
	return int(result)
}

func parseFile(input string) [][]Round {
	contentString := utils.InputToString(input)
	gamesStrings := strings.Split(contentString, "\n")
	var gamesString [][]string
	var parsedGames [][]Round

	for _, g := range gamesStrings {
		rounds := strings.Split(g, ":")

		sanitizedRounds := strings.Trim(rounds[1], " ")

		gamesString = append(gamesString, strings.Split(sanitizedRounds, "; "))
	}

	for _, g := range gamesString {
		var games []Round
		for _, r := range g {

			round := Round{
				red:   findColorNumber(r, "red"),
				blue:  findColorNumber(r, "blue"),
				green: findColorNumber(r, "green"),
			}

			games = append(games, round)
		}

		parsedGames = append(parsedGames, games)

	}

	return parsedGames
}

func Puzzle1() int {
	games := parseFile("./day_2/input.txt")

	LIMITS := Round{
		red:   12,
		green: 13,
		blue:  14,
	}

	var total int

	for i, g := range games {
		var passesLimit bool
		for _, r := range g {
			if LIMITS.blue < r.blue || LIMITS.green < r.green || LIMITS.red < r.red {
				passesLimit = true
			}
		}
		if !passesLimit {
			total += i + 1
		}
	}

	return total
}

func Puzzle2() int {
	games := parseFile("./day_2/input.txt")
	var maxGameRounds []Round

	for _, g := range games {
		var maxGameRound Round

		for _, r := range g {
			if r.blue > maxGameRound.blue {
				maxGameRound.blue = r.blue
			}
			if r.green > maxGameRound.green {
				maxGameRound.green = r.green
			}
			if r.red > maxGameRound.red {
				maxGameRound.red = r.red
			}
		}

		maxGameRounds = append(maxGameRounds, maxGameRound)
	}

	var total int

	for _, round := range maxGameRounds {
		total += round.blue * round.green * round.red
	}
	return total
}
