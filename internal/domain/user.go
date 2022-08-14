package domain

import (
	"time"

	"gorm.io/gorm"
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

type AccountForm struct {
	Uid        string
	Tp         string
	PageNumber string
}

type UserAccountForm struct {
	Uid      string
	Username string
	ID       string
}

type ReadTreeForm struct {
	Uid        string
	GardenID   string
	Tp         string
	PageNumber string
}

type ReadTreeUserForm struct {
	ID       string
	GardenID string
}

type ReadCommentForm struct {
	ID         string
	TreeID     string
	TagID      string
	UserID     string
	PageNumber string
	Uid        string
}

type ReadGardenForm struct {
	Uid        string
	UserID     string
	PageNumber string
	ID         string
}

type UserUseCase interface {
	Create(newUser *User) (int, error)
	SignIn(form *LoginForm) (UserResponse, int, error)
	Read(form AccountForm) ([]UserResponse, int, error)
	UserRead(form UserAccountForm) (UserResponse, int, error)
	Update(user *UserForm, uid string) (int, error)
	Delete(user *User, uid string) (int, error)

	CreateType(usertype *UserType, uid string) (int, error)
	ReadType(id string, uid string) ([]UserType, int, error)
	UpdateAccess(access *AccessForm, uid string) (int, error)
	UpdateType(usertype *UserTypeForm, uid string) (int, error)
	DeleteType(usertype *UserType, uid string) (int, error)
}

type TagUseCase interface {
	Create(tag *Tag, uid string) (int, error)
	Read(pageNumber string, uid string) ([]Tag, int, error)
	ReadID(id string) ([]Tag, int, error)
	Update(tag *TagForm, uid string) (int, error)
	Delete(tag *Tag, uid string) (int, error)
}

type GardenUseCase interface {
	Create(garden *Garden, uid string) (int, error)
	Read(form ReadGardenForm) ([]Garden, int, error)
	Update(garden *GardenForm, uid string) (int, error)
	Delete(garden *Garden, uid string) (int, error)

	CreateLocation(location *GardenLocation, uid string) (int, error)
	ReadLocation(id string, pageNumber string, uid string) ([]GardenLocation, int, error)
	UpdateLocation(loc *GardenLocationForm, uid string) (int, error)
	DeleteLocation(loc *GardenLocation, uid string) (int, error)

	CreateType(gardenType *GardenType, uid string) (int, error)
	ReadType(id string, uid string) ([]GardenType, int, error)
	UpdateType(gardenType *GardenTypeForm, uid string) (int, error)
	DeleteType(gardenType *GardenType, uid string) (int, error)
}

type TreeUseCase interface {
	Create(tree *Tree, uid string) (int, error)
	Read(form ReadTreeForm) ([]Tree, int, error)
	ReadUser(form ReadTreeUserForm) ([]Tree, int, error)
	Update(tree *TreeForm, uid string) (int, error)
	Delete(tree *Tree, uid string) (int, error)

	CreateType(treeType *TreeType, uid string) (int, error)
	ReadType(id string, uid string) ([]TreeType, int, error)
	UpdateType(treeType *TreeTypeForm, uid string) (int, error)
	DeleteType(treeType *TreeType, uid string) (int, error)
}

type CommentUseCase interface {
	Create(comment *Comment) (int, error)
	Read(form ReadCommentForm) ([]Comment, int, error)
	Update(comment *CommentForm, uid string) (int, error)
	Delete(comment *Comment, uid string) (int, error)
}

type ServiceUseCase interface {
	Create(service *Service) (int, error)
	Read(uid string) ([]Service, int, error)
	Update(usertype *ServiceForm, uid string) (int, error)
	Delete(service *Service, uid string) (int, error)
}

type UserRepository interface {
	Create(newUser *User) error
	SignIn(form *LoginForm) (User, error)
	Read(n int) ([]User, error)
	ReadUsername(username string) (User, error)
	ReadID(id uint) (User, error)
	ReadByType(n int, tp uint) ([]User, error)
	Update(user *UserForm) error
	Delete(id uint) error

	CreateType(usertype *UserType) error
	ReadType() ([]UserType, error)
	ReadTypeID(id uint) ([]UserType, error)
	UpdateType(userType *UserTypeForm) error
	DeleteType(id uint) error
	ReadTypeUser(id uint) (string, error)
}

type TagRepository interface {
	Create(tag *Tag) error
	Read(n int) ([]Tag, error)
	ReadID(u uint) ([]Tag, error)
	Update(tag *TagForm) error
	Delete(id uint) error
}

type GardenRepository interface {
	Create(garden *Garden) error
	Read(n int) ([]Garden, error)
	ReadID(u uint) ([]Garden, error)
	ReadUID(id uint) ([]Garden, error)
	Update(garden *GardenForm) error
	Delete(id uint) error

	CreateLocation(location *GardenLocation) error
	ReadLocation(n int) ([]GardenLocation, error)
	ReadLocationID(u uint) ([]GardenLocation, error)
	UpdateLocation(loc *GardenLocationForm) error
	DeleteLocation(id uint) error

	CreateType(gardenType *GardenType) error
	ReadType() ([]GardenType, error)
	ReadTypeID(u uint) ([]GardenType, error)
	UpdateType(gardenType *GardenTypeForm) error
	DeleteType(id uint) error
}

type TreeRepository interface {
	Create(tree *Tree) error
	Read(n int) ([]Tree, error)
	ReadID(id uint, q string) ([]Tree, error)
	ReadByType(t uint, n int) ([]Tree, error)
	Update(tree *TreeForm) error
	Delete(id uint) error

	CreateType(treeType *TreeType) error
	ReadType() ([]TreeType, error)
	ReadTypeID(u uint) ([]TreeType, error)
	UpdateType(treeType *TreeTypeForm) error
	DeleteType(id uint) error
}

type CommentRepository interface {
	Create(comment *Comment) error
	Read(n int) ([]Comment, error)
	ReadID(id uint, q string, span int) ([]Comment, error)
	Update(comment *CommentForm) error
	Delete(id uint) error
}

type ServiceRepository interface {
	Create(service *Service) error
	Read() ([]Service, error)
	ReadURL(url string) (Service, error)
	Update(service *ServiceForm) error
	Delete(id uint) error
}

type UseCases struct {
	User    UserUseCase
	Tag     TagUseCase
	Garden  GardenUseCase
	Tree    TreeUseCase
	Comment CommentUseCase
	Service ServiceUseCase
}

type Repositories struct {
	User    UserRepository
	Tag     TagRepository
	Garden  GardenRepository
	Tree    TreeRepository
	Comment CommentRepository
	Service ServiceRepository
}
