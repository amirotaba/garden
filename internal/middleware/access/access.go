package access

import (
	"garden/internal/domain"
	"garden/internal/features/comment/repository/mysql"
	"garden/internal/features/garden/repository/mysql"
	"garden/internal/features/service/repository/mysql"
	"garden/internal/features/tag/repository/mysql"
	"garden/internal/features/tree/repository/mysql"
	"garden/internal/features/user/repository/mysql"
	"garden/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type (
	Stats struct {
		Access bool `json:"access"`
	}
)

func NewStats() *Stats {
	return &Stats{}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//get uid from token
		uid := jwt.UserID(c)

		//uid := uint(1)
		//connect to database
		dbUser := "root"
		dbPassword := "97216017"
		dbName := "garden"
		dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Connecting to database failed")
		}

		repo := domain.Repositories{
			User:    userRepo.NewMysqlRepository(Db),
			Tag:     tagRepo.NewMysqlRepository(Db),
			Garden:  gardenRepo.NewMysqlRepository(Db),
			Tree:    treeRepo.NewMysqlRepository(Db),
			Comment: commentRepo.NewMysqlRepository(Db),
			Service: serviceRepo.NewMysqlRepository(Db),
		}

		//get the url
		url := c.Request().URL
		x := len(url.RawQuery) + 1
		rawurl := url.Path[x:]

		//get data from database
		sID, err := repo.Service.ReadURL(rawurl)
		if err != nil {
			return err
		}

		u, err := repo.User.ReadID(uid)
		if err != nil {
			return err
		}

		t, err := repo.User.ReadTypeID(u.Type)
		if err != nil {
			return err
		}
		var boolean bool
		List := strings.Split(t.AccessList, ",")
		for _, v := range List {
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			if uint(i) == sID.ID {
				boolean = true
				s.Access = true
			}
		}
		if !boolean {
			return echo.ErrUnauthorized
		}
		if err := next(c); err == nil {
			c.Error(err)
		}
		return nil
	}
}
