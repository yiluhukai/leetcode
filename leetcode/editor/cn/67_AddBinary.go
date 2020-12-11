//给你两个二进制字符串，返回它们的和（用二进制表示）。 
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 
//
// 提示： 
//
// 
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 
// 
// Related Topics 数学 字符串 
// 👍 523 👎 0

package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	length1,length2 := len(a),len(b)

	length  := length1
	res:=""

	//进位
	high := 0

	if length < length2 {
		length = length2
	}
	//  遍历字符串

	for i:=0;i<length;i++{
		num1,num2:=0,0

		if index1:=length1-1-i;index1 >= 0{
			num1 =  int(a[index1]-'0')
		}
		if index2:=length2-1-i;index2 >= 0{
			num2 =  int(b[index2]-'0')
		}
		sum := num1 + num2 + high
		high = sum / 2
		low := sum % 2
		res = fmt.Sprintf("%d%s",low,res)
	}

	//  如果最后加完有进位


	if high >0 {
		res = fmt.Sprintf("%d%s",high,res)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)


//func main(){
//	var s = "1010"
//	var s1 ="1011"
//
//
//	res:=addBinary(s,s1)
//
//
//	fmt.Printf("res=%s\n",res)
//}

// 1 1 0

//  1 1

//10 0 1