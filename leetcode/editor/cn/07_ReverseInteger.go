//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学
// 👍 2360 👎 0
package main

import (
	"math"
)

//  submit region begin(Prohibit modification and deletion)
func reverse(x int) int {

	/*
	 *解法1： 首位交换法
	 *
	 * 总结: 数据结构选择： 将正数 转为切片,利用切片去反序 在转成int
	 * 主流程： 全部当成正数去处理
	 * 边界处理：|math.MinInt32| 比 |math.math.MaxInt32| +1 ,转成正数会直接越界。
	 * 符合处理： 开始去掉符合，最后加上符合
	 */

	/*
		//  get flag  x>=0?1:-1
		var flag int
		if x<0{
			flag =-1
		}else{
			flag = 1
		}
		//  |math.MinInt32| 比 |math.math.MaxInt32| +1 ,转成正数会直接越界。

		if x == math.MinInt32 {
			return 0
		}
		//  --=>+,++=>+   int -> str
		str:= strconv.Itoa(x * flag)
		//str:=  string(x)
		//  str -> slice
		sli := []byte(str)
		// reverse slice
		for i,length:=0,len(sli);i<length/2;i++{
			// exchange location
			sli[i],sli[length-i-1] = sli[length-i-1],sli[i]
		}
		str = string(sli)

		num ,err:= strconv.ParseInt (str,10,64)

		if err!=nil{
			panic(err)
		}

		if num > math.MaxInt32 {
			return 0
		}else{
			return int(num)*flag
		}
	*/

	/*
	 *  最优解法： 数学运算法
	 *  思路：1. 首先还是对x去掉符号去处理
	 *  	 2. 然后 对x模10取余，得到的就是个位的数字,然后对用余数继续模10取余
	 *		 999%10 = 9 ，999/10 = 99    res = 0, res = res*10+9
	 *		 99%10=9 ，99 /10 = 9 		res = res*10+9
	 *		 9%10=9   9/10=0			res = res*10+9
	 *       3. 当 x%10 == x ||  x%10 == 0
	 *		 4.边界问题   到达最后一位， 补齐最后一位后需要判断是否移除，之前不会发生溢出，中间过程用 int64去计算防止溢出
	 *		 5。和 符号问题
	 *
	 */
	//  get flag  x>=0?1:-1
	var flag int
	if x<0{
		flag =-1
	}else{
		flag = 1
	}
	x = x * flag

	var res int64  = 0
	var last   =  x%10
	//  获取最后一位，然后累加
	for last != x  {
		res = res * 10 + int64(last)
		x  = x/10
		last = x%10
	}
	// last是最后一位
	if last!=0{
		res = res * 10 + int64(last)
		if res > math.MaxInt32 {
			return  0
		}
	}
	//  加上符号
	return int(res) * flag
}

//leetcode submit regionend(Prohibit modification and deletion)

//func main(){
//	var sli = []int{123,-123,120,math.MinInt32,math.MaxInt32,0}
//
//	for _,val := range sli{
//		res:= reverse(val)
//		fmt.Printf("reverse %d get %d\n", val,res)
//	}
//}
