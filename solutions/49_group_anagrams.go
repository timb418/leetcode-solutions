package solutions

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, v := range strs {
		key := getSortedWordLetters(v)
		groups[key] = append(groups[key], v)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

func getSortedWordLetters(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func GroupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, word := range strs {
		count := [26]int{}
		for i := range word {
			count[word[i]-'a']++
		}
		m[count] = append(m[count], word)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}


