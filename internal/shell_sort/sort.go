package shell_sort

// in-place
// space: O(1)
func Sort1(a []int) {
	l := len(a)
	if l < 2 {
		return
	}
	const T = 3
	n := T
	gap := l/n + 1
	for gap >= 1 {
		for k := 0; k < gap; k++ {
			for i := k + gap; i < l; i += gap {
				x := a[i]
				j := i - gap
				for j >= k && a[j] > x {
					a[j+gap] = a[j]
					j -= gap
				}
				a[j+gap] = x
			}
		}
		if gap == 1 {
			break
		}
		n *= T
		gap = l/n + 1
	}
}

// in-place
func Sort(a []int) {
	l := len(a)
	if l < 2 {
		return
	}
	const T = 3
	n := T
	gap := l/n + 1
	for gap >= 1 {
		for i := gap; i < l; i++ {
			x := a[i]
			j := i - gap
			for j >= 0 && a[j] > x {
				a[j+gap] = a[j]
				j -= gap
			}
			a[j+gap] = x
		}
		if gap == 1 {
			break
		}
		n *= T
		gap = l/n + 1
	}
}
