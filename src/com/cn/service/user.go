package service

import (
	"errors"
	"modules/pojo"
	"modules/respository"
	"modules/utils"

	"modules/query"
)

type UserService struct {
	Repo *respository.UserRespository
}

type UserSrv interface {
	List(req *query.ListQuery) (users []pojo.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user pojo.User) (*pojo.User, error)
	Exist(user pojo.User) *pojo.User
	Add(user pojo.User) (*pojo.User, error)
	Edit(user pojo.User) (bool, error)
	Delete(id string) (bool, error)
}

func (userService *UserService) List(req *query.ListQuery) (users []pojo.User, err error) {
	return userService.Repo.List(req)
}

func (userService *UserService) GetTotal(req *query.ListQuery) (total int64, err error) {
	return userService.Repo.GetTotal(req)
}

func (userService *UserService) Get(user pojo.User) (*pojo.User, error) {
	return userService.Repo.Get(user)
}

func (userService *UserService) Exist(user pojo.User) *pojo.User {
	return userService.Repo.Exist(user)
}

func (userService *UserService) Add(user pojo.User) (*pojo.User, error) {
	result := userService.Repo.ExistByMobile(user.Mobile)
	if result != nil {
		return result, errors.New("user already exist")
	}
	if user.Password == "" {
		user.Password = utils.MD5("123456")
	}
	user.UserId = "12345"
	user.IsDeleted = false
	user.IsLocked = false

	return userService.Repo.Add(user)

}

func (userService *UserService) Edit(user pojo.User) (bool, error) {
	if user.UserId == "" {
		return false, errors.New("argument wrong")
	}
	exist := userService.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("user does not exist")
	}
	return userService.Repo.Edit(user)
}

func (userService *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("argument wrong")
	}
	result := userService.Repo.ExistByUserID(id)
	if result != nil {
		return false, errors.New("user does not exist")
	}
	result.IsDeleted = true
	return userService.Repo.Delete(*result)
}
