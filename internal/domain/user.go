package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Type       uint   `json:"type"`
	Name       string `json:"name"`
	UserName   string `json:"user_name"`
	PassWord   string `json:"pass_word"`
	Email      string `json:"email"`
	IsActive   bool   `json:"is_active"`
	AccessList []uint `json:"access_list"`
}

type UserForm struct {
	ID       uint   `json:"id"`
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

type UserTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Garden struct {
	gorm.Model
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	UserID      uint    `json:"user_id"`
	Description string  `json:"description"`
}

type GardenForm struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	UserID      uint    `json:"user_id"`
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
	UserId   uint    `json:"user_id"`
}

type GardenLocationForm struct {
	ID       uint    `json:"id"`
	Lat1     float64 `json:"lat_1"`
	Lat2     float64 `json:"lat_2"`
	Lat3     float64 `json:"lat_3"`
	Lat4     float64 `json:"lat_4"`
	Long1    float64 `json:"long_1"`
	Long2    float64 `json:"long_2"`
	Long3    float64 `json:"long_3"`
	Long4    float64 `json:"long_4"`
	GardenId uint    `json:"garden_id"`
	UserId   uint    `json:"user_id"`
}

type GardenType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GardenTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
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

type TreeForm struct {
	ID          uint      `json:"id"`
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

type TreeTypeForm struct {
	ID          uint   `json:"id"`
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

type CommentForm struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	TreeID uint   `json:"tree_id"`
	TagID  uint   `json:"tag_id"`
	Image  string `json:"image"`
	UserID uint   `json:"user_id"`
}

type Tag struct {
	gorm.Model
	History        string `json:"history"`
	Detoxification string `json:"detoxification"`
	Image          string `json:"image"`
}

type TagForm struct {
	ID             uint   `json:"id"`
	History        string `json:"history"`
	Detoxification string `json:"detoxification"`
	Image          string `json:"image"`
}

type Service struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ServiceForm struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LoginForm struct {
	Type     uint   `json:"type"`
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type TypeStruct struct {
	ID   uint
	Name string
}

type UserResponse struct {
	UserName string
	Type     TypeStruct
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

type UserUsecase interface {
	SignUp(newuser *User) (int, error)
	SignIn(form *LoginForm) (UserResponse, error)
	Account(mp map[string]string) ([]UserResponse, error)
	UpdateUser(user *UserForm, uid string) error
	DeleteUser(user *User, uid string) error

	CreateGarden(garden *Garden, uid string) error
	ReadGarden(mp map[string]string) ([]Garden, error)
	UpdateGarden(garden *GardenForm, uid string) error
	DeleteGarden(garden *Garden, uid string) error

	CreateUserType(usertype *UserType, uid string) error
	ReadUserType(id string, uid string) ([]UserType, error)
	UpdateUserType(usertype *UserTypeForm, uid string) error
	DeleteUserType(usertype *UserType, uid string) error

	CreateTreeType(treeType *TreeType, uid string) error
	ReadTreeType(id string, uid string) ([]TreeType, error)
	UpdateTreeType(treeType *TreeTypeForm, uid string) error
	DeleteTreeType(treetype *TreeType, uid string) error

	CreateTag(tag *Tag, uid string) error
	ReadTag(id string, pageNumber string, uid string) ([]Tag, error)
	UpdateTag(tag *TagForm, uid string) error
	DeleteTag(tag *Tag, uid string) error

	CreateLocation(location *GardenLocation, uid string) error
	ReadLocation(id string, pageNumber string, uid string) ([]GardenLocation, error)
	UpdateLocation(loc *GardenLocationForm, uid string) error
	DeleteLocation(loc *GardenLocation, uid string) error

	CreateGardenType(gardenType *GardenType, uid string) error
	ReadGardenType(id string, uid string) ([]GardenType, error)
	UpdateGardenType(gardenType *GardenTypeForm, uid string) error
	DeleteGardenType(gardenType *GardenType, uid string) error

	CreateTree(tree *Tree, uid string) error
	ReadTree(mp map[string]string, pageNumber, uid string) ([]Tree, error)
	UpdateTree(tree *TreeForm, uid string) error
	DeleteTree(tree *Tree, uid string) error

	CreateComment(comment *Comment, uid string) error
	ReadComment(mp map[string]string, pageNumber, uid string) ([]Comment, error)
	UpdateComment(comment *CommentForm, uid string) error
	DeleteComment(comment *Comment, uid string) error

	CreateService(service *Service, uid string) error
	ReadService(uid string) ([]Service, error)
	UpdateService(usertype *ServiceForm, uid string) error
	DeleteService(service *Service, uid string) error
	ReadTagID(id string, number string, uid string) ([]Tag, error)
}

type UserRepository interface {
	CreateGarden(garden *Garden) error
	ReadGarden(n int) ([]Garden, error)
	ReadGardenID(u uint) ([]Garden, error)
	ReadGardenUID(id uint) ([]Garden, error)
	UpdateGarden(garden *GardenForm) error
	DeleteGarden(id uint) error

	SignUp(newuser *User) error
	SignIn(form *LoginForm) (User, error)
	Account(n int) ([]User, error)
	AccountUser(username string) (User, error)
	AccountID(id uint) (User, error)
	AccountType(n int, tp uint) ([]User, error)
	UpdateUser(user *UserForm) error
	DeleteUser(id uint) error

	CreateTreeType(treetype *TreeType) error
	ReadTreeType() ([]TreeType, error)
	ReadTreeTypeID(u uint) ([]TreeType, error)
	UpdateTreeType(treeType *TreeTypeForm) error
	DeleteTreeType(id uint) error

	CreateTag(tag *Tag) error
	ReadTag(n int) ([]Tag, error)
	ReadTagID(u uint) ([]Tag, error)
	UpdateTag(tag *TagForm) error
	DeleteTag(id uint) error

	CreateLocation(location *GardenLocation) error
	ReadLocation(n int) ([]GardenLocation, error)
	ReadLocationID(u uint) ([]GardenLocation, error)
	UpdateLocation(loc *GardenLocationForm) error
	DeleteLocation(id uint) error

	CreateUserType(usertype *UserType) error
	ReadUserType() ([]UserType, error)
	ReadUserTypeID(id uint) ([]UserType, error)
	UpdateUserType(userType *UserTypeForm) error
	DeleteUserType(id uint) error

	CreateGardenType(gardenType *GardenType) error
	ReadGardenType() ([]GardenType, error)
	ReadGardenTypeID(u uint) ([]GardenType, error)
	UpdateGardenType(gardenType *GardenTypeForm) error
	DeleteGardenType(id uint) error

	CreateTree(tree *Tree) error
	ReadTree(n int) ([]Tree, error)
	ReadTreeID(id uint, q string) ([]Tree, error)
	ReadTreeByType(t string, n int) ([]Tree, error)
	UpdateTree(tree *TreeForm) error
	DeleteTree(id uint) error

	CreateComment(comment *Comment) error
	ReadComment(n int) ([]Comment, error)
	ReadCommentID(id uint, q string, span int) ([]Comment, error)
	UpdateComment(comment *CommentForm) error
	DeleteComment(id uint) error

	CreateService(service *Service) error
	ReadService() ([]Service, error)
	ReadServiceUrl(url string) (Service, error)
	UpdateService(service *ServiceForm) error
	DeleteService(id uint) error

	UserType(id uint) (string, error)
}
