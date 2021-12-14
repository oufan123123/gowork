package resp

type User struct {
	Id       string `json:"id"`
	Key      string `json:"key"`
	UserId   string `json:"userId" gorm:"user_id"`
	UserName string `json:"userName" gorm:"user_name"`
	Mobile   string `json:"mobile" gorm:"mobile"`
	Address  string `json:"address" gorm:"address"`
	IsDelete bool   `json:"isDelete" gorm:"is_delete"`
	IsLocked bool   `json:"isLocked" gorm:"is_locked"`
}
