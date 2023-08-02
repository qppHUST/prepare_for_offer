package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 接雨水 NO42
func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 || len(height) == 2 {
		return 0
	}
	total := 0
	left := 0
	for {
		if height[left] == 0 {
			left++
		} else {
			break
		}
	}
	right := len(height)
	for i := len(height) - 1; i > left; i-- {
		if height[i] == 0 {
			right = i
		} else {
			break
		}
	}
	standardRight := right
	for right > left {
		isWall := true
		right--
		for i := left + 1; i < right; i++ {
			if height[i] > height[left] || height[i] > height[right] {
				isWall = false
				break
			}
		}
		if !isWall {
			continue
		}
		waterFace := int(math.Min(float64(height[left]), float64(height[right])))
		for i := left + 1; i < right; i++ {
			total = total + waterFace - height[i]
		}
		left = right
		right = standardRight
		if left == right || left == right-1 {
			break
		}
	}
	return total
}

// 长度最小的子数组(要求总和大于target并且长度最小) NO209
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 && nums[0] >= target {
		return 1
	}
	if len(nums) == 1 && nums[0] < target {
		return 0
	}
	min := len(nums)
	for i := 0; i < len(nums); i++ {
		tempTotal := nums[i]
		if tempTotal >= target {
			return 1
		}
		for j := i + 1; j < j+min; j++ {
			if j > len(nums)-1 {
				break
			}
			tempTotal = tempTotal + nums[j]
			if tempTotal >= target {
				min = int(math.Min(float64(j-i+1), float64(min)))
				break
			}
		}
		if tempTotal < target && i == 0 {
			return 0
		}
	}
	return min
}

// 长度最小的子数组(要求总和大于target并且长度最小) NO209 滑动窗口解法
func minSubArrayLenWithSlideWindow(target int, nums []int) int {
	left := 0
	right := -1
	tempTotal := 0
	min := len(nums)
	for {
		right++
		if right == len(nums) {
			break
		}
		tempTotal = tempTotal + nums[right]
		if tempTotal >= target {
			min = int(math.Min(float64(right-left+1), float64(min)))
		} else {
			continue
		}
		for {
			tempTotal = tempTotal - nums[left]
			left++
			if tempTotal < target {
				break
			} else if tempTotal >= target {
				min = int(math.Min(float64(right-left+1), float64(min)))
			}
		}
	}
	if left == 0 && tempTotal < target {
		return 0
	}
	return min
}

// 全排列 NO46
func permute(nums []int) [][]int {
	data := [][]int{}
	for _, item := range nums {
		list := []int{item}
		data = append(data, list)
	}
	ans := recur(data, nums, 1)
	return ans
}

func recur(ans [][]int, data []int, flag int) [][]int {
	if flag == len(data) {
		return ans
	}
	newAns := [][]int{}
	for _, list := range ans {
		index := findIndex(list, data)
		if index == len(data) {
			continue
		}
		for i := index; i < len(data); i++ {
			toInsert := data[i]
			if check(list, toInsert) != -1 {
				continue
			}
			for i := 0; i < len(list)+1; i++ {
				item := []int{}
				item = append(item, list[:i]...)
				item = append(item, toInsert)
				item = append(item, list[i:]...)
				newAns = append(newAns, item)
			}
		}
	}
	return recur(newAns, data, flag+1)
}

func check(data []int, target int) int {
	for index, item := range data {
		if item == target {
			return index
		}
	}
	return -1
}

func findIndex(list []int, data []int) int {
	for i := len(data) - 1; i >= 0; i-- {
		num := data[i]
		for _, item := range list {
			if item == num {
				return i + 1
			}
		}
	}
	return 1
}

// 反转链表 NO206
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	helper := &ListNode{
		Val:  0,
		Next: nil,
	}
	ans := helper
	flag := head
	for {
		if flag.Next != nil {
			flag = flag.Next
		} else {
			break
		}
	}
	helper.Next = flag
	helper = helper.Next
	for {
		cursor := head
		for {
			if cursor.Next == flag {
				cursor.Next = nil
				helper.Next = cursor
				helper = helper.Next
				flag = cursor
				break
			}
			cursor = cursor.Next
		}
		if cursor == head {
			break
		}
	}
	return ans.Next
}

// 反转链表二 NO92
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	helper := &ListNode{}
	leftEnd := &ListNode{}
	rightBegin := &ListNode{}
	listEnd := &ListNode{}

	ans := helper
	border := right
	for {
		index := 1
		cursor := head
		for {
			if index == right {
				rightBegin = cursor.Next
			}
			if index == left-1 {
				leftEnd = cursor
			}
			if index == border {
				cursor.Next = nil
				helper.Next = cursor
				helper = helper.Next
				break
			}
			index++
			cursor = cursor.Next
		}
		border = border - 1
		if index == left {
			listEnd = cursor
			break
		}
	}
	leftEnd.Next = ans.Next
	listEnd.Next = rightBegin

	return head
}

