package slices

// ZipInt combines 2 slices of int into one, by alternating elements of each slice, starting by the first one
func ZipInt(a, b []int) (result []int) {
	var short, long []int
	if len(a) < len(b) {
		short = a
		long = b
	} else {
		long = a
		short = b
	}

	for i := range short {
		result = append(result, a[i])
		result = append(result, b[i])
	}

	for i := len(short); i < len(long); i++ {
		result = append(result, long[i])
	}

	return result
}
