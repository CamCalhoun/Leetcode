package main

func vowelStrings(words []string, queries [][]int) []int {
	var numerizedStrings []int
	for _, word := range words {
		runes := []rune(word)
		first := runes[0]
		last := runes[len(runes)-1]
		numerizedStrings = append(numerizedStrings, areVowels(first, last))
	}
	prefixArr := prefixSum(numerizedStrings)

	ans := make([]int, len(queries))
	for i, query := range queries {
		l := query[0]
		r := query[1]

		if l == 0 {
			ans[i] = prefixArr[r+1]
		} else {
			ans[i] = prefixArr[r+1] - prefixArr[l]
		}
	}

	return ans
}

func areVowels(f rune, l rune) int {
	vowels := "aeiou"
	isInVowels := func(r rune) bool {
		for _, vowel := range vowels {
			if r == vowel {
				return true
			}
		}
		return false
	}
	if isInVowels(f) && isInVowels(l) {
		return 1
	}
	return 0
}

func prefixSum(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n+1)

	prefix[0] = 0

	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	return prefix
}
