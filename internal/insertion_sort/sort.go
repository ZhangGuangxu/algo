package insertion_sort

// in-place
// stable
// time: 最好O(n)，最坏O(n²)，平均O(n²)
// space: O(1)
func Sort(a []int) {
    l := len(a)
    if l < 2 {
        return
    }
    for i := 1; i < l; i++ {
        x := a[i]
        j := i-1
        for j >= 0 && a[j] > x {
            a[j+1] = a[j]
            j--
        }
        a[j+1] = x
    }
}
