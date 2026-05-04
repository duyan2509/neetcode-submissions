func lengthOfLongestSubstring(s string) int {
	exists := make(map[byte]int)
	start := 0
	rs := 0

	for i := 0; i < len(s); i++ {
		if idx, ok := exists[s[i]]; ok && idx >= start {
			start = idx + 1
		}

		exists[s[i]] = i

		if i-start+1 > rs {
			rs = i - start + 1
		}
	}

	return rs
}
