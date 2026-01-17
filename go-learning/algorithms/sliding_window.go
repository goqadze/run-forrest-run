package algorithms

// Sliding Window technique uses a window that slides through data
// to find subarrays/substrings that satisfy certain conditions.

// MaxSumSubarray finds the maximum sum of a subarray with size k.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSumSubarray(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// Slide the window
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k] // Add new element, remove old
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// MinSubArrayLen finds the minimum length subarray with sum >= target.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Shrink window while sum >= target
		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0 // No valid subarray found
	}
	return minLen
}

// LengthOfLongestSubstring finds the length of longest substring without repeating characters.
// Time Complexity: O(n)
// Space Complexity: O(min(m, n)) where m is charset size
func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // Maps character to its most recent index
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// If character was seen and is within current window
		if idx, exists := charIndex[char]; exists && idx >= left {
			left = idx + 1 // Move left pointer past the duplicate
		}

		charIndex[char] = right
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

// MaxConsecutiveOnes finds the max consecutive 1s if you can flip at most k 0s.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxConsecutiveOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		// Shrink window if we have more than k zeros
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

// CharacterReplacement finds the longest substring with same letters after k replacements.
// Time Complexity: O(n)
// Space Complexity: O(1) - only 26 letters
func CharacterReplacement(s string, k int) int {
	count := make([]int, 26)
	left := 0
	maxCount := 0 // Count of most frequent character in window
	maxLen := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++

		// Update max count of any character in window
		if count[s[right]-'A'] > maxCount {
			maxCount = count[s[right]-'A']
		}

		// Window size - max count = characters to replace
		// If this exceeds k, shrink window
		windowSize := right - left + 1
		if windowSize-maxCount > k {
			count[s[left]-'A']--
			left++
		}

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

// FindAnagrams finds all start indices of anagrams of pattern in string.
// Time Complexity: O(n)
// Space Complexity: O(1) - fixed size arrays
func FindAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	result := []int{}
	pCount := make([]int, 26)
	sCount := make([]int, 26)

	// Count characters in pattern
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		// Add current character to window
		sCount[s[i]-'a']++

		// Remove character that's no longer in window
		if i >= len(p) {
			sCount[s[i-len(p)]-'a']--
		}

		// Compare counts
		if i >= len(p)-1 && arraysEqual(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func arraysEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MinWindowSubstring finds the minimum window in s that contains all characters of t.
// Time Complexity: O(n + m)
// Space Complexity: O(m) where m is unique characters in t
func MinWindowSubstring(s, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Count characters needed
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	required := len(need)
	formed := 0

	left := 0
	minLen := len(s) + 1
	minStart := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		have[char]++

		// Check if current character satisfies its requirement
		if need[char] > 0 && have[char] == need[char] {
			formed++
		}

		// Try to shrink window
		for formed == required {
			windowLen := right - left + 1
			if windowLen < minLen {
				minLen = windowLen
				minStart = left
			}

			// Remove leftmost character
			leftChar := s[left]
			have[leftChar]--
			if need[leftChar] > 0 && have[leftChar] < need[leftChar] {
				formed--
			}
			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// SlidingWindowMaximum finds the maximum in each sliding window of size k.
// Uses a monotonic decreasing deque.
// Time Complexity: O(n)
// Space Complexity: O(k)
func SlidingWindowMaximum(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	result := make([]int, 0, len(nums)-k+1)
	deque := []int{} // Stores indices

	for i := 0; i < len(nums); i++ {
		// Remove elements outside the window
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove smaller elements (they can't be maximum)
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// Add to result once we have a full window
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
