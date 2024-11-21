package algorithm

import (
	"errors"
	"math"

	"github.com/happy-frost/supermarketfinder/dataStructure"
)

func NearestVertex(g *dataStructure.Graph, x float64, y float64) (int, int, float64, float64, error) {
	// takes a graph with vertex, as well as an input of x and y coordinate and find the nearest vertex
	// and the two vertex connected to the edge the point lies on
	// V1 is the one that the vertex is closest to and V2 is the other vertex

	// check that graph is not empty
	if len(g.Nodes) == 0 {
		return 0, 0, 0.0, 0.0, errors.New("graph is empty")
	}

	// brute force algorithm to find closest node
	min := 1000.0
	minID := 0

	for i := range g.Nodes {
		nx, ny, _ := g.GetNodeFromId(i)
		dist := calculateDistance(x, y, nx, ny)
		if dist < min {
			min = dist
			minID = i
		}
	}
	v1, v2, w1, w2, err := checkPointOnEdge(g, minID, x, y)
	return v1, v2, w1, w2, err
}

func calculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
}

func checkPointOnEdge(g *dataStructure.Graph, v int, x float64, y float64) (int, int, float64, float64, error) {
	nx1, ny1, _ := g.GetNodeFromId(v)
	neighbours, _ := g.GetNeighbours(v)
	for _, v2 := range neighbours {
		nx2, ny2, _ := g.GetNodeFromId(v2)
		valid, w1, w2 := pointOnLine(nx1, ny1, nx2, ny2, x, y)
		if valid {
			return v, v2, w1, w2, nil
		}
	}
	return -1, -1, 0.0, 0.0, errors.New("point does not lie on any edge connected to this node")
}

func pointOnLine(nx1 float64, ny1 float64, nx2 float64, ny2 float64, x1 float64, y1 float64) (bool, float64, float64) {
	// function takes 3 points, and determine if the last point lie on the line between the first two point (collinear but more specific)
	// It returns a boolean indicating if it fulfil the above condition, and two floats to indicate the distance from each point

	// If it is a horizontal line
	if ny1 == ny2 {
		small := min(nx1, nx2)
		big := max(nx1, nx2)
		if y1 == ny1 && x1 <= big && x1 >= small {
			w1 := math.Abs(nx1 - x1)
			w2 := math.Abs(nx2 - x1)
			return true, w1, w2
		}
	}
	// If it is a vertical line
	if nx1 == nx2 {
		small := min(ny1, ny2)
		big := max(ny1, ny2)
		if x1 == nx1 && y1 <= big && y1 >= small {
			w1 := math.Abs(ny1 - y1)
			w2 := math.Abs(ny2 - y1)
			return true, w1, w2
		}
	}
	// y = mx + c
	m := (ny1 - ny2) / (nx1 - nx2)
	c := ny2 - m*nx2
	small := min(nx1, nx2)
	big := max(nx1, nx2)
	w1 := calculateDistance(nx1, ny1, x1, y1)
	w2 := calculateDistance(nx2, ny2, x1, y1)
	// check it lies on the line and within the two vertex (i.e. is not an extrapolated point)
	return (y1 == m*x1+c && x1 <= big && x1 >= small), w1, w2
}
