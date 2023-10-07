package Graph

import (
	"fmt"
)

// GraphSlice Graph represents an undirected graph using a 2D slice.
type GraphSlice struct {
	vertices  int
	adjacency [][]int
}

// NewGraphSlice NewGraph creates a new empty graph with a given number of vertices.
func NewGraphSlice(vertices int) *GraphSlice {
	adjacency := make([][]int, vertices)
	for i := range adjacency {
		adjacency[i] = make([]int, vertices)
	}
	return &GraphSlice{
		vertices:  vertices,
		adjacency: adjacency,
	}
}

// AddEdge adds an edge between two vertices in the graph.
func (g *GraphSlice) AddEdge(u, v int) {
	if u >= 0 && u < g.vertices && v >= 0 && v < g.vertices {
		g.adjacency[u][v] = 1
		g.adjacency[v][u] = 1 // For an undirected graph
	}
}

// PrintGraphX prints the graph's adjacency matrix representation.
func (g *GraphSlice) PrintGraphX() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d ", g.adjacency[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Create a new graph with 5 vertices
	graph := NewGraphSlice(5)

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	// Print the graph
	graph.PrintGraphX()
}
