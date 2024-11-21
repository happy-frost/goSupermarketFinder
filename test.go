package main

import (
	"fmt"

	"github.com/happy-frost/supermarketfinder/algorithm"
	"github.com/happy-frost/supermarketfinder/dataStructure"
	"github.com/joho/godotenv"
)

func try() {
	// Calls to Graph

	// g := dataStructure.NewGraph()
	// g.AddNode(1, 1)
	// g.AddNode(1, 3)
	// g.AddEdge(0, 1)
	// g.PrintGraph()

	// test := []int{1, 2, 3}
	// for i := range test {
	// 	fmt.Println(test[i])
	// }

	godotenv.Load()
	g, errParse := dataStructure.ParseGraph()
	fmt.Println(errParse)
	g.PrintGraph()

	// v1, v2, w1, w2, err := algorithm.NearestVertex(g, 0, 0)
	// fmt.Println(err, v1, v2, w1, w2)

	fmt.Println("Shortest path search")
	q, a, err := algorithm.ShortestPath(g, 0, 0.1, 0, 0, 0.2)
	if err != nil {
		fmt.Println(err)
	}
	for !q.IsEmpty() {
		message, _ := q.Dequeue()
		fmt.Println(message)
	}
	fmt.Println(a)

	// Calls to BST
	// bst, err := dataStructure.TxtToItemBST("items.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// bst.Remove("Shampoo")
	// bst.InOrder()
	// bst.Insert(dataStructure.Item{
	// 	Name:  "Milk",
	// 	X:     1.2,
	// 	Y:     1.3,
	// 	Stock: 1})
	// bst.Insert("Cheese", 1.5, 1.3, 1)
	// bst.Insert("Egg", 1.2, 2.7, 1)
	// bst.Insert("Shampoo", 1.2, 8.3, 1)
	// bst.InOrder()

	// dataStructure.BSTToItemTXT(bst, "test.txt")
	// userBST, err := dataStructure.TxtToUserBST("user.txt")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// userBST.InOrder()
	// userResult, _ := userBST.Search("George")
	// userNode, ok := userResult.(dataStructure.User)
	// if ok {
	// 	err = userNode.LogIn()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// userNode.UpdateItemStock(bst, "Milk", 0)
	// bst.InOrder()

	// // Guest user
	// guest := dataStructure.GuestUser()
	// x, y, err := guest.FindItemLocation(bst, "Egg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(x, y)

}
