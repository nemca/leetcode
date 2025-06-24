package binary_search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}
