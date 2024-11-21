package dataStructure

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseGraph() (*Graph, error) {
	// Panic is handled by the defer function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("file is format incorrectly, please ensure correct file is selected and the formatting within the file is correct")
		}
	}()

	g := NewGraph()

	fileLoc := os.Getenv("GRAPH_LOCATION")
	// fmt.Println(fileLoc)
	file, err := os.Open(fileLoc)
	if err != nil {
		return g, errors.New("error opening file:" + err.Error())
	}
	defer file.Close() // Ensure the file is closed when we're done

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Boolean to determine what is being read
	readingNodes := false
	readingEdges := false

	// Loop through each line in the file
	for scanner.Scan() {
		// Start reading nodes
		if readingNodes {
			// Stop reading nodes at indicator ---
			if scanner.Text() == "---" {
				readingNodes = false
			} else {
				// extract the x and y coordinate from the nodes
				// splits the string into the node number and the coordinates
				nodeStr := strings.Split(scanner.Text(), ":")

				// index in txt is actually ignored
				// splits the coordinates into the x and y coordinates
				coor := strings.Split(
					strings.Trim(nodeStr[1], " "),
					",")

				// check that number of items inside coordinate is correct
				if len(coor) == 2 {
					x, errx := strconv.ParseFloat(coor[0], 64)
					y, erry := strconv.ParseFloat(coor[1], 64)
					if errx != nil || erry != nil {
						return g, errors.New("file formatting incorrect, unable to get x or y value")
					}
					g.AddNode(x, y)
				} else {
					return g, errors.New("file is format incorrectly, please ensure correct file is selected and the formatting within the file is correct")
				}

				// fmt.Printf("%v %v\n", x, y)
			}
		} else if scanner.Text() == "Nodes:" {
			readingNodes = true
		}

		if readingEdges {
			// Stop reading nodes at indicator ---
			if scanner.Text() == "---" {
				readingEdges = false
			} else if scanner.Text() != "\n" {
				edgeStr := strings.Split(scanner.Text(), " ")
				if len(edgeStr) == 3 {
					v1, errv1 := strconv.Atoi(edgeStr[0])
					v2, errv2 := strconv.Atoi(edgeStr[2])
					if errv1 != nil || errv2 != nil {
						return g, errors.New("file is formatted incorrectly, please ensure edges are formatted correctly")
					}
					errEdge := g.AddEdge(v1, v2)
					if errEdge != nil {
						return g, errors.New("error adding edge:" + errEdge.Error())
					}
				}
				// fmt.Printf("%v\n", scanner.Text())
			}
		} else if scanner.Text() == "Edges:" {
			readingEdges = true
		}
	}

	// Check if there was an error while reading
	if err := scanner.Err(); err != nil {
		return g, errors.New("error reading file:" + err.Error())
	}
	return g, nil
}

// Can add save graph feature in the future for this use case not necessary
