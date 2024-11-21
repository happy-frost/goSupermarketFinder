package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/happy-frost/supermarketfinder/algorithm"
	"github.com/happy-frost/supermarketfinder/dataStructure"
)

func login(userBST *dataStructure.BST[dataStructure.User], ch chan dataStructure.User, input chan string) {
	// defer wg.Done()
	// var username string
	fmt.Print("Username: ")
	// fmt.Scanln(&username)
	// <-input // To get rid of the inital command "login"
	username := <-input
	userResult, err := userBST.Search(strings.Trim(username, "\n"))
	if err != nil {
		fmt.Println("No such user")
		ch <- dataStructure.User{}
		input <- "q"
		return
	} else {
		userNode, ok := userResult.(dataStructure.User)
		if ok {
			fmt.Print("Password: ")
			password := <-input
			userNode.LogIn(strings.Trim(password, "\n"))
			ch <- userNode
			input <- "q"
			return
		}
	}
}

func saveSession(loggedInUser *[]dataStructure.User, ch chan dataStructure.User) {
	// defer wg.Done()
	newUser := <-ch
	if !newUser.EmptyUser() && newUser.LoggedIn() {
		*loggedInUser = append(*loggedInUser, newUser)
		fmt.Println("Logged in to session")
	} else {
		fmt.Println("Failed to log in, please check username and password")
	}
}

func findItem(g *dataStructure.Graph, itemBST *dataStructure.BST[dataStructure.Item], loggedInUser *[]dataStructure.User, input chan string, mu *sync.RWMutex) {
	// defer wg.Done()
	// <-input // To get rid of the inital command "login"
	var user dataStructure.User
	if len(*loggedInUser) == 0 {
		user = dataStructure.GuestUser()
	} else {
		user = (*loggedInUser)[0]
	}
	// var item string
	mu.RLock()
	itemBST.InOrder()
	mu.RUnlock()
	fmt.Println("Please enter item name")
	// fmt.Scanln(&item)
	item := <-input
	mu.RLock()
	x2, y2, err := user.FindItemLocation(itemBST, strings.ToLower(strings.Trim(item, "\n")))
	mu.RUnlock()
	if err != nil {
		fmt.Println(err)
		input <- "q"
		return
	}
	fmt.Println("Please enter current location in the following format (x,y): ")
	fmt.Println("x value must be between 0 and 4, y value must be between 0 and 2")
	coorInput := <-input
	x1, y1, err := getXYCoordinates(coorInput)
	if err != nil {
		fmt.Println(err)
		input <- "q"
		return
	}
	directions, _, err := algorithm.ShortestPath(g, x1, y1, 0, x2, y2)
	if err != nil {
		fmt.Println(err)
		return
	}
	input <- "q"
	fmt.Printf("%s: Directions\n", user.Username())
	for !directions.IsEmpty() {
		message, _ := directions.Dequeue()
		fmt.Println(message)
	}
}

func staffActions(itemBST *dataStructure.BST[dataStructure.Item], loggedInUser *[]dataStructure.User, input chan string, mu *sync.RWMutex) {
	var staffUser dataStructure.User
	// find if any of the user are staffs!
	for _, user := range *loggedInUser {
		if user.LoggedIn() && user.StaffCheck() {
			staffUser = user
		}
	}
	if staffUser.EmptyUser() {
		fmt.Println("You need to be a staff to add and edit items.")
		fmt.Println("If you are a staff please log in first.")
		input <- "q"
		return
	}
	fmt.Println("What action would you like to take?")
	userInput := <-input
	if userInput == "add item\n" {
		fmt.Println("Please provide the item name:")
		item := <-input
		fmt.Println("Please provide the item location in this format (x,y):")
		fmt.Println("x value must be between 0 and 4, y value must be between 0 and 2")
		coor := <-input
		x, y, err := getXYCoordinates(coor)
		if err != nil {
			input <- "q"
			return
		}
		input <- "q"
		mu.Lock()
		err = staffUser.AddItem(itemBST, strings.Trim(item, "\n"), x, y, 1) // default stock to 1, since > 0 is considered available for now
		mu.Unlock()
		if err != nil {
			fmt.Printf("%s:%v\n", staffUser.Username(), err)
			return
		}
		fmt.Printf("%s: Item successfully added\n", staffUser.Username())
		mu.Lock()
		dataStructure.BSTToItemTXT(itemBST, "items.txt")
		mu.Unlock()
		fmt.Printf("%s: Item saved to memory\n", staffUser.Username())
	} else if userInput == "remove item\n" {
		fmt.Println("Please provide the item name:")
		item := <-input
		input <- "q"
		mu.Lock()
		err := staffUser.RemoveItem(itemBST, strings.Trim(item, "\n"))
		mu.Unlock()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s: Item successfully removed\n", staffUser.Username())
			mu.Lock()
			dataStructure.BSTToItemTXT(itemBST, "items.txt")
			mu.Unlock()
			fmt.Printf("%s: Item removed from memory\n", staffUser.Username())
		}
	} else if userInput == "update stock\n" {
		fmt.Println("Please provide the item name:")
		item := <-input
		fmt.Println("Please provide new stock amount:")
		amt := <-input
		stock, err := strconv.Atoi(strings.Trim(amt, "\n"))
		if err != nil {
			fmt.Println(err)
			input <- "q"
			return
		}
		input <- "q"
		mu.Lock()
		err = staffUser.UpdateItemStock(itemBST, strings.Trim(item, "\n"), stock)
		mu.Unlock()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s: Stock updated\n", staffUser.Username())
		mu.Lock()
		dataStructure.BSTToItemTXT(itemBST, "items.txt")
		mu.Unlock()
		fmt.Printf("%s: Stock saved to memory\n", staffUser.Username())
	} else if userInput == "update item location\n" {
		fmt.Println("Please provide the item name:")
		item := <-input
		fmt.Println("Please provide the item location in this format (x,y):")
		fmt.Println("x value must be between 0 and 4, y value must be between 0 and 2")
		coor := <-input
		x, y, err := getXYCoordinates(coor)
		input <- "q"
		if err != nil {
			fmt.Println(err)
			return
		}
		mu.Lock()
		staffUser.UpdateItemLocation(itemBST, strings.Trim(item, "\n"), x, y)
		mu.Unlock()
		fmt.Printf("%s: Item location updated\n", staffUser.Username())
		mu.Lock()
		dataStructure.BSTToItemTXT(itemBST, "items.txt")
		mu.Unlock()
		fmt.Printf("%s: Item location saved to memory\n", staffUser.Username())
	} else {
		input <- "q"
		return
	}
}

func getXYCoordinates(input string) (float64, float64, error) {
	var x1, y1 float64
	coor := strings.Split(strings.Trim(input, "\n"), ",") // trim off the \n before splitting
	if len(coor) != 2 {
		return x1, y1, errors.New("incorrect format of coordinates, please input coordinate in the following format (x,y)")
	} else {
		x1, errx := strconv.ParseFloat(coor[0], 64)
		y1, erry := strconv.ParseFloat(coor[1], 64)
		if errx != nil || erry != nil {
			return x1, y1, errors.New("please input a float for the x and y value")
		}
		if x1 < 0 || x1 > 4 {
			return x1, y1, errors.New("x value out of range")
		}
		if y1 < 0 || y1 > 2 {
			return x1, y1, errors.New("y value out of range")
		}
		return x1, y1, nil
	}
}
