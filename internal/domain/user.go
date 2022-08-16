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

type CheckAccessForm struct {
	AccessList string
	ServiceID  uint
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
	Uid        uint
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

type GardenLocRead struct {
	GardenID   string
	PageNumber string
}

type UpdateCommentForm struct {
	Comment CommentForm
	Uid     uint
}

type UserReadForm struct {
	Span   int
	TypeID uint
}

type ReadTreeID struct {
	Query string
	ID    uint
}

type ReadTreeType struct {
	ID   uint
	Span int
}

type ReadComment struct {
	ID    uint
	Query string
	Span  int
}

type UserUseCase interface {
	Create(newUser User) (UserResponse, error)
	SignIn(form *LoginForm) (UserResponse, error)
	Read(form AccountForm) ([]UserResponse, error)
	UserRead(form UserAccountForm) (UserResponse, error)
	Update(user *UserForm) error
	Delete(user *User) error
}

type UserTypeUseCase interface {
	Create(usertype *UserType) error
	Read(id string) ([]UserType, error)
	UpdateAccess(access *AccessForm) error
	Update(usertype *UserTypeForm) error
	Delete(usertype *UserType) error
}

type TagUseCase interface {
	Create(tag *Tag) error
	Read(pageNumber string) ([]Tag, error)
	ReadID(id string) ([]Tag, error)
	Update(tag *TagForm) error
	Delete(tag *Tag) error
}

type GardenUseCase interface {
	Create(garden *Garden) error
	Read(form ReadGardenForm) ([]Garden, error)
	Update(garden *GardenForm) error
	Delete(garden *Garden) error
}

type GardenLocUseCase interface {
	Create(location *GardenLocation) error
	Read(form GardenLocRead) ([]GardenLocation, error)
	Update(loc *GardenLocationForm) error
	Delete(loc *GardenLocation) error
}

type GardenTypeUseCase interface {
	Create(gardenType *GardenType) error
	Read(id string) ([]GardenType, error)
	Update(gardenType *GardenTypeForm) error
	Delete(gardenType *GardenType) error
}

type TreeUseCase interface {
	Create(tree *Tree) error
	Read(form ReadTreeForm) ([]Tree, error)
	ReadUser(form ReadTreeUserForm) ([]Tree, error)
	Update(tree *TreeForm) error
	Delete(tree *Tree) error
}

type TreeTypeUseCase interface {
	Create(treeType *TreeType) error
	Read(id string) ([]TreeType, error)
	Update(treeType *TreeTypeForm) error
	Delete(treeType *TreeType) error
}

type CommentUseCase interface {
	Create(comment *Comment) error
	Read(form ReadCommentForm) ([]Comment, error)
	Update(form *UpdateCommentForm) error
	Delete(form *UpdateCommentForm) error
}

type ServiceUseCase interface {
	Create(service *Service) error
	Read() ([]Service, error)
	Update(usertype *ServiceForm) error
	Delete(service *Service) error
}

type UserRepository interface {
	Create(newUser User) error
	Read(n int) ([]User, error)
	ReadID(id uint) (User, error)
	ReadByType(readForm UserReadForm) ([]User, error)
	ReadUsername(username string) (User, error)
	Update(user *UserForm) error
	Delete(id uint) error
}

type UserTypeRepository interface {
	Create(usertype *UserType) error
	Read() ([]UserType, error)
	ReadID(id uint) (UserType, error)
	Update(userType *UserTypeForm) error
	Delete(id uint) error
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
}

type GardenLocRepository interface {
	Create(location *GardenLocation) error
	Read(n int) ([]GardenLocation, error)
	ReadID(u uint) ([]GardenLocation, error)
	Update(loc *GardenLocationForm) error
	Delete(id uint) error
}

type GardenTypeRepository interface {
	Create(gardenType *GardenType) error
	Read() ([]GardenType, error)
	ReadID(u uint) ([]GardenType, error)
	Update(gardenType *GardenTypeForm) error
	Delete(id uint) error
}

type TreeRepository interface {
	Create(tree *Tree) error
	Read(n int) ([]Tree, error)
	ReadID(readForm ReadTreeID) ([]Tree, error)
	ReadByType(readForm ReadTreeType) ([]Tree, error)
	Update(tree *TreeForm) error
	Delete(id uint) error
}

type TreeTypeRepository interface {
	Create(treeType *TreeType) error
	Read() ([]TreeType, error)
	ReadID(u uint) ([]TreeType, error)
	Update(treeType *TreeTypeForm) error
	Delete(id uint) error
}

type CommentRepository interface {
	Create(comment *Comment) error
	Read(n int) ([]Comment, error)
	ReadID(readForm ReadComment) ([]Comment, error)
	Update(comment CommentForm) error
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
	User       UserUseCase
	UserType   UserTypeUseCase
	Tag        TagUseCase
	Garden     GardenUseCase
	GardenLoc  GardenLocUseCase
	GardenType GardenTypeUseCase
	Tree       TreeUseCase
	TreeType   TreeTypeUseCase
	Comment    CommentUseCase
	Service    ServiceUseCase
}

type Repositories struct {
	User       UserRepository
	UserType   UserTypeRepository
	Tag        TagRepository
	Garden     GardenRepository
	GardenLoc  GardenLocRepository
	GardenType GardenTypeRepository
	Tree       TreeRepository
	TreeType   TreeTypeRepository
	Comment    CommentRepository
	Service    ServiceRepository
}
