package solutions

import "fmt"

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)
	for _, v := range s {
		map1[v]++
	}
	for _, v := range t {
		map2[v]++
	}

	return fmt.Sprint(map1) == fmt.Sprint(map2)
}

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
