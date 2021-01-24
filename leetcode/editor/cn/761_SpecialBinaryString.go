//特殊的二进制序列是具有以下两个性质的二进制序列： 
//
// 
// 0 的数量与 1 的数量相等。 
// 二进制序列的每一个前缀码中 1 的数量要大于等于 0 的数量。 
// 
//
// 给定一个特殊的二进制序列 S，以字符串形式表示。定义一个操作 为首先选择 S 的两个连续且非空的特殊的子串，然后将它们交换。（两个子串为连续的当且仅当第一
//个子串的最后一个字符恰好为第二个子串的第一个字符的前一个字符。) 
//
// 在任意次数的操作之后，交换后的字符串按照字典序排列的最大的结果是什么？ 
//
// 示例 1: 
//
// 
//输入: S = "11011000"
//输出: "11100100"
//解释:
//将子串 "10" （在S[1]出现） 和 "1100" （在S[3]出现）进行交换。
//这是在进行若干次操作后按字典序排列最大的结果。
// 
//
// 说明: 
//
// 
// S 的长度不超过 50。 
// S 保证为一个满足上述定义的特殊 的二进制序列。 
// 
// Related Topics 递归 字符串 
// 👍 56 👎 0

package main

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
/*
	1.找出符合要求的子串(子串有可能可以再分割成连续的子串，需要对子串进行遍历操作)
		我们可以对统计字符串中1，0的出现的次数，当出现次数相同时，就是一个特殊的二进制序列
		连续的二进制序列最多出现25个，因为S的长度最大为50，一个子串最小长度为2
		递归处理二进制序列子串的时候需要去除头为的1和0， 1110 1100 00 ->否则，只能匹配到整个子串
	2.对子串进行排序
	3.重新拼接字符串(字符串要么整个是一个特殊的二进制序列，要么可以拆分成多个连续的二进制序列)
 */
func makeLargestSpecial(S string) string {
	length:=len(S)
	if length <= 1{
		return S
	}
	// 存储二进制序列子串
	strArr := make([]string,0,25)
	// start二进制序列的开始位置，用于截取的时候定位开始位置
	start :=0
	// 使用count来统计10出现的次数，== 1 -> count++ ,== 0 => count-- ,
	// 当再次等于0的时候，便得到了一个子串
	count := 0
	// 查找二进制序列子串
	for i :=0;i<length;i++{
		if S[i] == '1'{
			count++
		}else{
			count--
		}
		// 得到了一个子串
		if count == 0{
			//  对字串进行头尾的1和0去除
			str:=S[start+1:i]
			// 对字串进行递归查找
			res:=makeLargestSpecial(str)

			strArr = append(strArr, "1"+res+"0")
			// 去找下一个
			start = i+1
		}
	}
	if len(strArr) == 1{
		return strArr[0]
	}
	//fmt.Printf("%v\n",strArr)
	// 获取了所有的子串，最子串进行排序
	bubbleSort(strArr)
	
	// 重新拼接字符串
	return strings.Join(strArr,"")
}

func bubbleSort(sli []string) {
	flag := true
	for i, length := 0, len(sli); i < length-1 && flag; i++ {
		flag = false
		for j := 0; j < length-1-i; j++ {
			if strings.Compare(sli[j],sli[j+1]) < 0   {
				// 发生了交换
				flag = true
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)


