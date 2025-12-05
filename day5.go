package main

import (
	"fmt"
	"strconv"
	"strings"
)

var inventory = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

type InventoryRange struct {
	low, high int
}

type ParsedInventory struct {
	freshRange []InventoryRange;
	ids []int

}

func parseInventory() ParsedInventory {
	separtorIndex := -1
	// fileBytes, _ := os.ReadFile("./day5-input.txt")
	// inventory := strings.Split(string(fileBytes), "\n")
	for i, v := range inventory {
		if len(v) == 0 {
			separtorIndex = i
			break
		}
	}
	if separtorIndex < 0 {
		panic("Separator not found")
	}
	rawRanges := inventory[:separtorIndex]
	var parsedRanges []InventoryRange = make([]InventoryRange, 0, len(rawRanges))
	var fresh []int = make([]int, 0, len(rawRanges))

	for _, rangeValue := range rawRanges {
		values := strings.Split(string(rangeValue), "-")
		low, _ := strconv.Atoi(values[0])
		high, _ := strconv.Atoi(values[1])
		if len(parsedRanges) == 0 {
			parsedRanges = append(parsedRanges, InventoryRange{low, high})
			continue
		}
		existingRangeChanged := false
		for _, parsed := range parsedRanges {
			if high > parsed.high && low < parsed.high {
				parsed.high = high
				existingRangeChanged = true
			}
			if low < parsed.low && high > parsed.low {
				parsed.low = low
				existingRangeChanged = true
			}
		}
		if existingRangeChanged {
			continue
		}
		parsedRanges = append(parsedRanges, InventoryRange{low, high})


		for i := low; i < high; i++ {
			itemExists := false
			for j := range fresh {
				if j == i {
					itemExists = true
					break
				}
			}
			if itemExists {
				continue
			}
			fresh = append(fresh, i)
		}
		parsedRanges = append(parsedRanges, InventoryRange{low, high})

		
	}
	fmt.Println("FRESH: ", len(fresh))
	parsedIds := make([]int, 0, len(inventory[separtorIndex + 1:]))
	for _, id := range inventory[separtorIndex + 1:] {
		intId, _ := strconv.Atoi(id)
		parsedIds = append(parsedIds, intId)
	}
	return ParsedInventory{freshRange: parsedRanges, ids: parsedIds}
}

func main() {
	freshItems := 0
	inventory := parseInventory()
		for _, freshRange := range inventory.freshRange {
			freshItems = freshRange.high - freshRange.low
	}
	fmt.Println("Fresh Items: ", inventory.freshRange)
	fmt.Println("Fresh Items: ", freshItems)
}