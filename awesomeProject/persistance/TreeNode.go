package persistance

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层次遍历打印出来
func (t *TreeNode) String() string {
	return fmt.Sprint(levelOrder(t))
}

func levelOrder(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

// todo 验证是否为平衡二叉树
// 一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1
func (t *TreeNode) IsBalanced() bool {
	return isBalanced(t)
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	leftH := maxDepth(root.Left) + 1
	rightH := maxDepth(root.Right) + 1
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// todo 验证是否为搜索二叉树 binary search tree
// 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
// 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
// 任意节点的左、右子树也分别为二叉查找树；
// 没有键值相等的节点。
func (t *TreeNode) IsBST() bool {
	return isBST(t, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}

// todo 在搜索二叉树中进行搜索
func (t *TreeNode) SearchBST(target int) *TreeNode {
	return searchBST2(t, target)
}

// 查找搜索二叉树 (递归
func searchBST1(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > target {
		return searchBST1(root.Left, target)
	} else if root.Val < target {
		return searchBST1(root.Right, target)
	}
	return root
}

// 查找搜索二叉树 (迭代
func searchBST2(root *TreeNode, target int) *TreeNode {
	for root != nil {
		if root.Val == target {
			return root
		} else if root.Val > target {
			root = root.Left
		} else if root.Val < target {
			root = root.Right
		}
	}
	return nil
}

// todo 删除二叉搜索数中某一个节点，并保持搜索二叉树性质不变
func (t *TreeNode) DeleteNode(target int) *TreeNode {
	return deleteNode2(t, target)
}

// 前驱
func deleteNode2(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > target {
		root.Left = deleteNode2(root.Left, target)
		return root
	} else if root.Val < target {
		root.Right = deleteNode2(root.Right, target)
		return root
	}
	// 已经找到，现在root就是要被删除的节点
	// 如果只存在一颗子数，那么就直接返回子数（相当于删除了root）
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	// 左右子树都存在
	maxNode := root.Left
	for maxNode.Right != nil {
		//查找前驱 (找到左子树最大的，将其替换成root
		maxNode = maxNode.Right
	}
	// 替换
	root.Val = maxNode.Val
	// 删除被替换成root的
	root.Left = deleteMaxNode(root.Left)
	return root
}

func deleteMaxNode(root *TreeNode) *TreeNode {
	if root.Right == nil {
		pLeft := root.Left
		root.Left = nil
		return pLeft
	}
	root.Right = deleteMaxNode(root.Right)
	return root
}

// 后继
func deleteNode1(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 首先要找到要删掉的节点对应的树 这里如果复用了searchBST会找不到 父节点关联=-=
	// 这里的return root 很精髓噢
	if root.Val > target {
		root.Left = deleteNode1(root.Left, target)
		return root
	} else if root.Val < target {
		root.Right = deleteNode1(root.Right, target)
		return root
	}

	// 已经找到，现在root就是要被删除的节点
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	minNode := root.Right
	for minNode.Left != nil {
		//查找后继 (右子树的最小
		minNode = minNode.Left
	}
	// 替换后继
	root.Val = minNode.Val
	// 删除后继（也就是最小的
	root.Right = deleteMinNode(root.Right)
	return root

}

func deleteMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteMinNode(root.Left)
	return root
}

// todo 完全二叉树的节点个数 (优化
// 如果二叉树中除了叶子结点，每个结点的度都为 2，则此二叉树称为满二叉树
// 如果二叉树中除去最后一层节点为满二叉树，且最后一层的结点依次从左到右分布，则此二叉树被称为完全二叉树
func (t *TreeNode) CountNodes() int {
	return countNodes(t)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countLevels(root.Left)
	right := countLevels(root.Right)
	if left == right {
		// 左右子树高度相等，证明左子树为满子树，接下来就对右子树查找
		return countNodes(root.Right) + 1<<left
	} else {
		return countNodes(root.Left) + 1<<right
	}
}

func countLevels(root *TreeNode) int {
	level := 0
	for root != nil {
		level += 1
		root = root.Left
	}
	return level
}

// todo 剪枝
// 只有0和1两种值，剪掉0的枝条

func (t *TreeNode) PruneTree() *TreeNode {
	return pruneTree(t)
}

func pruneTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = pruneTree(node.Left)
	node.Right = pruneTree(node.Right)
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
	return node
}
