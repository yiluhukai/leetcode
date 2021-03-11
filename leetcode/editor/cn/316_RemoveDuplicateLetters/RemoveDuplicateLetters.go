package  main

import (
	"bytes"
	"fmt"
)

//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct
//-characters 相同 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "bcabc"
//输出："abc"
// 
//
// 示例 2： 
//
// 
//输入：s = "cbacdcbc"
//输出："acdb" 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 由小写英文字母组成 
// 
// Related Topics 栈 贪心算法 字符串 
// 👍 457 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
	解题思路：遍历整个字符串，记录每个字符最后出现的位置,当每个字符最后出现的时候，代表了我们找到了一个字符，我们找到这个范围内的最小字符，
	最为第一个字符。
	1.查找当前范围内的最小字符
	2.删除剩余子串中的该字符串，然后继续在子串中进行查找下个最小字符

	最小字符串的查找方法：cbacdcbc
	先查找最小字符下标为0，遍历字符串，第一个字符是c不小于c且不是最后出现,所以继续查找下一个字符，
	当遍历到第二个字符是返现b小于c且不是最后一次出现，继续向后面查找，
	发现a小于当前的最小字符b，更新最小的下标为a的下标，由于a是最后一次出现，
	所以第一个字符为a,在剩余的字串中删除a后继续查找下一个最小的字符串
	解答成功:
					执行耗时:0 ms,击败了100.00% 的Go用户
					内存消耗:2.9 MB,击败了5.11% 的Go用户
	时间复杂度：每次执行递归函数都要遍历字符串，时间复杂度为O(n),最对执行26次
	空间复杂度: 主要是用来保存每个字符出现的最后下标，最多也是26个，最多递归26次，所以空间复杂度为O(1)
 */
//func removeDuplicateLetters(s string) string {
//	length := len(s)
//	// 递归的结束条件
//	if length <=1{
//		return s
//	}
//	// 遍历字符串找出字符最后一次出现的位置
//	lastAppearance := make([]int,26)
//	for i:=0;i<length;i++{
//		lastAppearance[s[i]-'a'] = i
//	}
//	// 查找当前范围内的最小字符
//	min:=0
//	for i:=0;i<length;i++{
//		if s[min] > s[i] {
//			min = i
//		}
//		// 字符是最后一次出现
//		if lastAppearance[s[i]-'a'] == i {
//			break
//		}
//	}
//	// 删除后续子串中的该字符
//	subStr := strings.ReplaceAll(s[min:],s[min:min+1],"")
//
//	return s[min:min+1] + removeDuplicateLetters(subStr)
//}
/**
	使用循环代替上面的递归
	解答成功:
				执行耗时:0 ms,击败了100.00% 的Go用户
				内存消耗:2.8 MB,击败了5.11% 的Go用户
 */
//func removeDuplicateLetters(s string) string {
//	length := len(s)
//	// 保存最后的结果
//	var res bytes.Buffer
//	// 递归的结束条件
//	for length >=1 {
//		// 遍历字符串找出字符最后一次出现的位置
//		lastAppearance := make([]int, 26)
//		for i := 0; i < length; i++ {
//			lastAppearance[s[i]-'a'] = i
//		}
//		// 查找当前范围内的最小字符
//		min := 0
//		for i := 0; i < length; i++ {
//			if s[min] > s[i] {
//				min = i
//			}
//			// 字符是最后一次出现
//			if lastAppearance[s[i]-'a'] == i {
//				break
//			}
//		}
//		// 删除后续子串中的该字符
//		err:=res.WriteByte(s[min])
//		if err!=nil{
//			panic(err)
//		}
//		s = strings.ReplaceAll(s[min:], s[min:min+1], "")
//		length = len(s)
//	}
//	return res.String()
//}
/**
	使用贪心算法配合栈来解决这个问题：
	使用栈存储最终要输入的字符串，每次进站都要保证栈中组成的字符串拥有最小字典序。
	所以当栈为空是进栈或当前字符大于栈顶元素时入栈，
	当字符已经在栈中存在时跳过，当字符小于栈顶元素且栈顶元素不是最后一个时，入栈，否则先对栈中
	大于当前字符的元素进行出栈，直到栈顶元素小于当前元素
	执行耗时:0 ms,击败了100.00% 的Go用户
    内存消耗:2.1 MB,击败了27.52% 的Go用户
 */

func removeDuplicateLetters(s string) string {
	length := len(s)
	// 保存字符最后一次出现的位置
	lastAppearance := make([]int,26)
	for i := 0; i < length; i++ {
		lastAppearance[s[i]-'a'] = i
	}
	// 保存最终结果的栈
	stack := make([]byte,26)
	// 栈顶
	top :=0
	// 使用一个数组去存储栈中是否已经出现了该元素
	seen:= make([]bool,26)
	//  遍历字符串对字符串进行入栈出栈操作
	for i:=0;i <length;i++{

	    if c:=s[i];!seen[c-'a'] {
				// 对栈中元素进行出栈操作
				for top>0 && c < stack[top-1] && lastAppearance[stack[top-1]-'a'] > i{
					// 出栈操作
					seen[stack[top-1]-'a'] = false
					top--
				}
				// 对当前的元素进行进栈操作
				// 栈为空
				// 栈不为空,且字符还未出现在栈中，栈顶元元素小于当前字符
				// 栈不为空且字符小于栈顶元素且栈顶元素是最后一次出现
				stack[top] = c
				top++
				seen[c-'a'] = true
		}
	}
	// 将栈中剩余的字符拼接成字符串
	var res bytes.Buffer
	for i:=0;i<top;i++{
		res.WriteByte(stack[i])
	}
	return res.String()
}
//leetcode submit region end(Prohibit modification and deletion)

func main(){
	res:=removeDuplicateLetters("bcabc")
	fmt.Printf("%v\n",res)
}

