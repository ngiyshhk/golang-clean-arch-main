package main

import (
	"fmt"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-infra/mock"
	usecase "github.com/ngiyshhk/golang-clean-arch-usecase"

	"github.com/ngiyshhk/golang-clean-arch-usecase/impl"
)

func main() {
	var hogeUsecase usecase.HogeUsecase = &impl.HogeUsecaseImpl{
		FugaRepository: &mock.FugaRepositoryImpl{},
	}
	fmt.Println(hogeUsecase.Create(&model.Fuga{Id: 123, Name: "aaa", Age: 23}))
}
