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
		if val == rollSpace || val == "R" {
			sum++
		}
	}
	return sum
}

func main() {
	accessibleRolls := 0
	rollsCanBeRemove := true
	rollsInput =  parseRolls()
	rollsSlice := make([][]byte, len(rollsInput))
	for i, s := range rollsInput {
		rollsSlice[i] = []byte(s)
	}
	checkCount := func (count int, row int, col int) {
		if count < 4 {
			accessibleRolls++
			rollsSlice[row][col] = 'R'
		}
	}
	for rollsCanBeRemove {

	
	for i, value := range rollsSlice {
		isTop := i == 0
		isBottom := i == len(rollsSlice) - 1
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
				rollsSlice[i][j] = 'R'
				continue
			}
			if isTop {
				right := value[j + 1]			
				bottomRight := rollsSlice[i + 1][j + 1]
				bottom := rollsSlice[i + 1][j]
				bottomLeft := rollsSlice[i + 1][j - 1]
				left := value[j - 1]			
				count := countRolls(string(right), string(bottomRight), string(bottom), string(bottomLeft),string(left))
				checkCount(count, i, j)
				continue
			}
			if isBottom {
				right := value[j + 1]			
				topRight := rollsSlice[i - 1][j + 1]
				top := rollsSlice[i - 1][j]
				topLeft := rollsSlice[i - 1][j - 1]
				left := value[j - 1]			
				count := countRolls(string(right), string(topRight), string(top), string(topLeft),string(left))
				checkCount(count, i, j)
				continue
			}
			if isLeft {
				right := value[j + 1]			
				topRight := rollsSlice[i - 1][j + 1]
				top := rollsSlice[i - 1][j]
				bottom := rollsSlice[i + 1][j]
				bottomRight := rollsSlice[i + 1][j + 1]			
				count := countRolls(string(right), string(topRight), string(top), string(bottomRight),string(bottom))
				checkCount(count, i, j)
				continue
			}
			if isRight {
				left := value[j - 1]			
				topLeft := rollsSlice[i - 1][j - 1]
				top := rollsSlice[i - 1][j]
				bottom := rollsSlice[i + 1][j]
				bottomLeft := rollsSlice[i + 1][j - 1]			
				count := countRolls(string(left), string(topLeft), string(top), string(bottomLeft),string(bottom))
				checkCount(count, i, j)
				continue
			}
			left := value[j - 1]			
			topLeft := rollsSlice[i - 1][j - 1]
			top := rollsSlice[i - 1][j]
			bottom := rollsSlice[i + 1][j]
			bottomLeft := rollsSlice[i + 1][j - 1]
			right := value[j + 1]			
			topRight := rollsSlice[i - 1][j + 1]	
			bottomRight := rollsSlice[i + 1][j + 1]			
		
			count := countRolls(
				string(left), string(topLeft), string(top), 
				string(bottomLeft),string(bottom), string(right),
				string(topRight),string(bottomRight))
				checkCount(count, i, j)


		}
		
	}
	rCount := 0
	for rowIndx, row := range rollsSlice {
		for colIdx, cell := range row {
			if string(cell) == "R" {
				rCount++
				rollsSlice[rowIndx][colIdx] = '.'
			}
		}
	}
	if rCount == 0 {
		rollsCanBeRemove = false
	}
	}


	fmt.Println("Accessible rolls: ", accessibleRolls)
}