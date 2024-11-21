package algorithm

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/happy-frost/supermarketfinder/dataStructure"
)

func ShortestPath(g *dataStructure.Graph, x1 float64, y1 float64, angle1 float64, x2 float64, y2 float64) (*dataStructure.Queue[string], float64, error) {
	// function takes input: graph, x1 and y1 is the starting point coordinate and angle is the angle user is starting off facing
	// x2 and y2 is the destination coordinate
	// angle is the angle from the x-axis

	directions := dataStructure.NewQueue[string]()
	// find the two node connected to the edge that the start point is on
	s1, s2, sw1, sw2, errstart := NearestVertex(g, x1, y1)
	// find the two node connected to the edge that the end point is on
	e1, e2, ew1, ew2, errend := NearestVertex(g, x2, y2)
	// path, distance, err := shortestPathBetweenNodes(g,v1,v2)
	if errstart != nil {
		return &directions, 0.0, errstart
	}
	if errend != nil {
		return &directions, 0.0, errstart
	}
	// If the two points are on the same edge
	// If there is a same node but they are on different edges then continue with rest of the program
	if s1 == e1 && s2 == e2 || s1 == e2 && s2 == e1 {
		message, currentAngle, err := instructionsInString(x1, y1, angle1, x2, y2, -1)
		if err != nil {
			return &directions, 0.0, err
		}
		directions.Enqueue(message)
		return &directions, currentAngle, nil
	}

	paths := make([]*dataStructure.Stack[int], 4)

	path1, distance1, err1 := shortestPathBetweenNodes(g, s1, e1)
	paths[0] = path1
	distance1 = distance1 + sw1 + ew1
	path2, distance2, err2 := shortestPathBetweenNodes(g, s1, e2)
	paths[1] = path2
	distance2 = distance2 + sw1 + ew2
	path3, distance3, err3 := shortestPathBetweenNodes(g, s2, e1)
	paths[2] = path3
	distance3 = distance3 + sw2 + ew1
	path4, distance4, err4 := shortestPathBetweenNodes(g, s2, e2)
	paths[3] = path4
	distance4 = distance4 + sw2 + ew2

	if err1 != nil {
		return &directions, 0.0, err1
	}
	if err2 != nil {
		return &directions, 0.0, err2
	}
	if err3 != nil {
		return &directions, 0.0, err3
	}
	if err4 != nil {
		return &directions, 0.0, err4
	}

	minimum := min(distance1, distance2, distance3, distance4)
	var shortestPathID int

	switch minimum {
	case distance1:
		shortestPathID = 0
	case distance2:
		shortestPathID = 1
	case distance3:
		shortestPathID = 2
	case distance4:
		shortestPathID = 3
	}

	var startNode int
	endNode, err := paths[shortestPathID].Pop()
	if err != nil {
		return &directions, 0.0, err
	}
	nx1, ny1, currentAngle := x1, y1, angle1
	nx2, ny2, err := g.GetNodeFromId(endNode)
	if err != nil {
		return &directions, 0.0, err
	}
	message, currentAngle, err := instructionsInString(nx1, ny1, currentAngle, nx2, ny2, endNode)
	if err != nil {
		return &directions, 0.0, err
	}
	if message != "" {
		directions.Enqueue(message)
	}
	// loop through to get all the nodes that needs to be passed
	for !paths[shortestPathID].IsEmpty() {
		startNode = endNode
		endNode, err = paths[shortestPathID].Pop()
		if err != nil {
			return &directions, 0.0, err
		}
		nx1, ny1, err = g.GetNodeFromId(startNode)
		if err != nil {
			return &directions, 0.0, err
		}
		nx2, ny2, err = g.GetNodeFromId(endNode)
		if err != nil {
			return &directions, 0.0, err
		}
		message, currentAngle, err = instructionsInString(nx1, ny1, currentAngle, nx2, ny2, endNode)
		if err != nil {
			return &directions, 0.0, err
		}

		directions.Enqueue(message)
	}

	// from last node to final destination
	nx1, ny1 = nx2, ny2
	nx2, ny2 = x2, y2
	message, angle, err := instructionsInString(nx1, ny1, currentAngle, nx2, ny2, -1)
	if err != nil {
		return &directions, angle, err
	}
	directions.Enqueue(message)

	return &directions, angle, nil
}

