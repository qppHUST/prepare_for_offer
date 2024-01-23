package leetcode

import "fmt"

type ErrorA struct {
	name string
}

type ErrorB struct {
	name string
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

type point struct {
	position string
	dist     int
	from     []*point
}

func (point *point) String() string {
	s := fmt.Sprintf("position:%s  dist:%d  from: %v", point.position, point.dist, point.from)
	return s
}
