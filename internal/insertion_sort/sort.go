package insertion_sort

// in-place
// stable
// time: 
//   最优：O(n)
//   最差：O(n²)
//   平均：O(n²)
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
