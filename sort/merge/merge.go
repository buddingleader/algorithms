package merge

// p <= q < r, 改变切片，就是改变原数组
func merge(a []int, p, q, r int) {
	a1, a2 := make([]int, q-p+1), make([]int, r-q)
	for index := 0; index < len(a1); index++ {
		a1[index] = a[index+p]
	}
	for index := 0; index < len(a2); index++ {
		a2[index] = a[q+1+index]
	}

	i, j, k := 0, 0, p
	for k <= r {
		if i == len(a1) || j == len(a2) {
			break
		}

		if a1[i] < a2[j] {
			a[k] = a1[i]
			i++
		} else {
			a[k] = a2[j]
			j++
		}
		k++
	}

	for index := i; index < len(a1); index++ {
		a[k] = a1[index]
		k++
	}
	for index := j; index < len(a2); index++ {
		a[k] = a2[index]
		k++
	}
}

// Sort for merge
func Sort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		Sort(a, p, q)
		Sort(a, q+1, r)
		merge(a, p, q, r)
	}
}
