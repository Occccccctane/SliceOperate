package main

import (
	"awesomeProject1/SliceOperate"
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s1:%v, len:%d, cap:%d\n", s1, len(s1), cap(s1))

	s1 = SliceOperate.SliceDel(2, s1)
	fmt.Printf("s1:%v, len:%d, cap:%d\n", s1, len(s1), cap(s1))
}
