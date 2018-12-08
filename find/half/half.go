package half

// Find ...
func Find(a []int, p, r, x int) (int, bool) {
	if p >= r {
		return -1, false
	}

	q := (p + r) / 2
	if a[q] == x {
		return q, true
	}

	if a[q] < x {
		return Find(a, q+1, r, x)
	}

	return Find(a, p, q-1, x)
}
