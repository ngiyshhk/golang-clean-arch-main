package impl

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

type LibraryUsecaseImpl struct {
	LibraryRepository repository.LibraryRepository
}

// Select is select hoge
func (usecase *LibraryUsecaseImpl) GetAll() (*model.Libraries, error) {
	return usecase.LibraryRepository.GetAll()
}
