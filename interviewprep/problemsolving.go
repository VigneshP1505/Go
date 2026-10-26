package main

func TwoSum(a []int, target int) (int, int) {
	if len(a) == 0 {
		return -1, -1
	}
	store := map[int]int{}
	for i, num := range a {
		diff := target - num
		if j, exists := store[diff]; exists {
			return j, i
		}
		store[num] = i
	}
	return -1, -1
}

func ContainsDuplicate(a []int) bool {
	seen := map[int]int{}
	for i, num := range a {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = i
	}
	return false
}

func ValidAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	count := map[rune]int{}
	for _, ch := range a {
		count[ch]++
	}
	for _, ch := range b {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return -1
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}

func MaxDifference(a []int) int {
	if len(a) == 0 {
		panic("Invalid input")
	}
	min := a[0]
	maxDiff := 0
	for _, num := range a {
		diff := num - min
		if diff > maxDiff {
			maxDiff = diff
		}
		if num < min {
			min = num
		}
	}
	return maxDiff
}

func MoveZeros(a []int) {
	idx1 := 0
	for i := range a {
		if a[i] != 0 {
			a[idx1], a[i] = a[i], a[idx1]
			idx1++
		}
	}
}

// dynamic programming
func Fibonacci(n int, table map[int]int) int {
	if n == 1 || n == 0 {
		return n
	}
	if _, exists := table[n]; exists {
		return table[n]
	}
	result := Fibonacci(n-1, table) + Fibonacci(n-2, table)
	table[n] = result
	return result
}
func DriveFibonacci(n int) {
	table := map[int]int{}
	Fibonacci(n, table)
}

func ClimbingStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return ClimbingStairs(n-1) + ClimbingStairs(n-2)
}

// You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the staircase, which is the position just past the last step (index cost.length).
func MinCostStairs(cost []int, step int, table map[int]int) int {
	if step == 0 || step == 1 {
		return cost[step]
	}
	if _, exists := table[step]; exists {
		return table[step]
	}
	left := MinCostStairs(cost, step-1, table)
	right := MinCostStairs(cost, step-2, table)
	result := cost[step] + min(left, right)
	table[step] = result
	return result
}

func MinCostToClimbStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[n-1], dp[n-2])
}

func MinCost(cost []int, step int) int {
	if step == 0 || step == 1 {
		return cost[step]
	}
	left := MinCost(cost, step-1)
	right := MinCost(cost, step-2)
	return cost[step] + min(left, right)
}

// Given an array of integers nums, find the maximum possible sum of elements such that no two selected elements are adjacent in the original array.
// You may choose zero or more elements, but you cannot select two neighboring elements.
func MaxSumNonAdjacent(num []int) int {
	n := len(num)
	dp := make([]int, n)
	dp[0] = max(0, num[0])

	if n == 1 {
		return dp[0]
	}

	dp[1] = max(dp[0], num[1])

	for i := 2; i < n; i++ {
		take := num[i] + dp[i-2]
		skip := dp[i-1]
		dp[i] = max(take, skip)
	}
	return dp[n-1]
}

// Given an integer array nums, find a non-empty contiguous subarray that has the largest product, and return that product.
// A subarray is a contiguous part of the array.
func MaxSubarrayProduct(num []int) int {
	if len(num) == 0 {
		return 0
	}
	maxEnding := num[0]
	minEnding := num[0]
	result := num[0]

	for i := 1; i < len(num); i++ {
		x := num[i]
		if x < 0 {
			maxEnding, minEnding = minEnding, maxEnding
		}
		maxEnding = max(x, x*maxEnding)
		minEnding = min(x, x*minEnding)
		result = max(result, maxEnding)
	}
	return result
}

func WordBreak(word string, dict map[string]bool) bool {
	n := len(word)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func LongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLength := 0
	left := 0
	var lastSeen [256]int

	for right := 0; right < n; right++ {
		ch := s[right]

		if lastSeen[ch] > left {
			left = lastSeen[right]
		}
		lastSeen[ch] = right + 1
		currLength := right - left + 1
		if currLength > maxLength {
			maxLength = currLength
		}
	}
	return maxLength
}
