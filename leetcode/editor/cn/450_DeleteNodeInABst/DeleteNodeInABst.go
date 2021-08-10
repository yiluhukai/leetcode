package  main
//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的
//根节点的引用。 
//
// 一般来说，删除节点可分为两个步骤： 
//
// 
// 首先找到需要删除的节点； 
// 如果找到了，删除它。 
// 
//
// 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。 
//
// 示例: 
//
// 
//root = [5,3,6,2,4,null,7]
//key = 3
//
//    5
//   / \
//  3   6
// / \   \
//2   4   7
//
//给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
//
//一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
//
//    5
//   / \
//  4   6
// /     \
//2       7
//
//另一个正确答案是 [5,2,6,null,4,null,7]。
//
//    5
//   / \
//  2   6
//   \   \
//    4   7
// 
// Related Topics 树 
// 👍 447 👎 0

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
//首先再二叉树中搜索要删除的节点
//解答成功:
//执行耗时:4 ms,击败了100.00% 的Go用户
//内存消耗:6.9 MB,击败了14.10% 的Go用户

// 删除节点的时候右三种情况：第一种是要删除的节点是叶子节点，直接从父节点上删除即可
// 第二种是右子树的情况，那么这个节点的中序遍历的后继节点一定可以替换当前节点的值，且后继节点位于当前节点的右子树上
// 第三种情况是有左子树，那么当前节点中序遍历的前驱节点一定可以替换当前节点的值而不影响树的平衡，且前驱节点也位于该节点的子树上。
// 替换后为了使树维持平衡，我们应该去子树上删除该节点

func deleteNode(root *TreeNode, key int) *TreeNode {
	x := root
	// x的父节点
	var p *TreeNode
	for x !=nil {
		if x.Val == key {
			// 找到了要删除的节点
			//如果是叶子节点，直接删除
			if x.Left == nil && x.Right == nil{
				if  p == nil{
					root = nil

				}else{
					if p.Left == x {
						p.Left = nil
					}else{
						p.Right = nil
					}
				}
				x =nil
			}else if x.Right !=nil{
				//如果该节点存在右子树，则找到当前节点中序遍历的后继节点
				s := successor(x)
				x.Val = s.Val
				x.Right=deleteNode(x.Right,s.Val)
			}else{
				// 右子树不存在且左子树存在，此时x的后继节点位于x的上方，我们并未保存，
				// x的前继节点位于他的下方
				p := predecessor(x)
				x.Val = p.Val
				x.Left = deleteNode(x.Left,p.Val)
			}
		}else if x.Val > key {
			p =x
			x = x.Left
		}else{
			p =x
			x = x.Right
		}
	}
	return root
}
// 获取后继节点
func successor(root *TreeNode) *TreeNode{
	if root == nil || root.Right == nil {
		return nil
	}
	// 找该节点中序遍历的下一个节点 右子树上最左节点

	r:= root.Right

	for r.Left!= nil{
		r = r.Left
	}
	return r
}
// 获取前驱节点:左子树上的最右节点
func predecessor(root *TreeNode)* TreeNode{
	if root == nil || root.Left == nil{
		return nil
	}
	l := root.Left

	for l.Right != nil {
		l =l.Right
	}
	return l
}
//leetcode submit region end(Prohibit modification and deletion)


