package  main

import (
	"fmt"
	"math"
)

//给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
//
// 二叉搜索树的定义如下： 
//
// 
// 任意节点的左子树中的键值都 小于 此节点的键值。 
// 任意节点的右子树中的键值都 大于 此节点的键值。 
// 任意节点的左子树和右子树都是二叉搜索树。 
// 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
//输出：20
//解释：键值为 3 的子树是和最大的二叉搜索树。
// 
//
// 示例 2： 
//
// 
//
// 
//输入：root = [4,3,null,1,2]
//输出：2
//解释：键值为 2 的单节点子树是和最大的二叉搜索树。
// 
//
// 示例 3： 
//
// 
//输入：root = [-4,-2,-5]
//输出：0
//解释：所有节点键值都为负数，和最大的二叉搜索树为空。
// 
//
// 示例 4： 
//
// 
//输入：root = [2,1,3]
//输出：6
// 
//
// 示例 5： 
//
// 
//输入：root = [5,4,8,3,null,6,3]
//输出：7
// 
//
// 
//
// 提示： 
//
// 
// 每棵树有 1 到 40000 个节点。 
// 每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。 
// 
// Related Topics 二叉搜索树 动态规划 
// 👍 49 👎 0
/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
      Val int
      Left *TreeNode
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

// 暴力解法：对整个树进行遍历(我们使用前序)，然后对遍历到的每一个节点使用：中序遍历去判断是不是一颗二叉搜索树，
// 是的话返回该节点为跟节点组成的二叉搜索树的节点和
func maxSumBST(root *TreeNode) int {
	max := math.MinInt64
	return preorderTraversal(root,max)
}

/**
	前序遍历二叉树
 */

func preorderTraversal(root *TreeNode,max int)int{
	// 前序遍历一颗二叉树
	if root == nil {
		return 0
	}
	fmt.Printf("%v\n",root.Val)
	valid,sum := isValidSearchTree(root,nil)

	fmt.Printf("%v,sum =%v\n",valid,sum)
	//if valid && *sum > max {
	//	max = *sum
	//}
	//preorderTraversal(root.Left,max)
	//if res > max {
	//	max= res
	//}
	// preorderTraversal(root.Right,max)
	//if res > max {
	//	max= res
	//}
	//return max
	return *sum
}

/*
	判断一棵树是不是二叉搜索树：利用二叉树中序遍历的结果是递增的来判断
 */
func isValidSearchTree(root *TreeNode,pre *TreeNode) (valid bool,sum *int){
	if root == nil {
		valid = true
		*sum = 0
		return
	}
	if pre ==nil {
		pre = new(TreeNode)
		pre.Val = math.MinInt64
	}
	// 不是空树
	valid,sum = isValidSearchTree(root.Left,pre)
	if  !valid || pre.Val >= root.Val {
		return
	}
	*sum += root.Val
	pre = root
	valid,sum = isValidSearchTree(root.Right,pre)
	if  !valid {
		return
	}
	*sum += root.Val
	return
}



//leetcode submit region end(Prohibit modification and deletion)
