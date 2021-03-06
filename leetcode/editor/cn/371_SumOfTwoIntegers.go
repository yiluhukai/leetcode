//不使用运算符 + 和 - ，计算两整数 a 、b 之和。 
//
// 示例 1: 
//
// 输入: a = 1, b = 2
//输出: 3
// 
//
// 示例 2: 
//
// 输入: a = -2, b = 3
//输出: 1 
// Related Topics 位运算 
// 👍 336 👎 0

package main

/*
	思路：两个二进制数异或的结果是 无进位的和
         0000 0011	->a
         0000 0001	->b

		 0000 0010  -> 无进位的和 a^b(忽略进位的情况下：两个数同一位数字相同的时候结果必定为0，不同的时候结果才为1 )
		 0000 0001	-> 左移一位 便是进位  a&b<<1  0000 00001 （忽略进位的情况下：两个数字都为1的情况下才有进位 ）

		 接一下来在进位和无进位的和求和,重复上面的过程（求无进制的和 和进位，然后求和，知道没有进位的时候，无进位和最终的和 ）

		 // 计算机中保存的补码，解决了不同符号的数字想加的问题 ，参照 下面-2+3的运算
         // 补码为啥可以呢 ：

		假设当前时针指向8点，而准确时间是6点，调整时间可有以下两种拨法：一种是倒拨2小时，即8-2=6；另一种是顺拨10小时，8+10=12+6=6，
        即8-2=8+10=8+12-2(mod 12)．在12为模的系统里，加10和减2效果是一样的，因此凡是减2运算，都可以用加10来代替。
        若用一般公式可表示为：a-b=a-b+mod=a+mod-b。对“模”而言，2和10互为补数。
	    实际上，以12为模的系统中，11和1，8和4，9和3，7和5，6和6都有这个特性，共同的特点是两者相加等于模。
		对于计算机，其概念和方法完全一样。n位计算机，设n=8，所能表示的最大数是11111111，若再加1成100000000(9位)，
        但因只有8位，最高位1自然丢失。又回到了 00000000，所以8位二进制系统的模为 2^8。
		在这样的系统中减法问题也可以化成加法问题，只需把减数用相应的补数表示就可以了。
		把补数用到计算机对数的处理上，就是补码 。

		//  求补码 -3
	    //   1000 0011（原码） ->  1111 1101(补码) == mod - 3  ?
		//	10000 0000  => mod
	    // - 0000 0011  => 3
        //   1 1 1 1   1101 ( == -3的补码)
		一个字节有符号数的范围  -128 ->127
		补码       -> 0
        0000 0000 -> 0 和 -0
		0111 1111 -> 127
        1000 0000 -> -128  (mod - 128 = 10000 0000 - 1000 0000 = 1000 0000) == 11000 0000（-128）-> 1 0111 1111(反码)->  110000 0000 %mod 取余 = 1000 0000
		1000 0001 -> -127
     	1111 1111 ->  -1
 */

//leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {
	 carry :=  a&b <<1
	 sum :=  a ^ b
     //  如果没有进位的时候
	 if carry == 0{
	 	return sum
	 }else{
	 	// 有进位的时候
	 	return getSum(sum,carry)
	 }
}
//leetcode submit region end(Prohibit modification and deletion)

//func main(){
//	a ,b:= 3,1
//
//	c,d:=-2, 3
//	//carry := (a & b) <<1
//
//	fmt.Printf("carry=%v\n",getSum(a,b))
//	fmt.Printf("carry=%v\n",getSum(c,d))
//}


// 0000 0010

// 0000 0010
// 0000 0000
// 0000 0100 进位

// 负数
//-2
// 1000 0010 原码



//-2
// 1111 1110 补码
// 3
// 0000 0011

//sum = a ^ b ->   1111 1101
// carry =  a&b<<1 0000 0100
//  在求和

//  sum  = sum ^ carry        ->  1111 1001
//  carry = sum & carry << 1  ->  0000 1000

// 在求和

//  sum  = sum ^ carry        ->  1111 0001
//  carry = sum & carry << 1  ->  0001 0000


// 在求和
//  sum  = sum ^ carry        ->  1110 0001
//  carry = sum & carry << 1  ->  0010 0000


// 在求和
//  sum  = sum ^ carry        ->  1100 0001
//  carry = sum & carry << 1  ->  0100 0000


// 在求和
//  sum  = sum ^ carry        ->  1000 0001
//  carry = sum & carry << 1  ->  1000 0000


// 在求和
//  sum  = sum ^ carry        ->  0 0000 0001
//  carry = sum & carry << 1  ->  1 0000 0000  溢出 -> 10000 0000 %mod = 0000 0000





