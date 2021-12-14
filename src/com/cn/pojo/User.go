package pojo

import "time"

type User struct {
	UserId    string    `json:"userId" gorm:"user_id"`
	UserName  string    `json:"userName" gorm:"user_name"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" binding:"required"`
	Password  string    `json:"password" gorm:"column:password"`
	IsLocked  bool      `json:isLocked gorm:"column:is_locked"`
	Address   string    `json:"address" gorm:"column:address"`
	IsDeleted bool      `json:"isDeleted" gorm:"column:is_deleted"`
	CreateAt  time.Time `json:"createAt" gorm:"column:createAt"`
	UpdateAt  time.Time `json:"updateAt" gorm:"column:updateAt"`
}
