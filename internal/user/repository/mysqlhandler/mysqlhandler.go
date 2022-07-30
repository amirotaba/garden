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

func (m *mysqlUserRepository) Account(n int) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Find(&user).Error; err != nil {
		return []domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountUsername(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountID(id uint) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) AccountType(n int, tp uint) ([]domain.User, error) {
	var user []domain.User
	if err := m.Conn.Limit(n).Where("type = ?", tp).Find(&user).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateUser(user *domain.UserForm) error {
	if err := m.Conn.Model(domain.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateUserType(userType *domain.UserTypeForm) error {
	if err := m.Conn.Model(domain.UserType{}).Where("id = ?", userType.ID).Updates(userType).Error; err != nil {
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

func (m *mysqlUserRepository) ReadTag(n int) ([]domain.Tag, error) {
	var tag []domain.Tag
	if err := m.Conn.Limit(n).Find(&tag).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateTag(tag *domain.TagForm) error {
	if err := m.Conn.Model(domain.Tag{}).Where("id = ?", tag.ID).Updates(tag).Error; err != nil {
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

func (m *mysqlUserRepository) ReadGarden(n int) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Limit(n).Find(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) ReadGardenID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) ReadGardenUID(id uint) ([]domain.Garden, error) {
	var garden []domain.Garden
	if err := m.Conn.Where("user_id = ?", id).First(&garden).Error; err != nil {
		return []domain.Garden{}, err
	}
	return garden, nil
}

func (m *mysqlUserRepository) UpdateGarden(garden *domain.GardenForm) error {
	if err := m.Conn.Model(domain.Garden{}).Where("id = ?", garden.ID).Updates(garden).Error; err != nil {
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

func (m *mysqlUserRepository) ReadLocation(n int) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Limit(n).Find(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlUserRepository) ReadLocationID(id uint) ([]domain.GardenLocation, error) {
	var loc []domain.GardenLocation
	if err := m.Conn.Where("garden_id = ?", id).First(&loc).Error; err != nil {
		return []domain.GardenLocation{}, err
	}
	return loc, nil
}

func (m *mysqlUserRepository) UpdateLocation(loc *domain.GardenLocationForm) error {
	if err := m.Conn.Model(domain.GardenLocation{}).Where("id = ?", loc.ID).Updates(loc).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateGardenType(gardenType *domain.GardenTypeForm) error {
	if err := m.Conn.Model(domain.GardenType{}).Where("id = ?", gardenType.ID).Updates(gardenType).Error; err != nil {
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

func (m *mysqlUserRepository) CreateTree(tree *domain.Tree) error {
	if err := m.Conn.Create(tree).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadTree(n int) ([]domain.Tree, error) {
	var tree []domain.Tree
	if err := m.Conn.Limit(n).Find(&tree).Error; err != nil {
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

func (m *mysqlUserRepository) ReadTreeByType(t uint, n int) ([]domain.Tree, error) {
	var tType []domain.Tree
	if err := m.Conn.Limit(n).Where("type = ?", t).First(&tType).Error; err != nil {
		return []domain.Tree{}, err
	}
	return tType, nil
}

func (m *mysqlUserRepository) UpdateTree(tree *domain.TreeForm) error {
	if err := m.Conn.Model(domain.Tree{}).Where("id = ?", tree.ID).Updates(tree).Error; err != nil {
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

func (m *mysqlUserRepository) UpdateTreeType(treeType *domain.TreeTypeForm) error {
	if err := m.Conn.Model(domain.TreeType{}).Where("id = ?", treeType.ID).Updates(treeType).Error; err != nil {
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

func (m *mysqlUserRepository) CreateComment(comment *domain.Comment) error {
	if err := m.Conn.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadComment(n int) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Limit(n).Find(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlUserRepository) ReadCommentID(id uint, q string, n int) ([]domain.Comment, error) {
	var comment []domain.Comment
	if err := m.Conn.Limit(n).Where(q, id).First(&comment).Error; err != nil {
		return []domain.Comment{}, err
	}
	return comment, nil
}

func (m *mysqlUserRepository) UpdateComment(comment *domain.CommentForm) error {
	if err := m.Conn.Model(domain.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error; err != nil {
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

func (m *mysqlUserRepository) CreateService(service *domain.Service) error {
	if err := m.Conn.Create(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) ReadService() ([]domain.Service, error) {
	var service []domain.Service
	if err := m.Conn.Find(&service).Error; err != nil {
		return []domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlUserRepository) ReadServiceUrl(url string) (domain.Service, error) {
	var service domain.Service
	if err := m.Conn.Where("url = ?", url).First(&service).Error; err != nil {
		return domain.Service{}, err
	}
	return service, nil
}

func (m *mysqlUserRepository) UpdateService(service *domain.ServiceForm) error {
	if err := m.Conn.Model(domain.Service{}).Where("id = ?", service.ID).Updates(service).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) DeleteService(id uint) error {
	var uType domain.Service
	if err := m.Conn.Where("id = ?", id).Delete(&uType).Error; err != nil {
		return err
	}
	return nil
}
