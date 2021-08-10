package main

import "fmt"

// TrieTree的节点定义
const length = 26

type TreeNode struct {
	IsWord bool //代表是不是一个单词
	Children  [length]*TreeNode // 假设单词只包含小写字母
	Val rune //节点中保存的字符
}


// 创建一个TrieTree
func NewTrieTree()*TreeNode{
	return new(TreeNode)
}


// 前序遍历
func (root *TreeNode)PreOrder(){
	if root == nil{
		return
	}
	//遍历它的子节点
	for i:=0;i<length;i++{
		if node:=root.Children[i];node != nil {
			fmt.Printf("nodeValue=%c\n",node.Val)
			node.PreOrder()
		}
	}
}

// 插入节点
func (root *TreeNode)InsertNode(word string){
	if root == nil {
		return
	}
	currentNode := root
	for _,char :=range word{
		// 找到char在子节点中的下标
		pos := char - 'a'
		if currentNode.Children[pos] == nil {
			newNode:= new(TreeNode)
			newNode.Val = char
			newNode.IsWord = false
			currentNode.Children[pos] = newNode
		}
		currentNode = currentNode.Children[pos]
	}
	currentNode.IsWord = true
}

// 搜索单词

func (root *TreeNode)Search(word string) bool{
	if root == nil {
		return false
	}
	currentNode := root
	for _,char :=range word{
		// 找到char在子节点中的下标
		pos := char - 'a'
		if currentNode.Children[pos] == nil {
			return false
		}
		currentNode = currentNode.Children[pos]
	}
	return currentNode.IsWord
}

// 允许去单词中的一个字符去搜索
func (root *TreeNode)FuzzySearch(word string) bool{
	if root == nil {
		return false
	}
	currentNode := root

	for index,ch :=range word{
		var i rune
		for i=0;i<26;i++{
			// 不用判断 i+'a' == ch 存在与否，放在最后去其他的查找结束后再去判断
			if currentNode.Children[i] == nil || i + 'a' == ch {
				continue
			}else{
				// 不等于该字符且单词存在，那么后序匹配到都应该是精准匹配
				currentNode :=  currentNode.Children[i]
				partWord := word[index+1:]
				if partWord == ""||currentNode.Search(partWord){
					return true
				}
			}
		}
		// 一轮循环结束，如果替换ch的字符都失败了，那么只能有一种情况还能成立 存在== ch的节点
		// 那么我们继续在该树的子树上查找
		// 不等于nil那么其中包含的元素一定等于ch
		if currentNode = currentNode.Children[ch-'a'];currentNode == nil {
			return false
		}
	}
	//return currentNode.IsWord
	return false
}


func main(){
	//
	root :=  NewTrieTree()
	cases:=[]string{"a","b","ab","abc","abcabacbababdbadbfaejfoiawfjaojfaojefaowjfoawjfoawj",
		"abcdefghijawefe","aefawoifjowajfowafjeoawjfaow","cba","cas","aaewfawi","babcda","bcd","awefj"}

	for _,str:= range cases {
		root.InsertNode(str)
	}
	//root.PreOrder()
	//word1:="hello"
	//exist := root.Search(word1)
	//if exist {
	//	fmt.Printf("%v  exist!\n",word1)
	//}else{
	//	fmt.Printf("%v doesn't exist!\n",word1)
	//}
	//word2:="w"
	//exist = root.Search(word2)
	//if exist {
	//	fmt.Printf("%v  exist!\n",word2)
	//}else{
	//	fmt.Printf("%v doesn't exist!\n",word2)
	//}
	word3:="ab"
	exist := root.FuzzySearch(word3)
	if exist {
		fmt.Printf("%v  exist!\n",word3)
	}else{
		fmt.Printf("%v doesn't exist!\n",word3)
	}
}