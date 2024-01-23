package leetcode

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
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
	var newAns [][]int
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
				var item []int
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

// 反转链表 NO206
func reverseListOfficial1(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return current
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

func removeElementsSimple(head *ListNode, val int) *ListNode {
	cursor := head
	for {
		if cursor.Val != val {
			break
		}
		cursor = cursor.Next
	}
	ans := cursor
	for cursor != nil {
		helper := cursor
		for helper.Next != nil {
			if helper.Next.Val == val {
				helper = helper.Next
			} else {
				break
			}
		}
		cursor.Next = helper
		cursor = cursor.Next
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

// N皇后 NO51
var queenAns = [][]string{}

func solveNQueens(n int) [][]string {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			originBoard := genBoard(n)
			flags := genFlag(n)
			originBoard[i][j] = "Q"
			editFlags(flags, i, j, n)
			insert(i, j, 2, n, originBoard, flags)
		}
	}
	return queenAns
}

func insert(i, j, num, n int, board [][]string, flags [][]bool) {
	if num > n {
		return
	}
	for {
		for {
			j = j + 1
			if j > n-1 {
				j = j % (n - 1)
				break
			}
			if !flags[i][j] {
				board[i][j] = "Q"
				if num == n {
					helper := []string{}
					for _, item := range board {
						line := strings.Join(item, "")
						helper = append(helper, line)
					}
					queenAns = append(queenAns, helper)
				}
				editFlags(flags, i, j, n)
				insert(i, j, num+1, n, board, flags)
			}
		}
		i++
		if i == n {
			break
		}
	}
}

func genBoard(n int) [][]string {
	board := [][]string{}
	for i := 0; i < n; i++ {
		line := []string{}
		for j := 0; j < n; j++ {
			line = append(line, ".")
		}
		board = append(board, line)
	}
	return board
}

func genFlag(n int) [][]bool {
	flags := [][]bool{}
	for i := 0; i < n; i++ {
		line := []bool{}
		for j := 0; j < n; j++ {
			line = append(line, false)
		}
		flags = append(flags, line)
	}
	return flags
}

func editFlags(flags [][]bool, x, y, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == y || i == x || (x-i)/(y-j) == 1 || (x-i)/(y-j) == -1 {
				flags[i][j] = true
			}
		}
	}
}

// 爬楼梯 NO70
var cacheStairs = make(map[int]int, 40)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	_, ok := cacheStairs[n]
	if !ok {
		cacheStairs[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return cacheStairs[n]
}

// 最小路径和 NO64
var minPathSumAns = math.MaxInt

func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])
	minPathSumHelper(grid, width-1, height-1, 0, 0, 0)
	return minPathSumAns
}

func minPathSumHelper(grid [][]int, width, height, i, j, helper int) {
	if i > height || j > width {
		return
	}
	fmt.Printf("当前位置 : %d : %d\n", i, j)
	helper = helper + grid[i][j]
	if width == j && height == i {
		if helper < minPathSumAns {
			minPathSumAns = helper
		}
	}
	minPathSumHelper(grid, width, height, i+1, j, helper)
	minPathSumHelper(grid, width, height, i, j+1, helper)
}

// 整数拆分 NO343
func integerBreak(n int) int {
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		helper := 1
		for j := 1; j < i; j++ {
			product1 := j * (i - j)
			product2 := j * dp[i-j]
			helper = max(product1, product2, helper)
		}
		dp = append(dp, helper)
	}
	return dp[n]
}

func max(list ...int) int {
	max := math.MinInt
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

// 完全平方数 NO279
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

// 解码方法 NO91
func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

// 不同路径 NO62
func uniquePaths(m int, n int) int {
	board := uniquePathsGen(m, n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			board[i][j] = board[i-1][j] + board[i][j-1]
		}
	}
	return board[m-1][n-1]
}

func uniquePathsGen(m, n int) [][]int {
	board := [][]int{}
	for i := 0; i < m; i++ {
		line := make([]int, n)
		if i == 0 {
			for j := 0; j < n; j++ {
				line[j] = 1
			}
		}
		line[0] = 1
		board = append(board, line)
	}
	return board
}

// 大疆
func calculateMaxTotalYie(x int, y int, field [][]int32) int32 {
	total := 0
	delta := 0
	for _, item := range field {
		for _, num := range item {
			total = total + int(num)
		}
	}
	fmt.Println("total is : ", total)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			helper := 0
			line := field[i]
			for _, item := range line {
				helper = helper + int(item)
			}
			for m := 0; m < y; m++ {
				if m == i {
					continue
				}
				helper = helper + int(field[m][j])
			}
			if helper > delta {
				delta = helper
			}
			fmt.Printf("%d行 ---- %d列 delta : %d \n", i+1, j+1, helper)
		}
	}
	return int32(total) + int32(delta)
}

