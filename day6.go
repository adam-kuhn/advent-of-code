package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rawFormulas = []string{
	"123 328  51 64 ",
 	" 45 64  387 23 ",
 	"  6 98  215 314",
	"*   +   *   +",
}

func parseForumalFile() []string {
	fileBytes, _ := os.ReadFile("./day6-input.txt")
	return strings.Split(string(fileBytes), "\n")
}

// only should filter on operator
func filter(items []string) []string {
	filteredItems := make([]string, 0)
	for _, v := range items {
		if v == "" {
			continue
		}
		filteredItems = append(filteredItems, v)
	}
	return filteredItems
}
// go through rows line by line
// split on " "
// then go across each element in the array and for each lower column
// if blank check below, if all blank start the next column of numbers and build again

func calculateFormula(vals []string, operator string) int {
	sum := 0
	for _, v := range vals {
		num, _ := strconv.Atoi(v)
		if sum == 0 || operator == "+" {
			sum += num
			continue
		}
		num += sum
	}
	return sum
}

func findSeparators(rawRows []string) []int {
	topRow := rawRows[0]
	indexes := make([]int, 0, 10)
	indexes = append(indexes, -1)
	for ei, e := range strings.Split(topRow, "") {
		blankCount := 0
		if e == " " {
			blankCount++
			for _, v := range rawRows[1:] {
				for bi, b := range strings.Split(v, "") {
					if ei == bi && b == " " {
						blankCount++
					}
				}
			}
		}
		if blankCount == len(rawRows) {
			indexes = append(indexes, ei)
			blankCount = 0
		}
	}
	return indexes
}

func madfgin() {
	sum := 0
	rawFormulas = parseForumalFile()
	operators := filter(strings.Split(rawFormulas[len(rawFormulas) - 1], " "))
	rawRows := rawFormulas[:len(rawFormulas) - 1]
	indexesOfColumns := findSeparators(rawRows)
	formulats := make([][]string, len(operators))
	for _, row := range rawRows {
		for i, v := range indexesOfColumns {
			rowElements := strings.Split(row, "")
			start := v + 1
			end := 0
			if i == len(indexesOfColumns) - 1 {
				end = len(rowElements)
			} else {
				end = indexesOfColumns[i + 1]
			}
			elements := rowElements[start:end]
			value := ""
			for _, e := range elements {
				value += e
			}
			formulats[i] = append(formulats[i], value)
		}
	}

	for o, operator := range operators {
		solution := 0
		formula := formulats[o]
		// find longest
		longestNumber := 0
		for _, d := range formula {
			if len(d) > longestNumber {
				longestNumber = len(d)
			}
		}
		// build numbers
		numbers := make([]int, 0, len(formula))
		for i := 0; i < longestNumber; i++ {
			builtNumber := ""
			for _, v := range formula {
				digit := string(v[i])
				if digit == " " {
					continue
				}
				builtNumber += digit
			}
			number, _ := strconv.Atoi(builtNumber)
			numbers = append(numbers, number)
		}
		for _, n := range numbers {
			if operator == "+" || solution == 0 {
				solution += n
				continue
			}
			solution *= n

		}
		sum += solution

	}



	fmt.Println("INDEXES", indexesOfColumns)
	fmt.Println("Form", formulats)
	fmt.Println("LE", len(formulats[0][2]))
	fmt.Println("LE", len(formulats[1][0]))
	fmt.Println("TOTAL", sum)
}
