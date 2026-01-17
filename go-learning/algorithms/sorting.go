package algorithms

// Sorting algorithms implementation

// QuickSort sorts a slice using the quicksort algorithm.
// Time Complexity: O(n log n) average, O(n²) worst case
// Space Complexity: O(log n) due to recursion stack
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose the last element as pivot
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort sorts a slice using the merge sort algorithm.
// Time Complexity: O(n log n) - always
// Space Complexity: O(n) for the temporary arrays
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// HeapSort sorts a slice using the heap sort algorithm.
// Time Complexity: O(n log n)
// Space Complexity: O(1) - in-place sorting
func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}

	return result
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// BubbleSort sorts a slice using the bubble sort algorithm.
// Time Complexity: O(n²)
// Space Complexity: O(1)
// Note: Simple but inefficient - mainly for educational purposes
func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		// Optimization: if no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}

	return result
}

// InsertionSort sorts a slice using the insertion sort algorithm.
// Time Complexity: O(n²) worst, O(n) best (nearly sorted)
// Space Complexity: O(1)
// Note: Efficient for small datasets or nearly sorted arrays
func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// SelectionSort sorts a slice using the selection sort algorithm.
// Time Complexity: O(n²)
// Space Complexity: O(1)
func SelectionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		result[i], result[minIdx] = result[minIdx], result[i]
	}

	return result
}

// CountingSort sorts a slice of non-negative integers.
// Time Complexity: O(n + k) where k is the range of input
// Space Complexity: O(k)
// Note: Efficient when the range of input is not significantly greater than n
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum element
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Create counting array
	count := make([]int, maxVal+1)
	for _, v := range arr {
		count[v]++
	}

	// Build sorted array
	result := make([]int, 0, len(arr))
	for i, c := range count {
		for j := 0; j < c; j++ {
			result = append(result, i)
		}
	}

	return result
}
