package main

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}

	if len(newInterval) == 0 {
		return intervals
	}

	var temp []int
	for i, interval := range intervals {
		if len(temp) > 0 {
			if interval[0] > temp[1] && interval[1] > temp[0] && interval[1] > temp[1] {
				newValue := []int{temp[0], temp[1]}
				result = append(result, newValue)
				result = append(result, interval)
				temp = []int{}
			} else if interval[0] > temp[0] && temp[1] >= interval[0] && interval[1] > temp[1] {
				temp[1] = interval[1]
			} else if interval[0] < temp[0] && temp[0] <= interval[1] && interval[1] < temp[1] {
				temp[0] = interval[0]
			} else if interval[0] < temp[0] && interval[1] < temp[1] {
				temp[0] = interval[0]
				temp[1] = interval[1]
			}
		} else if interval[0] > newInterval[1] && interval[1] > newInterval[0] {
			result = append(result, newInterval)
			result = append(result, intervals[i:len(intervals)]...)
			newInterval = []int{}
			break
		} else if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else {
			if interval[0] > newInterval[0] && newInterval[0] >= interval[1] && interval[1] > newInterval[1] {
				temp = append(temp, newInterval[0], interval[1])
				newInterval = []int{}
			} else if interval[0] < newInterval[0] && newInterval[0] <= interval[1] && interval[1] < newInterval[1] {
				temp = append(temp, interval[0], newInterval[1])
				newInterval = []int{}
			} else if interval[0] < newInterval[0] && interval[1] < newInterval[1] {
				temp = append(temp, interval[0], interval[1])
				newInterval = []int{}
			} else if newInterval[0] < interval[0] && newInterval[1] < interval[1] {
				temp = append(temp, newInterval[0], newInterval[1])
				newInterval = []int{}
			}
		}
	}

	if len(temp) > 0 {
		result = append(result, temp)
	}

	if len(newInterval) > 0 {
		result = append(result, newInterval)
	}

	return result
}
