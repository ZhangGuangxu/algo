package merge_sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	{
		src := []int{1}
		t.Logf("src=%v", src)
		l := len(src)
		dst := make([]int, l)
		var s int
		e := l - 1
		m := e / 2
		merge(src, dst, s, m, e)
		t.Logf("s=%d,m=%d,e=%d", s, m, e)
		t.Logf("dst=%v", dst)
	}
	{
		src := []int{101, 1}
		t.Logf("src=%v", src)
		l := len(src)
		dst := make([]int, l)
		var s int
		e := l - 1
		m := e / 2
		merge(src, dst, s, m, e)
		t.Logf("s=%d,m=%d,e=%d", s, m, e)
		t.Logf("dst=%v", dst)
	}
	{
		src := []int{99, 1, 88}
		t.Logf("src=%v", src)
		l := len(src)
		dst := make([]int, l)
		var s int
		e := l - 1
		m := e / 2
		merge(src, dst, s, m, e)
		t.Logf("s=%d,m=%d,e=%d", s, m, e)
		t.Logf("dst=%v", dst)
	}
	{
		src := []int{2, 8, 3, 6}
		t.Logf("src=%v", src)
		l := len(src)
		dst := make([]int, l)
		var s int
		e := l - 1
		m := e / 2
		merge(src, dst, s, m, e)
		t.Logf("s=%d,m=%d,e=%d", s, m, e)
		t.Logf("dst=%v", dst)
	}
	{
		src := []int{2, 8, 9, 5, 6}
		t.Logf("src=%v", src)
		l := len(src)
		dst := make([]int, l)
		var s int
		e := l - 1
		m := e / 2
		merge(src, dst, s, m, e)
		t.Logf("s=%d,m=%d,e=%d", s, m, e)
		t.Logf("dst=%v", dst)
	}
}

func TestMergePass(t *testing.T) {
	{
		var src []int
		var dst []int
		mergePass(src, dst, 1)
	}
	{
		src := []int{1}
		dst := []int{0}
		mergePass(src, dst, 1)
	}
	{
		src := []int{99, 1}
		dst := make([]int, len(src))
		mergePass(src, dst, 1)
		t.Logf("dst=%v", dst)
		for i := 0; i < len(dst); i++ {
			if i+1 <len(dst) && dst[i] > dst[i+1] {
				t.Error("failed")
				return
			}
		}
	}
	{
		src := []int{99, 1, 88}
		dst := make([]int, len(src))
		segLen := 1
		mergePass(src, dst, segLen)
		t.Logf("dst=%v", dst)
		segLen <<= 1
		mergePass(dst, src, segLen)
		t.Logf("src=%v", src)
		for i := 0; i < len(src); i++ {
			if i+1 <len(src) && src[i] > src[i+1] {
				t.Error("failed")
				return
			}
		}
	}
	{
		src := []int{99, 1, 88, 10}
		dst := make([]int, len(src))
		segLen := 1
		mergePass(src, dst, segLen)
		t.Logf("dst=%v", dst)
		segLen <<= 1
		mergePass(dst, src, segLen)
		t.Logf("src=%v", src)
		for i := 0; i < len(src); i++ {
			if i+1 <len(src) && src[i] > src[i+1] {
				t.Error("failed")
				return
			}
		}
	}
	{
		src := []int{99, 1, 88, 10, 21}
		l := len(src)
		dst := make([]int, l)
		segLen := 1
		for segLen < l {
			mergePass(src, dst, segLen)
			t.Logf("dst=%v", dst)
			segLen <<= 1
			mergePass(dst, src, segLen)
			t.Logf("src=%v", src)
			segLen <<= 1
		}
		for i := 0; i < len(src); i++ {
			if i+1 <len(src) && src[i] > src[i+1] {
				t.Error("failed")
				return
			}
		}
	}
	{
		src := []int{3051, 48964, 129, 26413, 37493, 98888, 36814}
		l := len(src)
		dst := make([]int, l)
		segLen := 1
		for segLen < l {
			mergePass(src, dst, segLen)
			t.Logf("dst=%v", dst)
			segLen <<= 1
			mergePass(dst, src, segLen)
			t.Logf("src=%v", src)
			segLen <<= 1
		}
		for i := 0; i < len(src); i++ {
			if i+1 <len(src) && src[i] > src[i+1] {
				t.Error("failed")
				return
			}
		}
	}
}

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
		a := []int{3051, 48964, 129, 26413, 37493, 98888, 36814}
		Sort(a)
		t.Log(a)
		for j := 0; j < len(a); j++ {
			if j+1 < len(a) && a[j] > a[j+1] {
				t.Error("failed")
				return
			}
		}
	}

	rand.Seed(time.Now().UnixNano())

	for l := 1; l < 1000; l++ {
		//t.Logf("length=%d", l)
		a := make([]int, 0, l)
		for i := 0; i < l; i++ {
			r := rand.Intn(99999)
			a = append(a, r)
		}
		//t.Logf("before, %v", a)
		Sort(a)
		//t.Logf("after, %v", a)
		for j := 0; j < len(a); j++ {
			if j+1 < len(a) && a[j] > a[j+1] {
				t.Errorf("failed, length=%d", l)
				return
			}
		}
	}
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UnixNano())
	const l = 1000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := make([]int, 0, l)
		for i := 0; i < l; i++ {
			r := rand.Intn(99999)
			a = append(a, r)
		}
		b.StartTimer()
		Sort(a)
	}
}
