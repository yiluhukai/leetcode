package heap

import "testing"

//        1
//     5      2
//   6    5  3  8
// 9 10 7
//
func TestBuildHeap(t *testing.T) {
	heap := BuildHeap([]int{7 ,1,3,10,5,2,8,9,6})

	t.Logf("%#v\n",heap)
	heap.Insert(5)
	t.Logf("%#v\n",heap)
	heap.Insert(0)
	t.Logf("%#v\n",heap)
	heap.Delete()
	t.Logf("%#v\n",heap)
	heap.Delete()
	t.Logf("%#v\n",heap)
}