// 移除链表元素 NO203
func removeElements(head *ListNode, val int) *ListNode {
	cursor := head
	for {
		if cursor.Val != val {
			break
		}
		cursor = cursor.Next
	}
	ans := cursor
	for {
		if cursor == nil || cursor.Next == nil {
			break
		}
		if cursor.Next.Val == val {
			helper := cursor.Next
			for {
				if helper == nil {
					cursor.Next = nil
					break
				}
				if helper.Val != val {
					cursor.Next = helper
					cursor = cursor.Next
					break
				}
				helper = helper.Next
			}
		} else {
			cursor = cursor.Next
		}
	}
	return ans
}

// 逆波兰表达式求值 NO150
func evalRPN(tokens []string) int {
	stack1 := tokens
	stack2 := [100]string{}
	cursor1 := len(stack1) - 1
	cursor2 := -1
	temp := []int{}
	item1 := ""
	item2 := ""
	ok1 := false
	ok2 := false
	for {
		if cursor1 == 0 {
			if cursor2 == -1 {
				break
			} else {
				if !ok1 && !ok2 {
					move := stack2[cursor2]
					cursor2--
					cursor1++
					stack1[cursor1] = move
				}
			}
		}
		item1 = stack1[cursor1]
		cursor1--
		if item1 == "+" || item1 == "-" || item1 == "*" || item1 == "/" {
			if ok1 {
				cursor2++
				// stack2 = append(stack2, strconv.Itoa(temp[0]))
				stack2[cursor2] = strconv.Itoa(temp[0])
				cursor2++
				// stack2 = append(stack2, item1)
				stack2[cursor2] = item1
				ok1 = false
				temp = []int{}
			} else {
				cursor2++
				// stack2 = append(stack2, item1)
				stack2[cursor2] = item1
			}
		} else {
			val, _ := strconv.Atoi(item1)
			temp = append(temp, val)
			if ok1 {
				ok2 = true
			} else if !ok1 && !ok2 {
				ok1 = true
			}
			if ok1 && ok2 {
				item2 = stack2[cursor2]
				cursor2--
				tempResult := 0
				if item2 == "+" {
					tempResult = temp[0] + temp[1]
				} else if item2 == "-" {
					tempResult = temp[1] - temp[0]
				} else if item2 == "*" {
					tempResult = temp[1] * temp[0]
				} else if item2 == "/" {
					tempResult = temp[1] / temp[0]
				}
				cursor1++
				stack1[cursor1] = strconv.Itoa(tempResult)
				ok1 = false
				ok2 = false
				temp = []int{}

			}
		}
	}
	val, _ := strconv.Atoi(stack1[0])
	return val
}

// 逆波兰表达式求值 NO150
func evalRPN1(tokens []string) int {
	stack := make([]string, len(tokens))
	cursor := -1
	for index, item := range tokens {
		//是符号的话
		if item == "+" || item == "-" || item == "*" || item == "/" {
			item1 := stack[cursor]
			cursor--
			item2 := stack[cursor]
			cursor--
			helper := -1
			num1, _ := strconv.Atoi(item1)
			num2, _ := strconv.Atoi(item2)
			if item == "+" {
				helper = num1 + num2
			}
			if item == "-" {
				helper = num2 - num1
			}
			if item == "*" {
				helper = num1 * num2
			}
			if item == "/" {
				helper = num2 / num1
			}
			cursor++
			stack[cursor] = strconv.Itoa(helper)
			if cursor == 1 && index == len(tokens)-1 {
				ans, _ := strconv.Atoi(stack[cursor])
				return ans
			}
		} else {
			//是数字的话
			cursor++
			stack[cursor] = item
		}
	}
	ans, _ := strconv.Atoi(stack[cursor])
	return ans
}

// 单词接龙 NO126
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	graph := initGraph(wordList)
	table := initTable(wordList)
	queue := []point{
		{position: beginWord, dist: 0, from: []*point{{position: beginWord, dist: 0, from: nil}}},
	}
	for len(queue) > 0 {
		dot := queue[0]
		insert := true
		points := table[dot.position]
		for _, item := range points {
			if dot.dist > item.dist {
				insert = false
				break
			}
		}
		if insert {
			points = append(points, dot)
			table[dot.position] = points
			queue = queue[1:]
			list := graph[dot.position]
			for _, item := range list {
				next := point{
					position: item,
					dist:     dot.dist + 1,
					from:     []*point{&dot},
				}
				queue = append(queue, next)
			}
		} else {
			queue = queue[1:]
		}

	}
	for k, v := range table {
		fmt.Printf("%s-----------%v\n", k, v)
	}
	return nil
}

func initGraph(list []string) map[string][]string {
	m := map[string][]string{}
	for _, item1 := range list {
		for _, item2 := range list {
			if item1 != item2 {
				if checkIsLine(item1, item2) {
					lines := m[item1]
					lines = append(lines, item2)
					m[item1] = lines
				}
			}
		}
	}
	return m
}

func initTable(list []string) map[string][]point {
	m := map[string][]point{}
	for _, item := range list {
		m[item] = []point{}
	}
	return m
}

