package algorithms

import "strconv"

// Stack implements a LIFO (Last-In-First-Out) data structure.
// Uses a slice for dynamic sizing.
type Stack struct {
	items []int
}

// NewStack creates and returns a new empty stack.
func NewStack() *Stack {
	return &Stack{items: []int{}}
}

// Push adds an element to the top of the stack.
// Time Complexity: O(1) amortized
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop removes and returns the top element from the stack.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(1)
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top element without removing it.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(1)
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack has no elements.
// Time Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
// Time Complexity: O(1)
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack.
// Time Complexity: O(1)
func (s *Stack) Clear() {
	s.items = []int{}
}

// ----------------------------------------------------------------------------
// MinStack - Stack that supports retrieving the minimum element in O(1)
// ----------------------------------------------------------------------------

// MinStack is a stack that supports push, pop, top, and retrieving
// the minimum element in constant time.
type MinStack struct {
	stack    []int // Main stack
	minStack []int // Auxiliary stack to track minimums
}

// NewMinStack creates and returns a new MinStack.
func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push adds an element to the stack.
// Time Complexity: O(1)
func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	// Push to minStack if it's empty or val is <= current minimum
	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

// Pop removes the top element from the stack.
// Time Complexity: O(1)
func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	// If popped element is the current minimum, pop from minStack too
	if top == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

// Top returns the top element.
// Time Complexity: O(1)
func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

// GetMin returns the minimum element in the stack.
// Time Complexity: O(1)
func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return 0
	}
	return ms.minStack[len(ms.minStack)-1]
}

// ----------------------------------------------------------------------------
// Stack-based Algorithm Examples
// ----------------------------------------------------------------------------

// IsValidParentheses checks if a string of brackets is valid.
// Valid means every opening bracket has a matching closing bracket
// in the correct order.
// Time Complexity: O(n)
// Space Complexity: O(n)
func IsValidParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// EvaluateRPN evaluates an expression in Reverse Polish Notation.
// Tokens are either integers or operators (+, -, *, /).
// Time Complexity: O(n)
// Space Complexity: O(n)
func EvaluateRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

// DailyTemperatures returns an array where answer[i] is the number of days
// you have to wait after day i to get a warmer temperature.
// Uses a monotonic decreasing stack.
// Time Complexity: O(n)
// Space Complexity: O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{} // Stack of indices

	for i := 0; i < n; i++ {
		// While current temperature is warmer than temperature at stack top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return answer
}

// NextGreaterElement finds the next greater element for each element.
// For each element, find the first greater element to its right.
// Returns -1 if no greater element exists.
// Time Complexity: O(n)
// Space Complexity: O(n)
func NextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}

	stack := []int{} // Stack of indices

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = nums[i]
		}
		stack = append(stack, i)
	}

	return result
}

// LargestRectangleInHistogram finds the largest rectangular area in a histogram.
// Time Complexity: O(n)
// Space Complexity: O(n)
func LargestRectangleInHistogram(heights []int) int {
	stack := []int{} // Stack of indices
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {
		// Use 0 as the height for the imaginary bar at the end
		h := 0
		if i < n {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}
