package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
type CodeData struct {
	direction string
	value int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// var input = []string {"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

func formatInput() []CodeData {
	bytes, err := os.ReadFile("./day1-input.txt")
	check(err)
	fileAsString := strings.Split(string(bytes), "\n")
	outputData := make([]CodeData, len(fileAsString))
	for idx, value := range fileAsString {
		// fmt.Println(value)
		direction := string(value[0])
		v, err := strconv.Atoi(value[1:])
		check(err)
		outputData[idx] = CodeData{direction, v}
	}
	// fmt.Println(outputData)
	return outputData
}




func main() {
	formattedData := formatInput()
	zeroCount := 0
	currentPosition := 50
	for _, data := range formattedData {
		direction := data.direction
		value := data.value
		fmt.Println(currentPosition)
	if direction == "L" {
	    // Count how many times we cross 0 going left
	   for range value {
		   currentPosition--
			if currentPosition == 0 {
				zeroCount++
			}
			if currentPosition == -1 {
				currentPosition = 99
			}
		}
	} else {
	    // Count how many times we cross 0 going right  
	     for range value {
		
		
			currentPosition++
			if currentPosition == 100 {
				currentPosition = 0
				zeroCount++
			}	
		}
	}
}
	fmt.Println("Hit zero this many times: ", zeroCount)
}
