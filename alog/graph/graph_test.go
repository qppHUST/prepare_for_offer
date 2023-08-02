package graph

import (
	"fmt"
	"testing"
)

func TestSingleSourceShortestPath(t *testing.T) {
	point1 := []int{4, 2}
	point2 := []int{4, 5}
	point3 := []int{1, 6}
	point4 := []int{3, 5, 6, 7}
	point5 := []int{7}
	point6 := []int{}
	point7 := []int{6}
	graph := AdjacencyList{}
	graph[1] = point1
	graph[2] = point2
	graph[3] = point3
	graph[4] = point4
	graph[5] = point5
	graph[6] = point6
	graph[7] = point7
	table := singleSourceShortestPathWithoutWeight(graph, 3)
	for _, list := range table {
		for _, item := range list {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}

func TestSingleSourceShortestPathWithWeight(t *testing.T) {
	graph := AdjacencyMatrix{
		1: []int{0, 2, 0, 1, 0, 0, 0},
		2: []int{0, 0, 0, 3, 10, 0, 0},
		3: []int{4, 0, 0, 0, 0, 5, 0},
		4: []int{0, 0, 2, 0, 2, 8, 4},
		5: []int{0, 0, 0, 0, 0, 0, 1},
		6: []int{0, 0, 0, 0, 0, 0, 0},
		7: []int{0, 0, 0, 0, 0, 1, 0},
	}
	table := singleSourceShortestPathWithWeight(graph, 3)
	for _, list := range table {
		for _, item := range list {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}
