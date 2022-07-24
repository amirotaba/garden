package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Type     uint   `json:"type"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type UserType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Garden struct {
	gorm.Model
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	UserId      uint    `json:"user_id"`
	TreesCount  uint    `json:"trees_count"`
	Description string  `json:"description"`
	//
	Trees     string `json:"trees"`
	DeletedBy string `json:"deleted_by"`
}

type GardenLocation struct {
	gorm.Model
	Lat1     float64 `json:"lat_1"`
	Lat2     float64 `json:"lat_2"`
	Lat3     float64 `json:"lat_3"`
	Lat4     float64 `json:"lat_4"`
	Long1    float64 `json:"long_1"`
	Long2    float64 `json:"long_2"`
	Long3    float64 `json:"long_3"`
	Long4    float64 `json:"long_4"`
	GardenId uint    `json:"garden_id"`
	//
	UserId uint `json:"user_id"`
}

type GardenType struct {
	gorm.Model
	Name        uint   `json:"name"`
	Description string `json:"description"`
}

type Tree struct {
	gorm.Model
	FullName string `json:"full_name"`
	//Age      uint   `json:"age"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        uint      `json:"type"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Qr          string
	Length      float64 `json:"length"`
	Image       string  `json:"image"`
	GardenId    uint    `json:"garden_id"`
	Description string  `json:"description"`
	//
	Attend   string `json:"attend"`
	FarmerID uint
}

type TreeType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Comment struct {
	gorm.Model
	Text   string `json:"text"`
	TreeId uint   `json:"tree_id"`
	TagId  uint   `json:"tag_id"`
	Image  string `json:"image"`
	UserId uint   `json:"user_id"`
}

type Tag struct {
	gorm.Model
	History        string `json:"history"`
	Detoxification string `json:"detoxification"`
	Image          string `json:"image"`
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
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type UserUsecase interface {
	SignUp(newuser *User) error
	USignIn(username, password string) (UserResponse, error)
	Account(username string) (UserResponse, error)
	Comment(comment *Comment, user_id string) error
}

type AdminUsecase interface {
	ASignIn(username, password string) (UserResponse, error)
	ShowGarden() ([]Garden, error)
	RemoveGarden(id string, u string) error
	AddGarden(gar *Garden, user_id string) error
	AddFarmer(far *User, user_id string) error
	AddLocation(location *GardenLocation, user_id string) error
}

type FarmerUsecase interface {
	FSignIn(username, password string) (UserResponse, error)
	ShowTrees(id string) ([]Tree, error)
	ShowComments(farmerid, id string) ([]Comment, error)
	AddTree(tree *Tree, user_id string) error
	RemoveTree(farmerid, treeid string) error
	AddAttend(form *AttendForm) error
}

type UserRepository interface {
	SignUp(newuser *User) error
	SignInUser(username, password string) (User, error)
	Account(username string) (User, error)
	Comment(comment *Comment) error
	SearchTree(id uint) (Tree, error)
}

type AdminRepository interface {
	SignInAdmin(username, password string) (User, error)
	ShowGarden() ([]Garden, error)
	DeletedBy(id, u uint) error
	RemoveGarden(id uint) error
	AddGarden(garden *Garden) error
	AddFarmer(farmer *User) error
	AddLocation(location *GardenLocation) error
	RemoveGardenLocation(id uint) error
}

type FarmerRepository interface {
	SignInFarmer(username, password string) (User, error)
	ShowTrees(id uint) ([]Tree, error)
	AddTree(tree *Tree) error
	RemoveTree(id uint) error
	AddAttend(tree Tree) error
	UpdateFarmer(id uint, trees string) error
	SearchFarmer(id uint) (User, error)
	SearchGarden(id uint) (Garden, error)
	LastTree() (Tree, error)
	UpdateGarden(id uint, trees uint) error
	SearchTree(id uint) (Tree, error)
	SearchComment(tid uint) ([]Comment, error)
}
