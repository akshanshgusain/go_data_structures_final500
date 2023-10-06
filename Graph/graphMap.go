package Graph

import "fmt"

// Graph represents an undirected graph using an adjacency list.
type GraphMap struct {
	adjacencyList map[int][]int
}

// NewGraph creates a new empty graph.
func NewGraph() *GraphMap {
	return &GraphMap{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge adds an edge between two vertices in the graph.
func (g *GraphMap) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u) // For an undirected graph
}

// PrintGraph prints the graph's adjacency list representation.
func (g *GraphMap) PrintGraph() {
	for vertex, neighbors := range g.adjacencyList {
		fmt.Printf("%d -> ", vertex)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	// Print the graph
	graph.PrintGraph()
}
