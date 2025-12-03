package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = []string{
	"11-22",
"95-115",
"998-1012",
"1188511880-1188511890",
"222220-222224", 
"1698522-1698528",
"446443-446449", 
"38593856-38593862", 
"565653-565659",
"824824821-824824827",
"2121212118-2121212124"}


func getInputFromFile() []string {
		bytes, err := os.ReadFile("./day2-input.txt")
	if (err != nil) {
			fmt.Println(err)
			panic(err)
	}
	fileAsString := strings.Split(string(bytes), ",")
	return fileAsString
}


func main() {
	sum := 0
	input = getInputFromFile()
	for _, value := range input {
		splitString := strings.Split(value, "-")
		startingNumber, _ := strconv.Atoi(splitString[0])
		endingNumber, err := strconv.Atoi(splitString[1])
		if (err != nil) {
			fmt.Println(err)
			return
		}
		for id := startingNumber; id <= endingNumber; id++ {
			stringId := strconv.Itoa(id)
			idLength := len(stringId)
			if idLength < 2 {
				continue
			}
			for idx := 0; idx < idLength; idx++ {
				digitToCheck := stringId[:idx + 1]
				digitLength := len(digitToCheck)
				digits := strings.Split(stringId, "")
				isValid := true
				for j := digitLength; j < len(digits); j+=digitLength {
					if j + digitLength > len(digits) {
						isValid = true;
						continue
					}
					nextChunk := strings.Join(digits[j:j+digitLength], "")
					if nextChunk == digitToCheck {
						isValid = false
						continue
					}
					isValid = true;
					break
				}
				if !isValid {
					sum += id;
					break;
				}				
			}
		}
	}
	fmt.Println("Invalid ids:", sum)
	// part1 13108371860
	// part2 22471660255
}
