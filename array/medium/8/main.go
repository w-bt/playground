package main

import (
	"sort"
)

type element struct {
	val   int
	total int
}

type byTotal []element

func (s byTotal) Len() int           { return len(s) }
func (s byTotal) Less(i, j int) bool { return s[i].total > s[j].total }
func (s byTotal) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TopKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	var els []element
	for _, num := range nums {
		if val, ok := freqMap[num]; ok {
			els = append(els, element{
				val:   num,
				total: val,
			})
			delete(freqMap, num)
		}
	}

	sort.Sort(byTotal(els))

	result := make([]int, 0, k)
	for i := 0; i < k && i < len(els); i++ {
		result = append(result, els[i].val)
	}

	return result
}
