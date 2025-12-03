package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}
func SliceIndex(limit int, predicate func(i int) bool) int {
    for i := 0; i < limit; i++ {
        if predicate(i) {
            return i
        }
    }
    return -1
}
func parseFile() []string {
	bytes, _ := os.ReadFile("./day3-input.txt")
	fileAsStrings := strings.Split(string(bytes), "\n")
	return fileAsStrings
}
func doSearch(bank string) []int {
	joltage := make([]int, 0, 12)

	lastBatterIndex := -1;
	for len(joltage) < 12 {
		nextBattery := 0;

		for searching := 9; searching > 0; searching-- {
			if nextBattery > 0 {
				break
			}
			for starterIndex := lastBatterIndex + 1; starterIndex < len(bank); starterIndex++ {
				potentialStarter, _ := strconv.Atoi(string(bank[starterIndex]))
				if (potentialStarter == searching) {
					currentBatteries := len(joltage)
					requiredBatteries := 12 - currentBatteries
					remainingInBank := len(bank) - starterIndex
					if remainingInBank < requiredBatteries {
						break
					}
					nextBattery = potentialStarter
					lastBatterIndex = starterIndex
					joltage = append(joltage, nextBattery)
					break
				}
			}
		}
	}
	return joltage
}

func main() {
	sum := 0
	input = parseFile()
	for _, value := range input {
	
		joltage := doSearch(value )

		var joltageString string
		for _, joltageNumber := range joltage {
			joltageString += strconv.Itoa(joltageNumber)
		}
		completeJoltage, _ := strconv.Atoi(string(joltageString))
		sum += completeJoltage
	}
	fmt.Println("SUM joltage: ", sum)
}