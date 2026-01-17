package algorithms

// DFS (Depth-First Search) explores as far as possible along each branch
// before backtracking. It uses a stack (or recursion) to track nodes.

// Graph represents an adjacency list graph
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new graph with the given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph (undirected)
func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

// AddDirectedEdge adds a directed edge from u to v
func (g *Graph) AddDirectedEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
}

// DFSRecursive performs DFS starting from the given vertex using recursion.
// Returns the order of visited vertices.
// Time Complexity: O(V + E) where V is vertices and E is edges
// Space Complexity: O(V) for the visited map and recursion stack
func (g *Graph) DFSRecursive(start int) []int {
	visited := make(map[int]bool)
	result := []int{}

	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true
		result = append(result, node)

		for _, neighbor := range g.AdjList[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return result
}

// DFSIterative performs DFS using an explicit stack instead of recursion.
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) DFSIterative(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	stack := []int{start}

	for len(stack) > 0 {
		// Pop from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		visited[node] = true
		result = append(result, node)

		// Add neighbors to stack (in reverse order to maintain order)
		neighbors := g.AdjList[node]
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !visited[neighbors[i]] {
				stack = append(stack, neighbors[i])
			}
		}
	}

	return result
}

// DFSMatrix performs DFS on a 2D grid/matrix.
// Useful for problems like island counting, maze solving, etc.
// Directions: up, down, left, right
func DFSMatrix(grid [][]int, row, col int, visited [][]bool) {
	// Check bounds and if already visited or is water (0)
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}
	if visited[row][col] || grid[row][col] == 0 {
		return
	}

	visited[row][col] = true

	// Explore all 4 directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		DFSMatrix(grid, row+dir[0], col+dir[1], visited)
	}
}

// CountIslands counts the number of islands in a 2D grid.
// An island is surrounded by water (0) and is formed by connecting
// adjacent lands (1) horizontally or vertically.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func CountIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	islands := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				DFSMatrix(grid, i, j, visited)
				islands++
			}
		}
	}

	return islands
}

// HasPathDFS checks if there's a path between source and destination.
// Time Complexity: O(V + E)
func (g *Graph) HasPathDFS(source, destination int) bool {
	visited := make(map[int]bool)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}

		visited[node] = true

		for _, neighbor := range g.AdjList[node] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}

		return false
	}

	return dfs(source)
}

// DetectCycleDirected detects if a directed graph has a cycle using DFS.
// Uses three states: unvisited (0), visiting (1), visited (2)
// Time Complexity: O(V + E)
func (g *Graph) DetectCycleDirected() bool {
	state := make(map[int]int) // 0: unvisited, 1: visiting, 2: visited

	var hasCycle func(node int) bool
	hasCycle = func(node int) bool {
		state[node] = 1 // Mark as visiting

		for _, neighbor := range g.AdjList[node] {
			if state[neighbor] == 1 {
				// Found a back edge (cycle)
				return true
			}
			if state[neighbor] == 0 {
				if hasCycle(neighbor) {
					return true
				}
			}
		}

		state[node] = 2 // Mark as visited
		return false
	}

	for i := 0; i < g.Vertices; i++ {
		if state[i] == 0 {
			if hasCycle(i) {
				return true
			}
		}
	}

	return false
}
