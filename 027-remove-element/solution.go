package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	n := 0
	for _, v := range nums {
		if v != val {
			nums[n] = v
			n++
		}
	}

	return len(nums[:n])
}
