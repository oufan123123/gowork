package respository

import (
	"errors"
	"modules/pojo"
	"modules/utils"

	"modules/query"

	"github.com/jinzhu/gorm"
)

type UserRespository struct {
	DB *gorm.DB
}

type UserRespoInterface interface {
	List(req *query.ListQuery) (users []pojo.User, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(user pojo.User) (*pojo.User, error)
	Exist(user pojo.User) *pojo.User
	ExistByUserID(id string) *pojo.User
	ExistByMobile(mobile string) *pojo.User
	Add(user pojo.User) (*pojo.User, error)
	Edit(user pojo.User) (bool, error)
	Deleted(user pojo.User) (bool, error)
}

func (repo *UserRespository) List(req *query.ListQuery) (users []pojo.User, err error) {
	db := repo.DB
	limit, offset := utils.Page(int64(req.PageSize), int64(req.Page))
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order("id desc").Limit(int(limit)).Offset(int(offset)).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRespository) GetTotal(req *query.ListQuery) (total int64, err error) {
	var users []pojo.User
	db := repo.DB
	if req.Where != "" {
		db = db.Where(req.Where)
	}

	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

func (repo *UserRespository) Get(user pojo.User) (*pojo.User, error) {
	db := repo.DB
	if err := db.Where(&user).Find(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (repo *UserRespository) Exist(user pojo.User) *pojo.User {
	db := repo.DB
	if err := db.Find(&user).Where("user_name = ?", user.UserName).Error; err != nil {
		return &user
	}
	return &user
}

func (repo *UserRespository) ExistByUserID(id string) *pojo.User {
	db := repo.DB
	var user pojo.User
	if err := db.Find(&user).Where("user_id = ?", user.UserId).Error; err != nil {
		return &user
	}
	return &user
}

func (repo *UserRespository) ExistByMobile(mobile string) *pojo.User {
	db := repo.DB
	var user pojo.User
	if err := db.Find(&user).Where("mobile = ?", user.Mobile); err != nil {
		return &user
	}
	return &user
}

func (repo *UserRespository) Add(user pojo.User) (*pojo.User, error) {
	db := repo.DB
	// check exist
	if err := repo.Exist(user); err != nil {
		return &user, errors.New("user already exist!")
	}
	if err := db.Create(user); err != nil {
		return nil, errors.New("user add fail")
	}
	return &user, nil
}

func (repo *UserRespository) Edit(user pojo.User) (bool, error) {
	db := repo.DB
	err := db.Model(&user).Where("user_id", user.UserId).Updates(map[string]interface{}{
		"user_name": user.UserName, "mobile": user.Mobile, "address": user.Address,
	}).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func (repo *UserRespository) Delete(user pojo.User) (bool, error) {
	db := repo.DB
	err := db.Model(&user).Where("user_id = ?", user.UserId).Update("is_delete=?", user.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
