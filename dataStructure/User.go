package dataStructure

import (
	"errors"
	"fmt"
)

type User struct {
	username string
	password string
	staff    bool // either staff or member
	loggedIn bool
}

func (u User) name() string {
	return u.username
}

func GuestUser() User {
	return User{
		username: "",
		password: "",
		staff:    false,
		loggedIn: true,
	}
}

func (u *User) AddUser(ubst *BST[User], user string, pass string, isStaff bool) error {
	newUser := User{
		username: user,
		password: pass,
		staff:    isStaff,
		loggedIn: false,
	}
	if isStaff && !u.staff {
		return errors.New("user must be staff to add user as staff")
	}
	err := ubst.Insert(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) LogIn() error {
	var password string
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	if password == u.password {
		u.loggedIn = true
		fmt.Println("Successfully logged in")
		return nil
	} else {
		return errors.New("incorrect password")
	}
}

func (u *User) LogOut() error {
	if u.loggedIn {
		u.loggedIn = false
		fmt.Println("Successfully logged out")
		return nil
	} else {
		return errors.New("user not logged in, cannot log out")
	}
}
func (u *User) staffCheck() error {
	if !u.loggedIn {
		return errors.New("please log in to edit the items in the store")
	}
	if !u.staff {
		return errors.New("only staff can edit items")
	}
	return nil
}
func (u *User) AddItem(bst *BST[Item], name string, x float64, y float64, s int) error {
	err := u.staffCheck()
	if err != nil {
		return err
	}
	err = bst.Insert(Item{
		Name:  name,
		X:     x,
		Y:     y,
		Stock: s,
	})
	return err
}

func (u *User) RemoveItem(bst *BST[Item], name string) error {
	err := u.staffCheck()
	if err != nil {
		return err
	}
	err = bst.Remove(name)
	return err
}

func (u *User) UpdateItemStock(bst *BST[Item], name string, stock int) error {
	err := u.staffCheck()
	if err != nil {
		return err
	}
	// Not sure if bst function should be exported, but feels safer as such!
	binaryNode, err := bst.searchFunction(bst.root, name)
	if err != nil {
		fmt.Println("error:", err)
	}
	binaryNode.data.Stock = stock
	return err
}

func (u *User) UpdateItemLocation(bst *BST[Item], name string, x float64, y float64) error {
	err := u.staffCheck()
	if err != nil {
		return err
	}
	// Not sure if bst function should be exported, but feels safer as such!
	// Members (not staff) can only access the Search function?
	binaryNode, err := bst.searchFunction(bst.root, name)
	if err != nil {
		fmt.Println("error:", err)
	}
	binaryNode.data.X = x
	binaryNode.data.Y = y
	return err
}

func (u *User) FindItemLocation(bst *BST[Item], name string) (x float64, y float64, e error) {
	// Non-members can use this function as well, is there a necessity for members then or is it only staff login
	result, err := bst.Search(name)
	if err != nil {
		return 0.0, 0.0, err
	}
	item, ok := result.(Item)
	if ok {
		if item.Stock > 0 {
			return item.X, item.Y, nil
		} else {
			return 0.0, 0.0, errors.New("item out of stock")
		}
		// fmt.Printf("%v %v %v %v\n", item.Name, item.X, item.Y, item.Stock)
	}
	return 0.0, 0.0, err
}
