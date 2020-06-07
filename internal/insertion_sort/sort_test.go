package insertion_sort

import "testing"

func TestSort(t *testing.T) {
    {
        var a []int
        Sort(a)
        t.Log(a)
    }
    {
        a := make([]int, 0)
        Sort(a)
        t.Log(a)
    }
    {
        a := []int{1}
        Sort(a)
        t.Log(a)
    }
    {
        a := []int{1,2,3}
        Sort(a)
        t.Log(a)
    }
    {
        a := []int{9,8,7,6,5,4,3,2,1}
        Sort(a)
        t.Log(a)
    }
    {
        a := []int{90,8,70,25,16,52,34,41,31,62,19,88}
        Sort(a)
        t.Log(a)
    }
}
