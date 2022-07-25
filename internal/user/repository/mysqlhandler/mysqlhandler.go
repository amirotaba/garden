package mysqlhandler

import (
	"garden/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

//func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
//	return &mysqlUserRepository{
//		Conn: Conn,
//	}
//}

func NewMysqlAdminRepository(Conn *gorm.DB) domain.AdminRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

//func NewMysqlFarmerRepository(Conn *gorm.DB) domain.FarmerRepository {
//	return &mysqlUserRepository{
//		Conn: Conn,
//	}
//}

func (m *mysqlUserRepository) Account() ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountUser(username string) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) SignUp(user *domain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SignIn(form *domain.LoginForm) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", form.Username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) UpdateUser(user *domain.User) error {
	if err := m.Conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteUser(id uint) error {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateUserType(usertype *domain.UserType) error {
	if err := m.Conn.Create(usertype).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadUserType() ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Find(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) ReadUserTypeID(id uint) ([]domain.UserType, error) {
	var uType []domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return []domain.UserType{}, err
	}
	return uType, nil
}

func (m *mysqlUserRepository) UpdateUserType(userType *domain.UserType) error {
	if err := m.Conn.Save(&userType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteUserType(id uint) error {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateTag(tag *domain.Tag) error {
	if err := m.Conn.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadTag() ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Find(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlUserRepository) ReadTagID(id uint) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Where("id = ?", id).First(&tag).Error; err != nil {
		return []domain.Tag{}, err
	}
	return tag, nil
}

func (m *mysqlUserRepository) UpdateTag(tag *domain.Tag) error {
	if err := m.Conn.Save(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteTag(id uint) error {
	var tag domain.Tag
	if err := m.Conn.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateGarden(garden *domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadGarden() ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) ReadGardenID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("user_id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) UpdateGarden(garden *domain.Garden) error {
	if err := m.Conn.Save(&garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteGarden(id uint) error {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateLocation(location *domain.GardenLocation) error {
	if err := m.Conn.Create(location).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadLocation() ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Find(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlUserRepository) ReadLocationID(id uint) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Where("id = ?", id).First(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlUserRepository) UpdateLocation(loc *domain.GardenLocation) error {
	if err := m.Conn.Save(&loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteLocation(id uint) error {
	var loc domain.GardenLocation
	if err := m.Conn.Where("id = ?", id).Delete(&loc).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateGardenType(gardenType *domain.GardenType) error {
	if err := m.Conn.Create(gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadGardenType() ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Find(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlUserRepository) ReadGardenTypeID(id uint) ([]domain.GardenType, error) {
	var gType []domain.GardenType
	if err := m.Conn.Where("id = ?", id).First(&gType).Error; err != nil {
		return []domain.GardenType{}, err
	}
	return gType, nil
}

func (m *mysqlUserRepository) UpdateGardenType(gardenType *domain.GardenType) error {
	if err := m.Conn.Save(&gardenType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteGardenType(id uint) error {
	var gType domain.GardenType
	if err := m.Conn.Where("id = ?", id).Delete(&gType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateTreeType(treeType *domain.TreeType) error {
	if err := m.Conn.Create(treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadTreeType() ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Find(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlUserRepository) ReadTreeTypeID(id uint) ([]domain.TreeType, error) {
	var tType []domain.TreeType
	if err := m.Conn.Where("id = ?", id).First(&tType).Error; err != nil {
		return []domain.TreeType{}, err
	}
	return tType, nil
}

func (m *mysqlUserRepository) UpdateTreeType(treeType *domain.TreeType) error {
	if err := m.Conn.Save(&treeType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteTreeType(id uint) error {
	var tType domain.TreeType
	if err := m.Conn.Where("id = ?", id).Delete(&tType).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateTree(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadTree() ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Find(&tree).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) ReadTreeID(id uint, q string) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Where(q, id).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlUserRepository) ReadTreeByType(t string) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Where("type = ?", t).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlUserRepository) UpdateTree(tree *domain.Tree) error {
	if err := m.Conn.Save(&tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteTree(id uint) error {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) CreateComment(comment *domain.Comment) error {
	if err := m.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadComment() ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Find(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlUserRepository) ReadCommentID(id uint, q string) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Where(q, id).First(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlUserRepository) UpdateComment(comment *domain.Comment) error {
	if err := m.Conn.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteComment(id uint) error {
	var comment domain.Comment
	if err := m.Conn.Where("id = ?", id).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) UserType(id uint) (string, error) {
	var uType domain.UserType
	if err := m.Conn.Where("id = ?", id).First(&uType).Error; err != nil {
		return "", err
	}
	return uType.Name, nil
}
