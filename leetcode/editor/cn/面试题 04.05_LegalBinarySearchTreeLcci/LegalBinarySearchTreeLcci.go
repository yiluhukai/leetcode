package main

import (
	"math"
)

//实现一个函数，检查一棵二叉树是否为二叉搜索树。示例 1: 输入:     2    / \   1   3 输出: true 示例 2: 输入:     5
//    / \   1   4      / \     3   6 输出: false 解释: 输入为: [5,1,4,null,null,3,6]。
//  根节点的值为 5 ，但是其右子节点值为 4 。 Related Topics 树 深度优先搜索
// 👍 46 👎 0

//Definition for a binary tree node.
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
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
// 这道题的定义应该是这样子的，因为[1,null,1]不是二叉搜索树
// 思路：搜索二叉树 中序遍历的结果是递增的,我们可以利用这条性质
// 如果节点的右子树只包含大于等于当前节点的数
// 不能利用将二叉树先中序遍历然后再比较值是否是递增的[1,1] 中序遍历的结果是[1,1]
// 但是它本身就不是一颗二叉搜索树
//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	preNode := int(math.MinInt64)
//	// 使用莫里斯算法中序遍历这棵树
//	for root != nil {
//		// 左子树存在
//		if root.Left !=nil {
//			// 找到左子树的最右节点
//			r := root.Left
//			for r.Right !=nil && r.Right !=root{
//				r = r.Right
//			}
//			if r.Right == nil {
//				// 使用这个记录它的下一个节点的值
//				r.Right = root
//				root = root.Left
//			}else{
//				// r.Right ==root,此时我们代表我们已经遍历完成了这棵树的左子树
//				if preNode >= root.Val{
//					return false
//				}
//				preNode = root.Val
//				root = root.Right
//			}
//		}else{
//			if preNode >= root.Val{
//				return false
//			}
//			preNode = root.Val
//			root = root.Right
//		}
//	}
//	return true
//}
// 使用递归中序遍历二叉树来判断
//解答成功:
//执行耗时:8 ms,击败了69.44% 的Go用户
//内存消耗:5.1 MB,击败了71.53% 的Go用户
func isValidBST(root *TreeNode) bool {
	// 中序遍历二叉树，比较相邻的两个节点
	preNode := int(math.MinInt64)
	return helper(root,&preNode)
}

func helper(root *TreeNode,preNode *int)bool{
	if root == nil {
		return true
	}
	res :=helper(root.Left,preNode)
	if res == false{
		return false
	}
	// fmt.Printf("%v %v\n",root.Val,preNode)
	if *preNode >= root.Val{
		return false
	}else{
		*preNode = root.Val
	}
	return helper(root.Right,preNode)
}

//leetcode submit region end(Prohibit modification and deletion)

