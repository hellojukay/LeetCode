package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	var strs = strings.Split(s, " ")
	if len(strs) == 0 {
		return 0
	}
	var length = len(strs)
	for i := length - 1; i >= 0; i-- {
		if func(str string) bool {
			if len(str) == 0 {
				return false
			}
			for index := range str {
				if !(('a' <= str[index] && str[index] <= 'z') || ('A' <= str[index] && str[index] <= 'Z')) {
					return false
				}
			}
			return true
		}(strs[i]) {
			return len(strs[i])
		}
	}
	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
