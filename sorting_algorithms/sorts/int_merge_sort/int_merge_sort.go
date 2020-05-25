package int_merge_sort

func IntMergeSort(a []int) {
	IntMergeSortHelper(a, 0, len(a)-1)
}

func IntMergeSortHelper(a []int, lo int, hi int) {
	if hi <= lo {
		return
	} else {
		med := lo + (hi-lo)/2
		IntMergeSortHelper(a, lo, med)
		IntMergeSortHelper(a, med+1, hi)
		IntMerge(a, lo, med, hi)
	}
}

func IntMerge(a []int, lo int, med int, hi int) {
	s := make([]int, 0, hi-lo)
	i := lo
	j := med + 1
	for i <= med || j <= hi {
		if i > med {
			s = append(s, a[j])
			j++
		} else if j > hi {
			s = append(s, a[i])
			i++
		} else if a[i] < a[j] {
			s = append(s, a[i])
			i++
		} else {
			s = append(s, a[j])
			j++
		}
	}
	i = lo
	for _, e := range s {
		a[i] = e
		i++
	}
}
