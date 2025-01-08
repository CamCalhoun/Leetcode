func minOperations(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)

	leftCost, leftBalls := 0, 0
	rightCost, rightBalls := 0, 0

	// Left-to-right pass
	for i := 0; i < n; i++ {
		result[i] += leftCost
		if boxes[i] == '1' {
			leftBalls++
		}
		leftCost += leftBalls
	}

	// Right-to-left pass
	for i := n - 1; i >= 0; i-- {
		result[i] += rightCost
		if boxes[i] == '1' {
			rightBalls++
		}
		rightCost += rightBalls
	}

	return result
}
