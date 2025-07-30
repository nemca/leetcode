package group_anagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	if len(strs) < 2 {
		return append(result, strs)
	}

	result = append(result, []string{strs[0]})

	for i := 1; i < len(strs); i++ {
		for k := range result {
			if len(strs[i]) == len(result[k][0]) {
				if isAnagrams(strs[i], result[k][0]) {
					result[k] = append(result[k], strs[i])
					break
				}
			}
			if k == len(result)-1 {
				result = append(result, []string{strs[i]})
			}
		}
	}

	return result
}

func isAnagrams(a, b string) bool {
	ma := make(map[rune]int)
	for _, r := range a {
		ma[r]++
	}

	mb := make(map[rune]int)
	for _, r := range b {
		mb[r]++
	}

	return mapsEqual(ma, mb)
}

func mapsEqual(m1, m2 map[rune]int) bool {
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	m := make(map[string][]string)

	for _, s := range strs {
		sorted := sortString(s)
		m[sorted] = append(m[sorted], s)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
