package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	result := ""
	for {
		tempResult := ""
		index := 0
		flagStop := false
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == index {
				flagStop = true
				break
			}

			if i == 0 {
				tempResult += string(strs[i][index])
			} else {
				if len(strs[i]) >= index+1 && tempResult[len(tempResult)-1] == strs[i][index] {
					if i == len(strs)-1 {
						result = tempResult
						if index+1 < len(strs[i]) {
							i = -1 // reset i to start checking from the beginning again
						} else {
							flagStop = true
							break
						}
						index++
					}
				} else {
					flagStop = true
					break
				}
			}
		}

		if flagStop {
			break
		}
	}

	return result
}
