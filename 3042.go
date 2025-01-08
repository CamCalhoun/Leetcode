package main

func countPrefixSuffixPairs(words []string) int {
	var n int = len(words)
	var count int = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}

	return count
}

func isPrefixAndSuffix(str1, str2 string) bool {
	var n int = len(str1)
	m := len(str2)

	if m < n {
		return false
	}

	for i := 0; i < n; i++ {
		if str1[i] != str2[i] {
			return false
		}

		if str1[i] != str2[m-n+i] {
			return false
		}
	}

	return true
}
