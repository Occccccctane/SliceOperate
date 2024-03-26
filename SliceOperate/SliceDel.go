package SliceOperate

import "fmt"

func SliceDel(index int, s []int) []int {
	println("传入容量:", cap(s), "传入长度:", len(s))
	rs := s[:index]
	fmt.Printf("rs:%v, len:%d, cap:%d\n", rs, len(rs), cap(rs))

	s2 := s[index+1:]
	fmt.Printf("s2:%v, len:%d, cap:%d\n", s2, len(s2), cap(s2))

	rs = append(rs, s2...)
	println("传出容量:", cap(rs), "传出长度:", len(rs))
	return rs
}
