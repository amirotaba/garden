package domain

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
}

type UserResponse struct {
	UserName string
	Email    string
	Token    string
}

type Farmer struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Trees    string `json:"trees"`
	GardenID int    `json:"garden_id"`
}

type Garden struct {
	gorm.Model
	Trees    string `json:"trees"`
	Location string `json:"location"`
}

type Tree struct {
	gorm.Model
	Name     string `json:"name"`
	FarmerID int    `json:"farmer_id"`
	Age      int    `json:"age"`
	Attend   string `json:"attend"`
	Comment  string `json:"comment"`
}

type CommentForm struct {
	ID     int    `json:"id"`
	TreeID int    `json:"tree_id"`
	Text   string `json:"text"`
}

type LoginForm struct {
	Type     string `json:"type"`
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type AuthMessage struct {
	Text     string
	UserInfo UserResponse
}

type SignUpMessage struct {
	Text     string
	UserName string
	Email    string
}

type AttendForm struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type UserUsecase interface {
	SignUp(newuser *User) error
	SignIn(tp string, password, username string) (UserResponse, error)
	Account(username string) (UserResponse, error)
	Comment(id int, treeID int, text string) error
	ShowGarden() ([]Garden, error)
	RemoveGarden(id int) error
	AddGarden(gar *Garden) error
	AddFarmer(far *Farmer) error
	ShowTrees(id string) ([]Tree, error)
	ShowComments(farmerid, id int) (string, error)
	AddTree(tree *Tree) error
	RemoveTree(farmerid, treeid int) error
	AddAttend(form *AttendForm) error
}

type UserRepository interface {
	SignUp(newuser *User) error
	SignInUser(password, username string) (User, error)
	Account(username string) (User, error)
	Comment(tree Tree) error
	SearchTree(id int) (Tree, error)
	SignInAdmin(password string, username string) (Admin, error)
	SignInFarmer(password string, username string) (Farmer, error)
	ShowGarden() ([]Garden, error)
	RemoveGarden(id int) error
	AddGarden(garden *Garden) error
	AddFarmer(farmer *Farmer) error
	ShowTrees(id int) ([]Tree, error)
	AddTree(tree *Tree) error
	RemoveTree(id int) error
	AddAttend(tree Tree) error
	UpdateFarmer(id int, trees string) error
	SearchFarmer(id int) (Farmer, error)
	SearchGarden(id int) (Garden, error)
	LastTree() (Tree, error)
	UpdateGarden(id int, tree string) error
}