// 打家劫舍 NO198
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		max := math.MinInt
		for j := i - 2; j >= 0; j-- {
			if dp[j] > max {
				max = dp[j]
			}
		}
		result1 := max + nums[i]
		result2 := dp[i-1]
		toInsert := 0
		if result1 > result2 {
			toInsert = result1
		} else {
			toInsert = result2
		}
		dp[i] = toInsert
	}
	return dp[len(dp)-1]
}

// N字形变换 NO6
func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := range mat {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i, ch := range s {
		mat[x][y] = byte(ch)
		if i%t < r-1 {
			x++ // 向下移动
		} else {
			x--
			y++ // 向右上移动
		}
	}
	ans := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}

// 整数反转 NO7
func reverse(x int) int {
	length := len(strconv.Itoa(x))
	if x < 0 {
		length--
	}
	flag := x > 0
	str := strconv.Itoa(x)
	if !flag {
		str = str[1:]
	}
	ans := reverseCodePoints(str)
	toReturn, _ := strconv.Atoi(ans)
	if !flag {
		toReturn = -toReturn
	}
	if toReturn > math.MinInt32 && toReturn < math.MaxInt32 {
		return toReturn
	}
	return 0
}

func reverseCodePoints(s string) string {
	r := make([]rune, len(s))
	start := len(s)
	for _, c := range s {
		// quietly skip invalid UTF-8
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs[0]) == 0 {
		return ""
	}
	var ans strings.Builder
	index := 0
	for {
		flag := true
		if index >= len(strs[0]) {
			return ans.String()
		}
		indexChar := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) {
				return ans.String()
			}
			if strs[i][index] != indexChar {
				flag = false
				break
			}
		}
		if !flag {
			return ans.String()
		}
		ans.WriteString(string(indexChar))
		index++
	}
}

// 两两交换链表的节点
func swapPairs(head *ListNode) *ListNode {
	ans := &ListNode{}
	helper := ans
	for {
		if head == nil {
			break
		}
		if head != nil && head.Next == nil {
			helper.Next = head
			break
		}
		nextHead := head.Next.Next
		head.Next.Next = nil
		helper.Next = head.Next
		helper.Next.Next = head
		helper.Next.Next.Next = nil
		head = nextHead
		helper = helper.Next.Next
	}
	if ans.Next == nil {
		return head
	}
	return ans.Next
}

// 分割等和子集
func canPartition(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 || len(nums) == 2 && nums[0] != nums[1] {
		return false
	}
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	leftSum := nums[0]
	rightSum := nums[len(nums)-1]
	for i < j {
		if leftSum == rightSum && i == j-1 {
			return true
		}
		if leftSum < rightSum {
			i++
			leftSum = leftSum + nums[i]
			continue
		}
		if leftSum >= rightSum {
			j--
			rightSum = rightSum + nums[j]
			continue
		}
	}
	return false
}

// 数美科技第一题 linux系统文件路径转化
func filePathParse(str string) string {
	items := strings.Split(str, "/")
	items = items[1:]
	stack := []string{}
	for _, item := range items {
		if item == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, item)
		}
	}
	ans := strings.Join(stack, "/")
	return "/" + ans
}

// 数美科技第二题 协程顺序打印
func printAbc() {
	advancer1 := make(chan struct{})
	advancer2 := make(chan struct{})
	advancer3 := make(chan struct{})
	finish := make(chan struct{})
	counter := int32(0)
	go func() {
		for {
			<-advancer1
			fmt.Print("协程1打印了 : A\n")
			advancer2 <- struct{}{}
		}
	}()
	go func() {
		for {
			<-advancer2
			fmt.Print("协程2打印了 :B\n")
			advancer3 <- struct{}{}
		}
	}()
	go func() {
		for {
			<-advancer3
			fmt.Print("协程3打印了 :C\n")
			counter++
			fmt.Println("打印第", counter, "轮")
			if counter == 10 {
				finish <- struct{}{}
			} else {
				advancer1 <- struct{}{}
			}
		}
	}()
	advancer1 <- struct{}{}
	<-finish
}

