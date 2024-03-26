package SliceOperate

func SliceDel[T collect](index int, s []T) []T {
	println("传入容量:", cap(s), "传入长度:", len(s))
	rs := s[:index]
	//fmt.Printf("rs:%v, len:%d, cap:%d\n", rs, len(rs), cap(rs))
	s2 := s[index+1:]
	//fmt.Printf("s2:%v, len:%d, cap:%d\n", s2, len(s2), cap(s2))

	rs = append(rs, s2...)

	newSlice := make([]T, len(rs), len(rs))
	copy(newSlice, rs)

	println("传出容量:", cap(newSlice), "传出长度:", len(newSlice))
	return newSlice
}

type collect interface {
	~int | float64 | float32 | byte | string
}
