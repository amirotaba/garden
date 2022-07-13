package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
}

type UserResponse struct {
	gorm.Model
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string
}

type Farmer struct {
	gorm.Model
	UserName string
	PassWord string
	Trees []int
}

type Garden struct {
	Id int
	Trees []int
	Locations string
}

type Comment struct {
	UserID int
	TreeID int
}

type Tree struct {
	ID int
	FarmerID int
	Age int
	Attend []string
	Comment map[int]string
}

type CommentForm struct {
	ID int
	TreeID int
	Text string
}

type LoginForm struct {
	Type string
	Username string
	Password string
}

type UserUsecase interface {
	SignUp(newuser *User) error
	SignIn(password, username string) (UserResponse, error)
	//user
	Account(username string) (UserResponse, error)
	Comment(id int, treeID int, text string) error
	//admin
	showGarden() (Garden)
	RemoveGarden(id int) error
	AddGarden(gar Garden) error
	AddFarmer(far Farmer) error
	//farmer
	ShowTrees(id int) Tree
	ShowComment(id int) string
	AddTree(tr Tree) error
	RemoveTree(id int) error
	AddAttend(id int) error
}

type UserRepository interface {
	SignUp(newuser *User) error
	SignIn(password, username string) (User, error)
	Account(username string) (UserResponse, error)
	Comment(tree Tree) (error)
	SearchTree(id int) (Tree, error)
	ShowGarden
	RemoveGarden
	AddGarden
	ShowTrees
	ShowComment
	AddTree
	RemoveTree
	ShowTree
	AddAttend
}
