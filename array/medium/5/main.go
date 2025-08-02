package main

func Merge(intervals [][]int) [][]int {
	result := [][]int{}
	sortInterval(intervals)
	if len(intervals) > 0 {
		result = append(result, intervals[0])
	}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= result[len(result)-1][1] {
			if intervals[i][1] >= result[len(result)-1][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func sortInterval(intervals [][]int) {
	for i := 0; i < len(intervals)-1; i++ {
		var isSwapped bool
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][0] > intervals[j+1][0] || (intervals[j][0] == intervals[j+1][0] && intervals[j][1] > intervals[j+1][1]) {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
