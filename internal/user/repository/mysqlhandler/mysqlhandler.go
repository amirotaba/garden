package mysqlhandler

import (
	"garden/internal/domain"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func NewMysqlAdminRepository(Conn *gorm.DB) domain.AdminRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func NewMysqlFarmerRepository(Conn *gorm.DB) domain.FarmerRepository {
	return &mysqlUserRepository{
		Conn: Conn,
	}
}

func (m *mysqlUserRepository) SignUp(user *domain.User) error {
	if err := m.Conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SignInUser(password, username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Account(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Comment(comment *domain.Comment) error {
	if err := m.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SearchTree(id uint) (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).First(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}

//admin

func (m *mysqlUserRepository) SignInAdmin(username, password string) (domain.Admin, error) {
	var admin domain.Admin
	if err := m.Conn.Where("user_name = ?", username).First(&admin).Error; err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

func (m *mysqlUserRepository) AddGarden(garden *domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) AddLocation(location *domain.GardenLocation) error {
	if err := m.Conn.Create(location).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) AddFarmer(farmer *domain.Farmer) error {
	if err := m.Conn.Create(farmer).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ShowGarden() ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil

}

func (m *mysqlUserRepository) RemoveGarden(id uint) error {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeletedBy(id uint, u uint) error {
	var garden domain.Garden
	if err := m.Conn.Model(garden).Where("id = ?", id).Update("user_id", u).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) RemoveGardenLocation(id uint) error {
	var gardenlocation domain.GardenLocation
	if err := m.Conn.Where("garden_id = ?", id).Delete(&gardenlocation).Error; err != nil {
		return err
	}
	return nil
}

//farmer

func (m *mysqlUserRepository) SignInFarmer(username, password string) (domain.Farmer, error) {
	var farmer domain.Farmer
	if err := m.Conn.Where("user_name = ?", username).First(&farmer).Error; err != nil {
		return domain.Farmer{}, err
	}
	return farmer, nil
}

func (m *mysqlUserRepository) AddTree(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) AddAttend(tree domain.Tree) error {
	if err := m.Conn.Model(&tree).Where("id = ?", tree.ID).Update("attend", tree.Attend).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ShowTrees(id uint) ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Where("farmer_id = ?", id).Find(&tree).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) RemoveTree(id uint) error {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) UpdateFarmer(id uint, trees string) error {
	var farmer domain.Farmer
	if err := m.Conn.Model(farmer).Where("id = ?", id).Update("trees", trees).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SearchFarmer(id uint) (domain.Farmer, error) {
	var farmer domain.Farmer
	if err := m.Conn.Where("id = ?", id).First(&farmer).Error; err != nil {
		return domain.Farmer{}, err
	}
	return farmer, nil
}

func (m *mysqlUserRepository) SearchGarden(id uint) (domain.Garden, error) {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) UpdateGarden(id, trees uint) error {
	var garden domain.Garden
	if err := m.Conn.Model(garden).Where("id = ?", id).Update("trees", trees).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) LastTree() (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Last(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) SearchComment(tid uint) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Where("tree_id = ?", tid).Find(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}
