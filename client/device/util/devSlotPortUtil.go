package util

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func rangeAppend(ranges *[]string, slot int, start int, end int) {
	if start == end {
		if slot > 0 {
			*ranges = append(*ranges, fmt.Sprintf("%d:%d", slot, start))
		} else {
			*ranges = append(*ranges, fmt.Sprintf("%d", start))
		}
	} else {
		if slot > 0 {
			*ranges = append(*ranges, fmt.Sprintf("%d:%d-%d", slot, start, end))
		} else {
			*ranges = append(*ranges, fmt.Sprintf("%d-%d", start, end))
		}
	}
}

//Takes a single slot number and an array of integer port numbers and outputs a string
//representing the slot:port ranges
//
//e.g. slot=3, []int{1,2,3,4,5,10,11,12]
// returns 3:1-5,3:10-12
//
//if slot=0, the output only contains the port numbers
//
//e.g. slot=0, []int{1,2,3,4,5,10,11,12]
// returns 1-5,10-12
func SlotSliceToString(slot int, n []int) (str string) {
	var start, end int
	ranges := make([]string, 0)

	// numbers must be in ascending order
	sort.Ints(n)
	delta := 0

	for i, v := range n {
		// compute the delta between the position in the slice and the value
		// if the values increase atomicly, the delta will be constant
		d := v - i
		if d != delta {
			if i > 0 {
				// package previous range
				rangeAppend(&ranges, slot, start, end)
			}
			start = v
			delta = d
		}
		end = v
	}
	// package last range
	rangeAppend(&ranges, slot, start, end)
	return strings.Join(ranges, ",")
}

//Converts a map of slots with []int ports into a single string
//    1 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40]
//    4 [40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90]
//    5 [1 2 3 4 5 9 10 20 21 40 50 51 100]
//
//Returns:
//1:1-40,4:40-60,4:71-90,5:1-5,5:9-10,5:20-21,5:40,5:50-51,5:100
func SlotMapToString(m map[int][]int) (str string) {
	strs := make([]string, 0)
	for k, v := range m {
		strs = append(strs, SlotSliceToString(k, v))
	}
	return strings.Join(strs, ",")
}

//Converts a string of slot:port-port,slot:port-port into map[int][]int
//If the string does not have slot numbers, all results are returned to map[0]
//
// e.g.
// 3:1-5,3:9-10,4:20-21,3:40,3:50-51,5:100
//
//Returns:
// map[3:[1 2 3 4 5 9 10 40 50 51] 4:[20 21] 5:[100]]
//
//e.g.
// 1-40
//
//Returns:
//map[0:[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]]
func StringToMap(s string) (m map[int][]int) {
	var portStr string
	m = make(map[int][]int)

	parts := strings.Split(s, ",")
	for _, p := range parts {
		slot := 0
		if strings.Contains(p, ":") {
			entryStr := strings.Split(p, ":")
			slot, _ = strconv.Atoi(entryStr[0])
			portStr = entryStr[1]
		} else {
			portStr = p
		}
		portRange := strings.Split(portStr, "-")
		start, _ := strconv.Atoi(portRange[0])
		end := start
		if len(portRange) > 1 {
			end, _ = strconv.Atoi(portRange[1])
		}
		for port := start; port <= end; port++ {
			m[slot] = append(m[slot], port)
		}
	}
	return
}
