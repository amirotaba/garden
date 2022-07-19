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

type Farmer struct {
	gorm.Model
	UserName  string `json:"user_name"`
	PassWord  string `json:"pass_word"`
	Trees     string `json:"trees"`
	GardenID  int    `json:"garden_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}

type Garden struct {
	gorm.Model
	Trees     string `json:"trees"`
	Location  string `json:"location"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
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
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type UserResponse struct {
	UserName string
	Email    string
	Token    string
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
	USignIn(username, password string) (UserResponse, error)
	Account(username string) (UserResponse, error)
	Comment(id int, treeID int, text string) error
}

type AdminUsecase interface {
	ASignIn(username, password string) (UserResponse, error)
	ShowGarden() ([]Garden, error)
	RemoveGarden(id string, u string) error
	AddGarden(gar *Garden) error
	AddFarmer(far *Farmer) error
}

type FarmerUsecase interface {
	FSignIn(username, password string) (UserResponse, error)
	ShowTrees(id string) ([]Tree, error)
	ShowComments(farmerid, id int) (string, error)
	AddTree(tree *Tree) error
	RemoveTree(farmerid, treeid int) error
	AddAttend(form *AttendForm) error
}

type UserRepository interface {
	SignUp(newuser *User) error
	SignInUser(username, password string) (User, error)
	Account(username string) (User, error)
	Comment(tree Tree) error
	SearchTree(id int) (Tree, error)
}

type AdminRepository interface {
	SignInAdmin(username, password string) (Admin, error)
	ShowGarden() ([]Garden, error)
	DeletedBy(id int, u string) error
	RemoveGarden(id int) error
	AddGarden(garden *Garden) error
	AddFarmer(farmer *Farmer) error
}

type FarmerRepository interface {
	SignInFarmer(username, password string) (Farmer, error)
	ShowTrees(id int) ([]Tree, error)
	AddTree(tree *Tree) error
	RemoveTree(id int) error
	AddAttend(tree Tree) error
	UpdateFarmer(id int, trees string) error
	SearchFarmer(id int) (Farmer, error)
	SearchGarden(id int) (Garden, error)
	LastTree() (Tree, error)
	UpdateGarden(id int, tree string) error
	SearchTree(id int) (Tree, error)
}
