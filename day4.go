package main

import (
	"fmt"
	"os"
	"strings"
)



var rollsInput = []string{
"..@@.@@@@.",
"@@@.@.@.@@",
"@@@@@.@.@@",
"@.@@@@..@.",
"@@.@@@@.@@",
".@@@@@@@.@",
".@.@.@.@@@",
"@.@@@.@@@@",
".@@@@@@@@.",
"@.@.@@@.@.",
}

func parseRolls() []string {
	bytes, _ := os.ReadFile("./day4-input.txt")
	fileAsString := strings.Split(string(bytes), "\n")
	return fileAsString

}

var emptySpace = "."
var rollSpace = "@"

func countRolls(args... string) int {
	sum := 0;
	for _, val := range args {
		if val == rollSpace {
			sum++
		}
	}
	return sum
}

func main() {
	accessibleRolls := 0
	rollsInput = parseRolls()
	checkCount := func (count int) {
		if count < 4 {
			accessibleRolls++
		}
	}
	
	for i, value := range rollsInput {
		isTop := i == 0
		isBottom := i == len(rollsInput) - 1
		for j, space := range value {
			spaceCharacter := string(space)
			if spaceCharacter == emptySpace {
				continue
			}

			// check surroundings
			isLeft := j == 0
			isRight := j == len(value) - 1
			// Corners are always accessible
			if ((isTop && isLeft) ||(isTop && isRight) || (isBottom && isLeft) || (isBottom && isRight)) {
				accessibleRolls++
				continue
			}
			if isTop {
				right := value[j + 1]			
				bottomRight := rollsInput[i + 1][j + 1]
				bottom := rollsInput[i + 1][j]
				bottomLeft := rollsInput[i + 1][j - 1]
				left := value[j - 1]			
				count := countRolls(string(right), string(bottomRight), string(bottom), string(bottomLeft),string(left))
				checkCount(count)
				continue
			}
			if isBottom {
				right := value[j + 1]			
				topRight := rollsInput[i - 1][j + 1]
				top := rollsInput[i - 1][j]
				topLeft := rollsInput[i - 1][j - 1]
				left := value[j - 1]			
				count := countRolls(string(right), string(topRight), string(top), string(topLeft),string(left))
				checkCount(count)
				continue
			}
			if isLeft {
				right := value[j + 1]			
				topRight := rollsInput[i - 1][j + 1]
				top := rollsInput[i - 1][j]
				bottom := rollsInput[i + 1][j]
				bottomRight := rollsInput[i + 1][j + 1]			
				count := countRolls(string(right), string(topRight), string(top), string(bottomRight),string(bottom))
				checkCount(count)
				continue
			}
			if isRight {
				left := value[j - 1]			
				topLeft := rollsInput[i - 1][j - 1]
				top := rollsInput[i - 1][j]
				bottom := rollsInput[i + 1][j]
				bottomLeft := rollsInput[i + 1][j - 1]			
				count := countRolls(string(left), string(topLeft), string(top), string(bottomLeft),string(bottom))
				checkCount(count)
				continue
			}
			left := value[j - 1]			
			topLeft := rollsInput[i - 1][j - 1]
			top := rollsInput[i - 1][j]
			bottom := rollsInput[i + 1][j]
			bottomLeft := rollsInput[i + 1][j - 1]
			right := value[j + 1]			
			topRight := rollsInput[i - 1][j + 1]	
			bottomRight := rollsInput[i + 1][j + 1]			
		
			count := countRolls(
				string(left), string(topLeft), string(top), 
				string(bottomLeft),string(bottom), string(right),
				string(topRight),string(bottomRight))
			checkCount(count)


		}
		
	}
	fmt.Println("Accessible rolls: ", accessibleRolls)
}