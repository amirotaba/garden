package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Type     uint   `json:"type"`
	Name     string `json:"name"`
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
	Description string  `json:"description"`
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
	FullName    string    `json:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Type        uint      `json:"type"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Qr          string
	Length      float64 `json:"length"`
	Image       string  `json:"image"`
	GardenId    uint    `json:"garden_id"`
	Description string  `json:"description"`
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
	Type     uint   `json:"type"`
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type UserResponse struct {
	UserName string
	Type     string
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

type AdminUsecase interface {
	SignUp(newuser *User) error
	SignIn(form *LoginForm) (UserResponse, error)
	Account(string) ([]UserResponse, error)
	UpdateUser(user *User) error
	DeleteUser(user *User) error

	CreateGarden(garden *Garden) error
	ReadGarden(id string) ([]Garden, error)
	UpdateGarden(garden *Garden) error
	DeleteGarden(garden *Garden) error

	CreateUserType(usertype *UserType) error
	ReadUserType(id string) ([]UserType, error)
	UpdateUserType(usertype *UserType) error
	DeleteUserType(usertype *UserType) error

	CreateTreeType(treeType *TreeType) error
	ReadTreeType(id string) ([]TreeType, error)
	UpdateTreeType(treeType *TreeType) error
	DeleteTreeType(treetype *TreeType) error

	CreateTag(tag *Tag) error
	ReadTag(id string) ([]Tag, error)
	UpdateTag(tag *Tag) error
	DeleteTag(tag *Tag) error

	CreateLocation(location *GardenLocation) error
	ReadLocation(id string) ([]GardenLocation, error)
	UpdateLocation(loc *GardenLocation) error
	DeleteLocation(loc *GardenLocation) error

	CreateGardenType(gardenType *GardenType) error
	ReadGardenType(id string) ([]GardenType, error)
	UpdateGardenType(gardenType *GardenType) error
	DeleteGardenType(gardenType *GardenType) error

	CreateTree(tree *Tree) error
	ReadTree(mp map[string]string) ([]Tree, error)
	UpdateTree(tree *Tree) error
	DeleteTree(tree *Tree) error

	CreateComment(comment *Comment) error
	ReadComment(mp map[string]string) ([]Comment, error)
	UpdateComment(comment *Comment) error
	DeleteComment(comment *Comment) error
}

type AdminRepository interface {
	CreateGarden(garden *Garden) error
	ReadGarden() ([]Garden, error)
	ReadGardenID(u uint) ([]Garden, error)
	UpdateGarden(garden *Garden) error
	DeleteGarden(id uint) error

	SignUp(newuser *User) error
	SignIn(form *LoginForm) (User, error)
	Account() ([]User, error)
	AccountUser(username string) ([]User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error

	CreateTreeType(treetype *TreeType) error
	ReadTreeType() ([]TreeType, error)
	ReadTreeTypeID(u uint) ([]TreeType, error)
	UpdateTreeType(treeType *TreeType) error
	DeleteTreeType(id uint) error

	CreateTag(tag *Tag) error
	ReadTag() ([]Tag, error)
	ReadTagID(u uint) ([]Tag, error)
	UpdateTag(tag *Tag) error
	DeleteTag(id uint) error

	CreateLocation(location *GardenLocation) error
	ReadLocation() ([]GardenLocation, error)
	ReadLocationID(u uint) ([]GardenLocation, error)
	UpdateLocation(loc *GardenLocation) error
	DeleteLocation(id uint) error

	CreateUserType(usertype *UserType) error
	ReadUserType() ([]UserType, error)
	ReadUserTypeID(id uint) ([]UserType, error)
	UpdateUserType(userType *UserType) error
	DeleteUserType(id uint) error

	CreateGardenType(gardenType *GardenType) error
	ReadGardenType() ([]GardenType, error)
	ReadGardenTypeID(u uint) ([]GardenType, error)
	UpdateGardenType(gardenType *GardenType) error
	DeleteGardenType(id uint) error

	CreateTree(tree *Tree) error
	ReadTree() ([]Tree, error)
	ReadTreeID(id uint, q string) ([]Tree, error)
	ReadTreeByType(t string) ([]Tree, error)
	UpdateTree(tree *Tree) error
	DeleteTree(id uint) error

	CreateComment(comment *Comment) error
	ReadComment() ([]Comment, error)
	ReadCommentID(id uint, q string) ([]Comment, error)
	UpdateComment(comment *Comment) error
	DeleteComment(id uint) error

	UserType(t uint) (string, error)
}
