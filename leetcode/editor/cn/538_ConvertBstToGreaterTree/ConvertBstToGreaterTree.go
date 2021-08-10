package  main

//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于
// node.val 的值之和。 
//
// 提醒一下，二叉搜索树满足下列约束条件： 
//
// 
// 节点的左子树仅包含键 小于 节点键的节点。 
// 节点的右子树仅包含键 大于 节点键的节点。 
// 左右子树也必须是二叉搜索树。 
// 
//
// 注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-s
//um-tree/ 相同 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
// 
//
// 示例 2： 
//
// 输入：root = [0,null,1]
//输出：[1,null,1]
// 
//
// 示例 3： 
//
// 输入：root = [1,0,2]
//输出：[3,3,2]
// 
//
// 示例 4： 
//
// 输入：root = [3,2,4,1]
//输出：[7,9,4,10]
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数介于 0 和 104 之间。 
// 每个节点的值介于 -104 和 104 之间。 
// 树中的所有值 互不相同 。 
// 给定的树为二叉搜索树。 
// 
// Related Topics 树 深度优先搜索 二叉搜索树 递归 
// 👍 524 👎 0
//Definition for a binary tree node.
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
//解答成功:
//执行耗时:12 ms,击败了93.10% 的Go用户
//内存消耗:6.8 MB,击败了78.17% 的Go用户

// 先获取整棵树的节点和
// 然后中序遍历去维护每个节点的累加和（大于等于该值的节点和）
//func convertBST(root *TreeNode) *TreeNode {
//	//获取总的键值和
//	sum := getTreeKeySum(root)
//	pre := new(TreeNode)
//	// 中序遍历
//	centerOrderTraversal(root,pre,&sum)
//
//	return root
//}
//
//func centerOrderTraversal(root,pre *TreeNode,sum *int){
//	if root == nil {
//		return
//	}
//	centerOrderTraversal(root.Left,pre,sum)
//	*sum -= pre.Val
//	pre.Val,root.Val = root.Val, *sum
//
//	//fmt.Printf("%v\n",root.Val)
//	centerOrderTraversal(root.Right,pre,sum)
//}
//
//
//func getTreeKeySum(root * TreeNode)int{
//	if root == nil{
//		return 0
//	}
//	return getTreeKeySum(root.Right) + getTreeKeySum(root.Left) + root.Val
//}

/*
	反向中序遍历：右中左，每次都用 sum + node.Val 来设置新的node.Val
	时间复杂度：O(n)
	解答成功:
	解答成功:
					执行耗时:8 ms,击败了98.88% 的Go用户
					内存消耗:6.8 MB,击败了77.85% 的Go用户
 */

//func convertBST(root *TreeNode) *TreeNode {
//	sum := 0
//	var reverseTraversal func (root *TreeNode)
//	reverseTraversal = func (root *TreeNode){
//		if root == nil {
//			return
//		}
//		reverseTraversal(root.Right)
//		sum +=  root.Val
//		root.Val = sum
//		reverseTraversal(root.Left)
//	}
//	reverseTraversal(root)
//	return root
//}

/*
  莫里斯算法:反向中序遍历
  时间复杂度O(n) 空间复杂度O(1)
  解答成功:
					执行耗时:8 ms,击败了98.88% 的Go用户
					内存消耗:6.8 MB,击败了99.11% 的Go用户
 */

func convertBST(root *TreeNode) *TreeNode {
	if root  == nil {
		return root
	}
	x:= root
	sum := 0
	for x!=nil {
		if r:=x.Right;r!=nil {
			// 找到左子树的最左节点，用这个节点左子树去记录上一个节点
			for r.Left !=nil && r.Left != x {
				r = r.Left
			}
			// 找到r右子树的最左节点
			if r.Left  == nil {
				r.Left = x
				x = x.Right
			}else{
				// r.Left == r 说明x的右边已经遍历完成了
				//fmt.Printf("%v\n",x.Val)
				sum += x.Val
				x.Val = sum
				// 将树的恢复成原来的样子
				r.Left = nil
				x = x.Left
			}
		}else{
			// x 没有右子树，遍历x,然后去x的左子树上遍历
			//fmt.Printf("%v\n",x.Val)
			sum += x.Val
			x.Val = sum
			x = x.Left
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)


