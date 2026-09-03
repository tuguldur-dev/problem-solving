/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	count(root, &maxDiameter)
	return maxDiameter
}

func count(root *TreeNode, maxD *int) int {
	if root == nil {
		return 0
	}
	left := count(root.Left, maxD)
	right :=  count(root.Right, maxD)
	*maxD = max(*maxD, left+right)
	return max(left, right) + 1
}
