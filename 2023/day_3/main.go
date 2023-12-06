package day_3

import (
	"aoc/utils"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1() int {
	fileString := utils.InputToString("./day_3/input.txt")

	lines := strings.Split(fileString, "\n")

	nums := make(map[int][]int)
	numsPos := make(map[int][][]int)
	specialPos := make(map[int][]int)

	for i, l := range lines {
		stringNums := regexp.MustCompile(`\d+`).FindAllString(l, -1)
		for _, strnum := range stringNums {
			num, err := strconv.Atoi(strnum)
			utils.Check(err)
			nums[i] = append(nums[i], num)
		}

		intIdx := regexp.MustCompile(`\d+`).FindAllStringIndex(l, -1)
		numsPos[i] = intIdx

		specialCharIdx := regexp.MustCompile(`[^0-9\.]`).FindAllStringIndex(l, -1)
		for _, s := range specialCharIdx {
			specialPos[i] = append(specialPos[i], s[0])
		}
	}

	var total int

	for line, num := range nums {
		for i, n := range num {
			start := numsPos[line][i][0]
			end := numsPos[line][i][1]
			if start != 0 {
				start = start - 1
			}
			specialChars := append(specialPos[line], specialPos[line-1]...)
			specialChars = append(specialChars, specialPos[line+1]...)
			for _, s := range specialChars {
				if start <= s && end >= s {
					total += n
					break
				}
			}
		}
	}

	return total
}

func Puzzle2() int {
	fileString := utils.InputToString("./day_3/input.txt")

	lines := strings.Split(fileString, "\n")

	type Number struct {
		num  int
		line int
		pos  []int
	}

	type SpecialNumber struct {
		line int
		pos  int
	}

	nums := make(map[int]Number)
	specialPos := make(map[int]SpecialNumber)

	for i, l := range lines {
		stringNums := regexp.MustCompile(`\d+`).FindAllString(l, -1)
		intIdx := regexp.MustCompile(`\d+`).FindAllStringIndex(l, -1)
		for numidx, strnum := range stringNums {
			num, err := strconv.Atoi(strnum)
			utils.Check(err)
			number := Number{line: i, num: num}

			number.pos = append(number.pos, intIdx[numidx]...)

			nums[len(nums)] = number
		}

		specialCharIdx := regexp.MustCompile(`\*`).FindAllStringIndex(l, -1)

		for _, s := range specialCharIdx {
			specialNumber := SpecialNumber{line: i, pos: s[0]}
			specialPos[len(specialPos)] = specialNumber
		}
	}

	var total int

	var numberAroundAll [][]int
	for _, sp := range specialPos {
		var numberAroundSp []int
		for _, info := range nums {
			if info.line >= sp.line-1 && info.line <= sp.line+1 {
				if info.pos[0]-1 <= sp.pos && info.pos[1] >= sp.pos {
					numberAroundSp = append(numberAroundSp, info.num)
				}
			}
		}

		numberAroundAll = append(numberAroundAll, numberAroundSp)
	}

	var productNums [][]int
	for _, adjacent := range numberAroundAll {
		if len(adjacent) > 1 {
			productNums = append(productNums, adjacent)
		}
	}

	for a := range productNums {
		product := 1
		for b := range productNums[a] {
			product = productNums[a][b] * product
		}

		total += product
	}

	// fmt.Println(specialPos)

	return total
}
