package main

func lengthOfLongestSubstring(s string) int {
	var m int
	length := len(s)
	if length <= 1 {
		return length
	}
	for i := 0; i < length; i++ {
		set := make(map[byte]bool)
		k := 0
		for j := i; j < length; j++ {
			ch := s[j]
			_, has := set[ch]
			if has {
				break
			} else {
				set[ch] = true
				k = k + 1
				if k > m {
					m = k
				}
			}
		}
	}
	return m
}
