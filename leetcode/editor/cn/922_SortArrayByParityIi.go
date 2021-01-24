//给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。 
//
// 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。 
//
// 你可以返回任何满足上述条件的数组作为答案。 
//
// 
//
// 示例： 
//
// 输入：[4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= A.length <= 20000 
// A.length % 2 == 0 
// 0 <= A[i] <= 1000 
// 
//
// 
// Related Topics 排序 数组 
// 👍 188 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(A []int) []int {
	length:= len(A)
	// 奇数
	oddArr := make([]int,0,length/2)
	// 偶数
	evenArr := make([]int,0,length/2)

	for _,val :=range A{
		if val % 2 == 0 {
			evenArr= append(evenArr,val)
		}else{
			oddArr = append(oddArr, val)
		}
	}
	oddIndex,evenIndex:=0,0
	for i:=0;i<length;i++{
		if i % 2 == 0{
			A[i] = evenArr[evenIndex]
			evenIndex++
		}else{
			A[i] = oddArr[oddIndex]
			oddIndex++
		}
	}
	return A
}
//leetcode submit region end(Prohibit modification and deletion)


