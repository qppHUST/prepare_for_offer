package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type point struct {
	position string
	dist     int
	from     []*point
}

func (point point) String() string {
	s := fmt.Sprintf("position:%s  dist:%d  from: %v", point.position, point.dist, point.from)
	return s
}
