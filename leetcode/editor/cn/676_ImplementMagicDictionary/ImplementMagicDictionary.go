package  main

//设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单
//词存在于你构建的字典中。 
//
// 实现 MagicDictionary 类： 
//
// 
// MagicDictionary() 初始化对象 
// void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字
//符串互不相同 
// bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得
//所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。 
// 
//
// 
//
// 
// 
// 
// 示例： 
//
// 
//输入
//["MagicDictionary", "buildDict", "search", "search", "search", "search"]
//[[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
//输出
//[null, null, false, true, false, false]
//
//解释
//MagicDictionary magicDictionary = new MagicDictionary();
//magicDictionary.buildDict(["hello", "leetcode"]);
//magicDictionary.search("hello"); // 返回 False
//magicDictionary.search("hhllo"); // 将第二个 'h' 替换为 'e' 可以匹配 "hello" ，所以返回 True
//magicDictionary.search("hell"); // 返回 False
//magicDictionary.search("leetcoded"); // 返回 False
// 
//
// 
//
// 提示： 
//
// 
// 1 <= dictionary.length <= 100 
// 1 <= dictionary[i].length <= 100 
// dictionary[i] 仅由小写英文字母组成 
// dictionary 中的所有字符串 互不相同 
// 1 <= searchWord.length <= 100 
// searchWord 仅由小写英文字母组成 
// buildDict 仅在 search 之前调用一次 
// 最多调用 100 次 search 
// 
// 
// 
// 
// Related Topics 字典树 哈希表 
// 👍 83 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
//解答成功:
//执行耗时:48 ms,击败了35.71% 的Go用户
//内存消耗:8.4 MB,击败了28.57% 的Go用户
// TrieTree的节点定义
const length = 26
type TreeNode struct {
	IsWord bool //代表是不是一个单词
	Children  [length]*TreeNode // 假设单词只包含小写字母
	Val rune //节点中保存的字符
}
//type MagicDictionary struct {
//	Tree *TreeNode
//}
//
//
//
//
//// 插入节点
//func (root *TreeNode)InsertNode(word string){
//	if root == nil {
//		return
//	}
//	currentNode := root
//	for _,char :=range word{
//		// 找到char在子节点中的下标
//		pos := char - 'a'
//		if currentNode.Children[pos] == nil {
//			newNode:= new(TreeNode)
//			newNode.Val = char
//			newNode.IsWord = false
//			currentNode.Children[pos] = newNode
//		}
//		currentNode = currentNode.Children[pos]
//	}
//	currentNode.IsWord = true
//}
//
//// 搜索单词
//
//func (root *TreeNode)Search(word string) bool{
//	if root == nil {
//		return false
//	}
//	currentNode := root
//	for _,char :=range word{
//		// 找到char在子节点中的下标
//		pos := char - 'a'
//		if currentNode.Children[pos] == nil {
//			return false
//		}
//		currentNode = currentNode.Children[pos]
//	}
//	return currentNode.IsWord
//}
//
//// 去替换单词中的一个字符去搜索
//func (root *TreeNode)FuzzySearch(word string) bool{
//	if root == nil {
//		return false
//	}
//	currentNode := root
//
//	for index,ch :=range word{
//		var i rune
//		for i=0;i<26;i++{
//			// 不用判断 i+'a' == ch 存在与否，放在最后去其他的查找结束后再去判断
//			if currentNode.Children[i] == nil || i + 'a' == ch {
//				continue
//			}else{
//				// 不等于该字符且单词存在，那么后序匹配到都应该是精准匹配
//				currentNode :=  currentNode.Children[i]
//				partWord := word[index+1:]
//				//当我们去替换了一个字符后，单词到达了结尾，这时需要判断是不是一个单词
//				if (partWord == ""&&currentNode.IsWord) ||currentNode.Search(partWord){
//					return true
//				}
//			}
//		}
//		// 一轮循环结束，如果替换ch的字符都失败了，那么只能有一种情况还能成立 存在== ch的节点
//		// 那么我们继续在该树的子树上查找
//		// 不等于nil那么其中包含的元素一定等于ch
//		if currentNode = currentNode.Children[ch-'a'];currentNode == nil {
//			return false
//		}
//	}
//	return false
//}
//
///** Initialize your data structure here. */
//func Constructor() MagicDictionary {
//	return MagicDictionary{
//		Tree: new(TreeNode),
//	}
//}
//
//
//func (this *MagicDictionary) BuildDict(dictionary []string)  {
//	for _,word :=range dictionary{
//		this.Tree.InsertNode(word)
//	}
//}
//
//
//func (this *MagicDictionary) Search(searchWord string) bool {
//	return this.Tree.FuzzySearch(searchWord)
//}

/**
	当我们替换单词中的一个字符然后去搜索。这个字符串的长度和未替换前还是一样的，基于此我们可以将字典中
	单词根据长度去分类，然后去字典中长度为这个单词长度的所以单词中查找
	解答成功:
					执行耗时:20 ms,击败了58.82% 的Go用户
					内存消耗:7 MB,击败了94.12% 的Go用户
 */
type MagicDictionary struct {
	maps map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		maps: make(map[int][]string,10),
	}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	for _,word :=range dictionary{
		length := len(word)
		this.maps[length] = append(this.maps[length], word)
	}
}


func (this *MagicDictionary) Search(searchWord string) bool {
	length:=len(searchWord)
	list, ok := this.maps[length]
	if !ok {
		return false
	}
	// 去队列中查找
	for _,word := range list{
		// 比较两个单词之间的差异是否只有一个字符的差别
	 	difference := compareWord(word,searchWord,length)
	 	if difference{
	 		return true
		}
	}
	return false
}

/**
	比较两个单词的是不是只有一个字符不同
 */
func compareWord(word,searchWord string,length int)bool{
	diffCount :=0
	for index:=0;index<length;index++{
		if word[index]!= searchWord[index]{
			diffCount++
			if diffCount >2 {
				return false
			}
		}
	}
	if diffCount == 1 {
		return true
	}
	return false
}
/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
//leetcode submit region end(Prohibit modification and deletion)


