package repository

import "github.com/ngiyshhk/golang-clean-arch-entity/model"

// UserRepository is
type UserRepository interface {
	Insert(entity *model.User) (bool, error)
	Update(entity *model.User) (bool, error)
	Select(ids []int) ([]*model.User, error)
	Delete(id int) (bool, error)
}

// UserRepoMock is
type UserRepoMock struct {
	UserRepository
	MockInsert func(entity *model.User) (bool, error)
	MockUpdate func(entity *model.User) (bool, error)
	MockSelect func(ids []int) ([]*model.User, error)
	MockDelete func(id int) (bool, error)
}

// Insert is
func (m *UserRepoMock) Insert(entity *model.User) (bool, error) {
	return m.MockInsert(entity)
}

// Update is
func (m *UserRepoMock) Update(entity *model.User) (bool, error) {
	return m.MockUpdate(entity)
}

// Select is
func (m *UserRepoMock) Select(ids []int) ([]*model.User, error) {
	return m.MockSelect(ids)
}

// Delete is
func (m *UserRepoMock) Delete(id int) (bool, error) {
	return m.MockDelete(id)
}
