package  main

//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼
//写检查。 
//
// 请你实现 Trie 类： 
//
// 
// Trie() 初始化前缀树对象。 
// void insert(String word) 向前缀树中插入字符串 word 。 
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false
// 。 
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否
//则，返回 false 。 
// 
//
// 
//
// 示例： 
//
// 
//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True
// 
//
// 
//
// 提示： 
//
// 
// 1 <= word.length, prefix.length <= 2000 
// word 和 prefix 仅由小写英文字母组成 
// insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次 
// 
// Related Topics 设计 字典树 
// 👍 786 👎 0
/*
	解答成功:
	执行耗时:64 ms,击败了75.19% 的Go用户
	内存消耗:18.6 MB,击败了12.03% 的Go用户
*/
 */
//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	Children [26]*Trie
	IsWord bool
	Val rune
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	var pos rune
	currentNode := this
	for _,ch :=range word{
		pos = ch - 'a'
		if currentNode.Children[pos] == nil {
			currentNode.Children[pos] = new(Trie)
		}
		currentNode.Children[pos].Val = ch
		currentNode = currentNode.Children[pos]
	}
	currentNode.IsWord =  true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	currentNode :=  this
	for _,ch :=range word{
		pos := ch -'a'
		if next:=currentNode.Children[pos];next == nil {
			return false
		}else{
			currentNode  = next
		}
	}
	return currentNode.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	currentNode :=  this
	for _,ch :=range prefix{
		pos := ch -'a'
		if next:=currentNode.Children[pos];next == nil {
			return false
		}else{
			currentNode  = next
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)


//func main(){
//	obj := Constructor()
//	obj.Insert("apple")
//	param_2 := obj.Search("apple")
//	param_3 := obj.StartsWith("app")
//	fmt.Printf("%v %v\n",param_2,param_3)
//}
