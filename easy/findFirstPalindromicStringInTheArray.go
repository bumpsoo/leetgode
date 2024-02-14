package easy

// https://leetcode.com/problems/find-first-palindromic-string-in-the-array
func firstPalindromeHelp(word string) bool {
	length := len(word)
	odd := length % 2
	half := length / 2
	s := make([]byte, 0, half)
	for i := 0; i < half; i++ {
		s = append(s, word[i])
	}
	for i := half + odd; i < length; i++ {
		v := s[len(s)-1]
		s = s[:len(s)-1]
		if v != word[i] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		if firstPalindromeHelp(word) {
			return word
		}
	}
	return ""
}
