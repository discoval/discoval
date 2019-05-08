package slice


func SliceDeleteByElement(s []int64, e int64) []int64 {
	var del_index int
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			del_index = i
			break
		}
	}

	if del_index == 0 && s[del_index] != e {
		return s
	}
	s = append(s[:del_index], s[del_index + 1:]...)
	return SliceDeleteByElement(s, e)
}