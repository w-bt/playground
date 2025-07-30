package main

import "sort"

type element struct {
	val   int
	total int
}

type byTotal []element

func (s byTotal) Len() int           { return len(s) }
func (s byTotal) Less(i, j int) bool { return s[i].total > s[j].total }
func (s byTotal) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TopKFrequent(nums []int, k int) []int {
	mapNums := make(map[int]int)
	for _, num := range nums {
		mapNums[num] += 1
	}

	var els []element
	for k, v := range mapNums {
		els = append(els, element{
			val:   k,
			total: v,
		})
	}
	sort.Sort(byTotal(els))

	result := []int{}
	for i := 0; i < k; i++ {
		if i > len(els)-1 {
			break
		}
		result = append(result, els[i].val)
	}

	return result
}
