package first_unique_character_in_a_string

func firstUniqChar(s string) int {
	index := -1
	if len(s) < 1 {
		return index
	}

	if len(s) == 1 {
		return 0
	}

	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}

	for i, ch := range s {
		if val, ok := m[ch]; ok {
			if val == 1 {
				return i
			}
		}
	}

	return index
}

func firstUniqChar2(s string) int {
	var count [256]byte
	n := len(s)

	for i := 0; i < n; i++ {
		if count[s[i]] < 2 {
			count[s[i]]++
		}
	}

	for i := 0; i < n; i++ {
		if count[s[i]] == 1 {
			return i
		}
	}

	return -1
}
