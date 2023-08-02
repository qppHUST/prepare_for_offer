package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	*TreeNode
	flag bool
}

// 树的前序遍历 NO144  根，左，右
func preorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

// 树的中序遍历 NO94
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// 树的后序遍历 NO145
func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

// 使用栈模拟递归
func preorderTraversalWithStack(root *TreeNode) (vals []int) {
	if root == nil {
		return []int{}
	}
	stack := []node{}
	head := root
	for {
		if len(stack) == 0 {
			stack = append(stack, node{TreeNode: head, flag: false})
			continue
		}
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !item.flag {
			node1 := node{TreeNode: item.TreeNode, flag: true}
			node2 := node{TreeNode: item.Left, flag: false}
			node3 := node{TreeNode: item.Right, flag: false}
			if node3.TreeNode != nil {
				stack = append(stack, node3)
			}
			if node2.TreeNode != nil {
				stack = append(stack, node2)
			}
			if node1.TreeNode != nil {
				stack = append(stack, node1)
			}
		} else {
			vals = append(vals, item.Val)
			if len(stack) == 0 {
				break
			}
		}
	}
	return vals
}
