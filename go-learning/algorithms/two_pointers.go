package algorithms

// Two Pointers technique uses two pointers to iterate through data structures.
// Common patterns: opposite ends, same direction (fast/slow), sliding window.

// TwoSum finds two numbers in a SORTED array that add up to target.
// Returns indices of the two numbers.
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil // No solution found
}

// ThreeSum finds all unique triplets that sum to zero.
// Time Complexity: O(n²)
// Space Complexity: O(1) excluding output
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	sorted := MergeSort(nums) // Sort the array first

	for i := 0; i < len(sorted)-2; i++ {
		// Skip duplicates
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}

		left, right := i+1, len(sorted)-1
		target := -sorted[i]

		for left < right {
			sum := sorted[left] + sorted[right]

			if sum == target {
				result = append(result, []int{sorted[i], sorted[left], sorted[right]})

				// Skip duplicates
				for left < right && sorted[left] == sorted[left+1] {
					left++
				}
				for left < right && sorted[right] == sorted[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// ContainerWithMostWater finds the maximum area of water that can be contained.
// Given heights of vertical lines, find two lines that together with x-axis
// forms a container that holds the most water.
// Time Complexity: O(n)
// Space Complexity: O(1)
func ContainerWithMostWater(heights []int) int {
	left, right := 0, len(heights)-1
	maxArea := 0

	for left < right {
		width := right - left
		height := min(heights[left], heights[right])
		area := width * height

		if area > maxArea {
			maxArea = area
		}

		// Move the pointer with smaller height
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// RemoveDuplicates removes duplicates from a sorted array in-place.
// Returns the new length.
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Slow pointer tracks position for unique elements
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// MoveZeroes moves all zeroes to the end while maintaining relative order.
// Time Complexity: O(n)
// Space Complexity: O(1)
func MoveZeroes(nums []int) {
	slow := 0

	// Move all non-zero elements to the front
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// IsPalindrome checks if a string is a palindrome using two pointers.
// Only considers alphanumeric characters and ignores cases.
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// ReverseString reverses a string in-place using two pointers.
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// SortColors sorts an array with values 0, 1, 2 (Dutch National Flag problem).
// Time Complexity: O(n)
// Space Complexity: O(1)
func SortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

// LinkedListCycle detects if a linked list has a cycle using Floyd's algorithm.
// Fast pointer moves 2 steps, slow pointer moves 1 step.
// If they meet, there's a cycle.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle detects if a linked list has a cycle.
// Time Complexity: O(n)
// Space Complexity: O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// FindMiddle finds the middle node of a linked list.
// Time Complexity: O(n)
// Space Complexity: O(1)
func FindMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
