package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var winner = ""

type lilisipoint struct {
	x int
	y int
}

func main() {
	index := 0
	boarder := [][]string{}
	reader := bufio.NewReader(os.Stdin)
	for index < 3 {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		str = str[:len(str)-2]
		boardLine := strings.Split(str, " ")
		boarder = append(boarder, boardLine)
		index++
	}
	for indexi, itemi := range boarder {
		for indexj, itemj := range itemi {
			point := lilisipoint{x: indexi, y: indexj}
			visit := [][]bool{}
			for i := 0; i < 3; i++ {
				visitLine := make([]bool, 3)
				visit = append(visit, visitLine)
			}
			lilisiTest1Dfs(boarder, itemj, 0, point, visit)
		}
	}
	if len(winner) == 0 {
		fmt.Println("No Winner!")
	} else {
		fmt.Println(winner)
	}
}

func lilisiTest1Dfs(boarder [][]string, name string, num int, point lilisipoint, visit [][]bool) {
	if boarder[point.x][point.y] != name {
		return
	}
	if num == 3 && boarder[point.x][point.y] == name {
		// 递归结束了
		winner = name
		return
	}
	//算出dfs的四个方向
	derect1 := lilisipoint{x: point.x, y: point.y - 1}
	derect2 := lilisipoint{x: point.x - 1, y: point.y}
	derect3 := lilisipoint{x: point.x + 1, y: point.y}
	derect4 := lilisipoint{x: point.x, y: point.y + 1}
	direct5 := lilisipoint{x: point.x - 1, y: point.y - 1}
	direct6 := lilisipoint{x: point.x - 1, y: point.y + 1}
	direct7 := lilisipoint{x: point.x + 1, y: point.y - 1}
	direct8 := lilisipoint{x: point.x + 1, y: point.y + 1}
	points := []lilisipoint{derect1, derect2, derect3, derect4, direct5, direct6, direct7, direct8}
	for _, item := range points {
		if item.x > 0 && item.x < 3 && item.y > 0 && item.y < 3 && !visit[item.x][item.y] {
			newVisit := [][]bool{}
			for _, item := range visit {
				newLine := []bool{}
				newLine = append(newLine, item...)
				newVisit = append(newVisit, newLine)
			}
			newVisit[point.x][point.y] = true
			lilisiTest1Dfs(boarder, name, num+1, item, newVisit)
		}
	}

}
