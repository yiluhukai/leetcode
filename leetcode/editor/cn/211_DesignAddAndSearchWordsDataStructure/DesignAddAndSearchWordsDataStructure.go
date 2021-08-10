package  main
//请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。 
//
// 实现词典类 WordDictionary ： 
//
// 
// WordDictionary() 初始化词典对象 
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配 
// bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些 '
//.' ，每个 . 都可以表示任何一个字母。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search","se
//arch"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//
//解释：
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // return False
//wordDictionary.search("bad"); // return True
//wordDictionary.search(".ad"); // return True
//wordDictionary.search("b.."); // return True
// 
//
// 
//
// 提示： 
//
// 
// 1 <= word.length <= 500 
// addWord 中的 word 由小写英文字母组成 
// search 中的 word 由 '.' 或小写英文字母组成 
// 最多调用 50000 次 addWord 和 search 
// 
// Related Topics 深度优先搜索 设计 字典树 回溯算法 
// 👍 247 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

/**
解答成功:
				执行耗时:72 ms,击败了86.36% 的Go用户
				内存消耗:14.6 MB,击败了53.89% 的Go用户
 */
type WordDictionary struct {
	Tree *Trie
}
type Trie struct {
	Children [26]*Trie
	IsWord bool
	Val rune
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		Tree: new(Trie),
	}
}


func (this *WordDictionary) AddWord(word string)  {
	currentNode:= this.Tree
	for _,ch :=range word{
		pos := ch - 'a'
		if currentNode.Children[pos] == nil {
			currentNode.Children[pos] = new(Trie)
		}
		currentNode.Children[pos].Val =  ch
		currentNode = currentNode.Children[pos]
	}
	currentNode.IsWord = true
}

/*

 */
func (this *WordDictionary) Search(word string) bool {
	return this.Tree.Search(word)
}


func (this *Trie) Search(word string)bool{
	currentNode:= this
	for index,ch :=range word{

		if ch == '.'{
			// ch可以任何小写字母
			for i:=0;i<26;i++{
				if target:= currentNode.Children[i];target !=nil {
					restChar := word[index+1:]
					// 在target.Search方法中可以搜索restChar == ""的情况
					if target.Search(restChar){
						return true
					}
				}
			}
			return false
		}else{
			pos := ch - 'a'
			if currentNode.Children[pos] == nil {
				return false
			}
			currentNode = currentNode.Children[pos]
		}
	}
	return currentNode.IsWord
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)


