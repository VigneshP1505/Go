package main

import (
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
)

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
			left = lastSeen[ch]
		}
		lastSeen[ch] = right + 1
		currLength := right - left + 1
		if currLength > maxLength {
			maxLength = currLength
		}
	}
	return maxLength
}

func LongestSubstringNonRepeating(s string) int {
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
			left = lastSeen[ch]
		}
		lastSeen[ch] = right + 1
		currLength := right - left + 1
		if currLength > maxLength {
			maxLength = currLength
		}
	}
	return maxLength
}

// max sum path
func MaxPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	//first row only
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	//first column only
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

func Jump(n int) *big.Int {
	dp := make([]big.Int, n+1)
	dp[0].SetInt64(1)

	for i := 1; i <= n; i++ {
		for jump := 1; jump <= 5 && jump <= i; jump++ {
			dp[i].Add(&dp[i], &dp[i-jump])
		}
	}
	return &dp[n]
}

func CountOpposites(a []int) int {
	occurrences := make(map[int]int)
	result := 0

	//pass 1: add the negates to the map and increment the value
	for num := range a {
		if num < 0 {
			occurrences[num]++
		}
	}

	//pass 2: consider the positives and check if negates are present in the map, if yes, decrement the occurrences and increment result
	for num := range a {
		if num > 0 && occurrences[-num] > 0 {
			occurrences[-num]--
			if occurrences[-num] == 0 {
				delete(occurrences, -num)
			}
			result++
		}
	}
	return result
}

func CountOppositesTwoPointer(a []int) int {
	sort.Ints(a)
	left, right := 0, len(a)
	result := 0
	for left < right {
		sum := a[left] + a[right]
		if sum < 0 {
			left++
		} else if sum > 0 {
			right--
		} else {
			if a[left] < 0 && a[right] > 0 {
				result++
			}
			left++
			right--
		}
	}
	return result
}

type Interval struct {
	Start time.Time
	End   time.Time
}

func LongestInterval(intervals []Interval) time.Duration {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start.Before(intervals[j].Start)
	})

	currentStart := intervals[0].Start
	currentEnd := intervals[0].End
	longest := time.Duration(0)

	for _, interval := range intervals[1:] {
		if !interval.Start.After(currentEnd) {
			if interval.End.After(currentEnd) {
				currentEnd = interval.End
			}
		} else {
			span := currentEnd.Sub(currentStart)
			if span > longest {
				longest = span
			}
			currentStart = interval.Start
			currentEnd = interval.End
		}
	}

	if currentEnd.Sub(currentStart) > longest {
		longest = currentEnd.Sub(currentStart)
	}
	return longest
}

func LargestEmptySquare(board [][]bool) int {
	m := len(board)
	if m == 0 {
		return 0
	}
	n := len(board[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	maxSide := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if !board[i][j] {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				top := dp[i-1][j]
				left := dp[i][j-1]
				diagonal := dp[i-1][j-1]
				smallest := top
				if left < smallest {
					smallest = left
				}
				if diagonal < smallest {
					smallest = diagonal
				}
				dp[i][j] = smallest + 1
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide
}

func ReverseWords(s string) string {
	words := strings.Fields(s)
	left, right := 0, len(words)-1
	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}

func ReverseWordsNoInbuilt(s string) string {
	i := len(s) - 1
	result := make([]byte, 0, len(s))
	for i >= 0 {
		//ignore the whitespace
		for i >= 0 && s[i] == ' ' {
			i--
		}

		if i < 0 {
			break
		}

		end := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		start := i + 1

		if len(result) > 0 {
			result = append(result, ' ')
		}

		result = append(result, s[start:end+1]...)

	}
	return string(result)
}

func ValidAnagrams(s string, w string) bool {
	if len(s) != len(w) {
		return false
	}
	var a [26]int
	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
		a[w[i]-'a']--
	}
	for _, c := range a {
		if c != 0 {
			return false
		}
	}
	return true
}

func StringCompression(s string) string {
	if len(s) == 0 {
		return ""
	}
	var result strings.Builder
	j := 0
	for i := 0; i < len(s); i++ {
		j = i
		for j < len(s) && s[i] == s[j] {
			j++
		}
		result.WriteByte(s[i])
		result.WriteString(strconv.Itoa(j - i))
		i = j - 1
	}
	return result.String()
}

func StringRotation(s string, w string) bool {
	if len(s) != len(w) {
		return false
	}
	if s == "" {
		return true
	}
	s = s + s
	return strings.Contains(s, w)
}

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, maxLength := 0, 1
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		length := right - left - 1
		if length >= maxLength {
			maxLength = length
			start = left + 1
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[start : start+maxLength]
}
