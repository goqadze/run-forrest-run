package algorithms

// BinarySearch searches for target in a sorted slice and returns its index.
// Returns -1 if target is not found.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// Calculate mid to avoid integer overflow
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive is the recursive version of binary search.
// Time Complexity: O(log n)
// Space Complexity: O(log n) due to recursion stack
func BinarySearchRecursive(nums []int, target int) int {
	return binarySearchHelper(nums, target, 0, len(nums)-1)
}

func binarySearchHelper(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearchHelper(nums, target, mid+1, right)
	}
	return binarySearchHelper(nums, target, left, mid-1)
}

// BinarySearchFirstOccurrence finds the first occurrence of target in a sorted slice.
// Useful when there are duplicate elements.
// Time Complexity: O(log n)
func BinarySearchFirstOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			right = mid - 1 // Continue searching in the left half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// BinarySearchLastOccurrence finds the last occurrence of target in a sorted slice.
// Time Complexity: O(log n)
func BinarySearchLastOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			result = mid
			left = mid + 1 // Continue searching in the right half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// LowerBound finds the index of the first element >= target.
// Returns len(nums) if all elements are less than target.
// Time Complexity: O(log n)
func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// UpperBound finds the index of the first element > target.
// Returns len(nums) if all elements are <= target.
// Time Complexity: O(log n)
func UpperBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