func checkIsLine(str1 string, str2 string) bool {
	b1 := []byte(str1)
	b2 := []byte(str2)
	flags := []struct{}{}
	sort.Slice(b1, func(i, j int) bool {
		return i < j
	})
	sort.Slice(b2, func(i, j int) bool {
		return i < j
	})
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			flags = append(flags, struct{}{})
		}
	}
	if len(flags) != 1 {
		return false
	}
	return true
}

// 二叉树的最大深度 NO104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 反转二叉树 NO226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 路经总和 NO112
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 二叉树的最小深度 NO111
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 1)
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return math.MaxInt64
	}
	if root.Left == nil && root.Right == nil {
		return depth
	}
	return min(helper(root.Left, depth+1), helper(root.Right, depth+1))
}

func min(list ...int) int {
	min := math.MaxInt64
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}

// 左叶子之和 NO404
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + root.Left.Val
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
}

// 二叉树的所有路径 NO257
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	constructPaths(root, "")
	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths(root.Left, pathSB)
			constructPaths(root.Right, pathSB)
		}
	}
}

// 简化路径
func simplifyPath(path string) string {
	if path == "/" {
		return "/"
	}
	data := ""
	if string(path[len(path)-1]) == "/" {
		data = path[1 : len(path)-1]
	}
	data = path[1:]
	list := strings.Split(data, "/")
	if len(list) == 0 {
		return ""
	}
	if len(list) == 1 && list[0] == ".." {
		return "/"
	}
	stack := make([]string, len(list))
	cursor := -1
	for _, item := range list {
		if item == ".." {
			if cursor >= 0 {
				cursor--
			}
		} else if item == "." {
			continue
		} else if len(item) == 0 {
			continue
		} else {
			cursor++
			stack[cursor] = item
		}
	}
	ans := strings.Join(stack[0:cursor+1], "/")
	return "/" + ans
}

// 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 路径总和2 NO113
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return
}

// 二叉搜索树的最近公共祖先 NO235
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val >= root.Val && q.Val <= root.Val || p.Val <= root.Val && q.Val >= root.Val {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return nil
}

// 二叉搜索树中第k小的元素 NO230
var ans = []int{}

func kthSmallest(root *TreeNode, k int) int {
	kthSmallestHelper(root)
	fmt.Println(ans)
	return ans[k-1]
}

func kthSmallestHelper(root *TreeNode) {
	if root == nil {
		return
	}
	kthSmallestHelper(root.Left)
	ans = append(ans, root.Val)
	kthSmallestHelper(root.Right)
}

// 组合 NO77
func combine(n int, k int) [][]int {
	previous := [][]int{}
	data := []int{}
	for i := 1; i <= n; i++ {
		previous = append(previous, []int{i})
		data = append(data, i)
	}
	fmt.Println(previous)
	fmt.Println(data)
	return combineHelper(previous, data, 1, k)
}

func combineHelper(previous [][]int, data []int, flag int, k int) [][]int {
	if flag == k {
		return previous
	}
	newPrevious := [][]int{}
	for _, list := range previous {
		for _, item := range data {
			if isInTheList(list, item) {
				continue
			}
			for i := 0; i <= len(list); i++ {
				toInsertList := []int{}
				toInsertList = append(toInsertList, list[:i]...)
				toInsertList = append(toInsertList, item)
				toInsertList = append(toInsertList, list[i:]...)
				newPrevious = append(newPrevious, toInsertList)
			}
		}
	}
	return combineHelper(newPrevious, data, flag+1, k)
}

func isInTheList(list []int, target int) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

// 单词搜索 NO79
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		// 标记该点已经被查看过
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 岛屿数量 NO200
func numIslands(grid [][]byte) int {
	h, w := len(grid), len(grid[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	queue := []pair{}
	ilandsNum := 0
	check := func(i, j int) {
		if vis[i][j] || grid[i][j] == '0' {
			return
		}
		queue = append(queue, pair{x: i, y: j})
		for len(queue) > 0 {
			item := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			vis[item.x][item.y] = true
			item1 := pair{x: item.x - 1, y: item.y} //左
			item2 := pair{x: item.x, y: item.y - 1} //上
			item3 := pair{x: item.x + 1, y: item.y} //右
			item4 := pair{x: item.x, y: item.y + 1} //下
			if item1.x >= 0 && item1.x < h && item1.y >= 0 && item1.y < w && grid[item1.x][item1.y] == '1' && !vis[item1.x][item1.y] {
				queue = append(queue, item1)
			}
			if item2.x >= 0 && item2.x < h && item2.y >= 0 && item2.y < w && grid[item2.x][item2.y] == '1' && !vis[item2.x][item2.y] {
				queue = append(queue, item2)
			}
			if item3.x >= 0 && item3.x < h && item3.y >= 0 && item3.y < w && grid[item3.x][item3.y] == '1' && !vis[item3.x][item3.y] {
				queue = append(queue, item3)
			}
			if item4.x >= 0 && item4.x < h && item4.y >= 0 && item4.y < w && grid[item4.x][item4.y] == '1' && !vis[item4.x][item4.y] {
				queue = append(queue, item4)
			}
		}
		ilandsNum++
	}
	for x, xitem := range grid {
		for y := range xitem {
			check(x, y)
		}
	}
	return ilandsNum
}