// 数美科技第三题 字符串脱敏
func strDesensitization(str string) string {
	if len(str) == 0 {
		return ""
	}
	items := []rune(str)
	length := len(items)
	if length < 4 {
		items[len(items)-1] = rune('*')
	} else if length == 4 {
		items[1] = rune('*')
		items[2] = rune('*')
	} else {
		helper := items[:2]
		for i := 0; i < length-4; i++ {
			helper = append(helper, rune('*'))
		}
		helper = append(helper, items[length-2:]...)
		items = helper
	}
	return string(items)
}

// 联想第一题
func lenovoTest1() {
	reader := bufio.NewReader(os.Stdin)
	lineNum, _ := reader.ReadString('\n')
	lines, _ := strconv.Atoi(lineNum)
	ans := make(map[string]int)
	helper := []string{}
	for lines > 0 {
		lines--
		line, _ := reader.ReadString('\n')
		items := strings.Split(line, " ")
		nums, _ := strconv.Atoi("item")
		for i := 1; i <= nums; i++ {
			writer := items[i]
			score, ok := ans[writer]
			if !ok {
				ans[writer] = 1
				continue
			}
			score = score + 4 - i
			ans[writer] = score
		}
	}
	for key, value := range ans {
		toInsert := key + " " + strconv.Itoa(value)
		helper = append(helper, toInsert)
	}
	sort.Slice(helper, func(i, j int) bool {
		return i < j
	})
}

// 莉莉丝第一题

var winner = ""

type lilisipoint struct {
	x int
	y int
}

func lilisiTest1() {
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
	points := []lilisipoint{derect1, derect2, derect3, derect4}
	for _, item := range points {
		if item.x > 0 && item.x < 3 && visit[item.x][item.y] == false {
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

// K个一组反转链表 NO25
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	totalHead, prev, toReverseHead, toReverseTail, next := &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
	prev = totalHead
	toReverseHead = head
	toReverseTail = head
	for {
		count := 0
		for toReverseHead != nil && toReverseTail.Next != nil && count < k-1 {
			count++
			toReverseTail = toReverseTail.Next
		}
		if count == k-1 {
			next = toReverseTail.Next
			toReverseTail.Next = nil
			newHead, newTail := reverseListOfficial(toReverseHead)
			prev.Next = newHead
			prev = newTail
			toReverseHead = next
			toReverseTail = next
		} else {
			//不够k个
			prev.Next = toReverseHead
			break
		}
	}
	return totalHead.Next
}

// 反转链表
func reverseListOfficial(head *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	curr := head
	tail := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev, tail
}

// 删除链表中的重复项
func removeDuplicates(nums []int) int {
	index := 1
	cursor := 1
	lastNum := nums[0]
	for cursor < len(nums) {
		if nums[cursor] == lastNum {
			cursor++
			continue
		} else {
			nums[index] = nums[cursor]
			lastNum = nums[cursor]
			index++
			cursor++
			if cursor >= len(nums) {
				break
			}
		}
	}
	return index
}

// 删除链表中的重复项 NO80
func removeDuplicates2(nums []int) int {
	index := 1
	cursor := 1
	lastNum := nums[0]
	count := 1
	for cursor < len(nums) {
		if nums[cursor] == lastNum {
			if count >= 2 {
				cursor++
				continue
			} else {
				count++
				nums[index] = nums[cursor]
				index++
				cursor++
			}
		} else {
			nums[index] = nums[cursor]
			lastNum = nums[index]
			index++
			cursor++
			count = 1
		}
	}
	return index
}

// 轮转数组 NO189
func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	k = k % len(nums)
	for k > 0 {
		rotateOnce(nums)
		k--
	}
}

func rotateOnce(nums []int) {
	one := nums[0]
	index := len(nums) - 1
	for index > 0 {
		newIndex := (index + 1) % len(nums)
		nums[newIndex] = nums[index]
		index--
	}
	nums[1] = one
}

// 买卖股票的最佳时机 NO121
func maxProfit(prices []int) int {
	minPrice := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		nowPrice := prices[i]
		if nowPrice > minPrice {
			helper := nowPrice - minPrice
			if helper > ans {
				ans = helper
			}
		} else {
			minPrice = nowPrice
		}
	}
	return ans
}

// 跳跃游戏 NO55
var canJumpAns = false

func canJump(nums []int) bool {
	canJumpHelper(nums, len(nums)-1)
	return canJumpAns
}

