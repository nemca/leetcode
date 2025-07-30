package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, n := range nums1 {
		if _, ok := m1[n]; !ok {
			m1[n]++
		}
	}

	for _, n := range nums2 {
		if _, ok := m1[n]; ok {
			if _, ok := m2[n]; !ok {
				result = append(result, n)
				m2[n]++
			}
		}
	}

	return result
}
