package dataStructure

import (
	"errors"
	"fmt"
	"time"
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

func (u User) EmptyUser() bool {
	return u.username == ""
}

func GuestUser() User {
	return User{
		username: "Guest",
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

func (u *User) LogIn(password string) error {
	// var password string
	// fmt.Print("Password: ")
	// fmt.Scanln(&password)
	if password == u.password {
		u.loggedIn = true
		fmt.Println("Successfully logged in")
		return nil
	} else {
		return errors.New("incorrect password")
	}
}

func (u *User) LoggedIn() bool {
	return u.loggedIn
}

func (u *User) Username() string {
	return u.username
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
func (u *User) StaffCheck() bool {
	return u.staff
}
func (u *User) AddItem(bst *BST[Item], name string, x float64, y float64, s int) error {
	if !u.staff {
		return errors.New("only staff can add items")
	}
	_, err := bst.Search(name)
	if err == nil {
		return errors.New("item exists, please edit item")
	}
	item := Item{
		Name:  name,
		X:     x,
		Y:     y,
		Stock: s,
	}
	err = bst.Insert(item)
	time.Sleep(15 * time.Second)
	return err
}

func (u *User) RemoveItem(bst *BST[Item], name string) error {
	if !u.staff {
		return errors.New("only staff can remove items")
	}
	err := bst.Remove(name)
	time.Sleep(5 * time.Second)
	return err
}

func (u *User) UpdateItemStock(bst *BST[Item], name string, stock int) error {
	if !u.staff {
		return errors.New("only staff can update items stock")
	}
	// Not sure if bst function should be exported, but feels safer as such!
	binaryNode, err := bst.searchFunction(bst.root, name)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	binaryNode.data.Stock = stock
	time.Sleep(5 * time.Second)
	return err
}

func (u *User) UpdateItemLocation(bst *BST[Item], name string, x float64, y float64) error {
	if !u.staff {
		return errors.New("only staff can update item location")
	}
	// Not sure if bst function should be exported, but feels safer as such!
	// Members (not staff) can only access the Search function?
	binaryNode, err := bst.searchFunction(bst.root, name)
	if err != nil {
		fmt.Println("error:", err)
	}
	binaryNode.data.X = x
	binaryNode.data.Y = y
	time.Sleep(5 * time.Second)
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
	// time.Sleep(10 * time.Second)
	return 0.0, 0.0, err
}
