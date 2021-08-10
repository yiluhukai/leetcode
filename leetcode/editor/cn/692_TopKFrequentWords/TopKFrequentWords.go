package  main

//给一非空的单词列表，返回前 k 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。 
//
// 示例 1： 
//
// 
//输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
// 
//
// 
//
// 示例 2： 
//
// 
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k
// = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
// 
//
// 
//
// 注意： 
//
// 
// 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。 
// 输入的单词均由小写字母组成。 
// 
//
// 
//
// 扩展练习： 
//
// 
// 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。 
// 
// Related Topics 字典树 哈希表 字符串 桶排序 计数 排序 堆（优先队列） 
// 👍 361 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
	解题思路：遍历数组，使用map来存储不同字符串的出现频率，然后对map中存储的元素
	先按照单词出现的频率排序，频率相同的时候按照字母序排列
	时间复杂度: 遍历 O（N) +  排序O(NlogN)  ==> O(NlogN)
	空间复杂度：map O(N) + 数组 O(N) ==> O(N)
	解答成功:
	执行耗时:8 ms,击败了77.33% 的Go用户
	内存消耗:4.4 MB,击败了38.71% 的Go用户
 */

//func topKFrequent(words []string, k int) []string {
//
//	mp := make(map[string]int)
//	for _,word :=range words{
//		mp[word]++
//	}
//	// 对map进行一个排序
//	wordArr := make([]string,0,len(words))
//	for word :=range mp {
//		wordArr = append(wordArr, word)
//	}
//	sort.Slice(wordArr, func(i, j int) bool {
//		word1,word2 := wordArr[i],wordArr[j]
//		// 频率不同
//		if mp[word1] != mp[word2]{
//			return mp[word1] > mp[word2]
//		}
//		// a < b == -1
//		return strings.Compare(word1,word2) == -1
//	})
//	return wordArr[:k]
//}

// map + 最小堆
// 先利用map来统计数组中单词出现的频率，然后使用最小堆来优化排序
// 每次当堆顶元素和当前元素做比较，如果堆顶元素出现的频次小于当前元素，删除堆顶元素并入堆
// 最后得到堆中元素依次出堆并保存到数组中，得到的就是出现频次最高的单词的反序
// 时间复杂度：O(N)遍历数组 + O(NlogK)
// 空间复杂度：O(N)
//解答成功:
//执行耗时:4 ms,击败了97.96% 的Go用户
//内存消耗:4.4 MB,击败了33.28% 的Go用户
//type Item struct {
//	word string
//	times int
//}
//type minHeap []Item
//
//// Len is the number of elements in the collection.
//func (h minHeap) Len() int{
//	return len(h)
//}
//// Less reports whether the element with
//// index i should sort before the element with index j.
//func (h minHeap)Less(i, j int) bool{
//	word1,word2 := h[i],h[j]
//	// 因为我们使用的是小顶堆，所以次数小的，字母序大的需要放到前面
//	return word1.times < word2.times || word1.times == word2.times && strings.Compare(word1.word,word2.word) == 1
//}
//// Swap swaps the elements with indexes i and j.
//func (h minHeap)Swap(i, j int){
//	h[i],h[j] = h[j],h[i]
//}
//// add x as element Len()
//func (h *minHeap)Push(x interface{}) {
//	*h = append(*h,x.(Item))
//}
//// remove and return element Len() - 1.
//func (h * minHeap)Pop() interface{}{
//	length := len(*h)
//	v:=(*h)[length-1]
//	*h = (*h)[:length-1]
//	return v
//}
//
//func topKFrequent(words []string, k int) []string {
//	mp := make(map[string]int)
//	for _,word :=range words{
//		mp[word]++
//	}
//	// 使用最小堆来优化排序
//	// 初始化最小堆
//	var h minHeap
//	heap.Init(&h)
//	for w,num :=range mp {
//		heap.Push(&h,Item{word: w ,times: num})
//		if h.Len() > k{
//			heap.Pop(&h)
//		}
//	}
//	//fmt.Printf("%#v\n",h)
//	// 取出堆中的元素，倒序排列
//	res :=  make([]string,k)
//	for i:=h.Len()-1;i>=0;i--{
//
//		res[i] = heap.Pop(&h).(Item).word
//	}
//	return res
//}

// 思路：使用桶排序 + trie树来优化排序
// 时间复杂度为 遍历单词为O(n) + 将所有的单词加入桶中O(mn) m代笔单词的总个数 取出单词的时间复杂度O(m*k)
// 总的时间复杂度为O(mn)

//	解答成功:
//						执行耗时:4 ms,击败了97.96% 的Go用户
//						内存消耗:6.6 MB,击败了5.00% 的Go用户
type Trie struct {
	word string
	children [26]*Trie
}
// add a word to trie tree
func (t *Trie)AddWord(word string){
	if t  == nil {
		return
	}
	cur := t
	for _,ch := range word{
		if cur.children[ch-'a'] == nil{
			cur.children[ch-'a'] = &Trie{}
		}
		cur = cur.children[ch-'a']
	}
	cur.word = word
}
// get all words from tire tree
func (t *Trie)GetWords()[]string{
	words := make([]string,0)
	if t == nil{
		return words
	}
	if t.word != ""{
		words = append(words,t.word)
	}
	for _,child := range t.children{
		if child !=nil{
			words = append(words, child.GetWords()...)
		}
	}
	return words
}



func topKFrequent(words []string, k int) []string {
	// 统计单词出现的次数
	mp :=  make(map[string]int,k)
	for _,word :=range words{
		mp[word]++
	}
	// 单词出现的次数最少为1次，最多为len(words)次
	// 根据单词出现的次数不同可以讲单词放入不同的桶中
	length := len(words)
	buckets := make([]*Trie,length)
	for word,times :=range mp{
		if buckets[times-1] == nil {
			buckets[times-1] = &Trie{}
		}
		buckets[times-1].AddWord(word)
	}
	// 桶中元素的排序是用字典树实现的，我们只要将字典树中包含的单词取出来即可
	res := make([]string,0,k)
	for i:=length-1;i>=0;i--{
		if bucket:=buckets[i];bucket != nil{
			// 代表这个桶有单词存在
			words := bucket.GetWords()
			if l:=len(words);l<= k {
				res = append(res, words...)
				k -= l
			}else{
				// 桶中的元素只有一部分(前k个)可以放入最终的结果中
				res=append(res, words[:k]...)
				break
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


