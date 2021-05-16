package  main

import (
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
// 时间复杂度O（你(n^2)）
// 空间复杂度 O(h^2)
// 执行的结果是是超时的
//func maxSumBST(root *TreeNode) int {
//	max := math.MinInt64
//	preorderTraversal(root,&max)
//	return max
//}
//
///**
//	前序遍历二叉树
// */
//
//func preorderTraversal(root *TreeNode,max *int){
//	// 前序遍历一颗二叉树
//	if root == nil {
//		if *max < 0 {
//			*max = 0
//		}
//		return
//	}
//	valid,sum := isValidSearchTree(root,nil)
//	if valid && sum > *max {
//		*max = sum
//	}
//	preorderTraversal(root.Left,max)
//	preorderTraversal(root.Right,max)
//}
//
///*
//	判断一棵树是不是二叉搜索树：利用二叉树中序遍历的结果是递增的来判断
// */
//func isValidSearchTree(root *TreeNode,pre *TreeNode) (valid bool,sum int){
//	if root == nil {
//		valid = true
//		return
//	}
//	if pre ==nil {
//		pre = new(TreeNode)
//		pre.Val = math.MinInt64
//	}
//
//	// 不是空树
//	valid,count := isValidSearchTree(root.Left,pre)
//	if  !valid || pre.Val >= root.Val {
//		valid = false
//		return
//	}
//	sum += count
//	sum += root.Val
//	pre.Val = root.Val
//	valid,count = isValidSearchTree(root.Right,pre)
//	if valid {
//		sum += count
//	}
//	return
//}
/**
	优化解法：判断一棵树是不是二叉搜索树(BST)：左子树是二叉搜索树，右子树也是BST.
	左子树的最大值小于根结点的，右子树的最小值大于等于(本题中只能是大于)根结点的值
	左右中-> 后序遍历
	1:46 上午	info
	解答成功:
	执行耗时:128 ms,击败了89.29% 的Go用户
	内存消耗:17.1 MB,击败了75.00% 的Go用户
	时间复杂度：O(n)
	空间复杂度：O(h) h代表二叉树的高度  h = 」log N」 + 1
 */
type Result struct {
	minValue int // 当前子树的最小节点值
	maxValue int // 当前紫薯的最大节点值
	isBST bool // 是否是BST
	sum  int // 节点的键值和
}
func maxSumBST(root *TreeNode) int {
	max := math.MinInt64
	postOrderTraversal(root,&max)
	return max
}

func postOrderTraversal(root *TreeNode,max *int) Result {
	if root == nil {
		if *max < 0{
			*max = 0
		}
		// 空树是BST,当为左子树是，它的minValue应该小于root.Val
		// 当为右子树时，他的maxValue应该小于root.value
		return Result{
			isBST: true,
			minValue: math.MaxInt64,
			maxValue: math.MinInt64,
		}
	}
	// 后序遍历
	leftResult := postOrderTraversal(root.Left,max)
	rightResult := postOrderTraversal(root.Right,max)

	if !leftResult.isBST || (leftResult.isBST && leftResult.maxValue >=  root.Val){
		return Result{
			isBST: false,
		}
	}
	if !rightResult.isBST || (rightResult.isBST && rightResult.minValue <= root.Val) {
		return Result{
			isBST: false,
		}
	}
	// 左右子树都是BST,且root为根的树也是BST
	// 更新最大键值和
	sum,minValue,maxValue :=  0,leftResult.minValue,rightResult.maxValue
	sum =leftResult.sum + rightResult.sum + root.Val
	//fmt.Printf("sum =%v,root =%v\n",sum,root.Val)
	if sum > *max{
		*max =sum
	}
	if root.Left == nil {
		minValue = root.Val
	}
	if root.Right == nil {
		maxValue =  root.Val
	}
	return Result{
		isBST: true,
		sum: sum,
		minValue: minValue,
		maxValue: maxValue,
	}
}
//leetcode submit region end(Prohibit modification and deletion)
