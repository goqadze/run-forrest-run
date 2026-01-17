package algorithms

// BFS (Breadth-First Search) explores all neighbors at the current depth
// before moving to nodes at the next depth level. It uses a queue.

// BFSTraversal performs BFS starting from the given vertex.
// Returns the order of visited vertices.
// Time Complexity: O(V + E) where V is vertices and E is edges
// Space Complexity: O(V) for the visited map and queue
func (g *Graph) BFSTraversal(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		// Dequeue
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// Visit all neighbors
		for _, neighbor := range g.AdjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// BFSShortestPath finds the shortest path between source and destination
// in an unweighted graph. Returns the path and its length.
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) BFSShortestPath(source, destination int) ([]int, int) {
	if source == destination {
		return []int{source}, 0
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := []int{source}
	visited[source] = true
	parent[source] = -1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.AdjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				queue = append(queue, neighbor)

				if neighbor == destination {
					// Reconstruct path
					path := []int{}
					current := destination
					for current != -1 {
						path = append([]int{current}, path...)
						current = parent[current]
					}
					return path, len(path) - 1
				}
			}
		}
	}

	return nil, -1 // No path found
}

// BFSLevelOrder returns nodes grouped by their level/distance from start.
// Useful for level-order traversal of trees.
// Time Complexity: O(V + E)
func (g *Graph) BFSLevelOrder(start int) [][]int {
	visited := make(map[int]bool)
	result := [][]int{}
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node)

			for _, neighbor := range g.AdjList[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// BFSMatrix performs BFS on a 2D grid starting from (startRow, startCol).
// Returns the minimum distance to reach (endRow, endCol), or -1 if unreachable.
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func BFSMatrix(grid [][]int, startRow, startCol, endRow, endCol int) int {
	if len(grid) == 0 || grid[startRow][startCol] == 0 || grid[endRow][endCol] == 0 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Queue holds {row, col, distance}
	type cell struct {
		row, col, dist int
	}
	queue := []cell{{startRow, startCol, 0}}
	visited[startRow][startCol] = true

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.row == endRow && current.col == endCol {
			return current.dist
		}

		for _, dir := range directions {
			newRow, newCol := current.row+dir[0], current.col+dir[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
				!visited[newRow][newCol] && grid[newRow][newCol] == 1 {
				visited[newRow][newCol] = true
				queue = append(queue, cell{newRow, newCol, current.dist + 1})
			}
		}
	}

	return -1 // No path found
}

// HasPathBFS checks if there's a path between source and destination using BFS.
// Time Complexity: O(V + E)
func (g *Graph) HasPathBFS(source, destination int) bool {
	if source == destination {
		return true
	}

	visited := make(map[int]bool)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.AdjList[node] {
			if neighbor == destination {
				return true
			}
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// IsBipartite checks if a graph is bipartite using BFS.
// A bipartite graph can be colored with two colors such that
// no two adjacent vertices have the same color.
// Time Complexity: O(V + E)
func (g *Graph) IsBipartite() bool {
	color := make(map[int]int) // 0: uncolored, 1: color1, 2: color2

	for start := 0; start < g.Vertices; start++ {
		if color[start] != 0 {
			continue
		}

		queue := []int{start}
		color[start] = 1

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range g.AdjList[node] {
				if color[neighbor] == 0 {
					// Color with opposite color
					if color[node] == 1 {
						color[neighbor] = 2
					} else {
						color[neighbor] = 1
					}
					queue = append(queue, neighbor)
				} else if color[neighbor] == color[node] {
					// Same color as parent - not bipartite
					return false
				}
			}
		}
	}

	return true
}
