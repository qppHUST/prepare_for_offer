package leetcode

import (
	"bufio"
	"fmt"
	"os"
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
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")
	fmt.Println(ans)
}

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	ans := numIslands(grid)
	fmt.Println(ans)
}

func TestSolveNQueens(t *testing.T) {
	ans := solveNQueens(4)
	fmt.Println(len(ans))
	for index, itemlist := range ans {
		fmt.Println("第 : ", index)
		for _, item := range itemlist {
			fmt.Println(item)
		}
	}
}

func TestClimbStairs(t *testing.T) {
	ans := climbStairs(44)
	fmt.Println(ans)
}

func TestMinPathSum(t *testing.T) {
	// grid1 := [][]int{
	// 	[]int{1, 3, 1},
	// 	[]int{1, 5, 1},
	// 	[]int{4, 2, 1},
	// }
	grid2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	ans := minPathSum(grid2)
	fmt.Println(ans)
}

func TestIntegerBreak(t *testing.T) {
	ans := integerBreak(10)
	fmt.Println(ans)
}

func TestNumSquares(t *testing.T) {
	ans := numSquares(26)
	fmt.Println(ans)
}

func TestUniquePaths(t *testing.T) {
	ans := uniquePaths(3, 7)
	fmt.Println(ans)
}

func TestCalculateMaxTotalYie(t *testing.T) {
	field := [][]int32{
		{2, 3, 1, 4},
		{1, 2, 0, 3},
		{4, 2, 1, 7},
		{3, 1, 4, 2},
	}
	ans := calculateMaxTotalYie(4, 4, field)
	fmt.Println(ans)
}

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	ans := rob(nums)
	fmt.Println(ans)
}

func TestReverse(t *testing.T) {
	ans := reverse(233)
	fmt.Println(ans)
}

func TestLongestCommonPrefix(t *testing.T) {
	// strs := []string{"flower", "flow", "flight"}
	strs1 := []string{"", ""}
	ans := longestCommonPrefix(strs1)
	fmt.Println(ans)
}

func TestSwapPairs(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	ans := swapPairs(node1)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func TestCanPartion(t *testing.T) {
	// nums1 := []int{1, 5, 11, 5}
	// nums2 := []int{1, 2, 3, 5}
	nums3 := []int{3, 3, 3, 4, 5}
	fmt.Println(canPartition(nums3))
}

func TestFilePathParser(t *testing.T) {
	// str1 := "/mnt/home/someone/../log"
	str2 := "/mnt/../mnt/engine/log/.."
	ans := filePathParse(str2)
	fmt.Println(ans)
}

func TestPrintAbc(t *testing.T) {
	printAbc()
}

func TestStrDesensitization(t *testing.T) {
	str1 := "abc"
	str2 := "中国ab"
	str3 := "这个是测试地址abc"
	ans1 := strDesensitization(str1)
	ans2 := strDesensitization(str2)
	ans3 := strDesensitization(str3)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
}

func TestLenovoTest1(t *testing.T) {
	lenovoTest1()
}

func TestLilisi1(t *testing.T) {
	lilisiTest1()
}

func TestReverseKGroup(t *testing.T) {
	// node5 := &ListNode{Val: 5, Next: nil}
	// node4 := &ListNode{Val: 4, Next: node5}
	// node3 := &ListNode{Val: 3, Next: }
	node2 := &ListNode{Val: 2, Next: nil}
	node1 := &ListNode{Val: 1, Next: node2}
	newHead := reverseKGroup(node1, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3}
	k := removeDuplicates(nums)
	for i := 0; i < k; i++ {
		fmt.Println(nums[i])
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates2(nums)
	for i := 0; i < k; i++ {
		fmt.Print(nums[i], " ")
	}
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}

func TestCanJump(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func TestTaoTianTest1(t *testing.T) {
	taoTianTest1()
}

func TestSortColor(t *testing.T) {
	list := []int{2, 0, 2, 1, 1, 0}
	sortColors(list)
	for _, i := range list {
		fmt.Println(i)
	}
}

func TestFree(t *testing.T) {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	a := &ListNode{
		Val:  1,
		Next: nil,
	}
	b := &ListNode{
		Val:  2,
		Next: nil,
	}
	helper := a
	head.Next = a
	a.Next = b
	*a = *b
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", helper)
	fmt.Println(".............")
	fmt.Println(a.Val)
	fmt.Println(b.Val)
	fmt.Println(".............")
	fmt.Println(a.Next)
	fmt.Println(b.Next)
}

func TestTypeSwitch(t *testing.T) {

}

func TestLru(t *testing.T) {
	lruMem := &lru{
		list: map[int]int{},
		cap:  3,
		num:  0,
	}
	lruMem.use(10)
	lruMem.use(10)
	lruMem.use(11)
	lruMem.use(12)
	lruMem.use(11)
	lruMem.use(14)
	fmt.Println(lruMem)
}

func TestCopy(t *testing.T) {
	arr1 := []int{1, 1, 1, 1, 1, 1, 1}
	arr2 := []int{2, 2, 2}
	copy(arr1[2:5], arr2)
	fmt.Println(arr1)
}

func TestGaoTu1(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader.ReadString('\n'))
}

func TestNumKLenSubstrRepeats(t *testing.T) {
	ans := numKLenSubstrRepeats("createfunonyoka", 4)
	fmt.Println(ans)
}

func TestRepeatedDNA(t *testing.T) {
	ans := repeatedDNA("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	fmt.Println(ans)
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 6, 5, 4, 9, 200, 103, -2, 8, 5, 5, 5, 7}
	quickSort(arr)
	fmt.Println(arr)
}

func TestFree1(t *testing.T) {
	var nums []int
	print(nums[10])
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case ErrorA:
				fmt.Println("errora")
				break
			case ErrorB:
				fmt.Println("errorb")
				break
			}

		}
	}()

	panic(ErrorA{name: "aaaa"})
}

func TestAfterFunc(t *testing.T) {
}
