package usecase

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

// UserUsecase is hogeusecase
type UserUsecase interface {
	Select(ids []int) ([]*model.User, error)
	Create(entity *model.User) (bool, error)
	Update(entity *model.User) (bool, error)
	Delete(id int) (bool, error)
	Get(id int) (*model.User, error)
}
