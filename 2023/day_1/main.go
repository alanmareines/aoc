package day_1

import (
	"aoc/2023/utils"
	"aoc/2023/utils/str"
	"regexp"
	"strconv"
	"strings"
)

func run(words []string) int {
	var total int

	for _, word := range words {

		onlyInt := regexp.MustCompile(`\d`)

		bytes := onlyInt.FindAll([]byte(word), -1)

		doubleDigitString := (string(bytes[0]) + string(bytes[len(bytes)-1]))
		intResult, err := strconv.Atoi(doubleDigitString)
		utils.Check(err)
		total += intResult
	}

	return total

}

func Puzzle1() int {
	contentString := utils.InputToString("./day_1/input.txt")

	return run(strings.Split(contentString, "\n"))

}

func Puzzle2() int {
	var wordToNumber = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	contentString := utils.InputToString("./day_1/input.txt")

	words := strings.Split(strings.Trim(contentString, " \n"), "\n")

	var total int

	for _, w := range words {

		regexStr := `one|two|three|four|five|six|seven|eight|nine`
		revRegex := str.Reverse(regexStr)
		revWord := str.Reverse(w)

		onlyInt := regexp.MustCompile(regexStr + `|\d`)
		onlyInt2 := regexp.MustCompile(revRegex + `|\d`)

		frontwardMatch := onlyInt.FindAllString(w, -1)
		backwardMatch := onlyInt2.FindAllString(revWord, -1)

		first := frontwardMatch[0]
		last := str.Reverse(backwardMatch[0])

		var numberString string

		if len(first) > 1 {
			numberString += wordToNumber[first]
		} else {
			numberString += first
		}

		if len(last) > 1 {
			numberString += wordToNumber[last]
		} else {
			numberString += last
		}

		num, err := strconv.Atoi(numberString)

		utils.Check(err)

		total += num
	}

	return total
}
