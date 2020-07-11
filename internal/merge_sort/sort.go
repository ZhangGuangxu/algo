package merge_sort

// not in-place
// stable
// time:
//    best: O(n*logn)，以2为底n的对数
//    avg: O(n*logn)
//    worst: O(n*logn)
// space: O(n)
func Sort(a []int) {
	l := len(a)
	if l < 2 {
		return
	}
	segLen := 1
	temp := make([]int, l)
	for segLen < l {
		mergePass(a, temp, segLen)
		segLen *= 2
		mergePass(temp, a, segLen)
		segLen *= 2
	}
}

// 将src中相邻的长度为segLen的子序列两两一组合并到dst中
func mergePass(src, dst []int, segLen int) {
	l := len(src)
	c := l / segLen / 2
	var s int
	for i := 0; i < c; i++ {
		m := s + segLen - 1
		e := m + segLen
		merge(src, dst, s, m, e)
		s = e + 1
	}
	r := l - c * 2 * segLen
	if r > segLen { // 处理剩余的两个子序列
		merge(src, dst, s, s+segLen-1, l-1)
	} else if r > 0 {
		// 复制未参与merge的部分
		for i := s; i < l; i++ {
			dst[i] = src[i]
		}
	}
}

// 将有序的src[s,m]和有序的src[m+1,e]合并到dst中
func merge(src, dst []int, s, m, e int) {
	i := s
	j := m + 1
	k := s
	for i <= m && j <= e {
		if src[i] <= src[j] {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
		k++
	}
	for i <= m {
		dst[k] = src[i]
		i++
		k++
	}
	for j <= e {
		dst[k] = src[j]
		j++
		k++
	}
}
