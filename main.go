package main

import (
	"fmt"

	"github.com/happy-frost/supermarketfinder/graph"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	g := graph.NewGraph()
	g.AddNode(1, 1)
	g.AddNode(1, 3)
	g.AddEdge(0, 1)
	g.PrintGraph()

	test := []int{1, 2, 3}
	for i := range test {
		fmt.Println(test[i])
	}

}
