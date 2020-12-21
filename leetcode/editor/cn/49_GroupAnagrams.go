//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。 
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串 
// 👍 609 👎 0
package main

import (
	"strconv"
)

/*
	根据特征分类的问题，我们应该首先想到哈希表
	解题思路：["eat", "tea", "tan", "ate", "nat", "bat"]，将但字符串按照字母序排列作为哈希表的key.将元素插入到哈希表中
	map[string][]string{"aet":["eat","tea","ate"],"ant":["nat","tan"],"abt":["abt"]}
	解答成功:
					执行耗时:28 ms,击败了76.19% 的Go用户
					内存消耗:7.7 MB,击败了71.37% 的Go用户
 */

//leetcode submit region begin(Prohibit modification and deletion)
// golang中没有直接对字符串排序的函数，但是提供了sort.Sort(data Interface).
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}*/

/*
type strSlice []rune

func (s strSlice)Len()int{
	return len(s)
}
func (s strSlice)Less(i,j int) bool{
	return s[i] < s[j]
}

func (s strSlice)Swap(i,j int){
	s[i],s[j] =s[j],s[i]
}


func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	strsMap:= make(map[string][]string,0)
	var res [][]string
	for _,str := range strs{
		//  str中的字符按照字母序排列生成一个key
		sli := []rune(str)
		sort.Sort(strSlice(sli))
		key := string(sli)
		//  将字符串保存到对应的[]string
		strArr:=strsMap[key]
		//  返回的新的切片
		strArr = append(strArr,str)
		//  map的key需要重新赋值
		strsMap[key]= strArr
	}
	// 遍历输入strMap

	for _,val := range strsMap{
		res =append(res,val)
	}
	//fmt.Printf("%#v\n",res)
	return res
}*/

/*

还是利用hash map,但是换一种生成key的方法，计数法:
用一个数组arr [26]int保存当前字符串的key a -> [0],z-[25],当为字母异位词是，那么这个数组肯定是相同的，
把数组中的值用#链接起来作为map的key

解答成功:
					执行耗时:52 ms,击败了13.24% 的Go用户
					内存消耗:10.1 MB,击败了5.01% 的Go用户

*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	strsMap:= make(map[string][]string,0)
	var res [][]string
	var count [26]int
	for _,str := range strs{
		//将字符串中的字母转为对应的数组下标，对应的值+1
		for _,s := range str{
			index := s - 'a'
			count[index]++
		}
		// 遍历数组生成key
		key := ""
		for index,val := range count{
			key += strconv.Itoa(val)+"#"
			// 对应的数字设置为0，为下次生成key作准备
			count[index] = 0
		}
		//  将字符串保存到对应的[]string
		strArr:=strsMap[key]
		//  返回的新的切片
		strArr = append(strArr,str)
		//  map的key需要重新赋值
		strsMap[key]= strArr
	}
	// 遍历输入strMap

	for _,val := range strsMap{
		res =append(res,val)
	}
	//fmt.Printf("%#v\n",res)
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


