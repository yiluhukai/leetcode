package main

import "fmt"

// 根据输入数据生成一颗二叉搜索树
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	data := []int{ 5,2,8,5,9 }
	// 生成一次二叉搜索树
	head:=createBinarySearchTree(data)
	// 中序遍历这棵树
	res := inorderTraversal(head)
	fmt.Printf("%#v\n",res)
	// 查找某个节点是否在树上
	//isExist := searchNodeFromTree(head,1)
	isExist := searchNodeFromTreeUseIterator(head,9)
	fmt.Printf("%v\n",isExist)
}


func createBinarySearchTree(data []int)*TreeNode{
	if data == nil || len(data) == 0 {
		return nil
	}
	var head *TreeNode
	for _,v := range data{
		head = insertNodeForTree(head,v)
		//fmt.Printf("%v\n",head==nil)
	}
	return head
}

func insertNodeForTree(head *TreeNode,val int) *TreeNode{
	if head == nil {
		head = new(TreeNode)
		head.Val =val
	}else{
		// 比较当前节点的值和val
		if head.Val > val {
			// 插入这个节点到左子树
			head.Left = insertNodeForTree(head.Left,val)
		}else{
			// 插入这个节点到右子树
			head.Right = insertNodeForTree(head.Right,val)
		}
	}
	return head
}
// 使用莫里斯算法中序遍历这棵树
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

// 从搜索二叉树上查找某个节点
func searchNodeFromTree(root *TreeNode,val int)bool{
	if root == nil {
		return false
	}
	// 树不为空，从跟节点开始查找
	if root.Val == val {
		return true
	}else if val < root.Val{
		// 在左子树上继续查找
		return searchNodeFromTree(root.Left,val)
	}else{
		return searchNodeFromTree(root.Right,val)
	}
}

func searchNodeFromTreeUseIterator(root *TreeNode,val int)bool {
	for root !=nil{
		if root.Val == val {
			return true
		}else if val  < root.Val {
			root = root.Left
		}else {
			root = root.Right
		}
	}
	return false
}

