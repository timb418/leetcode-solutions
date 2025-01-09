package solutions

func FindAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	lenS := len(s)
	lenP := len(p)
	pCharCount := [26]int{}
	for _, v := range p {
		pCharCount[v-'a']++
	}

	idx := []int{}
	for i := 0; i <= lenS-lenP; i++ {
		if isAnagram := isAnagramOf(s[i:i+lenP], pCharCount); isAnagram {
			idx = append(idx, i)
		}
	}
	return idx
}

func isAnagramOf(s string, pCharCount [26]int) bool {
	sCharCount := [26]int{}
	for _, v := range s {
		sCharCount[v-'a']++
	}

	return sCharCount == pCharCount
}
