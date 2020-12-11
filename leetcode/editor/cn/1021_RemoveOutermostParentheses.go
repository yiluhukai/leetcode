//有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"
//，"(())()" 和 "(()(()))" 都是有效的括号字符串。
//
// 如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
//
// 给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
//
// 对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
//
//
//
// 示例 1：
//
// 输入："(()())(())"
//输出："()()()"
//解释：
//输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
//删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。
//
// 示例 2：
//
// 输入："(()())(())(()(()))"
//输出："()()()()(())"
//解释：
//输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
//删除每个部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。
//
//
// 示例 3：
//
// 输入："()()"
//输出：""
//解释：
//输入字符串为 "()()"，原语化分解得到 "()" + "()"，
//删除每个部分中的最外层括号后得到 "" + "" = ""。
//
//
//
//
// 提示：
//
//
// S.length <= 10000
// S[i] 为 "(" 或 ")"
// S 是一个有效括号字符串
//
// Related Topics 栈
// 👍 143 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

//配合解法二使用，效率不如计数器
//type Stack struct {
//	data []rune
//	size int
//	top  int  // 代表下次插入元素的位置，  top ==  tail  || size ==0  是空栈
//	tail int
//}
//
//
//func (st *Stack) Push(elem rune) {
//	if st.size == len(st.data){
//		st.data = append(st.data, elem)
//	}else{
//		st.data[st.top] =  elem
//	}
//	st.size++
//	st.top++
//}
//
//func (st *Stack) Pop() rune {
//	elem := st.Peek()
//	st.top--
//	st.size--
//	return elem
//}
//
////  获取栈顶元素
//func (st *Stack) Peek() rune {
//	if st.top != st.tail {
//		return ' '
//	}
//	elem := st.data[st.top-1]
//	return elem
//}
//
////  判断一个栈是不是空栈
//func (st *Stack) IsEmpty() bool {
//	if st.size == 0 {
//		return true
//	} else {
//		return false
//	}
//}

func removeOuterParentheses(S string) string {
	/*
	 * 解法- :当遇到"（"时计数器加一，遇到"）"计数器-1。当计数器为0时，我们可以截取一个取出最外层"()"的有效括号字符串。最后将这些字符串拼接起来
	 * 返回
	 */
	//解答成功:
	//执行耗时:4 ms,击败了46.19% 的Go用户
	//内存消耗:6.7 MB,击败了26.98% 的Go用户
	//count := 0
	//start := 0
	//sli := []rune(S)
	//res := ""
	//
	//for index, val := range sli {
	//	if val == '(' {
	//		count++
	//	} else if val == ')' {
	//		count--
	//	}
	//	if count == 0 {
	//		//  匹配到一个有效括号字符串，进行截取
	//		res += string(sli[start+1 : index])
	//		start = index + 1
	//	}
	//}
	//return res

	/*
	 * 解法二：当遇到"（"时入栈，遇到"）"出栈。当计数器为0时，我们可以截取一个取出最外层"()"的有效括号字符串。最后将这些字符串拼接起来返回
	 *
	 */
	//解答成功:
	//执行耗时:4 ms,击败了46.19% 的Go用户
	//内存消耗:6.8 MB,击败了26.51% 的Go用户

	//start := 0
	//sli := []rune(S)
	//res := ""
	//
	//stack := new(Stack)
	//
	//for index, val := range sli {
	//	if val == '(' {
	//		stack.Push(val)
	//	} else if val == ')' {
	//		stack.Pop()
	//	}
	//	if stack.IsEmpty() {
	//		//  匹配到一个有效括号字符串，进行截取
	//		res += string(sli[start+1 : index])
	//		start = index + 1
	//	}
	//}
	//return res

	/*
	 * 解法三：在解法一的基础上优化，主要是不去截取字符串了：
	 * 		遇到"（"且计数器不为0，则这个字符属于一个有效括号字符串
	 *	    遇到"）"且计数器不为0，则这个字符属于一个有效括号字符串
	 */
	//解答成功:
	//执行耗时:0 ms,击败了100.00% 的Go用户
	//内存消耗:2.8 MB,击败了60.47% 的Go用户
	count:=0
	var res []rune
	for _,val := range S{
		if val == '('{
			if count !=0 {
				res = append(res,val)
			}
			count++
		}else{
			//  是"）"
			count--
			if count!=0 {
				res = append(res,val)
			}
		}
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	testCase := []string{
//		"(()())(())",
//		"(()())(())(()(()))",
//	}
//	for _, str := range testCase {
//		res := removeOuterParentheses(str)
//
//		fmt.Printf("res=%v\n", res)
//	}
//}
