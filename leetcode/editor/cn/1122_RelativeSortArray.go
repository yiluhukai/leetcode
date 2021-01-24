//给你两个数组，arr1 和 arr2， 
//
// 
// arr2 中的元素各不相同 
// arr2 中的每个元素都出现在 arr1 中 
// 
//
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末
//尾。 
//
// 
//
// 示例： 
//
// 
//输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
//输出：[2,2,2,1,4,3,3,9,6,7,19]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr1.length, arr2.length <= 1000 
// 0 <= arr1[i], arr2[i] <= 1000 
// arr2 中的元素 arr2[i] 各不相同 
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中 
// 
// Related Topics 排序 数组 
// 👍 153 👎 0

package main

/*
	该题是一个排序题,我们最容易想到的方法就是自定以排序函数，
	由于arg2规定了比较顺序，我们可以使用map对顺序进行映射，map的key是数组的元素，值为对应的下标。
	比较规则：args1的元素x,y,
		1。如果x,y都在map中，比较map[x],map[y]的大小。
		2。如果都不在则map中，比较x,y的大小
		3。其他情况的话，map中的小于没在map中的
	执行耗时:0 ms,击败了100.00% 的Go用户
	内存消耗:2.4 MB,击败了23.08% 的Go用户
*/



//leetcode submit region begin(Prohibit modification and deletion)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank:= make(map[int]int,len(arr2))
	for index,val :=range arr2{
		rank[val] = index
	}
	QuickSort(arr1,rank)
	return arr1
}

func QuickSort(sli []int,rank map[int]int){
	length := len(sli)
	if length <=1 {
		return
	}
	//选择base
	base := sli[0]
	start,end := 0,length-1
	// 开始调整base的位置
	for start < end {
		if  less(sli[end],sli[start],rank) {
			sli[start],sli[end] = sli[end],sli[start]
		}
		// 前移或者后移
		if sli[start] == base{
			end--
		}else{
			start++
		}
	}
	// 处理这个数字前后和后面的序列
	QuickSort(sli[:start],rank)
	QuickSort(sli[start+1:length],rank)
}
// x < y
func less(x,y int,mp map[int]int)bool{
	x_val,ok := mp[x]
	y_val,ok1 := mp[y]
	if ok && ok1{
		// 如果x,y都在map中，比较map[x],map[y]的大小。
		return x_val < y_val
	}else if ok || ok1{
		if ok {
			return true
		}else{
			return false
		}
	}else{
		return x < y
	}
}

//leetcode submit region end(Prohibit modification and deletion)


