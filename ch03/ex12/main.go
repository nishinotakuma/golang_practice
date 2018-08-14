package anagram

func anagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}

	if len(aFreq) != len(bFreq) {
		return false
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	return true
}
