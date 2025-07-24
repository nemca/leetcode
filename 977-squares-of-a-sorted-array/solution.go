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

func sortedSquares2(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1
	square := 0

	for i := n - 1; i >= 0; i-- {
		if abs(nums[left]) < abs(nums[right]) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}

		result[i] = square * square
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
