package main

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	difArray := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			difArray[start]++
			difArray[end+1]--
		} else {
			difArray[start]--
			difArray[end+1]++
		}
	}

	shiftNum := 0
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		shiftNum += difArray[i]
		shifted := (int(s[i]-'a') + shiftNum) % 26
		if shifted < 0 {
			shifted += 26
		}
		result[i] = byte(shifted) + 'a'
	}

	return string(result)

	return "fart"
}
