package graph

import "math"

//邻接表 以数字标识节点
type AdjacencyList map[int][]int

//邻接矩阵
type AdjacencyMatrix map[int][]int

type PointWithFrom struct {
	position int
	from     int
}

type PointWithDist struct {
	position int
	dist     int
	from     int
}

// 有向无权图的单源最短路径问题，根据给定的节点给出生成好的表
func singleSourceShortestPathWithoutWeight(graph AdjacencyList, s int) [][]int {
	queue := []PointWithFrom{
		{position: s, from: s},
	}
	//四维表格  点 是否查看 长度 from
	table := initTableForNoWeight(len(graph))
	for len(queue) > 0 {
		point := queue[0]
		if table[point.position-1][1] == 1 {
			queue = queue[1:]
			continue
		}
		if point.position == s {
			line := table[point.position-1]
			line[1] = 1
			line[2] = 0
			line[3] = s
		} else {
			line := table[point.position-1]
			line[1] = 1
			line[2] = table[point.from-1][2] + 1
			line[3] = point.from
		}
		queue = queue[1:]
		list := graph[point.position]
		for _, item := range list {
			toInsert := PointWithFrom{
				position: item,
				from:     point.position,
			}
			queue = append(queue, toInsert)
		}
	}
	return table
}

func initTableForNoWeight(size int) [][]int {
	ans := [][]int{}
	for i := 0; i < size; i++ {
		item := []int{}
		item = append(item, i+1, 0, 0, 0)
		ans = append(ans, item)
	}
	return ans
}

// 有向有权图的单源最短路径问题,根据给定的节点给出生成好的表
func singleSourceShortestPathWithWeight(graph AdjacencyMatrix, s int) [][]int {
	queue := []PointWithDist{
		{position: s, dist: 0, from: s},
	}
	table := initTableForWeight(len(graph))
	for len(queue) > 0 {
		point := queue[0]
		if point.dist < table[point.position-1][1] {
			line := table[point.position-1]
			line[1] = point.dist
			line[2] = point.from
			queue = queue[1:]
			dists := graph[point.position]
			for index, item := range dists {
				if item != 0 {
					toInsert := PointWithDist{
						position: index + 1,
						from:     point.position,
						dist:     point.dist + item,
					}
					queue = append(queue, toInsert)
				}
			}
			sortQueue(queue)
		} else {
			queue = queue[1:]
		}
	}
	return table
}

func initTableForWeight(size int) [][]int {
	ans := [][]int{}
	for i := 0; i < size; i++ {
		item := []int{i + 1, math.MaxInt64, 0}
		ans = append(ans, item)
	}
	return ans
}

//选择排序
func sortQueue(list []PointWithDist) {
	for i := 0; i < len(list)-1; i++ {
		min := i
		for j := i; j < len(list); j++ {
			if list[j].dist < list[min].dist {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}
