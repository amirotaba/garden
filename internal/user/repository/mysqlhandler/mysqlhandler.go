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

func (m *mysqlUserRepository) Comment(tree domain.Tree) error {
	if err := m.Conn.Where("id = ?", tree.ID).Update("comment", tree.Comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SearchTree(id int) (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).First(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) SignInFarmer(password, username string) (domain.Farmer, error) {
	var farmer domain.Farmer
	if err := m.Conn.Where("user_name = ?", username).First(&farmer).Error; err != nil {
		return domain.Farmer{}, err
	}
	return farmer, nil
}

func (m *mysqlUserRepository) SignInAdmin(password, username string) (domain.Admin, error) {
	var admin domain.Admin
	if err := m.Conn.Where("user_name = ?", username).First(&admin).Error; err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

func (m *mysqlUserRepository) ShowGarden() ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil

}

func (m *mysqlUserRepository) RemoveGarden(id int) error {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).Delete(&garden).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) AddGarden(garden *domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
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

func (m *mysqlUserRepository) ShowTrees(id int) ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Where("farmer_id = ?", id).Find(&tree).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tree, nil
}
func (m *mysqlUserRepository) AddTree(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}
func (m *mysqlUserRepository) RemoveTree(id int) error {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).Delete(&tree).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateFarmer(id int, trees string) error {
	var farmer domain.Farmer
	if err := m.Conn.Model(farmer).Where("id = ?", id).Update("trees", trees).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) SearchFarmer(id int) (domain.Farmer, error) {
	var farmer domain.Farmer
	if err := m.Conn.Where("id = ?", id).First(&farmer).Error; err != nil {
		return domain.Farmer{}, err
	}
	return farmer, nil
}

func (m *mysqlUserRepository) SearchGarden(id int) (domain.Garden, error) {
	var garden domain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) LastTree() (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Last(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) UpdateGarden(id int, trees string) error {
	var garden domain.Garden
	if err := m.Conn.Model(garden).Where("id = ?", id).Update("trees", trees).Error; err != nil {
		return err
	}
	return nil
}
