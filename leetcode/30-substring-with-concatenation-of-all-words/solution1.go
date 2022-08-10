package main

func hasSubstring(s string, words []string) bool {
	word_len := len(words[0])
	splitted := make([]string, len(words))
	for i := range splitted {
		splitted[i] = s[i*word_len : i*word_len+word_len]
	}

WORDS:
	for _, word := range words {
		for i, e := range splitted {
			if word == e {
				splitted = append(splitted[:i], splitted[i+1:]...)
				continue WORDS
			}
		}
		return false
	}
	return true
}

func findSubstring(s string, words []string) []int {
	ans := []int{}
	for i := 0; i+len(words)*len(words[0]) <= len(s); i++ {
		if hasSubstring(s[i:], words) {
			ans = append(ans, i)
		}
	}
	return ans
}

var solution1 = findSubstring
