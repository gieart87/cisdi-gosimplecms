package helper

func EqualUintSliceIgnoreOrder(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}

	count := make(map[uint]int)
	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}
