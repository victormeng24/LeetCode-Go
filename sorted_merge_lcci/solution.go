package sorted_merge_lcci

func merge(A []int, m int, B []int, n int)  {
	i := len(A) - 1
	for m > 0 && n > 0 {
		if A[m - 1] > B[n -1] {
			A[i] = A[m - 1]
			i--
			m--
		} else {
			A[i] = B[n - 1]
			i--
			n--
		}
	}
	for m > 0 {
		A[i] = A[m - 1]
		i--
		m--
	}
	for n > 0 {
		A[i] = B[n - 1]
		i--
		n--
	}
}