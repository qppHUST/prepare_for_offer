package leetcode

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	// height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height2 := []int{4, 2, 0, 3, 2, 5}
	ans := trap(height2)
	fmt.Println(ans)
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 15
	ans := minSubArrayLen(target, nums)
	fmt.Println(ans)
}

func TestMinSubArrayLenWithSlideWindow(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 11
	ans := minSubArrayLenWithSlideWindow(target, nums)
	fmt.Println(ans)
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	fmt.Printf("%v", ans)
}

func TestReverseList(t *testing.T) {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	ans := reverseList(node1)
	for {
		fmt.Println(ans.Val)
		if ans.Next != nil {
			ans = ans.Next
		} else {
			break
		}
	}
}

func TestReverseBetween(t *testing.T) {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	ans := reverseBetween(node1, 2, 4)
	for {
		fmt.Println(ans.Val)
		if ans.Next != nil {
			ans = ans.Next
		} else {
			break
		}
	}
}

func TestRemoveElements(t *testing.T) {
	node7 := &ListNode{Val: 6, Next: nil}
	node6 := &ListNode{Val: 5, Next: node7}
	node5 := &ListNode{Val: 4, Next: node6}
	node4 := &ListNode{Val: 3, Next: node5}
	node3 := &ListNode{Val: 6, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	ans := removeElements(node1, 6)
	for {
		fmt.Println(ans.Val)
		if ans.Next != nil {
			ans = ans.Next
		} else {
			break
		}
	}
}

func TestEvalRPN(t *testing.T) {
	// tokens1 := []string{"2", "1", "+", "3", "*"}
	// tokens2 := []string{"4", "13", "5", "/", "+"}
	tokens3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	ans := evalRPN(tokens3)
	fmt.Println(ans)
}

func TestEvalRPN1(t *testing.T) {
	// tokens1 := []string{"2", "1", "+", "3", "*"}
	tokens2 := []string{"4", "13", "5", "/", "+"}
	// tokens3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	ans := evalRPN1(tokens2)
	print(ans)
}

func TestFindLadders(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ans := findLadders(beginWord, endWord, wordList)
	for _, item := range ans {
		fmt.Printf("%v", item)
	}
}

func TestHasPathSum(t *testing.T) {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 8}
	node4 := &TreeNode{Val: 11}
	node5 := &TreeNode{Val: 13}
	node6 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 2}
	node9 := &TreeNode{Val: 1}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = nil
	node3.Left = node5
	node3.Right = node6
	node4.Left = node7
	node4.Right = node8
	node6.Right = node9
	ans := hasPathSum(node1, 22)
	fmt.Println(ans)
}

func TestSimplifyPath(t *testing.T) {
	ans := simplifyPath("/a//b////c/d//././/..")
	print(ans)
}

func TestKthSmallest(t *testing.T) {
	node1 := &TreeNode{Val: 4}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 3}
	node1.Left = node2
	node1.Right = node3
	node2.Right = node4
	ans := kthSmallest(node1, 1)
	println(ans)
}

func TestCombine(t *testing.T) {
	ans := combine(4, 2)
	fmt.Println(ans)
	fmt.Println(len(ans))
}

func TestExist(t *testing.T) {
	ans := exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCCED")
	fmt.Println(ans)
}

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	ans := numIslands(grid)
	fmt.Println(ans)
}