func canJumpHelper(nums []int, index int) {
	if index == 0 {
		canJumpAns = true
	}
	for i := index - 1; i >= 0; i-- {
		if nums[i] >= index-i {
			fmt.Println(index, "-------->", i)
			canJumpHelper(nums, i)
		}
	}
}

// 分发糖果 NO
func candy(ratings []int) int {
	ans := 0
	left := make([]int, len(ratings))
	for index, item := range ratings {
		if index > 0 && item > ratings[index-1] {
			left[index] = left[index-1] + 1
		} else {
			left[index] = 1
		}
	}
	right := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans = ans + max(right, left[i])
	}
	return ans
}

// Constructor 序列化和反序列化二叉搜索树 NO449
// Constructor
func Constructor() Codec {
	newCodec := Codec{
		//Tree: nil,
	}
	return newCodec
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var builder strings.Builder
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		builder.WriteString(strconv.Itoa(item.Val))
		builder.WriteString(",")
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[:len(data)-1]
	strs := strings.Split(data, ",")
	var queue []*TreeNode
	num0, _ := strconv.Atoi(strs[0])
	root := &TreeNode{Val: num0}
	queue = append(queue, root)
	for i := 1; i < len(strs); i++ {
		//item, _ := strconv.Atoi(strs[i])

	}
	return nil
}

// 淘天Test1
func taoTianTest1() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	var parseList []string
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == ')' {
			parseList = append(parseList, string(str[i]))
		} else if str[i] >= 65 && str[i] <= 90 {
			//是大写字母，开始找单词
			var builder strings.Builder
			builder.WriteString(string(str[i]))
			index := i + 1
			for {
				if str[index] >= 97 && str[index] <= 122 || str[index] >= 48 && str[index] <= 57 {
					//小写字母或者数字
					builder.WriteString(string(str[index]))
				} else if str[index] >= 65 && str[index] <= 90 {
					//大写字母，出去
					break
				}
				index++
			}
			parseList = append(parseList, builder.String())
			i = index
		} else if str[i] >= 48 && str[i] <= 57 {
			var builder strings.Builder
			index := i
			for {
				if str[index] >= 48 && str[index] <= 57 {
					builder.WriteString(string(str[index]))
				} else {
					break
				}
			}
			i = index
			parseList = append(parseList, builder.String())
		}
	}
	for _, item := range parseList {
		fmt.Println(item)
	}
}

// 淘天Test2
func taoTianTest2(str1, str2 string) {
	var list []int32
	var lock sync.Mutex
	chanelToAdvance := make(chan struct{})
	chanelToFinish := make(chan struct{})
	index := 0
	go func(str string) {
		ticker := time.NewTicker(time.Second)
		for _, c := range str {
			<-ticker.C
			lock.Lock()
			list = append(list, c)
			lock.Unlock()
			chanelToAdvance <- struct{}{}
		}
	}(str1)
	go func(str string) {
		ticker := time.NewTicker(time.Second)
		for _, c := range str {
			<-ticker.C
			lock.Lock()
			list = append(list, c)
			lock.Unlock()
			chanelToAdvance <- struct{}{}
		}
	}(str2)
	go func() {
		for {
			<-chanelToAdvance
			fmt.Print(string(list[index]))
			index++
			if index == len(str1)+len(str2) {
				chanelToFinish <- struct{}{}
			}
		}
	}()
	<-chanelToFinish
}

func sortColors(nums []int) {
	zero, two := -1, len(nums)
	index := 0
	for index < two {
		if nums[index] == 0 {
			zero++
			nums[zero], nums[index] = nums[index], nums[zero]
		} else if nums[index] == 2 {
			two--
			nums[two], nums[index] = nums[index], nums[two]
		}
		if nums[index] == 1 || zero == index {
			index++
		}
	}
}

func twoSum(numbers []int, target int) []int {
	i1, i2 := 0, len(numbers)-1
	for {
		if numbers[i1]+numbers[i2] > target {
			i2--
		}
		if numbers[i1]+numbers[i2] < target {
			i1++
		}
		if numbers[i1]+numbers[i2] == target {
			return []int{i1 + 1, i2 + 1}
		}
	}
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
}

func evalRPN2(tokens []string) int {
	stack := []string{tokens[0]}

	for i := range tokens {
		if i == 0 {
			continue
		}
		item := tokens[i]
		if item == "+" || item == "-" || item == "*" || item == "/" {
			num1Str := stack[len(stack)-1]
			num2Str := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			num1, _ := strconv.Atoi(num1Str)
			num2, _ := strconv.Atoi(num2Str)
			ans := -1
			if item == "+" {
				ans = num1 + num2
			} else if item == "-" {
				ans = num2 - num1
			} else if item == "*" {
				ans = num1 * num2
			} else if item == "/" {
				ans = num2 / num1
			}
			stack = append(stack, strconv.Itoa(ans))
		} else {
			stack = append(stack, item)
		}
	}

	toReturn, _ := strconv.Atoi(stack[0])

	return toReturn
}

