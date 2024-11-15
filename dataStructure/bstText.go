package dataStructure

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TxtToItemBST(fileLoc string) (*BST[Item], error) {
	// fileLoc := os.Getenv("ITEM_LOCATION")
	bst := NewBST[Item]()

	file, err := os.Open(fileLoc)
	if err != nil {
		return bst, errors.New("error opening file:" + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // To skip the first line
	for scanner.Scan() {
		// Ignore empty lines
		if scanner.Text() != "\n" && strings.Trim(scanner.Text(), " ") != "" {
			// Get the values (name, x, y, stock) from each lines
			vals := strings.Split(scanner.Text(), ",")
			// Check formatting is correct, split returns array with 4 items
			if len(vals) == 4 {
				name := vals[0]
				x, errx := strconv.ParseFloat(strings.Trim(vals[1], " "), 64)
				y, erry := strconv.ParseFloat(strings.Trim(vals[2], " "), 64)
				s, errs := strconv.Atoi(strings.Trim(vals[3], " "))
				// Check item in each position of the array is correct and can be parsed as float/int respectively
				if errx != nil || erry != nil || errs != nil {
					return bst, errors.New("data type does not match expected data type")
				}
				bst.Insert(Item{
					Name:  name,
					X:     x,
					Y:     y,
					Stock: s})
			} else {
				return bst, errors.New("item file formatting incorrect")
			}
		}
	}
	return bst, nil
}

func BSTToItemTXT(bst *BST[Item], fileLoc string) error {
	// Function takes bst and text file location and name as input to write the bst into that file
	// Preorder is used over in order as parsing a inorder one into the bst again will cause worst case scenario
	// for timecomplexity, unless a algorithm is in place to ensure the tree is balanced
	file, err := os.OpenFile(fileLoc, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "item, x, y, stock")
	slice := bst.PreOrderTraverseToSlice()
	for _, node := range *slice {
		item := node.data // No need to cast in this case as type of K is already defined in input!
		line := fmt.Sprintf("%v, %v, %v, %v", item.Name, item.X, item.Y, item.Stock)
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func TxtToUserBST(fileLoc string) (*BST[User], error) {
	// fileLoc := os.Getenv("ITEM_LOCATION")
	bst := NewBST[User]()

	file, err := os.Open(fileLoc)
	if err != nil {
		return bst, errors.New("error opening file:" + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	readingStaff := false
	readingMember := false

	for scanner.Scan() {
		if scanner.Text() == "" || scanner.Text() == "\n" {
			continue
		}
		if readingStaff {
			// Stop reading staff at indicator ---
			if scanner.Text() == "---" {
				readingStaff = false
			} else {
				vals := strings.Split(scanner.Text(), " ")
				if len(vals) == 2 {
					user := strings.Trim(vals[0], " ")
					pass := strings.Trim(vals[1], " ")
					bst.Insert(User{
						username: user,
						password: pass,
						staff:    true,
						loggedIn: false,
					})
				}
			}
		} else if scanner.Text() == "Staff:" {
			readingStaff = true
		}
		if readingMember {
			// Stop reading staff at indicator ---
			if scanner.Text() == "---" {
				readingMember = false
			} else {
				vals := strings.Split(scanner.Text(), " ")
				if len(vals) == 2 {
					user := strings.Trim(vals[0], " ")
					pass := strings.Trim(vals[1], " ")
					bst.Insert(User{
						username: user,
						password: pass,
						staff:    false,
						loggedIn: false,
					})
				}
			}
		} else if scanner.Text() == "Member:" {
			readingMember = true
		}
	}
	return bst, nil
}

// Functionality to add user not yet supported,
// func BSTToUserTXT(bst *BST[User], fileLoc string) error {
// 	return nil
// }
