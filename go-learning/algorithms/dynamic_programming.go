package algorithms

// Dynamic Programming solves problems by breaking them into overlapping subproblems
// and storing results to avoid redundant computation.

// Fibonacci calculates the nth Fibonacci number using DP.
// Time Complexity: O(n)
// Space Complexity: O(n)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// FibonacciOptimized calculates Fibonacci with O(1) space.
// Time Complexity: O(n)
// Space Complexity: O(1)
func FibonacciOptimized(n int) int {
	if n <= 1 {
		return n
	}

	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// ClimbingStairs calculates the number of ways to climb n stairs.
// You can climb 1 or 2 steps at a time.
// Time Complexity: O(n)
// Space Complexity: O(1)
func ClimbingStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2

	for i := 3; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// CoinChange finds the minimum number of coins needed to make the amount.
// Returns -1 if the amount cannot be made up.
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
	// dp[i] = minimum coins needed to make amount i
	dp := make([]int, amount+1)

	// Initialize with amount+1 (impossible value)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// LongestIncreasingSubsequence finds the length of the longest increasing subsequence.
// Time Complexity: O(n²) - can be optimized to O(n log n) with binary search
// Space Complexity: O(n)
func LongestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] = length of LIS ending at index i
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // Each element is an LIS of length 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

// LongestCommonSubsequence finds the length of the LCS of two strings.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)

	// dp[i][j] = LCS of text1[0:i] and text2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// EditDistance calculates minimum operations to convert word1 to word2.
// Operations: insert, delete, replace
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func EditDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)

	// dp[i][j] = min operations to convert word1[0:i] to word2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // Delete all characters
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // Insert all characters
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation needed
			} else {
				dp[i][j] = 1 + minThree(
					dp[i-1][j],   // Delete
					dp[i][j-1],   // Insert
					dp[i-1][j-1], // Replace
				)
			}
		}
	}

	return dp[m][n]
}

// Knapsack01 solves the 0/1 knapsack problem.
// Returns the maximum value that can be put in a knapsack of capacity W.
// Time Complexity: O(n * W)
// Space Complexity: O(n * W)
func Knapsack01(weights, values []int, W int) int {
	n := len(weights)

	// dp[i][w] = max value using items 0..i-1 with capacity w
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			// Don't take item i-1
			dp[i][w] = dp[i-1][w]

			// Take item i-1 if it fits
			if weights[i-1] <= w {
				withItem := values[i-1] + dp[i-1][w-weights[i-1]]
				if withItem > dp[i][w] {
					dp[i][w] = withItem
				}
			}
		}
	}

	return dp[n][W]
}

// UniquePaths finds the number of unique paths from top-left to bottom-right.
// You can only move right or down.
// Time Complexity: O(m * n)
// Space Complexity: O(n)
func UniquePaths(m, n int) int {
	// dp[j] = number of ways to reach column j in current row
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // First row: only one way to reach any cell
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1] // From top + from left
		}
	}

	return dp[n-1]
}

// HouseRobber finds the maximum amount of money you can rob.
// You cannot rob adjacent houses.
// Time Complexity: O(n)
// Space Complexity: O(1)
func HouseRobber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2, prev1 := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		curr := maxInt(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

// MaxSubArray finds the contiguous subarray with the largest sum (Kadane's algorithm).
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either start fresh from current element or extend previous subarray
		currentSum = maxInt(nums[i], currentSum+nums[i])
		maxSum = maxInt(maxSum, currentSum)
	}

	return maxSum
}

// WordBreak determines if s can be segmented into dictionary words.
// Time Complexity: O(n² * m) where m is average word length for comparison
// Space Complexity: O(n)
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// dp[i] = true if s[0:i] can be segmented
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// Helper functions
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minThree(a, b, c int) int {
	result := a
	if b < result {
		result = b
	}
	if c < result {
		result = c
	}
	return result
}
