package shell_sort

// in-place
// unstable
// time: 最差O(n²)，
//       最优情况以及平均值都优于O(n²)，例如：O(n^1.3)，O(n^1.5)
// space: O(1)
func Sort(a []int) {
	l := len(a)
	if l < 2 {
		return
	}
	const N = 3
	gap := l/N + 1
	for {
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
		gap = gap/N + 1
	}
}
