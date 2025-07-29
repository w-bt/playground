package main

func PrefixSuffixSum(input []int) ([]int, []int) {
	if len(input) == 0 {
		return nil, nil
	}

	prefixSum := make([]int, len(input))
	suffixSum := make([]int, len(input))

	prefixSum[0] = input[0]
	for i := 1; i < len(input); i++ {
		prefixSum[i] = prefixSum[i-1] + input[i]
	}

	suffixSum[len(input)-1] = input[len(input)-1]
	for i := len(input) - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + input[i]
	}

	return prefixSum, suffixSum
}
