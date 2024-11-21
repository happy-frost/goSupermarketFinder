package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/happy-frost/supermarketfinder/dataStructure"
	"github.com/joho/godotenv"
)

var (
	mu sync.RWMutex
	// wg sync.WaitGroup
)

func main() {
	// Start by parsing graph, user and item
	godotenv.Load()
	g, err := dataStructure.ParseGraph()
	if err != nil {
		fmt.Println(err)
	}
	itemBST, err := dataStructure.TxtToItemBST("./items.txt")
	if err != nil {
		fmt.Println(err)
	}
	userBST, err := dataStructure.TxtToUserBST("./user.txt")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)

	userChannel := make(chan dataStructure.User)
	inputChannel := make(chan string, 10)

	loggedInUser := make([]dataStructure.User, 0, 3)

	inputOngoing := false

	fmt.Println("Please choose an action: login, find item")
	// user input
	for {
		userInput, _ := reader.ReadString('\n')
		// fmt.Println(userInput)
		select {
		case reply := <-inputChannel:
			// fmt.Println(reply)
			if reply == "q" {
				fmt.Println("Please choose an action: login, find item, staff action")
				inputOngoing = false
			}
		default:
			if inputOngoing {
				inputChannel <- userInput
			} else {
				if userInput == "login\n" {
					inputOngoing = true
					go login(userBST, userChannel, inputChannel)
					go saveSession(&loggedInUser, userChannel)
				} else if userInput == "find item\n" {
					inputOngoing = true
					go findItem(g, itemBST, &loggedInUser, inputChannel, &mu)
				} else if userInput == "staff action\n" {
					inputOngoing = true
					go staffActions(itemBST, &loggedInUser, inputChannel, &mu)
				} else {
					fmt.Println("Invalid action")
					fmt.Println("Please choose an action: login, find item, staff action")
				}
			}
		}
	}

	// single line input

}
