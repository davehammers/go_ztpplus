package util

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 9, 10, 20, 21, 40, 50, 51, 100}
	t.Log(SlotSliceToString(0, n))
	t.Log(SlotSliceToString(3, n))

}

func TestDecode(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 9, 10, 20, 21, 40, 50, 51, 100}
	a := SlotSliceToString(1, n)
	b := SlotSliceToString(2, n)
	c := SlotSliceToString(3, n)
	d := SlotSliceToString(4, n)
	e := SlotSliceToString(5, n)
	f := SlotSliceToString(6, n)
	g := SlotSliceToString(7, n)
	h := SlotSliceToString(8, n)
	x := strings.Join([]string{a, b, c, d, e, f, g, h}, ",")
	y := StringToMap(x)
	for k, v := range y {
		t.Log(k, v)
	}
	y = StringToMap("4:100-200")
	for k, v := range y {
		t.Log(k, v)
	}

}
func TestDecode1(t *testing.T) {
	m := make(map[int][]int)
	n := []int{1, 2, 3, 4, 5, 9, 10, 20, 21, 40, 50, 51, 100}
	// slot==0 is no slot number in output
	m[0] = n
	t.Log(SlotMapToString(m))
}
func TestDecode2(t *testing.T) {
	m := make(map[int][]int)
	n := []int{1, 2, 3, 4, 5, 9, 10, 20, 21, 40, 50, 51, 100}
	s := "3:1-5,3:9-10,4:20-21,3:40,3:50-51,5:100"
	t.Log(s)
	t.Log(StringToMap(s))
	s = "1-20"
	t.Log(s)
	t.Log(StringToMap(s))
	m[1] = StringToMap(s)[0]
	m[4] = StringToMap("40-60,71-90")[0]
	m[5] = n
	for k, v := range m {
		t.Log(k, v)
	}
	t.Log(SlotMapToString(m))
}
