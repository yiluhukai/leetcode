package  main

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[2,1]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 哈希表 
// 👍 924 👎 0
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
// 使用递归完成
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.1 MB,击败了5.27% 的Go用户
// 时间复杂度O（n） n节点的个数，空间复杂度为O(n)
// 空间复杂度由存储节点内容的切片O(n) + 每层递归需要的栈空间0(h) h为树的深度
//func inorderTraversal(root *TreeNode) []int {
//	res := make([]int,0)
//	if root == nil {
//		return res
//	}
//	// 对左子树进行遍历
//	left := inorderTraversal(root.Left)
//	res = append(res, left...)
//	res = append(res, root.Val)
//	right := inorderTraversal(root.Right)
//	res = append(res, right...)
//	return res
//}

// 迭代实现中序遍历
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2 MB,击败了10.39% 的Go用户
// 时间复杂度O（n） n节点的个数，空间复杂度为O(n)
//func inorderTraversal(root *TreeNode) []int {
//	res := make([]int,0)
//	//我们使用list作为栈（先进先出）
//	parents := list.New()
//	for root != nil || parents.Len() != 0 {
//		// 找到最最左边的节点
//		for root !=nil {
//			parents.PushFront(root)
//			root = root.Left
//		}
//		// root ==nil
//		// 从栈里面取出一个节点
//		root = parents.Remove(parents.Front()).(*TreeNode)
//		res = append(res,root.Val)
//		root = root.Right
//	}
//	return res
//}

// 莫里斯算法：本质上是对迭代遍历二叉树的优化，减少存储节点的空间：
// 实现原理：中序遍历的前一个节点存在定义是左子树的最右节点。所以可以修改树的指针来
// 记录这个节点来代替保存到栈中
// 具体操作：从书的root节点开始遍历：
// 如果当前节点没有左子树，将当前节点保存，然后root指向当前节点的右子树
// 如果当前节点有左子树，找到当前子树的最右节点，将节点的指针指向当前的root节点，
// 然后root节点指向当前左子树。// 重复上面的步骤知道
func inorderTraversal(root *TreeNode) []int {
	res := make([]int,0)
	for root != nil {
		// 右左子树
		if root.Left !=nil {
			// 找到左子树的最右节点
			r :=root.Left
			for r.Right!=nil && r.Right != root{
				r =r.Right
			}
			// r就是当前子树的最右节点
			if r.Right == nil {
				r.Right = root
				root = root.Left
			}else{
				// r.Right == root ,代表root的左边节点已经遍历完成了
				// 恢复树的节点的原本指向
				r.Right = nil
				res = append(res, root.Val)
				root = root.Right
			}
		}else{
			// 无左子树
			res = append(res, root.Val)
			root = root.Right
		}
		// root ==nil
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
