package impl

import (
	"errors"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

// UserUsecaseImpl is
type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

// Select is select hoge
func (usecase *UserUsecaseImpl) Select(ids []int) ([]*model.User, error) {
	return usecase.UserRepository.Select(ids)
}

// Create is create hoge
func (usecase *UserUsecaseImpl) Create(entity *model.User) (bool, error) {
	return usecase.UserRepository.Insert(entity)
}

// Update is update hoge
func (usecase *UserUsecaseImpl) Update(entity *model.User) (bool, error) {
	return usecase.UserRepository.Update(entity)
}

// Delete is delete hoge
func (usecase *UserUsecaseImpl) Delete(id int) (bool, error) {
	return usecase.UserRepository.Delete(id)
}

// Get is get a hoge
func (usecase *UserUsecaseImpl) Get(id int) (*model.User, error) {
	entities, err := usecase.UserRepository.Select([]int{id})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, errors.New("not exist")
	}
	return entities[0], nil
}
