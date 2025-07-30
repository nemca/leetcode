package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	forward := make(map[rune]rune)
	reverse := make(map[rune]rune)

	for i, ch := range s {
		if _, ok := forward[ch]; ok {
			if forward[ch] != rune(t[i]) {
				return false
			}
			if _, ok := reverse[rune(t[i])]; ok {
				if ch != reverse[rune(t[i])] {
					return false
				}
			}
		} else {
			forward[ch] = rune(t[i])
			if _, ok := reverse[rune(t[i])]; ok {
				return false
			}
			reverse[rune(t[i])] = ch
		}
	}

	return true
}
