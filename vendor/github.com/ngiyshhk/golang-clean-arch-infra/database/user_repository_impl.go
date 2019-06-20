package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

type lUser struct {
	ID   int
	Name string
	Age  int
}

func (luser *lUser) toUser() *model.User {
	return &model.User{
		ID:   luser.ID,
		Name: luser.Name,
		Age:  luser.Age,
	}
}

// UserRepositoryImpl is
type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Insert is
func (repo *UserRepositoryImpl) Insert(entity *model.User) (bool, error) {
	res := repo.Db.Table("users").Create(&lUser{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	})
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

// Update is
func (repo *UserRepositoryImpl) Update(entity *model.User) (bool, error) {
	var cnt int
	res := repo.Db.Table("users").Where("id = ?", entity.ID).Update(&lUser{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}).Count(&cnt)
	if res.Error != nil {
		return false, res.Error
	}
	if cnt != 1 {
		return false, fmt.Errorf("update target is not exists. id = %d", entity.ID)
	}
	return true, nil
}

// Select is
func (repo *UserRepositoryImpl) Select(ids []int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	lusers := make([]*lUser, 0)

	res := repo.Db.Table("users")
	if len(ids) == 0 {
		res = res.Find(&lusers)
	} else {
		res = res.Where(ids).Find(&lusers)
	}

	if res.Error != nil {
		return users, res.Error
	}
	for _, luser := range lusers {
		users = append(users, luser.toUser())
	}
	return users, nil
}

// Delete is
func (repo *UserRepositoryImpl) Delete(id int) (bool, error) {
	res := repo.Db.Table("users").Where("id = ?", id).Delete(&lUser{})
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}
