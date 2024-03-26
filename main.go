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

	s2 := []string{
		"aaa",
		"bbb",
		"ccc",
		"ddd",
	}
	s2 = SliceOperate.SliceDel(3, s2)
	fmt.Printf("s2:%v, len:%d, cap:%d\n", s2, len(s2), cap(s2))

	s3 := []byte{'a', 2, 4, 'g', 'v', ';'}
	s3 = SliceOperate.SliceDel(4, s3)
	fmt.Printf("s3:%v, len:%d, cap:%d\n", s3, len(s3), cap(s3))
}
