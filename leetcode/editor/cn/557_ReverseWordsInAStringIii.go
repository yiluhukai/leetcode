//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。 
//
// 
//
// 示例： 
//
// 输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
// 
//
// 
//
// 提示： 
//
// 
// 在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。 
// 
// Related Topics 字符串 
// 👍 262 👎 0
package main

import "strings"

/*
	先对字符串切割成数组，对单词中进行反转
	解答成功:
					执行耗时:8 ms,击败了69.78% 的Go用户
					内存消耗:6.4 MB,击败了34.08% 的Go用户
 */
//func reverseWords(s string) string {
//	words := strings.Split(s," ")
//	for index,word := range words{
//		var chars []byte
//		for i:=len(word)-1;i>=0;i--{
//			chars = append(chars, word[i])
//		}
//		words[index] = string(chars)
//	}
//	return strings.Join(words," ")
//}
/*

 */

//leetcode submit region begin(Prohibit modification and deletion)

func reverseWords(s string) string {
	chars:=make([]rune,0)
	res := make([]string,0)
	for _,c :=range s{
		if c != ' '{
			chars = append(chars, c)
		}else{
			// 反转字符串
			res = append(res, resverseSli(chars))
			chars= chars[0:0]
		}
	}
	if len(chars) !=0{
		res = append(res, resverseSli(chars))
	}
	return strings.Join(res," ")
}

func resverseSli(sli []rune)string{
	for i,length:=0,len(sli);i<length/2;i++{
		sli[i],sli[length-i-1] = sli[length-i-1],sli[i]
	}
	return string(sli)
}
//leetcode submit region end(Prohibit modification and deletion)


