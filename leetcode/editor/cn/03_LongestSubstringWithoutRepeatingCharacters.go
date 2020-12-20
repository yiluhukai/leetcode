//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 示例 4: 
//
// 
//输入: s = ""
//输出: 0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 104 
// s 由英文字母、数字、符号和空格组成 
// 
// Related Topics 哈希表 双指针 字符串 Sliding Window 
// 👍 4728 👎 0
package main

/*
	找出当前字符串中的所有无重复字串，然后获取其中的最大长度
*/

//leetcode submit region begin(Prohibit modification and deletion)
//func lengthOfLongestSubstring(s string) int {
//	// 空字符串
//	length :=len(s)
//	if length == 0{
//		return 0
//	}
//	subStrs := make([]string,0)
//	// 字符串对应的切片
//	strSli := []byte(s)
//	// 字符串的长度大于1
//	for start:=0;start<length;start++{
//		for end:=start+1;end<=length;end++{
//			if subStr := string(strSli[start:end]);end == length || strings.Contains(subStr,string(s[end])){
//				subStrs = append(subStrs,subStr)
//				break
//			}
//		}
//	}
//	// 获取字串的最大长度
//	maxLen := 1
//	for _,str := range subStrs{
//		if length:=len(str);length>maxLen {
//			maxLen = length
//		}
//	}
//	return maxLen
//}
/*
	优化解法：不去保存子串，直接根据子串的长度决定是否最大的长度
*/
//func lengthOfLongestSubstring(s string) int {
//	// 空字符串
//	length :=len(s)
//	if length == 0{
//		return 0
//	}
//	maxLen := 1
//	// 字符串对应的切片
//	strSli := []byte(s)
//	// 字符串的长度大于1
//	for start:=0;start<length;start++{
//		for end:=start+1;end<=length;end++{
//			if subStr := string(strSli[start:end]);end == length || strings.Contains(subStr,string(s[end])){
//				if length:=len(subStr);length > maxLen {
//					maxLen = length
//				}
//				break
//			}
//		}
//	}
//	return maxLen
//}
/*
	使用hash表和双指针去判断字串是否重复且计算字串的长度
 */
// 自己实现一个简单的hash表
//type HashTable struct {
//	Array [128]byte
//}
//
//func hash(data byte)int{
//	return int(data)
//}
//
//func (ht * HashTable)GetData(data byte)(isExist bool){
//	index:= (len(ht.Array)-1)&hash(data)
//	if d := ht.Array[index];d != 0{
//		isExist =true
//	}else{
//		isExist =false
//	}
//	return
//}
//
//func (ht * HashTable)AddData(data byte){
//	index:= (len(ht.Array)-1)&hash(data)
//	ht.Array[index] =  data
//}
//
//func (ht * HashTable)DeleteData(data byte){
//	index:= (len(ht.Array)-1)&hash(data)
//	ht.Array[index] = 0
//}
//
//
//
//
//func lengthOfLongestSubstring(s string) int {
//	maxLen,left,right,length := 0,0,0,len(s)
//	ht := new(HashTable)
//	for right <  length {
//		// 检测元素在hashTable中是否存在
//		if isExist := ht.GetData(s[right]);isExist {
//			// 说明字符在字串中已经存在
//			ht.DeleteData(s[left])
//			left++
//		}else{
//			//  存入hash表中
//			ht.AddData(s[right])
//			right++
//			if  size :=  right - left;size > maxLen{
//				maxLen =size
//			}
//
//		}
//	}
//	return maxLen
//}

/*
	在上面代码的基础上进行优化，hash table中保存字符的位置信息，这样子当元素重复的时候可以快速的定位left的位置
 */


// 自己实现一个简单的hash表
type HashTable struct {
	Array [128]int
}

func hash(data byte)int{
	return int(data)
}
//  初始化hash table中的值
func (ht * HashTable)Init(val int){
	for index := range ht.Array{
		ht.Array[index] = val
	}
}

func (ht * HashTable)GetData(data byte)(val int,isExist bool){
	index:= (len(ht.Array)-1)&hash(data)
	if val = ht.Array[index];val != -1{
		isExist =true
	}else{
		isExist =false
	}
	return
}

func (ht * HashTable)AddData(data byte,position int){
	index:= (len(ht.Array)-1)&hash(data)
	ht.Array[index] =  position
}

func (ht * HashTable)DeleteData(data byte){
	index:= (len(ht.Array)-1)&hash(data)
	ht.Array[index] = 0
}




func lengthOfLongestSubstring(s string) int {
	maxLen,left,right,length := 0,0,0,len(s)
	ht := new(HashTable)
	//  将hash table 中的位置信息全部初始化为-1
	ht.Init(-1)
	for right <  length {
		// 检测元素在hashTable中是否存在
		if pos,isExist := ht.GetData(s[right]);isExist {
			// 说明字符在字串中已经存在，移动left指针到指定位置
			leftIndex := pos + 1
			if leftIndex > left{
				left =  leftIndex
			}
		}
		//  无论字符是否相等，我们都需要去保存right指向字符的位置
		ht.AddData(s[right],right)
		//  计算当前子串的长度
		if length:=right-left+1;length > maxLen{
			maxLen = length
		}
		right++
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)


