package repository

import "github.com/ngiyshhk/golang-clean-arch-entity/model"

type LibraryRepository interface {
	GetAll() (*model.Libraries, error)
}
