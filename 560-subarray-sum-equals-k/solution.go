package subarray_sum_equals_k

// Brute force O(n^2)
func subarraySumBrute(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			result++
		}

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				result++
			}
		}
	}

	return result
}

// Prefix sums and hash-map O(n)
func subarraySum(nums []int, k int) int {
	prefixSum := 0
	count := 0
	countMap := map[int]int{0: 1}

	for _, n := range nums {
		prefixSum += n
		if val, ok := countMap[prefixSum-k]; ok {
			count += val
		}
		countMap[prefixSum]++
	}

	return count
}
