package move_zeroes

const zero int = 0

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	n := 0
	for _, v := range nums {
		if v != zero {
			nums[n] = v
			n++
		}
	}

	for i := range nums[n:] {
		nums[n+i] = zero
	}
}
