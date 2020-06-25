package heap_sort

// in-place
// unstable
// time: 最好O(n*logn)，最坏O(n*logn)，平均O(n*logn)
// space: O(1)
func Sort(a []int) {
	if len(a) < 2 {
		return
	}
	for i := len(a)/2; i >= 0; i-- {
		heapAdjust(a, i, len(a)-1)
	}
	for i := len(a)-1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapAdjust(a, 0, i-1)
	}
}

func heapAdjust(a []int, start, end int) {
	temp := a[start]
	top := start
	for i := 2*top+1; i <= end; i = 2*top+1 {
		if i < end && a[i] < a[i+1] {
			i++
		}
		if temp > a[i] {
			break
		}
		a[top] = a[i]
		top = i
	}
	a[top] = temp
}
