package main

func waysToSplitArray(nums []int) int {
	result := 0
	prefix := (prefixSum(nums))
	totalSum := prefix[len(nums)]

	for i := 0; i < len(nums)-1; i++ {
		leftSum := prefix[i+1]
		rightSum := totalSum - leftSum
		if leftSum >= rightSum {
			result++
		}
	}
	return result
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
