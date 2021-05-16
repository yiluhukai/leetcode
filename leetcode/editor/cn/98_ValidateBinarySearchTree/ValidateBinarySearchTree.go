package  main

import "math"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 1039 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	preNode := int(math.MinInt64)
	// 使用莫里斯算法中序遍历这棵树
	for root != nil {
		// 左子树存在
		if root.Left !=nil {
			// 找到左子树的最右节点
			r := root.Left
			for r.Right !=nil && r.Right !=root{
				r = r.Right
			}
			if r.Right == nil {
				// 使用这个记录它的下一个节点的值
				r.Right = root
				root = root.Left
			}else{
				// r.Right ==root,此时我们代表我们已经遍历完成了这棵树的左子树
				if preNode >= root.Val{
					return false
				}
				preNode = root.Val
				root = root.Right
			}
		}else{
			if preNode >= root.Val{
				return false
			}
			preNode = root.Val
			root = root.Right
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
