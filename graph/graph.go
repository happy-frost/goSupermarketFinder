package graph

import (
	"errors"
	"fmt"
	"math"
)

type Graph struct {
	Nodes   []node
	adjList map[node][]edge
	// technically can just use two slice, but this is easier to find and implement
}

type node struct {
	id int
	x  float64
	y  float64
}

type edge struct {
	v1       node
	v2       node
	distance float64
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[node][]edge),
	}
}

func (g *Graph) addEdge(v1 node, v2 node) error {
	// check if the node already exists
	edgeList1, exists1 := g.adjList[v1]
	edgeList2, exists2 := g.adjList[v2]
	if exists1 && exists2 { // If both node exists add path
		// calculate the weight of the edge
		distance := math.Sqrt(math.Pow((v1.x-v2.x), 2) + math.Pow((v1.y-v2.y), 2))
		// create edge
		path1 := edge{v1, v2, distance}
		path2 := edge{v2, v1, distance}
		// append edge to both vertex
		edgeList1 = append(edgeList1, path1)
		edgeList2 = append(edgeList2, path2)
		// update the value in the map
		g.adjList[v1] = edgeList1
		g.adjList[v2] = edgeList2
		return nil
	} else {
		return errors.New("node does not exists yet")
	}
}

func (g *Graph) AddEdge(id1 int, id2 int) error {
	// Take the ID of the two nodes that the edge links
	v1 := g.Nodes[id1]
	v2 := g.Nodes[id2]
	err := g.addEdge(v1, v2)
	return err
}

func (g *Graph) addNode(v node) error {
	// create node if it does not already exists
	if _, exists := g.adjList[v]; !exists {
		g.Nodes = append(g.Nodes, v)
		g.adjList[v] = make([]edge, 0, 3) // Initialize to 3 as it most node will have 3 edges
		return nil
	} else {
		return errors.New("node already exists")
	}
}

func (g *Graph) AddNode(x float64, y float64) (int, error) {
	// Takes the x and y location and creates a node at the location

	// It should not not create node if another node is already at the x, y location.
	// Not implementing, as this implementation would then takes O(V) time to insert a node
	// Since it is not entirely necessary, since the nodes are only made when first initializing the map,
	// so as long as no duplicate exists in that file, the case of this happening is unlikely
	// for i := range g.Nodes {
	// 	if g.Nodes[i].X == x && g.Nodes[i].Y == y {
	// 		return -1, errors.New("A node already exists at this location")
	// 	}
	// }

	index := len(g.Nodes)
	v := node{index, x, y}
	err := g.addNode(v)
	return index, err // returns index and any error
}

func (g *Graph) GetNodeFromId(id int) (x float64, y float64, err error) {
	// From node id returns the x and y position of a node, given its index
	if id < len(g.Nodes) {
		v := g.Nodes[id]
		x = v.x
		y = v.y
		err = nil
	} else {
		err = errors.New("no such node exists yet")
	}
	return
}

// Can implement a heap sort function to search, to do later if have time.
// func (g *Graph) GetNodeFromLocation(x float64, y float64) (id int, err error) {
// 	// From x and y position of a node return node id if it exists, -1 and an error otherwise
// 	return
// }

// The GetNeighbours function returns the neighbouring nodea and the weight to get to that node
func (g *Graph) GetNeighbours(v1 node) map[node]float64 {
	edgeList := g.adjList[v1]
	output := make(map[node]float64)
	for i := 0; i < len(edgeList); i++ {
		edge := edgeList[i]
		output[edge.v2] = edge.distance
	}
	return output
}

func (g *Graph) PrintGraph() {
	for vertex, edgeList := range g.adjList {
		neighbours := make([]node, 0, len(edgeList))
		distance := make([]float64, 0, len(edgeList))
		for i := range edgeList {
			neighbours = append(neighbours, edgeList[i].v2)
			distance = append(distance, edgeList[i].distance)
		}
		fmt.Printf("%d: %v\n", vertex.id, neighbours)
	}
}
