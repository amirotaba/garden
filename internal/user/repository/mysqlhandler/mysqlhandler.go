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

func (m *mysqlUserRepository) SignIn(password, email string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Account(username string) (domain.UserResponse, error) {
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
}

func (m *mysqlUserRepository) SearchTree(id int) (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Where("id = ?", id).First(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}

func (m *mysqlUserRepository) ShowGarden() (domain.Garden, error) {
	var garden domain.Garden
	if err := m.Conn.First(&garden).Error; err != nil {
		return domain.Garden{}, err
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

func (m *mysqlUserRepository) AddGarden(garden domain.Garden) error {
	if err := m.Conn.Create(garden).Error; err != nil {
		return err
	}
	return nil
}
func (m *mysqlUserRepository) AddFarmer(farmer domain.Farmer) error {
	if err := m.Conn.Create(farmer).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ShowTrees(id int) (domain.Tree, error) {
	var tree domain.Tree
	if err := m.Conn.Where("farmerid = ?", id).First(&tree).Error; err != nil {
		return domain.Tree{}, err
	}
	return tree, nil
}
func (m *mysqlUserRepository) Addtree(tree domain.Tree) error {
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
}
func (m *mysqlUserRepository) AddAttend(tree domain.Tree) error {
	if err := m.Conn.Where("id = ?", tree.ID).Update("attend", tree.Attend).Error; err != nil {
		return err
	}
	return nil
}