func shortestPathBetweenNodes(g *dataStructure.Graph, v1 int, v2 int) (*dataStructure.Stack[int], float64, error) {
	// output variables
	path := dataStructure.NewStack[int]()
	if v1 == v2 {
		path.Push(v1)
		return &path, 0.0, nil
	}

	// array of distance from v1
	minDist := slices.Repeat([]float64{math.Inf(1)}, len(g.Nodes))
	// previous node on shortest patharray
	prevArray := slices.Repeat([]int{-1}, len(g.Nodes))

	// priority queue for dijstar algorithm
	pq := dataStructure.NewPriorityQueue[int]()
	err := pq.Enqueue(v1, 0)
	minDist[v1] = 0
	if err != nil {
		return &path, 0.0, err
	}
	for !pq.IsEmpty() {
		v, err := pq.Dequeue()
		if err != nil {
			return &path, 0.0, err
		}
		neighbours, distance := g.GetNeighbours(v)
		// problem here keep enqueuing not checking
		for i := 0; i < len(neighbours); i++ {
			newDist := minDist[v] + distance[i]
			if newDist < minDist[neighbours[i]] {
				minDist[neighbours[i]] = newDist
				pq.Enqueue(neighbours[i], newDist)
				prevArray[neighbours[i]] = v
			}
		}
	}

	// build a stack of the path
	path.Push(v2)
	prev := prevArray[v2]
	for prev != v1 {
		path.Push(prev)
		prev = prevArray[prev]
	}
	path.Push(v1)

	return &path, minDist[v2], nil
}

func instructionsInString(x1 float64, y1 float64, angle1 float64, x2 float64, y2 float64, v2 int) (string, float64, error) {
	if angle1 > math.Pi || angle1 < -math.Pi {
		return "", 0.0, errors.New("invalid angle, angle should be between pi and -pi")
	}
	destination := strconv.Itoa(v2)
	if v2 == -1 {
		destination = "your item"
	}
	vector := [2]float64{(x2 - x1), (y2 - y1)}
	if vector[0] == 0 && vector[1] == 0 {
		return "", angle1, nil
	}
	angleDesired := math.Atan2(vector[1], vector[0])
	// fmt.Println(x1, y1, x2, y2)
	// fmt.Println("vector:", vector)
	// fmt.Println("desired:", angleDesired)
	// fmt.Println("current:", angle1)
	turningAmount := angleDesired - angle1
	distance := calculateDistance(x1, y1, x2, y2)
	if turningAmount == 0 {
	} else if turningAmount > math.Pi {
		return fmt.Sprintf("Turn to your right %.2f degrees\nHead straight for %.1fm until you reach %s", (((math.Pi * 2) - turningAmount) / math.Pi * 180), distance, destination), angleDesired, nil
	} else if turningAmount > 0 {
		return fmt.Sprintf("Turn to your left %.2f degrees\nHead straight for %.1fm until you reach %s", (turningAmount / math.Pi * 180), distance, destination), angleDesired, nil
	} else if turningAmount < -math.Pi {
		return fmt.Sprintf("Turn to your left %.2f degrees\nHead straight for %.1fm until you reach %s", (((math.Pi * 2) + turningAmount) / math.Pi * 180), distance, destination), angleDesired, nil
	} else {
		return fmt.Sprintf("Turn to your right %.2f degrees\nHead straight for %.1fm until you reach %s", (-turningAmount / math.Pi * 180), distance, destination), angleDesired, nil
	}
	return fmt.Sprintf("Head straight for %.2fm until you reach %s", distance, destination), angleDesired, nil
}
