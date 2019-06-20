package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	infraApi "github.com/ngiyshhk/golang-clean-arch-infra/api"

	usecase "github.com/ngiyshhk/golang-clean-arch-usecase"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"

	"github.com/ngiyshhk/golang-clean-arch-usecase/impl"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ngiyshhk/golang-clean-arch-infra/database"

	"github.com/gin-gonic/gin"
)

func selectUsers(userUsecase usecase.UserUsecase) func(*gin.Context) {
	return func(c *gin.Context) {
		users, err := userUsecase.Select([]int{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func postUser(userUsecase usecase.UserUsecase) func(*gin.Context) {
	return func(c *gin.Context) {
		var json model.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if _, err := userUsecase.Create(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func putUser(userUsecase usecase.UserUsecase) func(*gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var json model.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if _, err := userUsecase.Update(&model.User{
			ID:   id,
			Name: json.Name,
			Age:  json.Age,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func deleteUser(userUsecase usecase.UserUsecase) func(*gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if _, err := userUsecase.Delete(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func getAllLibraries(libraryUsecase usecase.LibraryUsecase) func(*gin.Context) {
	return func(c *gin.Context) {
		libraries, err := libraryUsecase.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, libraries)
	}
}

func main() {
	dbserver := os.Getenv("CLEAN_ARCH_DB_HOST")
	db, err := gorm.Open("mysql", fmt.Sprintf("root:root@tcp(%s)/clean_arch?charset=utf8&parseTime=True", dbserver))
	userUsecase := &impl.UserUsecaseImpl{UserRepository: &database.UserRepositoryImpl{Db: db}}
	libraryUsecase := &impl.LibraryUsecaseImpl{LibraryRepository: &infraApi.LibraryRepositoryImpl{}}

	if err != nil {
		fmt.Printf("db connection error. got=%s\n", err.Error())
		return
	}

	r := gin.Default()
	r.GET("/users", selectUsers(userUsecase))
	r.POST("/users", postUser(userUsecase))
	r.PUT("/users/:id", putUser(userUsecase))
	r.DELETE("/users/:id", deleteUser(userUsecase))
	r.GET("/libraries", getAllLibraries(libraryUsecase))

	r.Run()
}
