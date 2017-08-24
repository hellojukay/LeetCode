package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs[0]) == 0 {
		return ""
	}
	var ch = strs[0][0]
	for index, str := range strs {
		if len(str) == 0 {
			return ""
		}
		if ch != str[0] {
			return ""
		}
		strs[index] = str[1:]
	}
	return string(ch) + longestCommonPrefix(strs)
}
