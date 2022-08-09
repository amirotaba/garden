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

type UserForm struct {
	ID       uint   `json:"id"`
	Type     uint   `json:"type"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active" gorm:"default:True"`
}

type UserType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessList  string `json:"access_list"`
}

type UserTypeForm struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessList  string `json:"access_list"`
}

type AccessForm struct {
	ID     uint   `json:"id"`
	TypeID string `json:"type_id"`
}

type Garden struct {
	gorm.Model
	Name        string  `json:"name"`
	Type        uint    `json:"type"`
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
	Name   string `json:"name"`
	Url    string `json:"url"`
	Method string `json:"method"`
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

type UserUseCase interface {
	SignUp(newUser *User) (int, error)
	SignIn(form *LoginForm) (UserResponse, int, error)
	Account(mp map[string]string) ([]UserResponse, int, error)
	UserAccount(mp map[string]string) (UserResponse, int, error)
	UpdateUser(user *UserForm, uid string) (int, error)
	DeleteUser(user *User, uid string) (int, error)

	CreateGarden(garden *Garden, uid string) (int, error)
	ReadGarden(mp map[string]string) ([]Garden, int, error)
	UpdateGarden(garden *GardenForm, uid string) (int, error)
	DeleteGarden(garden *Garden, uid string) (int, error)

	CreateUserType(usertype *UserType, uid string) (int, error)
	ReadUserType(id string, uid string) ([]UserType, int, error)
	UpdateAccess(access *AccessForm, uid string) (int, error)
	UpdateUserType(usertype *UserTypeForm, uid string) (int, error)
	DeleteUserType(usertype *UserType, uid string) (int, error)

	CreateTreeType(treeType *TreeType, uid string) (int, error)
	ReadTreeType(id string, uid string) ([]TreeType, int, error)
	UpdateTreeType(treeType *TreeTypeForm, uid string) (int, error)
	DeleteTreeType(treeType *TreeType, uid string) (int, error)

	CreateTag(tag *Tag, uid string) (int, error)
	ReadTag(pageNumber string, uid string) ([]Tag, int, error)
	ReadTagID(id string) ([]Tag, int, error)
	UpdateTag(tag *TagForm, uid string) (int, error)
	DeleteTag(tag *Tag, uid string) (int, error)

	CreateLocation(location *GardenLocation, uid string) (int, error)
	ReadLocation(id string, pageNumber string, uid string) ([]GardenLocation, int, error)
	UpdateLocation(loc *GardenLocationForm, uid string) (int, error)
	DeleteLocation(loc *GardenLocation, uid string) (int, error)

	CreateGardenType(gardenType *GardenType, uid string) (int, error)
	ReadGardenType(id string, uid string) ([]GardenType, int, error)
	UpdateGardenType(gardenType *GardenTypeForm, uid string) (int, error)
	DeleteGardenType(gardenType *GardenType, uid string) (int, error)

	CreateTree(tree *Tree, uid string) (int, error)
	ReadTree(mp map[string]string) ([]Tree, int, error)
	ReadTreeUser(mp map[string]string) ([]Tree, int, error)
	UpdateTree(tree *TreeForm, uid string) (int, error)
	DeleteTree(tree *Tree, uid string) (int, error)

	CreateComment(comment *Comment) (int, error)
	ReadComment(mp map[string]string, pageNumber, uid string) ([]Comment, int, error)
	UpdateComment(comment *CommentForm, uid string) (int, error)
	DeleteComment(comment *Comment, uid string) (int, error)

	CreateService(service *Service) (int, error)
	ReadService(uid string) ([]Service, int, error)
	UpdateService(usertype *ServiceForm, uid string) (int, error)
	DeleteService(service *Service, uid string) (int, error)
}

type UserRepository interface {
	CreateGarden(garden *Garden) error
	ReadGarden(n int) ([]Garden, error)
	ReadGardenID(u uint) ([]Garden, error)
	ReadGardenUID(id uint) ([]Garden, error)
	UpdateGarden(garden *GardenForm) error
	DeleteGarden(id uint) error

	SignUp(newUser *User) error
	SignIn(form *LoginForm) (User, error)
	Account(n int) ([]User, error)
	AccountUsername(username string) (User, error)
	AccountID(id uint) (User, error)
	AccountType(n int, tp uint) ([]User, error)
	UpdateUser(user *UserForm) error
	DeleteUser(id uint) error

	CreateTreeType(treeType *TreeType) error
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
	ReadTreeByType(t uint, n int) ([]Tree, error)
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

//1 , 2 , 3 , 4 , 5 , 6 , 7 , 8 , 9 , 10 , 11 , 12 , 13 , 14 , 15 , 16 , 17 , 18 , 19 , 20 , 21 , 22 , 23 , 24 , 25 , 26 , 27 , 28 , 29 , 30 , 31 , 32 , 33 , 34 , 35 , 36 , 37 , 38 , 39 , 40 , 41 , 42 , 43 , 44 , 45 , 46 , 47 , 48 , 49 , 50 , 51 , 52 , 53 , 54 , 55 , 56 , 57 , 58 , 59 , 60 , 61 , 62 , 63 , 64 , 65 , 66 , 67 , 68 , 69 , 70 , 71 , 72 , 73 , 74 , 75 , 76 , 77 , 78 , 79 , 80 , 81 , 82 , 83 , 84 , 85 , 86 , 87 , 88 , 89 , 90 , 91 , 92 , 93 , 94 , 95 , 96 , 97 , 98 , 99 , 100
