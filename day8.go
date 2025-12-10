package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// x, y, z
var positionInSpace = []string{
	"162,817,812",
	"57,618,57",
	"906,360,560",
	"592,479,940",
	"352,342,300",
	"466,668,158",
	"542,29,236",
	"431,825,988",
	"739,650,466",
	"52,470,668",
	"216,146,977",
	"819,987,18",
	"117,168,530",
	"805,96,715",
	"346,949,466",
	"970,615,88",
	"941,993,340",
	"862,61,35",
	"984,92,344",
	"425,690,689",
}
type Pos struct {
	id, x, y, z int;
}
type Edge struct {
	distance float64
	p1 int
	p2 int
}

func parsePosFile() []string {
	fileBytes, _ := os.ReadFile("./day8-input.txt")
	return strings.Split(string(fileBytes), "\n")
}

func toNumber(s string) int {
	n, e := strconv.Atoi(s)
	if e != nil {
        panic(e)
    }
	return n
}
func formatPositionInSpace() []Pos {
	positionInSpace = parsePosFile()
	pos := make([]Pos, 0, len(positionInSpace))
	for i, v := range positionInSpace {
		cords := strings.Split(v, ",")
		x := toNumber(cords[0])
		y := toNumber(cords[1])
		z := toNumber(cords[2])
		pos = append(pos, Pos{
			id: i,
			x: x,
			y: y, 
			z: z,
		})
	}
	return pos
}
func calculateDistance(p1 Pos, p2 Pos) float64 {
	distance := math.Sqrt(
		math.Pow(float64(p1.x - p2.x), 2) + 
		math.Pow(float64(p1.y - p2.y), 2) + 
		math.Pow(float64(p1.z - p2.z), 2))
	return distance

}


func calculateAllDistances(data []Pos) []Edge {
	edges := make([]Edge, 0)
	for i, vi := range data {
		for j := 1 + i; j < len(data); j++ {
			if i == j {
				continue
			}

			distanceBetween := calculateDistance(vi, data[j])
			edges = append(edges, Edge{
				distance: distanceBetween,
				p1: vi.id,
				p2: data[j].id,
			})
		}
	}
	return edges
}

func main() {
	connections := make([][]int, 0)
	data := formatPositionInSpace()
	// calculate all positions
	edges := calculateAllDistances(data)
	sort.Slice(edges,  func(i int, j int) bool {
		return edges[i].distance < edges[j].distance
	})
	var lastEdge Edge
	for i  := range edges {
		closestConnection := edges[i]
		if (len(connections) == 0) {
			connections = append(connections, []int{closestConnection.p1, closestConnection.p2})
			continue
		}

		circuit1Index := -1
		circuit2Index := -1
		for i, connection := range connections {
			for _, pointId := range connection {
				if pointId == closestConnection.p1 {
					circuit1Index = i
				}
				if pointId == closestConnection.p2{
					circuit2Index = i
				}
			}
		}
		if circuit1Index != -1 && circuit1Index == circuit2Index {
			continue
		}
		lastEdge = closestConnection
		// different connections need to combine them
		if circuit1Index != -1 && circuit2Index != -1 {
			connections[circuit1Index] = append(connections[circuit1Index], connections[circuit2Index]...)
			connections[circuit2Index] = []int{}
			firstSet := connections[:circuit2Index]
			secondSet := connections[circuit2Index + 1:]
			connections = append(firstSet, secondSet...)
		} else if circuit1Index != -1 {
			connections[circuit1Index] = append(connections[circuit1Index], closestConnection.p2)
		} else if circuit2Index != -1 {
			connections[circuit2Index] = append(connections[circuit2Index], closestConnection.p1)
		} else {
			connections = append(connections, []int{closestConnection.p1, closestConnection.p2})
		}
	}
	sort.Slice(connections,  func(i int, j int) bool {
		return len(connections[i]) > len(connections[j])
	})
	
	fmt.Println(len(connections[0]))
	fmt.Println(lastEdge)
	fmt.Println(data[lastEdge.p1])
	fmt.Println(data[lastEdge.p2])
	// fmt.Println(len(connections[2]))
	// fmt.Println(len(connections[3]))
	// fmt.Println(lastPoints)
}