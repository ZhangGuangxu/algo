package heap_sort

// in-place
// unstable
// 时间复杂度: 
//   最优：O(n*logn)
//   最差：O(n*logn)
//   平均：O(n*logn)
// 空间复杂度: O(1)
func Sort(a []int) {
	l := len(a)
	if l < 2 {
		return
	}
	for i := l/2-1; i >= 0; i-- {
		heapAdjust(a, i, l-1)
	}
	for i := l-1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapAdjust(a, 0, i-1)
	}
}

func heapAdjust(a []int, s, e int) {
	temp := a[s]
	i := s
	j := 2*i+1
	for i < e && j <= e {
		if j+1 <= e && a[j+1] > a[j] {
			j++
		}
		if temp >= a[j] {
			break
		}
		a[i] = a[j]
		i = j
		j = 2*i+1
	}
	a[i] = temp
}