func sumNumbers(root *TreeNode) int {
	var numList = []int{}
	var sumNumbersRecurv func(node *TreeNode, numStr string)
	sumNumbersRecurv = func(node *TreeNode, numStr string) {
		if node.Left == nil && node.Right == nil {
			atoi, _ := strconv.Atoi(numStr)
			numList = append(numList, atoi)
			return
		}
		if node.Left != nil {
			numStrLeft := numStr + strconv.Itoa(node.Left.Val)
			sumNumbersRecurv(node.Left, numStrLeft)
		}
		if node.Right != nil {
			numStrRight := numStr + strconv.Itoa(node.Right.Val)
			sumNumbersRecurv(node.Right, numStrRight)
		}
	}

	sumNumbersRecurv(root, strconv.Itoa(root.Val))
	ans := 0
	for _, item := range numList {
		fmt.Println(item)
		ans = ans + item
	}
	return ans
}

//func sumNumbersRecurv(node *TreeNode, numStr string) {
//	if node.Left == nil && node.Right == nil {
//		atoi, _ := strconv.Atoi(numStr)
//		numList = append(numList, atoi)
//		return
//	}
//	numStrLeft := numStr + strconv.Itoa(node.Left.Val)
//	sumNumbersRecurv(node.Left, numStrLeft)
//	numStrRight := numStr + strconv.Itoa(node.Right.Val)
//	sumNumbersRecurv(node.Right, numStrRight)
//}

type lru struct {
	list map[int]int
	cap  int
	num  int
}

func (this *lru) use(mem int) {
	if _, ok := this.list[mem]; !ok {
		this.list[mem] = 1
		this.num++
	} else {
		if this.num >= this.cap {
			//淘汰策略
			least := -1
			leastNum := math.MaxInt
			for k, v := range this.list {
				if v < leastNum {
					//这个可以
					least = k
					leastNum = v
				}
			}
			fmt.Println("淘汰了： ", least)
			delete(this.list, least)
		} else {
			this.num++
			this.list[mem] = 1
		}
	}
}

//func findSubstring(s string, words []string) []int {
//	view := make([]bool,len(words))
//	remaining := len(words)
//	ans
//	var findSubstringHelper func(s string, words []string)
//	findSubstringHelper(s,words)
//}

func meth55(list []int) {
	index, zero := 0, len(list)
	for index < len(list) {
		if list[index] == 0 {
			for i := index; i < zero; i++ {
				list[i] = list[i+1]
			}
			zero--
			list[zero] = 0
			if list[index] != 0 {
				index++
			}
		} else {
			index++
		}
	}
}

func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
}

func numKLenSubstrRepeats(s string, k int) int {
	ans := 0
	for m := 0; m < len(s)-k; m++ {
		begin := m
		end := m + k
		for i := begin; i < end; i++ {
			flag := false
			for j := i + 1; j < end; j++ {
				if s[i] == s[j] {
					ans++
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}
	return ans
}

func repeatedDNA(DNA string) []string {
	begin := -1
	var ans []string
	for {
		begin++
		end := begin + 10
		if end > len(DNA) {
			break
		}
		str := DNA[begin:end]
		cursor := begin
		for {
			cursor++
			cursorEnd := cursor + 10
			if cursorEnd > len(DNA) {
				break
			}
			cursorStr := DNA[cursor:cursorEnd]
			if str == cursorStr {
				ans = append(ans, str)
				break
			}
		}
	}
	return ans
}

func quickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	std := begin
	for i := begin + 1; i <= end; i++ {
		if arr[i] < arr[std] {
			numHelper := arr[std]
			arr[std] = arr[i]
			arr[i] = arr[std+1]
			arr[std+1] = numHelper
		}
	}
	quickSortHelper(arr, begin, std)
	quickSortHelper(arr, std+1, end)
}

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
//找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

func baiduTest(nums []int, target int) {

}

func recuricv(nums []int, view []bool, remaining int) {

}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func quickSortOfficial(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSortOfficial(arr, low, pivot-1)
		quickSortOfficial(arr, pivot+1, high)
	}
}
