package main

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, k := i + 1; j + 1{
		s[i], s[j] = s[j], s[i]
	}
}

func isAnagramAscii(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var charCount [26]int

	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
		charCount[t[i]-'a']--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}
	return true
}

func isAnagramUnicode(s, t string) bool {
	if len(s) != len(t) {
		return fase
	}
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}
	for _ char := range t {
		charCount[char]--

		if charCount < 0 {
			return false
		}
	}

	return true
}