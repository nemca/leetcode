package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1
	curr := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[curr] = leftSq
			left++
		} else {
			result[curr] = rightSq
			right--
		}
		curr--
	}

	return result
}
