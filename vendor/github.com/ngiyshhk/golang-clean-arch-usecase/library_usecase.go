package usecase

import "github.com/ngiyshhk/golang-clean-arch-entity/model"

type LibraryUsecase interface {
	GetAll() (*model.Libraries, error)
}
