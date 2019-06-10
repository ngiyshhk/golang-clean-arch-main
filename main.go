package main

import (
	"fmt"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

func main() {
	fmt.Println(app.HogeUsecase.Create(&model.Fuga{ID: 123, Name: "aaa", Age: 23}))
}
