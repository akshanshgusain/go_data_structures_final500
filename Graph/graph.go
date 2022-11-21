package Graph

import "fmt"

// Graph , adjacency list
type Graph struct {
	vertices []*Vertex // slice of vertices
}

// Vertex Structure
type Vertex struct {
	key      int
	adjacent []*Vertex // slice of vertices
}

// AddVertex , add vertex tot he graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v can not be added because its an exsisting key", k)
		fmt.Printf(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

// AddEdge Add edge, Directed graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// Check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v - > %v)", from, to)
		fmt.Printf(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("invalid edge (%v - > %v)", from, to)
		fmt.Printf(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// get vertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Contains
func contains(s []*Vertex, key int) bool {
	for _, v := range s {
		if v.key == key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Println(v)
	}
}

func Execute() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1, 2)
	test.Print()
}
