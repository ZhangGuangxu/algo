package quick_sort

import "github.com/ZhangGuangxu/algo/internal/shell_sort"

// 非稳定排序
// 时间复杂度：
//   最差：O(n^2)
//   最优：O(nlogn)
//   平均：O(nlogn)
// 空间复杂度：
//   非递归部分：O(1)
//   递归部分：
//     最优：O(logn)
//     最差：O(n)
func Sort(a []int) {
    l := len(a)
    if l < 2 {
        return
    }
    quickSort(a, 0, l-1)
}

func quickSort(a []int, low, high int) {
    if high - low > 16 { // high - low + 1 > 17
        for low < high {
            pivotIndex := doPivot(a, low, high)
            quickSort(a, low, pivotIndex-1)
            low = pivotIndex+1
        }
    } else {
        shell_sort.Sort(a[low:high+1])
    }
}

func doPivot(a []int, low, high int) int {
    if high - low > 1 { // high - low + 1 >= 3
        m := low + (high-low)/2
        if a[low] > a[high] {
            a[low], a[high] = a[high], a[low]
        }
        if a[m] > a[high] {
            a[m], a[high] = a[high], a[m]
        }
        if a[m] > a[low] {
            a[m], a[low] = a[low], a[m]
        }
        // a[m] <= a[low] <= a[high]
    }
    pivot := a[low]
    for low < high {
        for low < high && a[high] >= pivot {
            high--
        }
        a[low] = a[high]
        for low < high && a[low] <= pivot {
            low++
        }
        a[high] = a[low]
    }
    a[low] = pivot
    return low
}
