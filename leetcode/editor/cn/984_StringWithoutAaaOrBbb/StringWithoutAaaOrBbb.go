package  main

import "fmt"

//给定两个整数 A 和 B，返回任意字符串 S，要求满足：
//
// 
// S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母； 
// 子串 'aaa' 没有出现在 S 中； 
// 子串 'bbb' 没有出现在 S 中。 
// 
//
// 
//
// 示例 1： 
//
// 输入：A = 1, B = 2
//输出："abb"
//解释："abb", "bab" 和 "bba" 都是正确答案。
// 
//
// 示例 2： 
//
// 输入：A = 4, B = 1
//输出："aabaa" 
//
// 
//
// 提示： 
//
// 
// 0 <= A <= 100 
// 0 <= B <= 100 
// 对于给定的 A 和 B，保证存在满足要求的 S。 
// 
// Related Topics 贪心算法 
// 👍 53 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/*

思路
直观感觉，我们应该先选择当前所剩最多的待写字母写入字符串中。举一个例子，如果 A = 6, B = 2，
那么我们期望写出 'aabaabaa'。进一步说，设当前所剩最多的待写字母为 x，只有前两个已经写下的字母都是 x 的时候，
下一个写入字符串中的字母才不应该选择它。
	执行耗时:0 ms,击败了100.00% 的Go用户
				内存消耗:2.6 MB,击败了5.00% 的Go用户
*/


func strWithout3a3b(a int, b int) string {
	res,current:="",""
	// 获取a 表示a,b中的较大值
	a_times,b_times:=0,0
	for total:=a+b;len(res) < total; {
		// 选择a的情况
		if a >= b&& a_times < 2 || b_times >= 2 {
			current = "a"
			res += current
			a_times++
			b_times =0
			a--
		}else{
			// 选择b的情况
			current = "b"
			res += current
			b_times++
			a_times =0
			b--
		}

	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)

func main(){
	samples:=[][2]int{{1,2},{4,1}}

	for _,sample := range samples{
		res:=strWithout3a3b(sample[0],sample[1])
		fmt.Printf("%v\n",res)
	}
}

