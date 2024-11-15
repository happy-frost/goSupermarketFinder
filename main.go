package main

import (
	"fmt"

	"github.com/happy-frost/supermarketfinder/dataStructure"
)

func main() {
	// Calls to Graph
	// godotenv.Load()

	// g := graph.NewGraph()
	// g.AddNode(1, 1)
	// g.AddNode(1, 3)
	// g.AddEdge(0, 1)
	// g.PrintGraph()

	// test := []int{1, 2, 3}
	// for i := range test {
	// 	fmt.Println(test[i])
	// }

	// g, errParse := graph.ParseGraph()
	// fmt.Println(errParse)
	// g.PrintGraph()

	// Calls to BST
	bst, err := dataStructure.TxtToItemBST("items.txt")
	if err != nil {
		fmt.Println(err)
	}
	// bst.Remove("Shampoo")
	bst.InOrder()
	// bst.Insert(dataStructure.Item{
	// 	Name:  "Milk",
	// 	X:     1.2,
	// 	Y:     1.3,
	// 	Stock: 1})
	// bst.Insert("Cheese", 1.5, 1.3, 1)
	// bst.Insert("Egg", 1.2, 2.7, 1)
	// bst.Insert("Shampoo", 1.2, 8.3, 1)
	// bst.InOrder()

	dataStructure.BSTToItemTXT(bst, "test.txt")
	userBST, err := dataStructure.TxtToUserBST("user.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	userBST.InOrder()
	userResult, _ := userBST.Search("George")
	userNode, ok := userResult.(dataStructure.User)
	if ok {
		err = userNode.LogIn()
		if err != nil {
			fmt.Println(err)
		}
	}
	userNode.UpdateItemStock(bst, "Milk", 0)
	bst.InOrder()

	// Guest user
	guest := dataStructure.GuestUser()
	x, y, err := guest.FindItemLocation(bst, "Egg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x, y)

}
