package main

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for {

		if left >= right {
			return true
		}
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}
		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
}

func isAlphaNumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || isUpperAlpha(c) || c >= '0' && c <= '9'
}

func isUpperAlpha(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func toLower(c byte) byte {
	if isUpperAlpha(c) {
		return c + 'a' - 'A'
	}
	return c
}
