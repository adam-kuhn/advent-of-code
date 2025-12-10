package main

import (
	"fmt"
	"os"
	"strings"
)

var beamLayout = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}
func removeDuplicate[T comparable](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func parseBeamFile() []string {
	fileBytes, _ := os.ReadFile("./day7-input.txt")
	return strings.Split(string(fileBytes), "\n")
}
type PointCount struct {
	point, count int
}
func main() {
	beamLayout = parseBeamFile()
	startingPoint := strings.Index(beamLayout[0], "S")
	beamPositions := make([]int, 0)
	beamPositions = append(beamPositions, startingPoint)
	pathC := make(map[int]int)
	pathC[startingPoint] = 1
	for _, row := range beamLayout[1:] {
        newCounts := make(map[int]int)
		for pos, count := range pathC {
			if string(row[pos]) == "^" {
				left := pos - 1;
				right := pos + 1;
	            newCounts[left] += count
                newCounts[right] += count
			} else {
                newCounts[pos] += count
            }
		}
		pathC = newCounts
	}
	verses := 0
	
    
    // Sum all timeline counts
    for _, count := range pathC {
        verses += count
    }
	// could just do this once
	// pathCounts := make([]PointCount, 0)
	// pathC := make(map[int] int)
	fmt.Println(len(beamPositions))

	fmt.Println("VERSES ", verses)
	// 10733529153890
}