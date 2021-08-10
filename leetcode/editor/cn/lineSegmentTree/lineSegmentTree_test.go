package lineSegmentTree

import (
	"fmt"
	"testing"
)

func TestNewLineSegmentTree(t *testing.T) {
	l := NewLineSegmentTree(5, []int{0,1,3,4,6})
	fmt.Printf("%#v\n",l.container)

	sum := l.Query(0,0,4, 2,4)
	fmt.Printf("%v\n",sum)
	l.Update(0,0,4,0,2)
	fmt.Printf("%#v\n",l.container)
}
